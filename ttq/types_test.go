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

package ttq

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOptions(t *testing.T) {
	t.Parallel()

	Convey("Validate", t, func() {
		valid := Options{}
		Convey("Valid", func() {
			So(valid.Validate(), ShouldBeNil)
			So(valid.Shards, ShouldEqual, 16)
		})
		Convey("Allow non default", func() {
			valid.Shards = 17
			So(valid.Validate(), ShouldBeNil)
			So(valid.Shards, ShouldEqual, 17)
		})
	})
}
