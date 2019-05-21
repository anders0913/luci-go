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

package main

import (
	"flag"
	"os"

	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/server"
	"go.chromium.org/luci/server/router"
)

func main() {
	opts := server.Options{}
	opts.Register(flag.CommandLine)
	flag.Parse()

	srv := server.New(opts)

	srv.Routes.GET("/", router.MiddlewareChain{}, func(c *router.Context) {
		logging.Infof(c.Context, "Hello info world")
		logging.Warningf(c.Context, "Hello warning world")
		c.Writer.Write([]byte("Hello, world"))
	})

	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}