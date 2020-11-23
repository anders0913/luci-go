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

package updater

import (
	"context"
	"testing"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/gerrit/gerritfake"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestRelatedChangeProcessing(t *testing.T) {
	t.Parallel()

	Convey("setGitDeps works", t, func() {
		ctx := context.Background()
		f := fetcher{
			change:      111,
			host:        "host",
			newSnapshot: &changelist.Snapshot{Kind: &changelist.Snapshot_Gerrit{Gerrit: &changelist.Gerrit{}}},
		}

		Convey("No related changes", func() {
			err := f.setGitDeps(ctx, nil)
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldBeNil)

			err = f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{})
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Just itself", func() {
			// This isn't happening today, but CV shouldn't choke if Gerrit changes.
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(111, 3, 3), // No parents.
			})
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldBeNil)

			err = f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(111, 3, 3, "107_2"),
			})
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Has related, but no deps", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(111, 3, 3, "107_2"),
				gerritfake.RelatedChange(114, 1, 3, "111_3"),
				gerritfake.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Has related, but lacking this change", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(114, 1, 3, "111_3"),
				gerritfake.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldErrLike, "Unexpected Gerrit.GetRelatedChangesResponse")
		})
		Convey("Has related, and several times itself", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(111, 2, 2, "107_2"),
				gerritfake.RelatedChange(111, 3, 3, "107_2"),
				gerritfake.RelatedChange(114, 1, 3, "111_3"),
			})
			So(err, ShouldErrLike, "Unexpected Gerrit.GetRelatedChangesResponse")
		})

		Convey("1 parent", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(107, 1, 3, "104_2"),
				gerritfake.RelatedChange(111, 3, 3, "107_1"),
				gerritfake.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 107, Immediate: true},
			})
		})

		Convey("Diamond", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(103, 2, 2),
				gerritfake.RelatedChange(104, 2, 2, "103_2"),
				gerritfake.RelatedChange(107, 1, 3, "104_2"),
				gerritfake.RelatedChange(108, 1, 3, "104_2"),
				gerritfake.RelatedChange(111, 3, 3, "107_1", "108_1"),
				gerritfake.RelatedChange(114, 1, 3, "111_3"),
				gerritfake.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 107, Immediate: true},
				{Change: 108, Immediate: true},
				{Change: 104, Immediate: false},
				{Change: 103, Immediate: false},
			})
		})

		Convey("Same revision, different changes", func() {
			c104 := gerritfake.RelatedChange(104, 1, 1, "103_2")
			c105 := gerritfake.RelatedChange(105, 1, 1, "103_2")
			c105.GetCommit().Id = c104.GetCommit().GetId()
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(103, 2, 2),
				c104,
				c105, // should be ignored, somewhat arbitrarily.
				gerritfake.RelatedChange(111, 3, 3, "104_1"),
			})
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 104, Immediate: true},
				{Change: 103, Immediate: false},
			})
		})

		Convey("2 parents which are the same change at different revisions", func() {
			// Actually happened, see https://crbug.com/988309.
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gerritfake.RelatedChange(104, 1, 2, "long-ago-merged1"),
				gerritfake.RelatedChange(107, 1, 1, "long-ago-merged2"),
				gerritfake.RelatedChange(104, 2, 2, "107_1"),
				gerritfake.RelatedChange(111, 3, 3, "104_1", "104_2"),
			})
			So(err, ShouldBeNil)
			So(f.newSnapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 104, Immediate: true},
				{Change: 107, Immediate: false},
			})
		})
	})
}