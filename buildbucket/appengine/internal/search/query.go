// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import (
	"container/heap"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"

	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/data/strpair"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/grpc/appstatus"

	"go.chromium.org/luci/buildbucket/appengine/internal/buildid"
	"go.chromium.org/luci/buildbucket/appengine/internal/perm"
	"go.chromium.org/luci/buildbucket/appengine/model"
	pb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/buildbucket/protoutil"
)

const (
	defaultPageSize = 100
	maxPageSize     = 1000
)

var (
	PageTokenRegex = regexp.MustCompile(`^id>\d+$`)
)

// Query is the intermediate to store the arguments for ds search query.
type Query struct {
	Builder             *pb.BuilderID
	Tags                strpair.Map
	Status              pb.Status
	CreatedBy           identity.Identity
	StartTime           time.Time
	EndTime             time.Time
	IncludeExperimental bool
	BuildIDHigh         int64
	BuildIDLow          int64
	Canary              *bool
	PageSize            int32
	PageToken           string
}

// NewQuery builds a Query from pb.SearchBuildsRequest.
// It assumes req is valid, otherwise may panic.
func NewQuery(req *pb.SearchBuildsRequest) *Query {
	if req.GetPredicate() == nil {
		return &Query{
			PageSize:  fixPageSize(req.GetPageSize()),
			PageToken: req.GetPageToken(),
		}
	}

	p := req.Predicate
	q := &Query{
		Builder:             p.GetBuilder(),
		Tags:                protoutil.StringPairMap(p.Tags),
		Status:              p.Status,
		CreatedBy:           identity.Identity(fixCreatedBy(p.CreatedBy)),
		StartTime:           mustTimestamp(p.CreateTime.GetStartTime()),
		EndTime:             mustTimestamp(p.CreateTime.GetEndTime()),
		IncludeExperimental: p.IncludeExperimental,
		PageSize:            fixPageSize(req.PageSize),
		PageToken:           req.PageToken,
	}

	// Filter by gerrit changes.
	for _, change := range p.GerritChanges {
		q.Tags.Add("buildset", protoutil.GerritBuildSet(change))
	}

	// Filter by build range.
	// BuildIds less or equal to 0 means no boundary.
	// Convert BuildRange to [buildLow, buildHigh).
	// Note that unlike buildLow/buildHigh, BuildRange in req encapsulates the fact
	// that build ids are decreasing. So we need to reverse the order.
	if p.Build.GetStartBuildId() > 0 {
		// Add 1 because startBuildId is inclusive and buildHigh is exclusive.
		q.BuildIDHigh = p.Build.GetStartBuildId() + 1
	}
	if p.Build.GetEndBuildId() > 0 {
		// Subtract 1 because endBuildId is exclusive and buildLow is inclusive.
		q.BuildIDLow = p.Build.GetEndBuildId() - 1
	}

	// Filter by canary.
	if p.GetCanary() != pb.Trinary_UNSET {
		q.Canary = proto.Bool(p.GetCanary() == pb.Trinary_YES)
	}
	return q
}

// IndexedTags returns the indexed tags.
func IndexedTags(tags strpair.Map) []string {
	set := make(stringset.Set)
	for k, vals := range tags {
		if k != "buildset" && k != "build_address" {
			continue
		}
		for _, val := range vals {
			set.Add(strpair.Format(k, val))
		}
	}
	return set.ToSortedSlice()
}

// Fetch performs main build search logic.
func (q *Query) Fetch(ctx context.Context) (*pb.SearchBuildsResponse, error) {
	if !buildid.MayContainBuilds(q.StartTime, q.EndTime) {
		return &pb.SearchBuildsResponse{}, nil
	}

	// Verify bucket ACL permission.
	if q.Builder != nil && q.Builder.Bucket != "" {
		if err := perm.HasInBuilder(ctx, perm.BuildsList, q.Builder); err != nil {
			return nil, err
		}
	}

	cpy := *q
	q = &cpy
	q.PageSize = fixPageSize(q.PageSize)
	// Determine which subflow - directly query on Builds or on TagIndex.
	if len(IndexedTags(q.Tags)) != 0 {
		switch res, err := q.fetchOnTagIndex(ctx); {
		case model.TagIndexIncomplete.In(err):
			logging.Warningf(ctx, "Falling back to querying search on builds")
		case err != nil:
			return nil, err
		default:
			return res, nil
		}
	}

	return q.fetchOnBuild(ctx)
}

