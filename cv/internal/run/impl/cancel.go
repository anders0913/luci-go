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

package impl

import (
	"context"

	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/cv/internal/eventbox"
	"go.chromium.org/luci/cv/internal/run"
)

func cancel(ctx context.Context, runID run.ID, s *state) (eventbox.SideEffectFn, *state, error) {
	switch status := s.Run.Status; {
	case status == run.Status_STATUS_UNSPECIFIED:
		err := errors.Reason("CRITICAL: can't cancel a Run with unspecified status").Err()
		errors.Log(ctx, err)
		panic(err)
	case status == run.Status_FINALIZING:
		logging.Debugf(ctx, "can't cancel run as it is currently finalizing")
		return nil, s, nil
	case run.IsEnded(status):
		logging.Debugf(ctx, "can't cancel an already ended run")
		return nil, s, nil
	}

	ret := s.shallowCopy()
	ret.Run.Status = run.Status_CANCELLED
	now := clock.Now(ctx).UTC()
	ret.Run.EndTime = now
	if ret.Run.StartTime.IsZero() {
		// This run has never started but already gets a cancelled event.
		ret.Run.StartTime = now
	}
	return nil, ret, nil
}
