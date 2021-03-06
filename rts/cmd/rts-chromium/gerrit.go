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
	"fmt"
	"sort"
	"strconv"
	"time"

	"golang.org/x/time/rate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/common/errors"
	gerritpb "go.chromium.org/luci/common/proto/gerrit"
	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/grpc/grpcutil"

	evalpb "go.chromium.org/luci/rts/presubmit/eval/proto"
)

type gerritClient struct {
	// listFilesRPC makes a Gerrit RPC to fetch the list of changed files.
	// Mockable.
	listFilesRPC  func(ctx context.Context, host string, req *gerritpb.ListFilesRequest) (*gerritpb.ListFilesResponse, error)
	limiter       *rate.Limiter
	fileListCache cache
}

type changedFiles struct {
	Names []string `json:"names"`
}

// ChangedFiles returns the list of files changed in the given patchset.
// Each file is a relative path, e.g. "chrome.cc".
// If the patchset does not exist, returns empty list.
func (c *gerritClient) ChangedFiles(ctx context.Context, ps *evalpb.GerritPatchset) ([]string, error) {
	cacheKey := fmt.Sprintf("%s-%d-%d", ps.Change.Host, ps.Change.Number, ps.Patchset)

	value, err := c.fileListCache.GetOrCreate(ctx, cacheKey, func() (interface{}, error) {
		files, err := c.fetchChangedFiles(ctx, ps)
		if err != nil {
			return nil, err
		}
		return &changedFiles{Names: files}, nil
	})
	if err != nil {
		return nil, err
	}
	return value.(*changedFiles).Names, nil
}

func (c *gerritClient) fetchChangedFiles(ctx context.Context, ps *evalpb.GerritPatchset) ([]string, error) {
	var files []string
	err := retry.Retry(ctx, transient.Only(retry.Default), func() (err error) {
		files, err = c.listFilesWithQuotaErrorsRetries(ctx, ps.Change.Host, &gerritpb.ListFilesRequest{
			Project:    ps.Change.Project,
			Number:     int64(ps.Change.Number),
			RevisionId: strconv.Itoa(int(ps.Patchset)),
		})
		if grpcutil.IsTransientCode(statusCode(err)) {
			err = transient.Tag.Apply(err)
		}
		return
	}, retry.LogCallback(ctx, fmt.Sprintf("read %s", psURL(ps))))
	return files, err
}

// listFilesWithQuotaErrorsRetries fetches the list of changed files.
// If the request fails with quota exhaustion, retries the request in a second,
// up to 5 times.
// Does not retry other transient errors, e.g. internal errors.
func (c *gerritClient) listFilesWithQuotaErrorsRetries(ctx context.Context, host string, req *gerritpb.ListFilesRequest) ([]string, error) {
	// Retry ResourceExhausted errors with an increased delay.
	iter := func() retry.Iterator {
		base := retry.Limited{
			Delay:   5 * time.Second, // short-term quota resets at most every 5s.
			Retries: 5,
		}
		return retry.NewIterator(func(ctx context.Context, err error) time.Duration {
			if statusCode(err) == codes.ResourceExhausted {
				return base.Next(ctx, err)
			}
			return retry.Stop
		})
	}

	var files []string
	err := retry.Retry(ctx, iter, func() (err error) {
		files, err = c.callListFiles(ctx, host, req)
		return
	}, nil)
	return files, err
}

// callListFiles makes a ListFiles RPC.
func (c *gerritClient) callListFiles(ctx context.Context, host string, req *gerritpb.ListFilesRequest) ([]string, error) {
	// Make an RPC.
	if err := c.limiter.Wait(ctx); err != nil {
		return nil, err
	}
	res, err := c.listFilesRPC(ctx, host, req)
	switch {
	case statusCode(err) == codes.NotFound:
		return nil, nil
	case err != nil:
		return nil, err
	}

	files := make([]string, 0, len(res.Files))
	for name := range res.Files {
		if name != "/COMMIT_MSG" {
			files = append(files, name)
		}
	}
	sort.Strings(files)
	return files, nil
}

func statusCode(err error) codes.Code {
	return status.Code(errors.Unwrap(err))
}

// psURL returns the patchset URL.
func psURL(p *evalpb.GerritPatchset) string {
	return fmt.Sprintf("https://%s/c/%d/%d", p.Change.Host, p.Change.Number, p.Patchset)
}
