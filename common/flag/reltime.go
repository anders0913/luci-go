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

// Package flag provides the utilities for implementing command line flags.
package flag

import (
	"strconv"
	"time"

	"go.chromium.org/luci/common/errors"
)

// RelativeTime is an implementation of flag.Value for parsing a time
// by a relative day offset.
type RelativeTime struct {
	T   *time.Time
	now func() time.Time
}

// String implements the flag.Value interface.
func (f RelativeTime) String() string {
	if f.T == nil {
		return "<empty>"
	}
	return f.T.Format(time.RFC1123Z)
}

// Set implements the flag.Value interface.
func (f RelativeTime) Set(s string) error {
	if f.T == nil {
		return errors.Reason("set RelativeTime: nil time pointer").Err()
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return errors.Annotate(err, "set RelativeTime").Err()
	}
	if f.now == nil {
		f.now = time.Now
	}
	*f.T = f.now().Add(time.Duration(n*24) * time.Hour)
	return nil
}
