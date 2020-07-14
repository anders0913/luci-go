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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestValidateQueueName(t *testing.T) {
	t.Parallel()

	Convey("ValidateQueueName", t, func() {
		So(ValidateQueueName("projects/example-project/locations/us-central1/queues/ttq"),
			ShouldBeNil)
		So(ValidateQueueName(""), ShouldErrLike, "name not given")
		So(ValidateQueueName("bad"), ShouldErrLike, "must be in format")
	})
}
