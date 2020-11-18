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

package config

import (
	"context"
	"testing"
	"time"

	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"

	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/config"
	"go.chromium.org/luci/config/cfgclient"
	cfgmemory "go.chromium.org/luci/config/impl/memory"
	pb "go.chromium.org/luci/cv/api/config/v2"
	"go.chromium.org/luci/gae/filter/txndefer"
	gaememory "go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server/tq"
	"go.chromium.org/luci/server/tq/tqtesting"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

var testNow = testclock.TestTimeUTC.Round(1 * time.Millisecond)

func TestSubmitRefreshTasks(t *testing.T) {
	t.Parallel()

	Convey("Submit refresh task", t, func() {
		ctx, _, tqScheduler := mkTestingCtx()

		Convey("for a new project", func() {
			ctx = cfgclient.Use(ctx, cfgmemory.New(map[config.Set]cfgmemory.Files{
				config.ProjectSet("chromium"): {configFileName: ""},
			}))
			// Project chromium doesn't exist in datastore.
			err := SubmitRefreshTasks(ctx)
			So(err, ShouldBeNil)
			So(tqScheduler.Tasks().Payloads(), ShouldResembleProto, []*RefreshProjectConfigTask{
				{Project: "chromium"},
			})
		})

		Convey("for an existing project", func() {
			ctx = cfgclient.Use(ctx, cfgmemory.New(map[config.Set]cfgmemory.Files{
				config.ProjectSet("chromium"): {configFileName: ""},
			}))
			So(datastore.Put(ctx, &ProjectConfig{
				Project: "chromium",
				Enabled: true,
			}), ShouldBeNil)
			err := SubmitRefreshTasks(ctx)
			So(err, ShouldBeNil)
			So(tqScheduler.Tasks().Payloads(), ShouldResembleProto, []*RefreshProjectConfigTask{
				{Project: "chromium"},
			})
		})

		Convey("Disable project", func() {
			Convey("that doesn't have CV config", func() {
				ctx = cfgclient.Use(ctx, cfgmemory.New(map[config.Set]cfgmemory.Files{
					config.ProjectSet("chromium"): {"other.cfg": ""},
				}))
				So(datastore.Put(ctx, &ProjectConfig{
					Project: "chromium",
					Enabled: true,
				}), ShouldBeNil)
				err := SubmitRefreshTasks(ctx)
				So(err, ShouldBeNil)
				So(tqScheduler.Tasks().Payloads(), ShouldResembleProto, []*RefreshProjectConfigTask{
					{Project: "chromium", Disable: true},
				})
			})
			Convey("that doesn't exist in LUCI Config", func() {
				ctx = cfgclient.Use(ctx, cfgmemory.New(map[config.Set]cfgmemory.Files{}))
				So(datastore.Put(ctx, &ProjectConfig{
					Project: "chromium",
					Enabled: true,
				}), ShouldBeNil)
				err := SubmitRefreshTasks(ctx)
				So(err, ShouldBeNil)
				So(tqScheduler.Tasks().Payloads(), ShouldResembleProto, []*RefreshProjectConfigTask{
					{Project: "chromium", Disable: true},
				})
			})

			Convey("Skip already disabled Project", func() {
				ctx = cfgclient.Use(ctx, cfgmemory.New(map[config.Set]cfgmemory.Files{}))
				So(datastore.Put(ctx, &ProjectConfig{
					Project: "foo",
					Enabled: false,
				}), ShouldBeNil)
				err := SubmitRefreshTasks(ctx)
				So(err, ShouldBeNil)
				So(tqScheduler.Tasks(), ShouldBeEmpty)
			})
		})

	})
}

func TestUpdateProject(t *testing.T) {
	Convey("Update Project", t, func() {
		ctx, testClock, _ := mkTestingCtx()
		chromiumConfig := &pb.Config{
			DrainingStartTime: "2017-12-23T15:47:58Z",
			SubmitOptions: &pb.SubmitOptions{
				MaxBurst:   100,
				BurstDelay: durationpb.New(1 * time.Second),
			},
			ConfigGroups: []*pb.ConfigGroup{
				{
					Name: "branch_m100",
					Gerrit: []*pb.ConfigGroup_Gerrit{
						{
							Url: "https://chromium-review.googlesource.com/",
							Projects: []*pb.ConfigGroup_Gerrit_Project{
								{
									Name:      "chromium/src",
									RefRegexp: []string{"refs/heads/branch_m100"},
								},
							},
						},
					},
				},
				{
					Gerrit: []*pb.ConfigGroup_Gerrit{
						{
							Url: "https://chromium-review.googlesource.com/",
							Projects: []*pb.ConfigGroup_Gerrit_Project{
								{
									Name:      "chromium/src",
									RefRegexp: []string{"refs/heads/main"},
								},
							},
						},
					},
				},
			},
		}
		verifyEntitiesInDatastore := func(ctx context.Context, expectedEVersion int64) {
			cfg, meta := &pb.Config{}, &config.Meta{}
			err := cfgclient.Get(ctx, config.ProjectSet("chromium"), configFileName, cfgclient.ProtoText(cfg), meta)
			So(err, ShouldBeNil)
			localHash := computeHash(cfg)
			projKey := datastore.MakeKey(ctx, projectConfigKind, "chromium")
			cgNames := make([]string, len(cfg.GetConfigGroups()))
			// Verify ConfigGroups
			for i, cgpb := range cfg.GetConfigGroups() {
				cgNames[i] = makeConfigGroupName(cgpb.GetName(), i)
				cg := ConfigGroup{
					ID:      makeConfigGroupID(localHash, cgNames[i], i),
					Project: projKey,
				}
				err := datastore.Get(ctx, &cg)
				So(err, ShouldBeNil)
				So(cg.DrainingStartTime, ShouldEqual, cfg.GetDrainingStartTime())
				So(cg.SubmitOptions, ShouldResembleProto, cfg.GetSubmitOptions())
				So(cg.Content, ShouldResembleProto, cfg.GetConfigGroups()[i])
			}
			// Verify ProjectConfig
			pc := ProjectConfig{Project: "chromium"}
			err = datastore.Get(ctx, &pc)
			So(err, ShouldBeNil)
			So(pc, ShouldResemble, ProjectConfig{
				Project:          "chromium",
				Enabled:          true,
				EVersion:         expectedEVersion,
				Hash:             localHash,
				ExternalHash:     meta.ContentHash,
				UpdateTime:       datastore.RoundTime(testClock.Now()),
				ConfigGroupNames: cgNames,
			})
			// Verify ConfigHashInfo
			hashInfo := ConfigHashInfo{Hash: localHash, Project: projKey}
			err = datastore.Get(ctx, &hashInfo)
			So(err, ShouldBeNil)
			So(hashInfo, ShouldResemble, ConfigHashInfo{
				Hash:             localHash,
				Project:          projKey,
				ProjectEVersion:  expectedEVersion,
				UpdateTime:       datastore.RoundTime(testClock.Now()),
				ConfigGroupNames: cgNames,
			})
		}

		Convey("Creates new ProjectConfig", func() {
			ctx = cfgclient.Use(ctx, cfgmemory.New(map[config.Set]cfgmemory.Files{
				config.ProjectSet("chromium"): {
					configFileName: toProtoText(chromiumConfig),
				},
			}))
			err := updateProject(ctx, "chromium")
			So(err, ShouldBeNil)
			verifyEntitiesInDatastore(ctx, 1)
			testClock.Add(10 * time.Minute)

			Convey("Noop if config is up-to-date", func() {
				err := updateProject(ctx, "chromium")
				So(err, ShouldBeNil)
				pc := ProjectConfig{Project: "chromium"}
				err = datastore.Get(ctx, &pc)
				So(err, ShouldBeNil)
				So(pc.EVersion, ShouldEqual, 1)
				prevUpdatedTime := testClock.Now().Add(-10 * time.Minute)
				So(pc.UpdateTime, ShouldResemble, prevUpdatedTime)
			})

			Convey("Update existing ProjectConfig", func() {
				updatedConfig := proto.Clone(chromiumConfig).(*pb.Config)
				updatedConfig.ConfigGroups = append(updatedConfig.ConfigGroups, &pb.ConfigGroup{
					Name: "experimental",
					Gerrit: []*pb.ConfigGroup_Gerrit{
						{
							Url: "https://chromium-review.googlesource.com/",
							Projects: []*pb.ConfigGroup_Gerrit_Project{
								{
									Name:      "chromium/src/experimental",
									RefRegexp: []string{"refs/heads/main"},
								},
							},
						},
					},
				})
				ctx = cfgclient.Use(ctx, cfgmemory.New(map[config.Set]cfgmemory.Files{
					config.ProjectSet("chromium"): {
						configFileName: toProtoText(updatedConfig),
					},
				}))
				err := updateProject(ctx, "chromium")
				So(err, ShouldBeNil)
				verifyEntitiesInDatastore(ctx, 2)
				testClock.Add(10 * time.Minute)

				Convey("Roll back to previous version", func() {
					ctx = cfgclient.Use(ctx, cfgmemory.New(map[config.Set]cfgmemory.Files{
						config.ProjectSet("chromium"): {
							configFileName: toProtoText(chromiumConfig),
						},
					}))

					err := updateProject(ctx, "chromium")
					So(err, ShouldBeNil)
					verifyEntitiesInDatastore(ctx, 3)
				})
			})
		})
	})
}

func TestDisableProject(t *testing.T) {
	Convey("Disable", t, func() {
		ctx, testClock, _ := mkTestingCtx()
		writeProjectConfig := func(enabled bool) {
			pc := ProjectConfig{
				Project:          "chromium",
				Enabled:          enabled,
				EVersion:         100,
				Hash:             "hash",
				ExternalHash:     "externalHash",
				UpdateTime:       datastore.RoundTime(testClock.Now()),
				ConfigGroupNames: []string{"default"},
			}
			So(datastore.Put(ctx, &pc), ShouldBeNil)
			testClock.Add(10 * time.Minute)
		}

		Convey("currently enabled Project", func() {
			writeProjectConfig(true)
			err := disableProject(ctx, "chromium")
			So(err, ShouldBeNil)
			actual := ProjectConfig{Project: "chromium"}
			So(datastore.Get(ctx, &actual), ShouldBeNil)
			So(actual.Enabled, ShouldBeFalse)
			So(actual.EVersion, ShouldEqual, 101)
			So(actual.UpdateTime, ShouldResemble, datastore.RoundTime(testClock.Now()))
		})

		Convey("currently disabled Project", func() {
			writeProjectConfig(false)
			err := disableProject(ctx, "chromium")
			So(err, ShouldBeNil)
			actual := ProjectConfig{Project: "chromium"}
			So(datastore.Get(ctx, &actual), ShouldBeNil)
			So(actual.Enabled, ShouldBeFalse)
			So(actual.EVersion, ShouldEqual, 100)
		})

		Convey("non-existing Project", func() {
			err := disableProject(ctx, "non-existing")
			So(err, ShouldBeNil)
			So(datastore.Get(ctx, &ProjectConfig{Project: "non-existing"}), ShouldErrLike, datastore.ErrNoSuchEntity)
		})
	})
}

func mkTestingCtx() (context.Context, testclock.TestClock, *tqtesting.Scheduler) {
	ctx, clock := testclock.UseTime(context.Background(), testNow)
	ctx = txndefer.FilterRDS(gaememory.Use(ctx))
	datastore.GetTestable(ctx).AutoIndex(true)
	datastore.GetTestable(ctx).Consistent(true)

	ctx, scheduler := tq.TestingContext(ctx, nil)
	return ctx, clock, scheduler
}

func toProtoText(msg proto.Message) string {
	bs, err := prototext.Marshal(msg)
	So(err, ShouldBeNil)
	return string(bs)
}