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

package internal

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/common/data/rand/cryptorand"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/common/tsmon"
	"go.chromium.org/luci/common/tsmon/distribution"
	"go.chromium.org/luci/common/tsmon/types"
	"go.chromium.org/luci/ttq"
	"go.chromium.org/luci/ttq/internal/partition"

	"github.com/golang/mock/gomock"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestMakeReminder(t *testing.T) {
	t.Parallel()

	Convey("makeReminder", t, func() {
		ctx := cryptorand.MockForTest(context.Background(), 111)
		ctx, _ = testclock.UseTime(ctx, testclock.TestRecentTimeLocal)
		req := taskspb.CreateTaskRequest{
			Parent: "projects/example-project/locations/us-central1/queues/q",
			Task: &taskspb.Task{
				MessageType: &taskspb.Task_HttpRequest{HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_GET,
					Url:        "https://example.com/user/exec",
				}},
				ScheduleTime: google.NewTimestamp(clock.Now(ctx).Add(time.Minute)),
			},
		}

		Convey("typical request with deadline", func() {
			deadline := testclock.TestRecentTimeLocal.Add(5 * time.Second)
			reqCtx, cancel := clock.WithDeadline(ctx, deadline)
			defer cancel()
			r, clonedReq, err := makeReminder(reqCtx, &req)
			So(err, ShouldBeNil)
			So(r.Id, ShouldResemble, "1408d4aaccbdf3f4b03972992f179ce4")
			So(len(r.Id), ShouldEqual, 2*keySpaceBytes)
			So(r.FreshUntil, ShouldEqual, deadline.UTC())
			So(clonedReq.Task.Name, ShouldResemble,
				"projects/example-project/locations/us-central1/queues/q/tasks/1408d4aaccbdf3f4b03972992f179ce4")

			deserialized, err := r.createTaskRequest()
			So(err, ShouldBeNil)
			So(clonedReq, ShouldResembleProto, deserialized)

			clonedReq.Task.Name = ""
			So(clonedReq, ShouldResembleProto, &req)
		})

		Convey("reasonable FreshUntil without deadline", func() {
			r, _, err := makeReminder(ctx, &req)
			So(err, ShouldBeNil)
			So(r.FreshUntil, ShouldEqual, testclock.TestRecentTimeLocal.Add(happyPathMaxDuration).UTC())
		})

		Convey("named tasks not supported", func() {
			req.Task.Name = req.Parent + "/tasks/deadbeef"
			_, _, err := makeReminder(ctx, &req)
			So(err, ShouldErrLike, "Named tasks are not supported")
		})
	})
}

