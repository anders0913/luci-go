// Code generated by svcdec; DO NOT EDIT.

package api

import (
	"context"

	proto "github.com/golang/protobuf/proto"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type DecoratedAdmin struct {
	// Service is the service to decorate.
	Service AdminServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(ctx context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(ctx context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedAdmin) LaunchJob(ctx context.Context, req *JobConfig) (rsp *JobID, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "LaunchJob", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.LaunchJob(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "LaunchJob", rsp, err)
	}
	return
}

func (s *DecoratedAdmin) AbortJob(ctx context.Context, req *JobID) (rsp *emptypb.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "AbortJob", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.AbortJob(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "AbortJob", rsp, err)
	}
	return
}

func (s *DecoratedAdmin) GetJobState(ctx context.Context, req *JobID) (rsp *JobState, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "GetJobState", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.GetJobState(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "GetJobState", rsp, err)
	}
	return
}

func (s *DecoratedAdmin) FixMarkedTags(ctx context.Context, req *JobID) (rsp *TagFixReport, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "FixMarkedTags", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.FixMarkedTags(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "FixMarkedTags", rsp, err)
	}
	return
}
