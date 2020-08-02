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

package tqtesting

import (
	"container/heap"
	"context"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"

	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/stringset"
)

// ClockTag tags the clock used in scheduler's sleep.
const ClockTag = "tq-scheduler-sleep"

// Scheduler knows how to execute submitted tasks when they are due.
//
// This is a very primitive in-memory version of Cloud Tasks service that can be
// used in tests and on localhost.
//
// Scheduler implements tq.Submitter interface.
type Scheduler struct {
	// MaxAttempts is the maximum number of attempts for a task, including the
	// first attempt.
	//
	// If negative the number of attempts is unlimited.
	//
	// Default is 20.
	MaxAttempts int

	// MinBackoff is an initial retry delay for failed tasks.
	//
	// It is doubled after each failed attempt until it reaches MaxBackoff after
	// which it stays constant.
	//
	// Default is 1 sec.
	MinBackoff time.Duration

	// MaxBackoff is an upper limit on a retry delay.
	//
	// Default is 5 min.
	MaxBackoff time.Duration

	// TaskSucceeded is called from within the executor's `done` callback whenever
	// a task finishes successfully, perhaps after a bunch of retries.
	//
	// Receives the same context as passed to Run.
	TaskSucceeded func(ctx context.Context, task *Task)

	// TaskFailed is called from within the executor's `done` callback whenever
	// a task fails after being attempted MaxAttempts times.
	//
	// Receives the same context as passed to Run.
	TaskFailed func(ctx context.Context, task *Task)

	m         sync.Mutex         // a global lock protecting everything
	clock     clock.Clock        // used to make sure only one clock is used
	nextID    int64              // for generating task names
	seen      stringset.Set      // names of all tasks scheduled ever
	tasks     tasksHeap          // scheduled tasks, earliest to execute first
	executing map[*Task]struct{} // tasks being executed right now
	wg        sync.WaitGroup     // tracks 'executing' set
	wakeUp    chan struct{}      // used to wake up Run
}

// Task represents an enqueued or executing task.
type Task struct {
	Task      *taskspb.Task // a clone of original Task proto as passed to CreateTask
	Name      string        // full task name (perhaps generated)
	ETA       time.Time     // when the task is due, always set at now or in future
	Attempts  int           // 0 initially, incremented before each execution attempt
	Executing bool          // true if executing right now

	index int // index in tasksHeap
}

// Copy makes a shallow copy of the task.
func (t *Task) Copy() *Task {
	cpy := *t
	return &cpy
}

// Executor knows how to execute tasks when their ETA arrives.
type Executor interface {
	// Execute is called from Run to execute the task.
	//
	// The executor may execute the task right away in a blocking way or dispatch
	// it to some other goroutine. Either way it must call `done` callback when it
	// is done executing the task, indicating whether the task should be
	// reenqueued for a retry.
	//
	// It is safe to call Scheduler's CreateTask from inside Execute.
	//
	// Receives the exact same context as Run(...), in particular this context
	// is canceled when Run is done.
	Execute(ctx context.Context, t *Task, done func(retry bool))
}

// CreateTask scheduler a task for later execution.
func (s *Scheduler) CreateTask(ctx context.Context, req *taskspb.CreateTaskRequest) error {
	// Note: this validation is pretty sloppy. It validates only things Scheduler
	// depends on. It doesn't validate full conformance to Cloud Tasks API.
	if req.Parent == "" {
		return status.Errorf(codes.InvalidArgument, "no Parent in the request")
	}
	if req.Task == nil {
		return status.Errorf(codes.InvalidArgument, "no Task in the request")
	}
	if req.Task.Name != "" && !strings.HasPrefix(req.Task.Name, req.Parent+"/tasks/") {
		return status.Errorf(codes.InvalidArgument, "bad task name")
	}

	task := &Task{
		Task: proto.Clone(req.Task).(*taskspb.Task),
		Name: req.Task.Name,
		ETA:  req.Task.ScheduleTime.AsTime(),
	}
	if now := clock.Now(ctx); task.ETA.Before(now) {
		task.ETA = now
	}

	s.m.Lock()
	defer s.m.Unlock()

	s.checkClockLocked(ctx)

	if s.seen == nil {
		s.seen = stringset.New(1)
	}
	if s.executing == nil {
		s.executing = make(map[*Task]struct{}, 1)
	}

	if task.Name == "" {
		task.Name = fmt.Sprintf("%s/tasks/generated-task-id-%08d", req.Parent, s.nextID)
		s.nextID++
	} else if !s.seen.Add(task.Name) {
		return status.Errorf(codes.AlreadyExists, "task %q already exists", task.Name)
	}

	s.enqueueLocked(task)
	return nil
}

// Tasks returns a snapshot of the scheduler state.
//
// Recalculates it from scratch, so it is a pretty expensive call.
//
// Tasks are ordered by ETA: currently executing tasks first, then scheduled
// tasks.
func (s *Scheduler) Tasks() []*Task {
	s.m.Lock()
	defer s.m.Unlock()

	tasks := make([]*Task, 0, len(s.tasks)+len(s.executing))
	for _, t := range s.tasks {
		tasks = append(tasks, t.Copy())
	}
	for t := range s.executing {
		tasks = append(tasks, t.Copy())
	}

	sort.Slice(tasks, func(i, j int) bool {
		switch l, r := tasks[i], tasks[j]; {
		case l.Executing && !r.Executing:
			return true
		case !r.Executing && r.Executing:
			return false
		case l.ETA.Equal(r.ETA):
			return l.Name < r.Name
		default:
			return l.ETA.Before(r.ETA)
		}
	})

	return tasks
}

