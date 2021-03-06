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
	"testing"

	structpb "github.com/golang/protobuf/ptypes/struct"

	"go.chromium.org/luci/gae/impl/memory"
	"go.chromium.org/luci/gae/service/datastore"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"

	"go.chromium.org/luci/buildbucket/appengine/model"
	pb "go.chromium.org/luci/buildbucket/proto"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestScheduleBuild(t *testing.T) {
	t.Parallel()

	Convey("scheduleRequestFromTemplate", t, func() {
		ctx := memory.Use(context.Background())
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)
		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: "user:caller@example.com",
		})

		So(datastore.Put(ctx, &model.Bucket{
			ID:     "bucket",
			Parent: model.ProjectKey(ctx, "project"),
			Proto: pb.Bucket{
				Acls: []*pb.Acl{
					{
						Identity: "user:caller@example.com",
						Role:     pb.Acl_READER,
					},
				},
			},
		}), ShouldBeNil)

		Convey("nil", func() {
			ret, err := scheduleRequestFromTemplate(ctx, nil)
			So(err, ShouldBeNil)
			So(ret, ShouldBeNil)
		})

		Convey("empty", func() {
			req := &pb.ScheduleBuildRequest{}
			ret, err := scheduleRequestFromTemplate(ctx, req)
			So(err, ShouldBeNil)
			So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{})
			So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{})
		})

		Convey("not found", func() {
			req := &pb.ScheduleBuildRequest{
				TemplateBuildId: 1,
			}
			ret, err := scheduleRequestFromTemplate(ctx, req)
			So(err, ShouldErrLike, "not found")
			So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
				TemplateBuildId: 1,
			})
			So(ret, ShouldBeNil)
		})

		Convey("permission denied", func() {
			ctx = auth.WithState(ctx, &authtest.FakeState{
				Identity: "user:unauthorized@example.com",
			})
			So(datastore.Put(ctx, &model.Build{
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				},
			}), ShouldBeNil)
			req := &pb.ScheduleBuildRequest{
				TemplateBuildId: 1,
			}
			ret, err := scheduleRequestFromTemplate(ctx, req)
			So(err, ShouldErrLike, "not found")
			So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
				TemplateBuildId: 1,
			})
			So(ret, ShouldBeNil)
		})

		Convey("canary", func() {
			Convey("false default", func() {
				So(datastore.Put(ctx, &model.Build{
					Proto: pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary: false,
					},
				}), ShouldBeNil)

				Convey("merge", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Canary:          pb.Trinary_YES,
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Canary:          pb.Trinary_YES,
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_YES,
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
					})
				})

				Convey("ok", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
					})
				})
			})

			Convey("true default", func() {
				So(datastore.Put(ctx, &model.Build{
					Proto: pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary: true,
					},
				}), ShouldBeNil)

				Convey("merge", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Canary:          pb.Trinary_NO,
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Canary:          pb.Trinary_NO,
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
					})
				})

				Convey("ok", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_YES,
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
					})
				})
			})
		})

		Convey("critical", func() {
			So(datastore.Put(ctx, &model.Build{
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Critical: pb.Trinary_YES,
				},
			}), ShouldBeNil)

			Convey("merge", func() {
				req := &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
					Critical:        pb.Trinary_NO,
				}
				ret, err := scheduleRequestFromTemplate(ctx, req)
				So(err, ShouldBeNil)
				So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
					Critical:        pb.Trinary_NO,
				})
				So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Canary:       pb.Trinary_NO,
					Critical:     pb.Trinary_NO,
					Experimental: pb.Trinary_NO,
					Properties:   &structpb.Struct{},
				})
			})

			Convey("ok", func() {
				req := &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				}
				ret, err := scheduleRequestFromTemplate(ctx, req)
				So(err, ShouldBeNil)
				So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				})
				So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Canary:       pb.Trinary_NO,
					Critical:     pb.Trinary_YES,
					Experimental: pb.Trinary_NO,
					Properties:   &structpb.Struct{},
				})
			})
		})

		Convey("exe", func() {
			So(datastore.Put(ctx, &model.Build{
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Exe: &pb.Executable{
						CipdPackage: "package",
						CipdVersion: "version",
					},
				},
			}), ShouldBeNil)

			Convey("merge", func() {
				Convey("empty", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Exe:             &pb.Executable{},
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Exe:             &pb.Executable{},
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Exe:          &pb.Executable{},
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
					})
				})

				Convey("non-empty", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Exe: &pb.Executable{
							CipdPackage: "package",
							CipdVersion: "new",
						},
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Exe: &pb.Executable{
							CipdPackage: "package",
							CipdVersion: "new",
						},
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary: pb.Trinary_NO,
						Exe: &pb.Executable{
							CipdPackage: "package",
							CipdVersion: "new",
						},
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
					})
				})
			})

			Convey("ok", func() {
				req := &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				}
				ret, err := scheduleRequestFromTemplate(ctx, req)
				So(err, ShouldBeNil)
				So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				})
				So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Canary: pb.Trinary_NO,
					Exe: &pb.Executable{
						CipdPackage: "package",
						CipdVersion: "version",
					},
					Experimental: pb.Trinary_NO,
					Properties:   &structpb.Struct{},
				})
			})
		})

		Convey("gerrit changes", func() {
			So(datastore.Put(ctx, &model.Build{
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Input: &pb.Build_Input{
						GerritChanges: []*pb.GerritChange{
							{
								Host:     "example.com",
								Project:  "project",
								Change:   1,
								Patchset: 1,
							},
						},
					},
				},
			}), ShouldBeNil)

			Convey("merge", func() {
				Convey("empty", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						GerritChanges:   []*pb.GerritChange{},
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						GerritChanges:   []*pb.GerritChange{},
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						GerritChanges: []*pb.GerritChange{
							{
								Host:     "example.com",
								Project:  "project",
								Change:   1,
								Patchset: 1,
							},
						},
						Properties: &structpb.Struct{},
					})
				})

				Convey("non-empty", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						GerritChanges: []*pb.GerritChange{
							{
								Host:     "example.com",
								Project:  "project",
								Change:   1,
								Patchset: 2,
							},
						},
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						GerritChanges: []*pb.GerritChange{
							{
								Host:     "example.com",
								Project:  "project",
								Change:   1,
								Patchset: 2,
							},
						},
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						GerritChanges: []*pb.GerritChange{
							{
								Host:     "example.com",
								Project:  "project",
								Change:   1,
								Patchset: 2,
							},
						},
						Properties: &structpb.Struct{},
					})
				})
			})

			Convey("ok", func() {
				req := &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				}
				ret, err := scheduleRequestFromTemplate(ctx, req)
				So(err, ShouldBeNil)
				So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				})
				So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Canary:       pb.Trinary_NO,
					Experimental: pb.Trinary_NO,
					GerritChanges: []*pb.GerritChange{
						{
							Host:     "example.com",
							Project:  "project",
							Change:   1,
							Patchset: 1,
						},
					},
					Properties: &structpb.Struct{},
				})
			})
		})

		Convey("gitiles commit", func() {
			So(datastore.Put(ctx, &model.Build{
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Input: &pb.Build_Input{
						GitilesCommit: &pb.GitilesCommit{
							Host:    "example.com",
							Project: "project",
							Ref:     "refs/heads/master",
						},
					},
				},
			}), ShouldBeNil)
			req := &pb.ScheduleBuildRequest{
				TemplateBuildId: 1,
			}
			ret, err := scheduleRequestFromTemplate(ctx, req)
			So(err, ShouldBeNil)
			So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
				TemplateBuildId: 1,
			})
			So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
				Builder: &pb.BuilderID{
					Project: "project",
					Bucket:  "bucket",
					Builder: "builder",
				},
				Canary:       pb.Trinary_NO,
				Experimental: pb.Trinary_NO,
				GitilesCommit: &pb.GitilesCommit{
					Host:    "example.com",
					Project: "project",
					Ref:     "refs/heads/master",
				},
				Properties: &structpb.Struct{},
			})
		})

		Convey("input properties", func() {
			So(datastore.Put(ctx, &model.Build{
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				},
			}), ShouldBeNil)

			Convey("empty", func() {
				So(datastore.Put(ctx, &model.BuildInputProperties{
					Build: datastore.MakeKey(ctx, "Build", 1),
				}), ShouldBeNil)

				Convey("merge", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"input": {
									Kind: &structpb.Value_StringValue{
										StringValue: "input value",
									},
								},
							},
						},
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"input": {
									Kind: &structpb.Value_StringValue{
										StringValue: "input value",
									},
								},
							},
						},
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"input": {
									Kind: &structpb.Value_StringValue{
										StringValue: "input value",
									},
								},
							},
						},
					})
				})

				Convey("ok", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
					})
				})
			})

			Convey("non-empty", func() {
				So(datastore.Put(ctx, &model.BuildInputProperties{
					Build: datastore.MakeKey(ctx, "Build", 1),
					Proto: model.DSStruct{
						Struct: structpb.Struct{
							Fields: map[string]*structpb.Value{
								"input": {
									Kind: &structpb.Value_StringValue{
										StringValue: "input value",
									},
								},
							},
						},
					},
				}), ShouldBeNil)

				Convey("merge", func() {
					Convey("empty", func() {
						req := &pb.ScheduleBuildRequest{
							TemplateBuildId: 1,
							Properties:      &structpb.Struct{},
						}
						ret, err := scheduleRequestFromTemplate(ctx, req)
						So(err, ShouldBeNil)
						So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
							TemplateBuildId: 1,
							Properties:      &structpb.Struct{},
						})
						So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
							Builder: &pb.BuilderID{
								Project: "project",
								Bucket:  "bucket",
								Builder: "builder",
							},
							Canary:       pb.Trinary_NO,
							Experimental: pb.Trinary_NO,
							Properties:   &structpb.Struct{},
						})
					})

					Convey("non-empty", func() {
						req := &pb.ScheduleBuildRequest{
							TemplateBuildId: 1,
							Properties: &structpb.Struct{
								Fields: map[string]*structpb.Value{
									"other": {
										Kind: &structpb.Value_StringValue{
											StringValue: "other value",
										},
									},
								},
							},
						}
						ret, err := scheduleRequestFromTemplate(ctx, req)
						So(err, ShouldBeNil)
						So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
							TemplateBuildId: 1,
							Properties: &structpb.Struct{
								Fields: map[string]*structpb.Value{
									"other": {
										Kind: &structpb.Value_StringValue{
											StringValue: "other value",
										},
									},
								},
							},
						})
						So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
							Builder: &pb.BuilderID{
								Project: "project",
								Bucket:  "bucket",
								Builder: "builder",
							},
							Canary:       pb.Trinary_NO,
							Experimental: pb.Trinary_NO,
							Properties: &structpb.Struct{
								Fields: map[string]*structpb.Value{
									"other": {
										Kind: &structpb.Value_StringValue{
											StringValue: "other value",
										},
									},
								},
							},
						})
					})
				})

				Convey("ok", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						Properties: &structpb.Struct{
							Fields: map[string]*structpb.Value{
								"input": {
									Kind: &structpb.Value_StringValue{
										StringValue: "input value",
									},
								},
							},
						},
					})
				})
			})
		})

		Convey("tags", func() {
			So(datastore.Put(ctx, &model.Build{
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				},
				Tags: []string{
					"key:value",
				},
			}), ShouldBeNil)

			Convey("merge", func() {
				Convey("empty", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Tags:            []*pb.StringPair{},
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Tags:            []*pb.StringPair{},
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
						Tags: []*pb.StringPair{
							{
								Key:   "key",
								Value: "value",
							},
						},
					})
				})

				Convey("non-empty", func() {
					req := &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Tags: []*pb.StringPair{
							{
								Key:   "other",
								Value: "other",
							},
						},
					}
					ret, err := scheduleRequestFromTemplate(ctx, req)
					So(err, ShouldBeNil)
					So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
						TemplateBuildId: 1,
						Tags: []*pb.StringPair{
							{
								Key:   "other",
								Value: "other",
							},
						},
					})
					So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
						Canary:       pb.Trinary_NO,
						Experimental: pb.Trinary_NO,
						Properties:   &structpb.Struct{},
						Tags: []*pb.StringPair{
							{
								Key:   "other",
								Value: "other",
							},
						},
					})
				})
			})

			Convey("ok", func() {
				req := &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				}
				ret, err := scheduleRequestFromTemplate(ctx, req)
				So(err, ShouldBeNil)
				So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				})
				So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
					Canary:       pb.Trinary_NO,
					Experimental: pb.Trinary_NO,
					Properties:   &structpb.Struct{},
					Tags: []*pb.StringPair{
						{
							Key:   "key",
							Value: "value",
						},
					},
				})
			})
		})

		Convey("ok", func() {
			So(datastore.Put(ctx, &model.Build{
				Proto: pb.Build{
					Id: 1,
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				},
			}), ShouldBeNil)
			req := &pb.ScheduleBuildRequest{
				TemplateBuildId: 1,
			}
			ret, err := scheduleRequestFromTemplate(ctx, req)
			So(err, ShouldBeNil)
			So(req, ShouldResembleProto, &pb.ScheduleBuildRequest{
				TemplateBuildId: 1,
			})
			So(ret, ShouldResembleProto, &pb.ScheduleBuildRequest{
				Builder: &pb.BuilderID{
					Project: "project",
					Bucket:  "bucket",
					Builder: "builder",
				},
				Canary:       pb.Trinary_NO,
				Experimental: pb.Trinary_NO,
				Properties:   &structpb.Struct{},
			})
		})
	})

	Convey("ScheduleBuild", t, func() {
		srv := &Builds{}

		ctx := memory.Use(context.Background())
		datastore.GetTestable(ctx).AutoIndex(true)
		datastore.GetTestable(ctx).Consistent(true)
		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: "user:caller@example.com",
		})

		Convey("builder", func() {
			Convey("not found", func() {
				req := &pb.ScheduleBuildRequest{
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				}
				rsp, err := srv.ScheduleBuild(ctx, req)
				So(err, ShouldErrLike, "not found")
				So(rsp, ShouldBeNil)
			})

			Convey("permission denied", func() {
				So(datastore.Put(ctx, &model.Build{
					Proto: pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
					},
				}), ShouldBeNil)
				req := &pb.ScheduleBuildRequest{
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				}
				rsp, err := srv.ScheduleBuild(ctx, req)
				So(err, ShouldErrLike, "not found")
				So(rsp, ShouldBeNil)
			})

			Convey("ok", func() {
				So(datastore.Put(ctx, &model.Bucket{
					ID:     "bucket",
					Parent: model.ProjectKey(ctx, "project"),
					Proto: pb.Bucket{
						Acls: []*pb.Acl{
							{
								Identity: "user:caller@example.com",
								Role:     pb.Acl_SCHEDULER,
							},
						},
					},
				}), ShouldBeNil)
				So(datastore.Put(ctx, &model.Build{
					Proto: pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
					},
				}), ShouldBeNil)
				req := &pb.ScheduleBuildRequest{
					Builder: &pb.BuilderID{
						Project: "project",
						Bucket:  "bucket",
						Builder: "builder",
					},
				}
				rsp, err := srv.ScheduleBuild(ctx, req)
				So(err, ShouldBeNil)
				So(rsp, ShouldBeNil)
			})
		})

		Convey("template build ID", func() {
			Convey("not found", func() {
				req := &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				}
				rsp, err := srv.ScheduleBuild(ctx, req)
				So(err, ShouldErrLike, "not found")
				So(rsp, ShouldBeNil)
			})

			Convey("permission denied", func() {
				So(datastore.Put(ctx, &model.Build{
					Proto: pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
					},
				}), ShouldBeNil)
				req := &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				}
				rsp, err := srv.ScheduleBuild(ctx, req)
				So(err, ShouldErrLike, "not found")
				So(rsp, ShouldBeNil)
			})

			Convey("ok", func() {
				So(datastore.Put(ctx, &model.Bucket{
					ID:     "bucket",
					Parent: model.ProjectKey(ctx, "project"),
					Proto: pb.Bucket{
						Acls: []*pb.Acl{
							{
								Identity: "user:caller@example.com",
								Role:     pb.Acl_SCHEDULER,
							},
						},
					},
				}), ShouldBeNil)
				So(datastore.Put(ctx, &model.Build{
					Proto: pb.Build{
						Id: 1,
						Builder: &pb.BuilderID{
							Project: "project",
							Bucket:  "bucket",
							Builder: "builder",
						},
					},
				}), ShouldBeNil)
				req := &pb.ScheduleBuildRequest{
					TemplateBuildId: 1,
				}
				rsp, err := srv.ScheduleBuild(ctx, req)
				So(err, ShouldBeNil)
				So(rsp, ShouldBeNil)
			})
		})
	})

	Convey("validateSchedule", t, func() {
		Convey("nil", func() {
			err := validateSchedule(nil)
			So(err, ShouldErrLike, "builder or template_build_id is required")
		})

		Convey("empty", func() {
			req := &pb.ScheduleBuildRequest{}
			err := validateSchedule(req)
			So(err, ShouldErrLike, "builder or template_build_id is required")
		})

		Convey("request ID", func() {
			req := &pb.ScheduleBuildRequest{
				RequestId:       "request/id",
				TemplateBuildId: 1,
			}
			err := validateSchedule(req)
			So(err, ShouldErrLike, "request_id cannot contain")
		})

		Convey("builder ID", func() {
			req := &pb.ScheduleBuildRequest{
				Builder: &pb.BuilderID{},
			}
			err := validateSchedule(req)
			So(err, ShouldErrLike, "project must match")
		})

		Convey("exe", func() {
			Convey("empty", func() {
				req := &pb.ScheduleBuildRequest{
					Exe:             &pb.Executable{},
					TemplateBuildId: 1,
				}
				err := validateSchedule(req)
				So(err, ShouldBeNil)
			})

			Convey("package", func() {
				req := &pb.ScheduleBuildRequest{
					Exe: &pb.Executable{
						CipdPackage: "package",
					},
					TemplateBuildId: 1,
				}
				err := validateSchedule(req)
				So(err, ShouldErrLike, "cipd_package must not be specified")
			})

			Convey("version", func() {
				Convey("invalid", func() {
					req := &pb.ScheduleBuildRequest{
						Exe: &pb.Executable{
							CipdVersion: "invalid!",
						},
						TemplateBuildId: 1,
					}
					err := validateSchedule(req)
					So(err, ShouldErrLike, "cipd_version")
				})

				Convey("valid", func() {
					req := &pb.ScheduleBuildRequest{
						Exe: &pb.Executable{
							CipdVersion: "valid",
						},
						TemplateBuildId: 1,
					}
					err := validateSchedule(req)
					So(err, ShouldBeNil)
				})
			})
		})

		Convey("gitiles commit", func() {
			req := &pb.ScheduleBuildRequest{
				GitilesCommit: &pb.GitilesCommit{
					Host: "example.com",
				},
				TemplateBuildId: 1,
			}
			err := validateSchedule(req)
			So(err, ShouldErrLike, "gitiles_commit")
		})

		Convey("tags", func() {
			req := &pb.ScheduleBuildRequest{
				Tags: []*pb.StringPair{
					{
						Key: "key:value",
					},
				},
				TemplateBuildId: 1,
			}
			err := validateSchedule(req)
			So(err, ShouldErrLike, "tags")
		})

		Convey("priority", func() {
			Convey("negative", func() {
				req := &pb.ScheduleBuildRequest{
					Priority:        -1,
					TemplateBuildId: 1,
				}
				err := validateSchedule(req)
				So(err, ShouldErrLike, "priority must be in")
			})

			Convey("excessive", func() {
				req := &pb.ScheduleBuildRequest{
					Priority:        256,
					TemplateBuildId: 1,
				}
				err := validateSchedule(req)
				So(err, ShouldErrLike, "priority must be in")
			})
		})
	})
}
