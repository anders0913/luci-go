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
	"sort"
	"sync/atomic"
	"testing"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"

	cfgpb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/cv/internal/changelist"
	"go.chromium.org/luci/cv/internal/cvtesting"
	gf "go.chromium.org/luci/cv/internal/gerrit/gerritfake"
	"go.chromium.org/luci/cv/internal/gerrit/gobmap"
	"go.chromium.org/luci/cv/internal/prjmanager/pmtest"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestRelatedChangeProcessing(t *testing.T) {
	t.Parallel()

	Convey("setGitDeps works", t, func() {
		ctx := context.Background()
		f := fetcher{
			change: 111,
			host:   "host",
			toUpdate: changelist.UpdateFields{
				Snapshot: &changelist.Snapshot{Kind: &changelist.Snapshot_Gerrit{Gerrit: &changelist.Gerrit{}}},
			},
		}

		Convey("No related changes", func() {
			err := f.setGitDeps(ctx, nil)
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)

			err = f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Just itself", func() {
			// This isn't happening today, but CV shouldn't choke if Gerrit changes.
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(111, 3, 3), // No parents.
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)

			err = f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(111, 3, 3, "107_2"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Has related, but no deps", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(111, 3, 3, "107_2"),
				gf.RelatedChange(114, 1, 3, "111_3"),
				gf.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
		})

		Convey("Has related, but lacking this change", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(114, 1, 3, "111_3"),
				gf.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldErrLike, "Unexpected Gerrit.GetRelatedChangesResponse")
		})
		Convey("Has related, and several times itself", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(111, 2, 2, "107_2"),
				gf.RelatedChange(111, 3, 3, "107_2"),
				gf.RelatedChange(114, 1, 3, "111_3"),
			})
			So(err, ShouldErrLike, "Unexpected Gerrit.GetRelatedChangesResponse")
		})

		Convey("1 parent", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(107, 1, 3, "104_2"),
				gf.RelatedChange(111, 3, 3, "107_1"),
				gf.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 107, Immediate: true},
			})
		})

		Convey("Diamond", func() {
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(103, 2, 2),
				gf.RelatedChange(104, 2, 2, "103_2"),
				gf.RelatedChange(107, 1, 3, "104_2"),
				gf.RelatedChange(108, 1, 3, "104_2"),
				gf.RelatedChange(111, 3, 3, "107_1", "108_1"),
				gf.RelatedChange(114, 1, 3, "111_3"),
				gf.RelatedChange(117, 2, 2, "114_1"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 107, Immediate: true},
				{Change: 108, Immediate: true},
				{Change: 104, Immediate: false},
				{Change: 103, Immediate: false},
			})
		})

		Convey("Same revision, different changes", func() {
			c104 := gf.RelatedChange(104, 1, 1, "103_2")
			c105 := gf.RelatedChange(105, 1, 1, "103_2")
			c105.GetCommit().Id = c104.GetCommit().GetId()
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(103, 2, 2),
				c104,
				c105, // should be ignored, somewhat arbitrarily.
				gf.RelatedChange(111, 3, 3, "104_1"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 104, Immediate: true},
				{Change: 103, Immediate: false},
			})
		})

		Convey("2 parents which are the same change at different revisions", func() {
			// Actually happened, see https://crbug.com/988309.
			err := f.setGitDeps(ctx, []*gerritpb.GetRelatedChangesResponse_ChangeAndCommit{
				gf.RelatedChange(104, 1, 2, "long-ago-merged1"),
				gf.RelatedChange(107, 1, 1, "long-ago-merged2"),
				gf.RelatedChange(104, 2, 2, "107_1"),
				gf.RelatedChange(111, 3, 3, "104_1", "104_2"),
			})
			So(err, ShouldBeNil)
			So(f.toUpdate.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto, []*changelist.GerritGitDep{
				{Change: 104, Immediate: true},
				{Change: 107, Immediate: false},
			})
		})
	})
}