// Run executes the scheduler's loop until the context is canceled.
//
// Upon exit:
//   * The context is canceled.
//   * All executing tasks has finished, there still may be pending tasks.
//
// Panics if Run is already running (perhaps in another goroutine).
func (s *Scheduler) Run(ctx context.Context, e Executor) {
	func() {
		s.m.Lock()
		defer s.m.Unlock()
		s.checkClockLocked(ctx)
		if s.wakeUp != nil {
			panic("Run is already running")
		}
		s.wakeUp = make(chan struct{}, 1)
	}()

	defer func() {
		s.m.Lock()
		defer s.m.Unlock()
		close(s.wakeUp)
		s.wakeUp = nil
	}()

	// Waits for all initiated executing tasks to finish before returning.
	defer s.wg.Wait()

	for ctx.Err() == nil {
		switch task, nextETA, taskDone := s.tryDequeueTask(ctx); {
		case task != nil:
			// Pass the task to the executor. It may either execute it right away
			// or asynchronously later. Either way, when it is done it will call
			// the finalization callback.
			e.Execute(ctx, task, taskDone)
		case !nextETA.IsZero():
			select {
			case <-s.wakeUp:
			case <-clock.After(clock.Tag(ctx, ClockTag), nextETA.Sub(clock.Now(ctx))):
			}
		default:
			select {
			case <-s.wakeUp:
			case <-ctx.Done():
			}
		}
	}
}

// enqueueLocked adds the task to the task heap and wakes up the scheduler.
func (s *Scheduler) enqueueLocked(task *Task) {
	heap.Push(&s.tasks, task)

	// This would wake up Run if it is listening or does nothing if wakeUp is nil
	// (i.e. Run is not running).
	select {
	case s.wakeUp <- struct{}{}:
	default:
	}
}

// tryDequeueTask pops the earliest task if it is ready for execution.
//
// A task is executable if it has ETA <= now. If no tasks are ready, returns
// ETA of the earliest task or time.Time{} if the queue is empty.
//
// If pops a task, returns a callback that must be called (perhaps
// asynchronously) when the task finishes execution.
func (s *Scheduler) tryDequeueTask(ctx context.Context) (t *Task, eta time.Time, done func(retry bool)) {
	s.m.Lock()
	defer s.m.Unlock()

	if len(s.tasks) == 0 {
		return nil, time.Time{}, nil
	}
	if eta := s.tasks[0].ETA; eta.After(clock.Now(ctx)) {
		return nil, eta, nil
	}

	task := heap.Pop(&s.tasks).(*Task)
	task.Attempts++
	task.Executing = true
	s.executing[task] = struct{}{}
	s.wg.Add(1)

	return task, time.Time{}, func(retry bool) {
		defer s.wg.Done()

		reenqueued := false

		s.m.Lock()
		defer func() {
			s.m.Unlock()
			if !reenqueued {
				switch {
				case !retry && s.TaskSucceeded != nil:
					s.TaskSucceeded(ctx, task)
				case retry && s.TaskFailed != nil:
					s.TaskFailed(ctx, task)
				}
			}
		}()

		task.Executing = false
		delete(s.executing, task)

		if retry {
			if ok, delay := s.evalRetryLocked(task); ok {
				task.ETA = clock.Now(ctx).Add(delay)
				s.enqueueLocked(task)
				reenqueued = true
			}
		}
	}
}

// evalRetryLocked decides if a task should be retried and when.
func (s *Scheduler) evalRetryLocked(t *Task) (retry bool, delay time.Duration) {
	maxAttempts := s.MaxAttempts
	if maxAttempts == 0 {
		maxAttempts = 20
	}

	minBackoff := s.MinBackoff
	if minBackoff == 0 {
		minBackoff = time.Second
	}

	maxBackoff := s.MaxBackoff
	if maxBackoff == 0 {
		maxBackoff = 5 * time.Minute
	}

	if maxAttempts > 0 && t.Attempts >= maxAttempts {
		return false, 0
	}

	delay = time.Duration(math.Pow(2, float64(t.Attempts))) * minBackoff
	if delay > maxBackoff {
		delay = maxBackoff
	}
	return true, delay
}

// checkClockLocked panics if `ctx` uses an unexpected clock.
func (s *Scheduler) checkClockLocked(ctx context.Context) {
	clock := clock.Get(ctx)
	if s.clock == nil {
		s.clock = clock
	} else if s.clock != clock {
		panic("multiple clocks used with a single Scheduler, this is dangerous")
	}
}

////////////////////////////////////////////////////////////////////////////////

// tasksHeap is a heap of scheduled tasks, the implementation is copy-pasted
// from the godoc.
type tasksHeap []*Task

func (th tasksHeap) Len() int { return len(th) }

func (th tasksHeap) Less(i, j int) bool {
	l, r := th[i], th[j]
	if l.ETA.Equal(r.ETA) {
		return l.Name < r.Name
	}
	return l.ETA.Before(r.ETA)
}

func (th tasksHeap) Swap(i, j int) {
	th[i], th[j] = th[j], th[i]
	th[i].index = i
	th[j].index = j
}

func (th *tasksHeap) Push(x interface{}) {
	n := len(*th)
	item := x.(*Task)
	item.index = n
	*th = append(*th, item)
}

func (th *tasksHeap) Pop() interface{} {
	old := *th
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*th = old[0 : n-1]
	return item
}