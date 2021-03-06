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
	"math/rand"
	"time"

	"go.chromium.org/luci/rts/presubmit/eval"
)

func main() {
	ctx := context.Background()
	rand.Seed(time.Now().Unix())
	eval.Main(ctx, func(ctx context.Context, in eval.Input, out *eval.Output) error {
		for i := range in.TestVariants {
			out.TestVariantAffectedness[i].Distance = rand.Float64()
		}
		return nil
	})
}
