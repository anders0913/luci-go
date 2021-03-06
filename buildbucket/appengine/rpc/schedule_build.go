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

package rpc

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"

	"go.chromium.org/luci/cipd/common"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/proto/mask"
	"go.chromium.org/luci/grpc/appstatus"

	"go.chromium.org/luci/buildbucket/appengine/internal/perm"
	"go.chromium.org/luci/buildbucket/appengine/model"
	pb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/buildbucket/protoutil"
)

// validateExecutable validates the given executable.
func validateExecutable(exe *pb.Executable) error {
	var err error
	switch {
	case exe.GetCipdPackage() != "":
		return errors.Reason("cipd_package must not be specified").Err()
	case exe.GetCipdVersion() != "" && teeErr(common.ValidateInstanceVersion(exe.CipdVersion), &err) != nil:
		return errors.Annotate(err, "cipd_version").Err()
	default:
		return nil
	}
}

// validateSchedule validates the given request.
func validateSchedule(req *pb.ScheduleBuildRequest) error {
	var err error
	switch {
	case strings.Contains(req.GetRequestId(), "/"):
		return errors.Reason("request_id cannot contain '/'").Err()
	case req.GetBuilder() == nil && req.GetTemplateBuildId() == 0:
		return errors.Reason("builder or template_build_id is required").Err()
	case req.Builder != nil && teeErr(protoutil.ValidateRequiredBuilderID(req.Builder), &err) != nil:
		return errors.Annotate(err, "builder").Err()
	case teeErr(validateExecutable(req.Exe), &err) != nil:
		return errors.Annotate(err, "exe").Err()
	case req.GitilesCommit != nil && teeErr(validateCommitWithRef(req.GitilesCommit), &err) != nil:
		return errors.Annotate(err, "gitiles_commit").Err()
	case teeErr(validateTags(req.Tags, TagNew), &err) != nil:
		return errors.Annotate(err, "tags").Err()
	case req.Priority < 0 || req.Priority > 255:
		return errors.Reason("priority must be in [0, 255]").Err()
	}

	// TODO(crbug/1042991): Validate Properties, Gerrit Changes, Dimensions, Notify.
	return nil
}

// templateBuildMask enumerates properties to read from template builds. See
// scheduleRequestFromTemplate.
var templateBuildMask = mask.MustFromReadMask(
	&pb.Build{},
	"builder",
	"canary",
	"critical",
	"exe",
	"input.experimental",
	"input.gerrit_changes",
	"input.gitiles_commit",
	"input.properties",
	"tags",
)

// scheduleRequestFromTemplate returns a request with fields populated by the
// given template_build_id if there is one. Fields set in the request override
// fields populated from the template. Does not modify the incoming request.
func scheduleRequestFromTemplate(ctx context.Context, req *pb.ScheduleBuildRequest) (*pb.ScheduleBuildRequest, error) {
	if req.GetTemplateBuildId() == 0 {
		return req, nil
	}

	bld, err := getBuild(ctx, req.TemplateBuildId)
	if err != nil {
		return nil, err
	}
	if err := perm.HasInBuilder(ctx, perm.BuildsGet, bld.Proto.Builder); err != nil {
		return nil, err
	}

	b := bld.ToSimpleBuildProto(ctx)
	if err := model.LoadBuildDetails(ctx, templateBuildMask, b); err != nil {
		return nil, err
	}

	ret := &pb.ScheduleBuildRequest{
		Builder:       b.Builder,
		Critical:      b.Critical,
		Exe:           b.Exe,
		GerritChanges: b.Input.GerritChanges,
		GitilesCommit: b.Input.GitilesCommit,
		Properties:    b.Input.Properties,
		Tags:          b.Tags,
	}

	// Convert bool to the corresponding pb.Trinary values.
	ret.Canary = pb.Trinary_NO
	ret.Experimental = pb.Trinary_NO
	if b.Canary {
		ret.Canary = pb.Trinary_YES
	}
	if b.Input.Experimental {
		ret.Experimental = pb.Trinary_YES
	}

	// proto.Merge concatenates repeated fields. Here the desired behavior is replacement,
	// so clear slices from the return value before merging, if specified in the request.
	if req.Exe != nil {
		ret.Exe = nil
	}
	if len(req.GerritChanges) > 0 {
		ret.GerritChanges = nil
	}
	if req.Properties != nil {
		ret.Properties = nil
	}
	if len(req.Tags) > 0 {
		ret.Tags = nil
	}
	proto.Merge(ret, req)
	ret.TemplateBuildId = 0
	return ret, nil
}

// ScheduleBuild handles a request to schedule a build. Implements pb.BuildsServer.
func (*Builds) ScheduleBuild(ctx context.Context, req *pb.ScheduleBuildRequest) (*pb.Build, error) {
	var err error
	if err = validateSchedule(req); err != nil {
		return nil, appstatus.BadRequest(err)
	}
	if req, err = scheduleRequestFromTemplate(ctx, req); err != nil {
		return nil, err
	}
	if err = perm.HasInBucket(ctx, perm.BuildsAdd, req.Builder.Project, req.Builder.Bucket); err != nil {
		return nil, err
	}

	return nil, nil
}
