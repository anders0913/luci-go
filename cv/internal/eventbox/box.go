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

// Package eventbox batches incoming events for a single Datastore entity
// for processing.
package eventbox

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/gae/service/datastore"
	"golang.org/x/sync/errgroup"

	"go.chromium.org/luci/cv/internal/dsset"
)

// Emit emits a new event with provided value and auto-generated unique ID.
func Emit(ctx context.Context, value []byte, to *datastore.Key) error {
	d := dsset.Set{Parent: to} // TombstonesDelay doesn't matter for Add.
	id := uuid.New().String()
	if err := d.Add(ctx, []dsset.Item{{ID: id, Value: value}}); err != nil {
		return errors.Annotate(err, "failed to send event").Err()
	}
	return nil
}

const tombstonesDelay = 5 * time.Minute

// List returns unprocessed events. For use in tests only.
func List(ctx context.Context, recipient *datastore.Key) (Events, error) {
	d := dsset.Set{
		Parent:          recipient,
		TombstonesDelay: tombstonesDelay,
	}
	switch l, err := d.List(ctx); {
	case err != nil:
		return nil, err
	default:
		return toEvents(l.Items), nil
	}
}

// ProcessBatch reliably processes events, while transactionally modifying state
// and performing arbitrary side effects.
func ProcessBatch(ctx context.Context, recipient *datastore.Key, p Processor) error {
	var state State
	var expectedEV EVersion
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		state, expectedEV, err = p.LoadState(ectx)
		return
	})
	d := dsset.Set{
		Parent:          recipient,
		TombstonesDelay: tombstonesDelay,
	}
	var listing *dsset.Listing
	eg.Go(func() (err error) {
		if listing, err = d.List(ectx); err == nil {
			err = dsset.CleanupGarbage(ectx, listing.Garbage)
		}
		return
	})
	if err := eg.Wait(); err != nil {
		return err
	}

	// Compute resulting state before transaction.
	transitions, err := p.Mutate(ctx, toEvents(listing.Items), state)
	switch {
	case err != nil:
		return err
	case len(transitions) == 0:
		return nil // nothing to do.
	}

	err = datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		switch latestEV, err := p.FetchEVersion(ctx); {
		case err != nil:
			return err
		case latestEV != expectedEV:
			return errors.Reason("Concurrent modification: EVersion read %d, but expected %d",
				latestEV, expectedEV).Tag(transient.Tag).Err()
		}
		popOp, err := d.BeginPop(ctx, listing)
		if err != nil {
			return err
		}
		var newState State
		eventsConsumed := 0
		for _, t := range transitions {
			switch err := t.apply(ctx, popOp); {
			case err != nil && !transient.Tag.In(err):
				logging.Errorf(ctx, "FIXME: only transient errors expected: %s", err)
				fallthrough
			case err != nil:
				return err
			default:
				newState = t.TransitionTo
				eventsConsumed += len(t.Events)
			}
		}

		logging.Debugf(ctx, "%d transitions, %d events", len(transitions), eventsConsumed)
		if newState != state {
			if err := p.SaveState(ctx, newState, expectedEV+1); err != nil {
				return err
			}
		} else {
			logging.Debugf(ctx, "state wasn't modified")
		}
		return dsset.FinishPop(ctx, popOp)
	}, nil)
	if err != nil {
		// Unconditionally mark error as transient. If Transition.Apply needs to
		// return non-transient errors, this code needs changing.
		return errors.Annotate(err, "failed to commit mutation").Tag(transient.Tag).Err()
	}
	return nil
}

// Processor defines safe way to process events in a batch.
type Processor interface {
	// LoadState is called to load the state before a transaction.
	LoadState(context.Context) (State, EVersion, error)
	// Mutate is called before a transaction to compute transitions based on a
	// batch of events.
	//
	// All actions that must be done atomically with updating state must be
	// encapsulated inside Transition.SideEffectFn callback.
	Mutate(context.Context, Events, State) ([]Transition, error)
	// FetchEVersion is called at the beginning of a transaction.
	//
	// The returned EVersion is compared against the one associated with a state
	// loaded via GetState. If different, the transaction is aborted and new state
	// isn't saved.
	FetchEVersion(ctx context.Context) (EVersion, error)
	// SaveState is called in a transaction to save the state if it has changed.
	//
	// The passed eversion is incremented value of eversion of what GetState
	// returned before.
	SaveState(context.Context, State, EVersion) error
}

// Event is an incoming event.
type Event dsset.Item

// Events are incoming events.
type Events []Event

// toEvents is an annoying redundant malloc to avoid exposing dsset.Item :(
func toEvents(items []dsset.Item) Events {
	es := make(Events, len(items))
	for i, item := range items {
		es[i] = Event(item)
	}
	return es
}

// State is an arbitrary object.
//
// Use a pointer to an actual state.
type State interface{}

// EVersion is recipient entity version.
type EVersion int

// SideEffectFn performs side effects with a Datastore transaction context.
// See Transition.SideEffectFn doc.
type SideEffectFn func(context.Context) error

// Chain combines several SideEffectFn.
func Chain(fs ...SideEffectFn) SideEffectFn {
	return func(ctx context.Context) error {
		for _, f := range fs {
			if err := f(ctx); err != nil {
				return err
			}
		}
		return nil
	}
}

// Transition is a state transition.
type Transition struct {
	// SideEffectFn is called in a transaction to atomically with the state change
	// execute any side effects of a state transition.
	//
	// Typical use is notifying other CV components via TQ tasks.
	// Can be nil, meaning there no side effects to execute.
	//
	// TODO(tandrii): introduce error tag to indicate that failure was clean and
	// should be treated as if Transition wasn't started, s.t. progress of all
	// transitions before can be saved.
	SideEffectFn SideEffectFn
	// Events to consume with this transition.
	Events Events
	// TransitionTo is a state to transition to.
	//
	// It's allowed to transition to the exact same state.
	TransitionTo State
}

func (t *Transition) apply(ctx context.Context, p *dsset.PopOp) error {
	if t.SideEffectFn != nil {
		if err := t.SideEffectFn(ctx); err != nil {
			return err
		}
	}
	for _, e := range t.Events {
		_ = p.Pop(e.ID) // silently ignore if event has already been consumed.
	}
	return nil
}