func TestUpdateCLWorks(t *testing.T) {
	t.Parallel()

	Convey("Updating CL works", t, func() {
		ct := cvtesting.Test{}
		ctx, cancel := ct.SetUp()
		defer cancel()
		const lProject = "infra"
		const gHost = "chromium-review.example.com"
		const gHostInternal = "internal-review.example.com"
		const gRepo = "depot_tools"

		ct.Cfg.Create(ctx, lProject, singleRepoConfig(gHost, gRepo))
		gobmap.Update(ctx, lProject)

		Convey("No access or permission denied", func() {
			assertDependentMetaOnly := func(change int) {
				cl := getCL(ctx, gHost, change)
				So(cl.Snapshot, ShouldBeNil)
				So(cl.ApplicableConfig, ShouldBeNil)
				So(cl.DependentMeta.GetByProject()[lProject].GetUpdateTime().AsTime(),
					ShouldResemble, ct.Clock.Now().UTC())
			}
			So(refreshExternal(ctx, lProject, gHost, 404, time.Time{}, 0), ShouldBeNil)
			assertDependentMetaOnly(404)
			So(refreshExternal(ctx, lProject, gHost, 403, time.Time{}, 0), ShouldBeNil)
			assertDependentMetaOnly(403)
		})

		Convey("Unhandled Gerrit error results in no CL update", func() {
			ci500 := gf.CI(500, gf.Project(gRepo), gf.Ref("refs/heads/main"))
			Convey("fail to fetch change details", func() {
				ct.GFake.AddFrom(gf.WithCIs(gHost, err5xx, ci500))
				So(refreshExternal(ctx, lProject, gHost, 500, time.Time{}, 0), ShouldErrLike, "boo")
				cl := getCL(ctx, gHost, 500)
				So(cl, ShouldBeNil)
			})

			Convey("fail to get filelist", func() {
				ct.GFake.AddFrom(gf.WithCIs(gHost, okThenErr5xx(), ci500))
				So(refreshExternal(ctx, lProject, gHost, 500, time.Time{}, 0), ShouldErrLike, "boo")
				cl := getCL(ctx, gHost, 500)
				So(cl, ShouldBeNil)
			})
		})

		Convey("CL hint must actually exist", func() {
			So(refreshExternal(ctx, lProject, gHost, 123, time.Time{}, 848484881),
				ShouldErrLike, "clidHint 848484881 doesn't refer to an existing CL")
		})

		Convey("Fetch for the first time", func() {
			ci := gf.CI(123, gf.Project(gRepo), gf.Ref("refs/heads/main"),
				gf.Files("a.cpp", "c/b.py"), gf.Desc("T.\n\nCq-Depend: 101"))
			ciParent := gf.CI(122, gf.Desc("Z\n\nCq-Depend: must-be-ignored:47"))
			ciGrandpa := gf.CI(121, gf.Desc("Z\n\nCq-Depend: must-be-ignored:46"))
			ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLPublic(), ci, ciParent, ciGrandpa))
			ct.GFake.SetDependsOn(gHost, ci, ciParent)
			ct.GFake.SetDependsOn(gHost, ciParent, ciGrandpa)

			So(refreshExternal(ctx, lProject, gHost, 123, time.Time{}, 0), ShouldBeNil)
			cl := getCL(ctx, gHost, 123)
			So(cl.ApplicableConfig.GetUpdateTime().AsTime(), ShouldResemble, ct.Clock.Now().UTC())
			So(cl.ApplicableConfig.HasOnlyProject(lProject), ShouldBeTrue)
			So(cl.Snapshot.GetGerrit().GetHost(), ShouldEqual, gHost)
			So(cl.Snapshot.GetGerrit().Info.GetProject(), ShouldEqual, gRepo)
			So(cl.Snapshot.GetGerrit().Info.GetRef(), ShouldEqual, "refs/heads/main")
			So(cl.Snapshot.GetGerrit().GetFiles(), ShouldResemble, []string{"a.cpp", "c/b.py"})
			So(cl.Snapshot.GetLuciProject(), ShouldEqual, lProject)
			So(cl.Snapshot.GetExternalUpdateTime(), ShouldResembleProto, ci.GetUpdated())
			So(cl.Snapshot.GetGerrit().GetGitDeps(), ShouldResembleProto,
				[]*changelist.GerritGitDep{
					{Change: 122, Immediate: true},
					{Change: 121},
				})
			So(cl.Snapshot.GetGerrit().GetSoftDeps(), ShouldResembleProto,
				[]*changelist.GerritSoftDep{
					{Change: 101, Host: gHost},
				})

			// Each of the dep should have an existing CL + a task schedule.
			expectedDeps := []*changelist.Dep{
				{Clid: int64(getCL(ctx, gHost, 122).ID), Kind: changelist.DepKind_HARD},
				{Clid: int64(getCL(ctx, gHost, 121).ID), Kind: changelist.DepKind_SOFT},
				{Clid: int64(getCL(ctx, gHost, 101).ID), Kind: changelist.DepKind_SOFT},
			}
			sort.Slice(expectedDeps, func(i, j int) bool {
				return expectedDeps[i].GetClid() < expectedDeps[j].GetClid()
			})
			So(cl.Snapshot.GetDeps(), ShouldResembleProto, expectedDeps)
			expectedTasks := []*RefreshGerritCL{
				{
					LuciProject: lProject,
					Host:        gHost,
					Change:      101,
					ClidHint:    int64(getCL(ctx, gHost, 101).ID),
				},
				{
					LuciProject: lProject,
					Host:        gHost,
					Change:      121,
					ClidHint:    int64(getCL(ctx, gHost, 121).ID),
				},
				{
					LuciProject: lProject,
					Host:        gHost,
					Change:      122,
					ClidHint:    int64(getCL(ctx, gHost, 122).ID),
				},
			}
			So(sortedRefreshTasks(ct), ShouldResembleProto, expectedTasks)
			// Project Manager should be notified.
			// NOTE: since PM won't be consuming these notifications in this test, we
			// keep a list of notifications and assert against the accumulated list.
			pmNotifications := []string{lProject}
			So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, pmNotifications)

			// Simulate Gerrit change being updated with +1s timestamp.
			ct.GFake.MutateChange(gHost, 123, func(c *gf.Change) {
				c.Info.Updated.Seconds++
			})

			Convey("skips update with updatedHint", func() {
				updatedHint := cl.Snapshot.GetExternalUpdateTime().AsTime()
				So(refreshExternal(ctx, lProject, gHost, 123, updatedHint, 0), ShouldBeNil)
				So(getCL(ctx, gHost, 123).EVersion, ShouldEqual, cl.EVersion)
				So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, pmNotifications)
			})

			Convey("Don't update iff fetched less recent than updatedHint ", func() {
				// Set expectation that Gerrit serves change with >=+1m timestamp.
				updatedHint := cl.Snapshot.GetExternalUpdateTime().AsTime().Add(time.Minute)
				err := refreshExternal(ctx, lProject, gHost, 123, updatedHint, 0)
				So(err, ShouldErrLike, "stale Gerrit data")
				So(transient.Tag.In(err), ShouldBeTrue)
				So(getCL(ctx, gHost, 123).EVersion, ShouldEqual, cl.EVersion)
				So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, pmNotifications)
			})

			Convey("Heeds updatedHint and updates the CL", func() {
				// Set expectation that Gerrit serves change with >=+1ms timestamp.
				updatedHint := cl.Snapshot.GetExternalUpdateTime().AsTime().Add(time.Millisecond)
				ct.GFake.MutateChange(gHost, 123, func(c *gf.Change) {
					// Only ChangeInfo but not ListFiles and GetRelatedChanges RPCs should
					// be called. So, ensure 2+ RPCs return 5xx.
					c.ACLs = okThenErr5xx()
				})
				So(refreshExternal(ctx, lProject, gHost, 123, updatedHint, 0), ShouldBeNil)
				cl2 := getCL(ctx, gHost, 123)
				So(cl2.EVersion, ShouldEqual, cl.EVersion+1)
				So(cl2.Snapshot.GetExternalUpdateTime().AsTime(), ShouldResemble,
					cl.Snapshot.GetExternalUpdateTime().AsTime().Add(time.Second))
				pmNotifications = append(pmNotifications, lProject)
				So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, pmNotifications)

				Convey("New revision doesn't re-use files & related changes", func() {
					// Stay within the same blindRefreshInterval for de-duping refresh
					// tasks of dependencies.
					ct.Clock.Add(blindRefreshInterval - 2*time.Second)
					ct.GFake.MutateChange(gHost, 123, func(c *gf.Change) {
						c.ACLs = gf.ACLPublic()
						// Simulate new patchset which no longer has GerritGitDeps.
						gf.PS(10)(c.Info)
						gf.Files("z.zz")(c.Info)
						// 101 is from before, internal:477 is new.
						gf.Desc("T\n\nCq-Depend: 101,internal:477")(c.Info)
						gf.Updated(ct.Clock.Now())(c.Info)
					})

					So(refreshExternal(ctx, lProject, gHost, 123, time.Time{}, 0), ShouldBeNil)
					cl3 := getCL(ctx, gHost, 123)
					So(cl3.EVersion, ShouldEqual, cl2.EVersion+1)
					So(cl3.Snapshot.GetExternalUpdateTime().AsTime(), ShouldResemble, ct.Clock.Now().UTC())
					So(cl3.Snapshot.GetGerrit().GetFiles(), ShouldResemble, []string{"z.zz"})
					So(cl3.Snapshot.GetGerrit().GetGitDeps(), ShouldBeNil)
					So(cl3.Snapshot.GetGerrit().GetSoftDeps(), ShouldResembleProto,
						[]*changelist.GerritSoftDep{
							{Change: 101, Host: gHost},
							{Change: 477, Host: gHostInternal},
						})
					// For each dep, a task should have been created, but 101 should have
					// been de-duped with an earlier one. So, only 1 new task for 477:
					So(sortedRefreshTasks(ct), ShouldResembleProto, append(expectedTasks,
						&RefreshGerritCL{
							LuciProject: lProject,
							Host:        gHostInternal,
							Change:      477,
							ClidHint:    int64(getCL(ctx, gHostInternal, 477).ID),
						},
					))
					pmNotifications = append(pmNotifications, lProject)
					So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, pmNotifications)
				})
			})

			Convey("No longer watched", func() {
				ct.Clock.Add(time.Second)
				ct.Cfg.Update(ctx, lProject, singleRepoConfig(gHost, "another/repo"))
				gobmap.Update(ctx, lProject)
				So(refreshExternal(ctx, lProject, gHost, 123, time.Time{}, 0), ShouldBeNil)
				cl2 := getCL(ctx, gHost, 123)
				So(cl2.EVersion, ShouldEqual, cl.EVersion+1)
				// Snapshot is preserved (handy, if this is temporal misconfiguration).
				So(cl2.Snapshot, ShouldResembleProto, cl.Snapshot)
				So(cl2.ApplicableConfig, ShouldResembleProto, &changelist.ApplicableConfig{
					UpdateTime: timestamppb.New(ct.Clock.Now()),
				})
				// PM is still notified.
				pmNotifications = append(pmNotifications, lProject)
				So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, pmNotifications)
			})

			Convey("Watched by a diff project", func() {
				ct.Clock.Add(time.Second)
				const lProject2 = "proj-2"
				ct.Cfg.Update(ctx, lProject, singleRepoConfig(gHost, "another repo"))
				ct.Cfg.Create(ctx, lProject2, singleRepoConfig(gHost, gRepo))
				gobmap.Update(ctx, lProject)
				gobmap.Update(ctx, lProject2)

				// Use a hint that'd normally prevent an update.
				updateHint := cl.Snapshot.GetExternalUpdateTime().AsTime()

				Convey("with access", func() {
					So(refreshExternal(ctx, lProject2, gHost, 123, updateHint, 0), ShouldBeNil)
					cl2 := getCL(ctx, gHost, 123)
					So(cl2.EVersion, ShouldEqual, cl.EVersion+1)
					So(cl2.Snapshot.GetLuciProject(), ShouldEqual, lProject2)
					So(cl2.Snapshot.GetExternalUpdateTime(), ShouldResemble,
						ct.GFake.GetChange(gHost, 123).Info.GetUpdated())
					So(cl2.ApplicableConfig.HasOnlyProject(lProject2), ShouldBeTrue)
					// A different PM is notified.
					pmNotifications = append(pmNotifications, lProject2)
					So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, pmNotifications)
				})

				Convey("without access", func() {
					ct.GFake.MutateChange(gHost, 123, func(c *gf.Change) {
						c.ACLs = gf.ACLRestricted("not-lProject2")
					})
					So(refreshExternal(ctx, lProject2, gHost, 123, updateHint, 0), ShouldBeNil)
					cl2 := getCL(ctx, gHost, 123)
					So(cl2.EVersion, ShouldEqual, cl.EVersion+1)
					// Snapshot is kept as is, including its ExternalUpdateTime.
					So(cl2.Snapshot, ShouldResembleProto, cl.Snapshot)
					So(cl2.ApplicableConfig.HasOnlyProject(lProject2), ShouldBeTrue)
					So(cl2.DependentMeta.GetByProject()[lProject2].GetNoAccess(), ShouldBeTrue)
					// A different PM is notified anyway.
					pmNotifications = append(pmNotifications, lProject2)
					So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, pmNotifications)
				})
			})
		})

		Convey("Fetch dep after bare CL was crated", func() {
			eid, err := changelist.GobID(gHost, 101)
			So(err, ShouldBeNil)
			cl, err := eid.GetOrInsert(ctx, func(cl *changelist.CL) {})
			So(err, ShouldBeNil)
			So(cl.EVersion, ShouldEqual, 1)

			ci := gf.CI(101, gf.Project(gRepo), gf.Ref("refs/heads/main"))
			ct.GFake.AddFrom(gf.WithCIs(gHost, gf.ACLPublic(), ci))
			So(refreshExternal(ctx, lProject, gHost, 101, time.Time{}, cl.ID), ShouldBeNil)

			cl2 := getCL(ctx, gHost, 101)
			So(cl2.EVersion, ShouldEqual, 2)
			So(cl2.Snapshot.GetGerrit().GetInfo(), ShouldResembleProto, ci)
			So(pmtest.Projects(ct.TQ.Tasks()), ShouldResemble, []string{lProject})
		})
	})
}