func TestAddTask(t *testing.T) {
	t.Parallel()

	Convey("AddTask works", t, func() {
		ctx := cryptorand.MockForTest(context.Background(), 111)
		ctx, tclock := testclock.UseTime(ctx, testclock.TestRecentTimeLocal)
		ctx, _ = tsmon.WithDummyInMemory(ctx)

		latestOf := func(m types.Metric, fieldVals ...interface{}) interface{} {
			return tsmon.GetState(ctx).Store().Get(ctx, m, time.Time{}, fieldVals)
		}

		req := taskspb.CreateTaskRequest{
			Parent: "projects/example-project/locations/us-central1/queues/q",
			Task: &taskspb.Task{
				MessageType: &taskspb.Task_HttpRequest{HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_GET,
					Url:        "https://example.com/user/exec",
				}},
				ScheduleTime: google.NewTimestamp(clock.Now(ctx).Add(time.Minute)),
			},
		}
		opts := ttq.Options{}
		So(opts.Validate(), ShouldBeNil)
		db := FakeDB{}
		fakeCT := fakeCloudTasks{}
		impl := Impl{Options: opts, DB: &db, TasksClient: &fakeCT}

		Convey("with deadline", func() {
			pp, err := impl.AddTask(ctx, &req)
			So(err, ShouldBeNil)
			So(pp, ShouldNotBeNil)
			So(len(db.reminders), ShouldEqual, 1)

			Convey("no time to complete happy path", func() {
				for _, r := range db.reminders {
					tclock.Set(r.FreshUntil.Add(time.Second))
				}
				pp(ctx) // noop now
				So(len(fakeCT.calls), ShouldEqual, 0)
				So(len(db.reminders), ShouldEqual, 1)
				So(latestOf(metricPostProcessedAttempts, "OK", "happy", "FakeDB"), ShouldBeNil)
			})

			Convey("happy path", func() {
				pp(ctx)
				So(len(db.reminders), ShouldEqual, 0)
				So(len(fakeCT.calls), ShouldEqual, 1)
				So(latestOf(metricPostProcessedAttempts, "OK", "happy", "FakeDB"), ShouldEqual, 1)
				So(latestOf(metricPostProcessedDurationsMS, "OK", "happy", "FakeDB").(*distribution.Distribution).Count(), ShouldEqual, 1)

				Convey("calling 2nd time is a noop", func() {
					pp(ctx)
					So(len(fakeCT.calls), ShouldEqual, 1)
					So(latestOf(metricPostProcessedAttempts, "OK", "happy", "FakeDB"), ShouldEqual, 1)
				})
			})
			Convey("happy path transient task creation error", func() {
				fakeCT.errs = []error{status.Errorf(codes.Unavailable, "please retry")}
				pp(ctx)
				So(len(db.reminders), ShouldEqual, 1)
				So(latestOf(metricPostProcessedAttempts, "fail-task", "happy", "FakeDB"), ShouldEqual, 1)
			})
			Convey("happy path permanent task creation error", func() {
				fakeCT.errs = []error{status.Errorf(codes.InvalidArgument, "forever")}
				pp(ctx)
				So(latestOf(metricPostProcessedAttempts, "ignored-bad-task", "happy", "FakeDB"), ShouldEqual, 1)
				So(latestOf(metricTasksCreated, "InvalidArgument", "happy", "FakeDB"), ShouldEqual, 1)
				// Must remove reminder despite the error.
				So(len(db.reminders), ShouldEqual, 0)
			})
		})
	})
}

func TestSweepAll(t *testing.T) {
	t.Parallel()

	Convey("SweepAll works", t, func() {
		impl := Impl{Options: ttq.Options{Shards: 2}}
		items := impl.SweepAll()
		So(len(items), ShouldEqual, 2)
		So(items[0].Partition.High, ShouldResemble, items[1].Partition.Low)
		items[0].Partition = nil
		items[1].Partition = nil
		So(items, ShouldResemble, []ScanItem{
			{Shard: 0, Level: 0},
			{Shard: 1, Level: 0},
		})
	})
}

