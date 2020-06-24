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

// Package ttqspanner implements task enqueueing into Cloud Tasks with
// transactional semantics from inside a Cloud Spanner read-write transaction.
//
// Limitations:
//   * Does NOT support named tasks, for which Cloud Tasks provides
//     de-duplication. This is also a limitation of AppEngine Classic
//     transactional Tasks enqueueing.
//     Therefore, if you need to de-duplicate, you need to do this in your
//     application code yourself.
//   * All limits of Cloud Tasks apply, see
//     https://cloud.google.com/tasks/docs/quotas.
package ttqspanner

import (
	"context"
	"errors"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/spanner"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"

	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/ttq"
	"go.chromium.org/luci/ttq/internal"
)

// TTQ implements transaction task enqueueing with Spanner backend.
type TTQ struct {
	impl internal.Impl
}

// New creates a new TTQ for Spanner.
// You must also call InstallRoutes in at least one of your app's microservices.
func New(c *cloudtasks.Client, opts ttq.Options, sp *spanner.Client) *TTQ {
	// TODO(tandrii): implement.
	return &TTQ{impl: internal.Impl{
		Options: opts,
		DB:      &db{client: sp}},
	}
}

// InstallRoutes installs handlers for sweeping to ensure correctness.
//
// Users must ensure at least one of their microservices calls InstallRoutes.
//
// Requires a SweepOptions.Queue to be available.
// Panics if called twice.
// Reserves for its own use all path under the given pathPrefix.
//
// TODO(tandrii): describe cron setup.
func (t *TTQ) InstallRoutes(r *router.Router, pathPrefix string, mw router.MiddlewareChain) {
	// TODO(tandrii): implement.
}

// AddTask guarantees eventual creation of a task in Cloud Tasks if the current
// transaction completes successfully.
//
// The returned ttq.PostProcess should be called after the successful
// transaction. See ttq.PostProcess
//   TODO(tandrii): link
// documentation for more info.
func (t *TTQ) AddTask(ctx context.Context, txn *spanner.ReadWriteTransaction, req *taskspb.CreateTaskRequest) (ttq.PostProcess, error) {
	// TODO(tandrii): implement.
	return nil, errors.New("not implemented")
}