// fetchOnBuild fetches directly on Build entity.
func (q *Query) fetchOnBuild(ctx context.Context) (*pb.SearchBuildsResponse, error) {
	dq := datastore.NewQuery(model.BuildKind)

	for _, tag := range q.Tags.Format() {
		dq = dq.Eq("tags", tag)
	}

	switch {
	case q.Status == pb.Status_ENDED_MASK:
		dq = dq.Eq("incomplete", false)
	case q.Status != pb.Status_STATUS_UNSPECIFIED:
		dq = dq.Eq("status_v2", q.Status)
	}

	if q.CreatedBy != "" {
		dq = dq.Eq("created_by", q.CreatedBy)
	}
	if q.Canary != nil {
		dq = dq.Eq("canary", *q.Canary)
	}
	if !q.IncludeExperimental {
		dq = dq.Eq("experimental", false)
	}

	idLow, idHigh := q.idRange()
	if idLow != 0 {
		dq = dq.Gte("__key__", datastore.KeyForObj(ctx, &model.Build{ID: idLow}))
	}
	if idHigh != 0 {
		dq = dq.Lt("__key__", datastore.KeyForObj(ctx, &model.Build{ID: idHigh}))
	}

	var queries []*datastore.Query
	var buckets []string
	var err error
	switch {
	case q.Builder.GetBuilder() != "":
		queries = append(queries, dq.Eq("builder_id", protoutil.FormatBuilderID(q.Builder)))
	case q.Builder.GetBucket() != "":
		buckets = []string{protoutil.FormatBucketID(q.Builder.Project, q.Builder.Bucket)}
	default:
		switch buckets, err = perm.BucketsByPerm(ctx, perm.BuildersList, q.Builder.GetProject()); {
		case err != nil:
			return nil, errors.Annotate(err, "error fetching accessible buckets").Err()
		case len(buckets) == 0:
			return &pb.SearchBuildsResponse{}, nil
		}
	}

	for _, bucket := range buckets {
		queries = append(queries, dq.Eq("bucket_id", bucket))
	}

	rsp := &pb.SearchBuildsResponse{}
	logging.Debugf(ctx, "datastore query for FetchOnBuild: %v", queries)
	err = datastore.RunMulti(ctx, queries, func(b *model.Build) error {
		if len(rsp.Builds) >= int(q.PageSize) {
			return datastore.Stop
		}

		// Check the build status again, as the index might be stale.
		if q.Status != pb.Status_STATUS_UNSPECIFIED &&
			q.Status != pb.Status_ENDED_MASK &&
			q.Status != b.Status {
			return nil
		}

		rsp.Builds = append(rsp.Builds, b.ToSimpleBuildProto(ctx))
		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(rsp.Builds) == int(q.PageSize) {
		rsp.NextPageToken = fmt.Sprintf("id>%d", rsp.Builds[q.PageSize-1].Id)
	}

	return rsp, nil
}

// fetchOnTagIndex searches for builds using the TagIndex entities.
func (q *Query) fetchOnTagIndex(ctx context.Context) (*pb.SearchBuildsResponse, error) {
	// Have checked earlier that len(IndexedTags) > 0.
	// Choose the most selective tag to search by.
	indexedTag := IndexedTags(q.Tags)[0]
	k, v := strpair.Parse(indexedTag)

	// Load tag index entries and put them to a min-heap, sorted by build ID.
	entries, err := model.SearchTagIndex(ctx, k, v)
	if err != nil {
		return nil, err
	}

	var eHeap minHeap
	switch filteredEntries, err := q.filterEntries(ctx, entries); {
	case err != nil:
		return nil, err
	case len(filteredEntries) == 0:
		return &pb.SearchBuildsResponse{}, nil
	default:
		eHeap = filteredEntries
	}
	heap.Init(&eHeap)

	// Find the builds.
	results := make([]*pb.Build, 0, q.PageSize) // Ordered by build id by ascending.
	var lastConsideredEntry *model.TagIndexEntry
	inconsistentEntries := 0
	var entriesToFetch []*model.TagIndexEntry
	tags := q.Tags.Format()
	for len(results) < int(q.PageSize) {
		toFetchCount := int(q.PageSize) - len(results)
		entriesToFetch = entriesToFetch[:0]
		for eHeap.Len() > 0 && len(entriesToFetch) < toFetchCount {
			entry := heap.Pop(&eHeap).(*model.TagIndexEntry)
			prev := lastConsideredEntry
			lastConsideredEntry = entry
			// Tolerate duplicates.
			if prev != nil && prev.BuildID == entry.BuildID {
				continue
			}
			entriesToFetch = append(entriesToFetch, entry)
		}

		if len(entriesToFetch) == 0 {
			break
		}

		// Fetch builds
		builds := make([]*model.Build, len(entriesToFetch))
		for i, e := range entriesToFetch {
			builds[i] = &model.Build{ID: e.BuildID}
		}
		// The non-existent builds will be filtered out in the filtering builds for-loop as they have no tags.
		if err := model.GetIgnoreMissing(ctx, builds); err != nil {
			logging.Errorf(ctx, "error fetching builds on fetchOnTagIndex code path : %s", err)
			return nil, errors.Annotate(err, "error fetching builds").Err()
		}

		// Filter builds
		for i, b := range builds {
			buildTags := stringset.NewFromSlice(b.Tags...)
			// Check for inconsistent entries.
			if b.BucketID != entriesToFetch[i].BucketID || !buildTags.Has(indexedTag) {
				logging.Warningf(ctx, "entry with build_id %d is inconsistent", b.ID)
				inconsistentEntries++
				continue
			}
			// Check user-supplied filters.
			if !buildTags.HasAll(tags...) ||
				(q.Status == pb.Status_ENDED_MASK && b.Incomplete) ||
				(q.Status != pb.Status_STATUS_UNSPECIFIED && q.Status != pb.Status_ENDED_MASK && q.Status != b.Status) ||
				(q.CreatedBy != "" && q.CreatedBy != b.CreatedBy) ||
				(q.Builder.GetBuilder() != "" && b.Proto.Builder.Builder != q.Builder.Builder) ||
				(q.Builder.GetProject() != "" && b.Proto.Builder.Project != q.Builder.Project) ||
				(b.Experimental && !q.IncludeExperimental) ||
				(q.Canary != nil && *q.Canary != b.Canary) {
				continue
			}
			results = append(results, b.ToSimpleBuildProto(ctx))
		}
	}
	// TODO(crbug/1090540): add metrics for inconsistentEntries.
	rsp := &pb.SearchBuildsResponse{
		Builds: results,
	}
	if len(results) == int(q.PageSize) && lastConsideredEntry != nil {
		rsp.NextPageToken = fmt.Sprintf("id>%d", lastConsideredEntry.BuildID)
	}
	return rsp, nil
}

// filterEntries filters tag index entries by the build ID ranges and buckets
// conditions in the Query.
func (q *Query) filterEntries(ctx context.Context, entries []*model.TagIndexEntry) ([]*model.TagIndexEntry, error) {
	idLow, idHigh := q.idRange()
	if idHigh == 0 {
		idHigh = int64(uint64(1)<<63 - 1)
	}
	if idLow >= idHigh {
		return nil, nil
	}

	bucketID := protoutil.FormatBucketID(q.Builder.GetProject(), q.Builder.GetBucket())
	preprocessed := make([]*model.TagIndexEntry, 0, len(entries))
	// A cache whether the user has the access permission to buckets.
	hasAccessCache := map[string]bool{}
	for _, e := range entries {
		if e.BuildID < idLow || e.BuildID >= idHigh {
			continue
		}
		// If the bucket in query is not specified, the permission was not checked earlier.
		// In this case, check the permission.
		if q.Builder.GetBucket() == "" {
			has, ok := hasAccessCache[e.BucketID]
			if !ok {
				proj, bkt, _ := protoutil.ParseBucketID(e.BucketID)
				if err := perm.HasInBucket(ctx, perm.BuildsList, proj, bkt); err == nil {
					has = true
				} else {
					status, ok := appstatus.Get(err)
					if !ok || (status.Code() != codes.PermissionDenied && status.Code() != codes.NotFound) {
						return nil, err
					}
				}
				hasAccessCache[e.BucketID] = has
			}
			if !has {
				continue
			}
		} else if bucketID != e.BucketID {
			continue
		}
		preprocessed = append(preprocessed, e)
	}
	return preprocessed, nil
}

// idRange computes the id range from q.BuildIdLow/q.BuildIdHigh, q.StartTime/q.EndTime and q.StartCursor.
// Returning 0 means no boundary.
func (q *Query) idRange() (idLow, idHigh int64) {
	if q.BuildIDLow != 0 || q.BuildIDHigh != 0 {
		idLow, idHigh = q.BuildIDLow, q.BuildIDHigh
	} else {
		idLow, idHigh = buildid.IDRange(q.StartTime, q.EndTime)
	}

	if q.PageToken != "" {
		if minExclusiveID, _ := strconv.ParseInt(q.PageToken[len("id>"):], 10, 64); minExclusiveID+1 > idLow {
			idLow = minExclusiveID + 1
		}
	}
	return
}

// fixPageSize ensures the size is positive and less than or equal to maxPageSize.
func fixPageSize(size int32) int32 {
	switch {
	case size <= 0:
		return defaultPageSize
	case size > maxPageSize:
		return maxPageSize
	default:
		return size
	}
}

// fixPageSize ensures the createdBy identity string is the format "kind:value".
func fixCreatedBy(createdBy string) string {
	if createdBy != "" && !strings.Contains(createdBy, ":") {
		createdBy = fmt.Sprintf("user:%s", createdBy)
	}
	return createdBy
}

// mustTimestamp converts a protobuf timestamp to a time.Time and panics on failures.
// It returns zero time for nil timestamp.
func mustTimestamp(ts *timestamp.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}

	t, err := ptypes.Timestamp(ts)
	if err != nil {
		panic(err)
	}
	return t
}

// minHeap holds a slice of TagIndexEntry and implements heap.Interface.
type minHeap []*model.TagIndexEntry

var _ heap.Interface = &minHeap{}

func (m minHeap) Len() int { return len(m) }

func (m minHeap) Less(i, j int) bool { return m[i].BuildID < m[j].BuildID }

func (m minHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(*model.TagIndexEntry))
}

func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	item := old[n-1]
	*m = old[0 : n-1]
	return item
}