func TestScan(t *testing.T) {
	t.Parallel()

	Convey("Scan works", t, func() {
		ctx := cryptorand.MockForTest(context.Background(), 111)
		epoch := testclock.TestRecentTimeLocal
		ctx, tclock := testclock.UseTime(ctx, epoch)
		ctx, _ = tsmon.WithDummyInMemory(ctx)
		ctx = gologger.StdConfig.Use(ctx)
		ctx = logging.SetLevel(ctx, logging.Debug)

		// Use first shard spanning 0000000...0 to 1000000...0
		// such that Reminders in tests can use short IDs "0*".
		part := partition.Universe(keySpaceBytes).Split(16)[0]

		Convey("Normal operation", func() {
			db := FakeDB{}
			impl := Impl{Options: ttq.Options{TasksPerScan: 2048}, DB: &db}

			Convey("Empty", func() {
				more, res, err := impl.Scan(ctx, ScanItem{Partition: part})
				So(len(more), ShouldEqual, 0)
				So(err, ShouldBeNil)
				So(len(res.Reminders), ShouldEqual, 0)
			})

			stale := epoch.Add(30 * time.Second)
			fresh := epoch.Add(90 * time.Second)
			savedReminders := []*Reminder{
				&Reminder{Id: "01", FreshUntil: fresh},
				&Reminder{Id: "02", FreshUntil: stale},
				&Reminder{Id: "03", FreshUntil: fresh},
				&Reminder{Id: "04", FreshUntil: stale},
			}
			for _, r := range savedReminders {
				So(db.SaveReminder(ctx, r), ShouldBeNil)
			}
			tclock.Set(epoch.Add(60 * time.Second))

			Convey("Scan complete partition", func() {
				impl.Options.TasksPerScan = 10
				more, res, err := impl.Scan(ctx, ScanItem{Partition: part})
				So(err, ShouldBeNil)
				So(len(more), ShouldEqual, 0)
				So(res.Reminders, ShouldResemble, []*Reminder{
					&Reminder{Id: "02", FreshUntil: stale},
					&Reminder{Id: "04", FreshUntil: stale},
				})
			})

			Convey("Scan reaches TasksPerScan limit", func() {
				// Only Reminders 01..03 will be fetched. 02 is stale.
				// Follow up scans should start from 04.
				impl.Options.TasksPerScan = 3
				impl.Options.SecondaryScanShards = 2
				scanItem := ScanItem{Shard: 1, Level: 1, Partition: part}
				more, res, err := impl.Scan(ctx, scanItem)

				So(err, ShouldBeNil)
				So(res.Reminders, ShouldResemble, []*Reminder{
					&Reminder{Id: "02", FreshUntil: stale},
				})

				Convey("Scan returns correct follow up ScanItems", func() {
					So(len(more), ShouldEqual, impl.Options.SecondaryScanShards)
					So(more[0].Partition.Low, ShouldResemble, *big.NewInt(3 + 1))
					So(more[0].Partition.High, ShouldResemble, more[1].Partition.Low)
					So(more[1].Partition.High, ShouldResemble, part.High)
					So(more[0].Shard, ShouldEqual, scanItem.Shard)
					So(more[1].Shard, ShouldEqual, scanItem.Shard)
					So(more[0].Level, ShouldEqual, scanItem.Level+1)
					So(more[1].Level, ShouldEqual, scanItem.Level+1)
				})
			})
		})
		Convey("Abnormal operation", func() {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db := NewMockDatabase(ctrl)
			db.EXPECT().Kind().AnyTimes().Return("mockdb")
			impl := Impl{Options: ttq.Options{TasksPerScan: 2048, SecondaryScanShards: 2}, DB: db}

			Convey("Level-3 not allowed", func() {
				_, _, err := impl.Scan(ctx, ScanItem{Level: 3, Partition: part})
				So(err, ShouldErrLike, "level-3 scan")
			})

			stale := epoch.Add(30 * time.Second)
			fresh := epoch.Add(90 * time.Second)
			someReminders := []*Reminder{
				&Reminder{Id: "01", FreshUntil: fresh},
				&Reminder{Id: "02", FreshUntil: stale},
			}
			tclock.Set(epoch.Add(60 * time.Second))

			// Simulate context expiry by using deadline in the past.
			// TODO(tandrii): find a way to expire ctx withing FetchRemindersMeta for a
			// realistic test.
			ctx, cancel := clock.WithDeadline(ctx, epoch)
			defer cancel()
			So(ctx.Err(), ShouldResemble, context.DeadlineExceeded)

			Convey("Timeout without anything fetched", func() {
				errExpected := errors.Annotate(ctx.Err(), "failed to fetch all").Err()
				db.EXPECT().FetchRemindersMeta(ctx, gomock.Any(), gomock.Any(), gomock.Any()).
					Return(nil, errExpected)

				more, res, err := impl.Scan(ctx, ScanItem{Partition: part})
				So(err, ShouldResemble, errExpected)
				So(len(res.Reminders), ShouldEqual, 0)
				So(len(more), ShouldEqual, impl.Options.SecondaryScanShards)
				So(more[0].Partition.Low, ShouldResemble, part.Low)
				So(more[1].Partition.High, ShouldResemble, part.High)
			})

			Convey("Timeout after fetching some", func() {
				db.EXPECT().FetchRemindersMeta(ctx, gomock.Any(), gomock.Any(), gomock.Any()).
					Return(someReminders, ctx.Err())

				more, res, err := impl.Scan(ctx, ScanItem{Partition: part})
				So(err, ShouldResemble, context.DeadlineExceeded)
				So(res.Reminders, ShouldResemble, []*Reminder{
					&Reminder{Id: "02", FreshUntil: stale},
				})
				So(len(more), ShouldEqual, impl.Options.SecondaryScanShards)
				So(more[0].Partition.Low, ShouldResemble, *big.NewInt(2 + 1))
				So(more[1].Partition.High, ShouldResemble, part.High)
			})
		})
	})
}

