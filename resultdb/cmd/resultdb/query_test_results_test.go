// Copyright 2019 The LUCI Authors.
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

package main

import (
	"fmt"
	"testing"

	"cloud.google.com/go/spanner"
	durpb "github.com/golang/protobuf/ptypes/duration"
	pb "go.chromium.org/luci/resultdb/proto/rpc/v1"

	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/resultdb/internal/span"
	"go.chromium.org/luci/resultdb/internal/testutil"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func insertInv(id span.InvocationID, included ...span.InvocationID) []*spanner.Mutation {
	t := testclock.TestRecentTimeUTC
	ms := []*spanner.Mutation{span.InsertMap("Invocations", map[string]interface{}{
		"InvocationId":                      id,
		"State":                             pb.Invocation_COMPLETED,
		"Realm":                             "",
		"UpdateToken":                       "",
		"InvocationExpirationTime":          t,
		"InvocationExpirationWeek":          t,
		"ExpectedTestResultsExpirationTime": t,
		"ExpectedTestResultsExpirationWeek": t,
		"CreateTime":                        t,
		"Deadline":                          t,
		"FinalizeTime":                      t,
	})}
	for _, incl := range included {
		ms = append(ms, testutil.InsertInclusion(id, incl))
	}
	return ms
}
func TestFetchGraph(t *testing.T) {
	Convey(`TestInclude`, t, func() {
		ctx := testutil.SpannerTestContext(t)

		read := func(roots ...span.InvocationID) (map[span.InvocationID]*pb.Invocation, error) {
			txn := span.Client(ctx).ReadOnlyTransaction()
			defer txn.Close()
			return span.ReadReachableInvocations(ctx, txn, roots...)
		}

		mustReadIDs := func(roots ...span.InvocationID) []span.InvocationID {
			invs, err := read(roots...)
			So(err, ShouldBeNil)
			ids := make([]span.InvocationID, 0, len(invs))
			for id := range invs {
				ids = append(ids, id)
			}
			span.SortInvocationIDs(ids)
			return ids
		}

		Convey(`fetch nothing`, func() {
			So(mustReadIDs(), ShouldBeEmpty)
		})

		Convey(`not found`, func() {
			_, err := read("inv")
			So(err, ShouldErrLike, `"invocations/inv" not found`)
		})

		Convey(`a -> []`, func() {
			testutil.MustApply(ctx, insertInv("a")...)
			So(mustReadIDs("a"), ShouldResemble, []span.InvocationID{"a"})
		})

		Convey(`a -> [b, c]`, func() {
			testutil.MustApply(ctx, insertInv("b")...)
			testutil.MustApply(ctx, insertInv("c")...)
			testutil.MustApply(ctx, insertInv("a", "b", "c")...)
			So(mustReadIDs("a"), ShouldResemble, []span.InvocationID{"a", "b", "c"})
		})

		Convey(`a -> b -> c`, func() {
			testutil.MustApply(ctx, insertInv("c")...)
			testutil.MustApply(ctx, insertInv("b", "c")...)
			testutil.MustApply(ctx, insertInv("a", "b")...)
			So(mustReadIDs("a"), ShouldResemble, []span.InvocationID{"a", "b", "c"})
		})
	})
}

// BenchmarkChainFetch measures performance of a fetching a graph
// with a 10 linear inclusions.
func BenchmarkChainFetch(b *testing.B) {
	ctx := testutil.SpannerTestContext(b)
	client := span.Client(ctx)

	var ms []*spanner.Mutation
	var prev span.InvocationID
	for i := 0; i < 10; i++ {
		var included []span.InvocationID
		if prev != "" {
			included = append(included, prev)
		}
		id := span.InvocationID(fmt.Sprintf("inv%d", i))
		prev = id
		ms = append(ms, insertInv(id, included...)...)
	}

	if _, err := client.Apply(ctx, ms); err != nil {
		b.Fatal(err)
	}

	read := func() {
		txn := span.Client(ctx).ReadOnlyTransaction()
		defer txn.Close()
		_, err := span.ReadReachableInvocations(ctx, txn, prev)
		if err != nil {
			b.Fatal(err)
		}
	}

	// Run fetch a few times before starting measuring.
	for i := 0; i < 5; i++ {
		read()
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		read()
	}
}
func TestValidateQueryTestResultsRequest(t *testing.T) {
	t.Parallel()
	Convey(`Valid`, t, func() {
		err := validateQueryTestResultsRequest(&pb.QueryTestResultsRequest{
			Predicate: &pb.TestResultPredicate{
				Invocation: &pb.InvocationPredicate{
					RootPredicate: &pb.InvocationPredicate_Name{Name: "invocations/x"},
				},
			},
			PageSize:     50,
			MaxStaleness: &durpb.Duration{Seconds: 60},
		})
		So(err, ShouldBeNil)
	})

	Convey(`invalid predicate`, t, func() {
		err := validateQueryTestResultsRequest(&pb.QueryTestResultsRequest{
			Predicate: &pb.TestResultPredicate{
				Invocation: &pb.InvocationPredicate{
					RootPredicate: &pb.InvocationPredicate_Name{Name: "xxxxxxxxxxxxx"},
				},
			},
		})
		So(err, ShouldErrLike, `predicate: invocation: name: does not match`)
	})
}