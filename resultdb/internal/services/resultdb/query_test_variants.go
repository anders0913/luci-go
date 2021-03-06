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

package resultdb

import (
	"context"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/server/span"

	"go.chromium.org/luci/resultdb/internal/invocations"
	"go.chromium.org/luci/resultdb/internal/pagination"
	"go.chromium.org/luci/resultdb/internal/permissions"
	uipb "go.chromium.org/luci/resultdb/internal/proto/ui"
	"go.chromium.org/luci/resultdb/internal/testvariants"
)

// QueryTestVariants implements uipb.UIServer.
func (s *uiServer) QueryTestVariants(ctx context.Context, in *uipb.QueryTestVariantsRequest) (*uipb.QueryTestVariantsResponse, error) {
	if err := permissions.VerifyInvNames(ctx, permListTestResults, in.Invocations...); err != nil {
		return nil, err
	}

	if err := validateQueryRequest(in); err != nil {
		return nil, err
	}

	// Open a transaction.
	ctx, cancel := span.ReadOnlyTransaction(ctx)
	defer cancel()

	// Get the transitive closure.
	invs, err := invocations.Reachable(ctx, invocations.MustParseNames(in.Invocations))
	if err != nil {
		return nil, errors.Annotate(err, "failed to read the reach").Err()
	}

	// Query test variants.
	q := testvariants.Query{
		InvocationIDs: invs,
		PageSize:      pagination.AdjustPageSize(in.PageSize),
		PageToken:     in.PageToken,
	}
	tvs, token, err := q.Fetch(ctx)
	if err != nil {
		return nil, errors.Annotate(err, "failed to read test variants").Err()
	}

	return &uipb.QueryTestVariantsResponse{
		TestVariants:  tvs,
		NextPageToken: token,
	}, nil
}