func TestPostProcessBatch(t *testing.T) {
	t.Parallel()

	Convey("PostProcessBatch works", t, func() {
		ctx := cryptorand.MockForTest(context.Background(), 111)
		ctx, tclock := testclock.UseTime(ctx, testclock.TestRecentTimeLocal)
		ctx, _ = tsmon.WithDummyInMemory(ctx)

		db := FakeDB{}
		fakeCT := fakeCloudTasks{}
		impl := Impl{Options: ttq.Options{}, DB: &db, TasksClient: &fakeCT}

		batch := make([]*Reminder, 10)
		for i, _ := range batch {
			r := genReminder(ctx)
			So(db.SaveReminder(ctx, r), ShouldBeNil)
			batch[i] = &Reminder{Id: r.Id, FreshUntil: r.FreshUntil}
		}
		tclock.Add(time.Hour)

		Convey("Normal operation", func() {
			err := impl.PostProcessBatch(ctx, batch, nil)
			So(err, ShouldBeNil)
			So(len(db.AllReminders()), ShouldEqual, 0)
		})
		Convey("Concurrent deletion", func() {
			So(db.DeleteReminder(ctx, batch[0]), ShouldBeNil)
			So(db.DeleteReminder(ctx, batch[9]), ShouldBeNil)
			err := impl.PostProcessBatch(ctx, batch, nil)
			So(err, ShouldBeNil)
			So(len(db.AllReminders()), ShouldEqual, 0)
		})
		Convey("Task Creation failures", func() {
			fakeCT.errs = []error{
				status.Errorf(codes.Unavailable, "please retry"),
				status.Errorf(codes.InvalidArgument, "user error"),
			}
			err := impl.PostProcessBatch(ctx, batch, nil)
			So(err, ShouldNotBeNil)
			merr := err.(errors.MultiError)
			So(len(merr), ShouldEqual, 1)
			So(merr[0], ShouldErrLike, "please retry")
			So(len(db.AllReminders()), ShouldEqual, 1)
		})
	})
}

func TestLeaseHelpers(t *testing.T) {
	t.Parallel()

	Convey("Lease Helpers", t, func() {

		Convey("onlyLeased", func() {
			reminders := []*Reminder{
				// Each key be exactly 2*keySpaceBytes chars long.
				&Reminder{Id: "00000000000000000000000000000001"},
				&Reminder{Id: "00000000000000000000000000000005"},
				&Reminder{Id: "00000000000000000000000000000009"},
				&Reminder{Id: "0000000000000000000000000000000f"}, // ie 15
			}
			leased := partition.SortedPartitions{partition.FromInts(5, 9)}
			So(onlyLeased(reminders, leased), ShouldResemble, []*Reminder{
				&Reminder{Id: "00000000000000000000000000000005"},
			})
		})
	})
}

// helpers

var genReminderCounter = int32(0)

func genReminder(ctx context.Context) *Reminder {
	id := atomic.AddInt32(&genReminderCounter, 1)
	req := taskspb.CreateTaskRequest{
		Parent: fmt.Sprintf("projects/example-project/locations/us-central1/queues/q-%d", id),
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{HttpRequest: &taskspb.HttpRequest{
				HttpMethod: taskspb.HttpMethod_GET,
				Url:        fmt.Sprintf("https://example.com/user/exec/%d", id),
			}},
			ScheduleTime: google.NewTimestamp(clock.Now(ctx).Add(time.Minute)),
		},
	}
	r, _, err := makeReminder(ctx, &req)
	if err != nil {
		panic(err)
	}
	return r
}

type fakeCloudTasks struct {
	mu    sync.Mutex
	errs  []error
	calls []*taskspb.CreateTaskRequest
}

func (f *fakeCloudTasks) CreateTask(_ context.Context, req *taskspb.CreateTaskRequest, _ ...gax.CallOption) (*taskspb.Task, error) {
	// NOTE: actual implementation returns the created Task, but the TTQ library
	// doesn't care.
	f.mu.Lock()
	defer f.mu.Unlock()
	f.calls = append(f.calls, req)
	if len(f.errs) > 0 {
		var err error
		f.errs, err = f.errs[1:], f.errs[0]
		return nil, err
	}
	return nil, nil
}
