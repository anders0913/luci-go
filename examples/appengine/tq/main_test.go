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

package main

import (
	"context"
	"testing"
	"time"

	"go.chromium.org/gae/filter/txndefer"
	"go.chromium.org/gae/impl/memory"
	"go.chromium.org/gae/service/datastore"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/examples/appengine/tq/taskspb"

	"go.chromium.org/luci/server/tq"
	"go.chromium.org/luci/server/tq/tqtesting"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestQueue(t *testing.T) {
	t.Parallel()

	Convey("Chain works", t, func() {
		var epoch = time.Unix(1500000000, 0).UTC()

		// Need the test clock to emulate delayed tasks. Tick it whenever TQ waits.
		ctx, tc := testclock.UseTime(context.Background(), epoch)
		tc.SetTimerCallback(func(d time.Duration, t clock.Timer) {
			if testclock.HasTags(t, tqtesting.ClockTag) {
				tc.Add(d)
			}
		})

		// Need the datastore fake with txndefer filter installed. This filter is
		// required when using server/tq with transactional tasks. AddTask calls
		// will panic otherwise. It is installed in production server contexts by
		// default.
		ctx = txndefer.FilterRDS(memory.Use(ctx))

		// Put a Cloud Tasks scheduler fake to be used by AddTask.
		ctx, sched := tq.TestingContext(ctx, nil)

		var succeeded tqtesting.TaskList

		// Can tweak it more, if necessary.
		sched.TaskSucceeded = tqtesting.TasksCollector(&succeeded)
		sched.TaskFailed = func(ctx context.Context, task *tqtesting.Task) { panic("should not fail") }

		// Enqueue the first task.
		So(EnqueueCountDown(ctx, 5), ShouldBeNil)

		// Examine currently enqueue tasks.
		So(sched.Tasks().Payloads(), ShouldResembleProto, []*taskspb.CountDownTask{
			{Number: 5},
		})

		// Simulate the Cloud Tasks run loop until there's no more pending or
		// executing tasks left
		sched.Run(ctx, tqtesting.StopWhenDrained())

		// Verify all expected entities have been created, and when expected.
		numbers := map[int64]time.Duration{}
		datastore.GetTestable(ctx).CatchupIndexes()
		datastore.Run(ctx, datastore.NewQuery("ExampleEntity"), func(e *ExampleEntity) {
			numbers[e.ID] = e.LastUpdate.Sub(epoch)
		})
		So(numbers, ShouldResemble, map[int64]time.Duration{
			5: 100 * time.Millisecond,
			4: 200 * time.Millisecond,
			3: 300 * time.Millisecond,
			2: 400 * time.Millisecond,
			1: 500 * time.Millisecond,
		})

		// Can also examine all executed tasks.
		So(succeeded.Payloads(), ShouldResembleProto, []*taskspb.CountDownTask{
			{Number: 5},
			{Number: 4},
			{Number: 3},
			{Number: 2},
			{Number: 1},
			{Number: 0},
		})
	})
}
