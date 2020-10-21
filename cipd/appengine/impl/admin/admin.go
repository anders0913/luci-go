// Copyright 2018 The LUCI Authors.
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

package admin

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/retry/transient"

	"go.chromium.org/luci/appengine/mapper"
	"go.chromium.org/luci/appengine/tq"

	api "go.chromium.org/luci/cipd/api/admin/v1"
	"go.chromium.org/luci/cipd/appengine/impl/rpcacl"
)

// AdminAPI returns an ACL-protected implementation of cipd.AdminServer that can
// be exposed as a public API (i.e. admins can use it via external RPCs).
func AdminAPI(d *tq.Dispatcher) api.AdminServer {
	impl := &adminImpl{
		acl: rpcacl.CheckAdmin,
		tq:  d,
	}
	impl.init()
	return impl
}

// adminImpl implements cipd.AdminServer.
type adminImpl struct {
	api.UnimplementedAdminServer

	acl func(context.Context) error
	tq  *tq.Dispatcher
	ctl *mapper.Controller
}

// init initializes mapper controller and registers mapping tasks.
func (impl *adminImpl) init() {
	impl.ctl = &mapper.Controller{
		MapperQueue:  "mappers", // see queue.yaml
		ControlQueue: "default",
	}
	for _, m := range mappers { // see mappers.go
		impl.ctl.RegisterFactory(m.mapperID(), m.newMapper)
	}
	impl.ctl.Install(impl.tq)
}

// toStatus converts an error from mapper.Controller to an grpc status.
//
// Passes nil as is.
func toStatus(err error) error {
	switch {
	case err == mapper.ErrNoSuchJob:
		return status.Errorf(codes.NotFound, "no such mapping job")
	case transient.Tag.In(err):
		return status.Errorf(codes.Internal, err.Error())
	case err != nil:
		return status.Errorf(codes.InvalidArgument, err.Error())
	default:
		return nil
	}
}

// LaunchJob implements the corresponding RPC method, see the proto doc.
func (impl *adminImpl) LaunchJob(ctx context.Context, cfg *api.JobConfig) (*api.JobID, error) {
	if err := impl.acl(ctx); err != nil {
		return nil, err
	}

	def, ok := mappers[cfg.Kind] // see mappers.go
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "unknown mapper kind")
	}

	cfgBlob, err := proto.Marshal(cfg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to marshal JobConfig - %s", err)
	}

	launchCfg := def.Config
	launchCfg.Mapper = def.mapperID()
	launchCfg.Params = cfgBlob

	jid, err := impl.ctl.LaunchJob(ctx, &launchCfg)
	if err != nil {
		return nil, toStatus(err)
	}

	return &api.JobID{JobId: int64(jid)}, nil
}

// AbortJob implements the corresponding RPC method, see the proto doc.
func (impl *adminImpl) AbortJob(ctx context.Context, id *api.JobID) (*empty.Empty, error) {
	if err := impl.acl(ctx); err != nil {
		return nil, err
	}
	_, err := impl.ctl.AbortJob(ctx, mapper.JobID(id.JobId))
	if err != nil {
		return nil, toStatus(err)
	}
	return &empty.Empty{}, nil
}

// GetJobState implements the corresponding RPC method, see the proto doc.
func (impl *adminImpl) GetJobState(ctx context.Context, id *api.JobID) (*api.JobState, error) {
	if err := impl.acl(ctx); err != nil {
		return nil, err
	}
	job, err := impl.ctl.GetJob(ctx, mapper.JobID(id.JobId))
	if err != nil {
		return nil, toStatus(err)
	}
	cfg := &api.JobConfig{}
	if err := proto.Unmarshal(job.Config.Params, cfg); err != nil {
		return nil, toStatus(errors.Annotate(err, "failed to unmarshal JobConfig").Err())
	}
	info, err := job.FetchInfo(ctx)
	if err != nil {
		return nil, toStatus(err)
	}
	return &api.JobState{Config: cfg, Info: info}, nil
}

// GetJobState implements the corresponding RPC method, see the proto doc.
func (impl *adminImpl) FixMarkedTags(ctx context.Context, id *api.JobID) (*api.TagFixReport, error) {
	if err := impl.acl(ctx); err != nil {
		return nil, err
	}
	tags, err := fixMarkedTags(ctx, mapper.JobID(id.JobId))
	if err != nil {
		return nil, toStatus(err)
	}
	return &api.TagFixReport{Fixed: tags}, nil
}