func getCL(ctx context.Context, host string, change int) *changelist.CL {
	eid, err := changelist.GobID(host, int64(change))
	So(err, ShouldBeNil)
	cl, err := eid.Get(ctx)
	if err == datastore.ErrNoSuchEntity {
		return nil
	}
	So(err, ShouldBeNil)
	return cl
}

func singleRepoConfig(gHost string, gRepos ...string) *cfgpb.Config {
	projects := make([]*cfgpb.ConfigGroup_Gerrit_Project, len(gRepos))
	for i, gRepo := range gRepos {
		projects[i] = &cfgpb.ConfigGroup_Gerrit_Project{
			Name:      gRepo,
			RefRegexp: []string{"refs/heads/main"},
		}
	}
	return &cfgpb.Config{
		ConfigGroups: []*cfgpb.ConfigGroup{
			{
				Name: "main",
				Gerrit: []*cfgpb.ConfigGroup_Gerrit{
					{
						Url:      "https://" + gHost + "/",
						Projects: projects,
					},
				},
			},
		},
	}
}

func err5xx(gf.Operation, string) *status.Status {
	return status.New(codes.Internal, "boo")
}

func okThenErr5xx() gf.AccessCheck {
	calls := int32(0)
	return func(o gf.Operation, p string) *status.Status {
		if atomic.AddInt32(&calls, 1) == 1 {
			return status.New(codes.OK, "")
		} else {
			return err5xx(o, p)
		}
	}
}

func sortedRefreshTasks(ct cvtesting.Test) []*RefreshGerritCL {
	ret := make([]*RefreshGerritCL, 0, len(ct.TQ.Tasks().Payloads()))
	for _, m := range ct.TQ.Tasks().Payloads() {
		v, ok := m.(*RefreshGerritCL)
		if ok {
			ret = append(ret, v)
		}
	}
	sort.SliceStable(ret, func(i, j int) bool { return ret[i].less(ret[j]) })
	return ret
}

func (l *RefreshGerritCL) less(r *RefreshGerritCL) bool {
	switch {
	case l.GetHost() < r.GetHost():
		return true
	case l.GetHost() > r.GetHost():
		return false
	case l.GetChange() < r.GetChange():
		return true
	case l.GetChange() > r.GetChange():
		return false
	case l.GetLuciProject() < r.GetLuciProject():
		return true
	case l.GetLuciProject() > r.GetLuciProject():
		return false
	default:
		return l.GetUpdatedHint().AsTime().Before(r.GetUpdatedHint().AsTime())
	}
}
