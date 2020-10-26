// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package scheduler

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulerClient interface {
	// GetJobs fetches all jobs satisfying JobsRequest and visibility ACLs.
	// If JobsRequest.project is specified but the project doesn't exist, empty
	// list of Jobs is returned.
	GetJobs(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (*JobsReply, error)
	// GetInvocations fetches invocations of a given job, most recent first.
	GetInvocations(ctx context.Context, in *InvocationsRequest, opts ...grpc.CallOption) (*InvocationsReply, error)
	// GetInvocation fetches a single invocation.
	GetInvocation(ctx context.Context, in *InvocationRef, opts ...grpc.CallOption) (*Invocation, error)
	// PauseJob will prevent automatic triggering of a job. Manual triggering such
	// as through this API is still allowed. Any pending or running invocations
	// are still executed. PauseJob does nothing if job is already paused.
	//
	// Requires OWNER Job permission.
	PauseJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*empty.Empty, error)
	// ResumeJob resumes paused job. ResumeJob does nothing if job is not paused.
	//
	// Requires OWNER Job permission.
	ResumeJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*empty.Empty, error)
	// AbortJob resets the job to scheduled state, aborting a currently pending or
	// running invocation if any.
	//
	// Note, that this is similar to AbortInvocation except that AbortInvocation
	// requires invocation ID and doesn't ensure that the invocation aborted is
	// actually latest triggered for the job.
	//
	// Requires OWNER Job permission.
	AbortJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*empty.Empty, error)
	// AbortInvocation aborts a given job invocation.
	// If an invocation is final, AbortInvocation does nothing.
	//
	// If you want to abort a specific hung invocation, use this request instead
	// of AbortJob.
	//
	// Requires OWNER Job permission.
	AbortInvocation(ctx context.Context, in *InvocationRef, opts ...grpc.CallOption) (*empty.Empty, error)
	// EmitTriggers puts one or more triggers into pending trigger queues of the
	// specified jobs.
	//
	// This eventually causes jobs to start executing. The scheduler may merge
	// multiple triggers into one job execution, based on how the job is
	// configured.
	//
	// If at least one job doesn't exist or the caller has no permission to
	// trigger it, the entire request is aborted. Otherwise, the request is NOT
	// transactional: if it fails midway (e.g by returning internal server error),
	// some triggers may have been submitted and some may not. It is safe to retry
	// the call, supplying the same trigger IDs. Triggers with the same IDs will
	// be deduplicated. See Trigger message for more details.
	//
	// Requires TRIGGERER Job permission.
	EmitTriggers(ctx context.Context, in *EmitTriggersRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) GetJobs(ctx context.Context, in *JobsRequest, opts ...grpc.CallOption) (*JobsReply, error) {
	out := new(JobsReply)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/GetJobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetInvocations(ctx context.Context, in *InvocationsRequest, opts ...grpc.CallOption) (*InvocationsReply, error) {
	out := new(InvocationsReply)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/GetInvocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) GetInvocation(ctx context.Context, in *InvocationRef, opts ...grpc.CallOption) (*Invocation, error) {
	out := new(Invocation)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/GetInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) PauseJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/PauseJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) ResumeJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/ResumeJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AbortJob(ctx context.Context, in *JobRef, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/AbortJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) AbortInvocation(ctx context.Context, in *InvocationRef, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/AbortInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) EmitTriggers(ctx context.Context, in *EmitTriggersRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/EmitTriggers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
// All implementations must embed UnimplementedSchedulerServer
// for forward compatibility
type SchedulerServer interface {
	// GetJobs fetches all jobs satisfying JobsRequest and visibility ACLs.
	// If JobsRequest.project is specified but the project doesn't exist, empty
	// list of Jobs is returned.
	GetJobs(context.Context, *JobsRequest) (*JobsReply, error)
	// GetInvocations fetches invocations of a given job, most recent first.
	GetInvocations(context.Context, *InvocationsRequest) (*InvocationsReply, error)
	// GetInvocation fetches a single invocation.
	GetInvocation(context.Context, *InvocationRef) (*Invocation, error)
	// PauseJob will prevent automatic triggering of a job. Manual triggering such
	// as through this API is still allowed. Any pending or running invocations
	// are still executed. PauseJob does nothing if job is already paused.
	//
	// Requires OWNER Job permission.
	PauseJob(context.Context, *JobRef) (*empty.Empty, error)
	// ResumeJob resumes paused job. ResumeJob does nothing if job is not paused.
	//
	// Requires OWNER Job permission.
	ResumeJob(context.Context, *JobRef) (*empty.Empty, error)
	// AbortJob resets the job to scheduled state, aborting a currently pending or
	// running invocation if any.
	//
	// Note, that this is similar to AbortInvocation except that AbortInvocation
	// requires invocation ID and doesn't ensure that the invocation aborted is
	// actually latest triggered for the job.
	//
	// Requires OWNER Job permission.
	AbortJob(context.Context, *JobRef) (*empty.Empty, error)
	// AbortInvocation aborts a given job invocation.
	// If an invocation is final, AbortInvocation does nothing.
	//
	// If you want to abort a specific hung invocation, use this request instead
	// of AbortJob.
	//
	// Requires OWNER Job permission.
	AbortInvocation(context.Context, *InvocationRef) (*empty.Empty, error)
	// EmitTriggers puts one or more triggers into pending trigger queues of the
	// specified jobs.
	//
	// This eventually causes jobs to start executing. The scheduler may merge
	// multiple triggers into one job execution, based on how the job is
	// configured.
	//
	// If at least one job doesn't exist or the caller has no permission to
	// trigger it, the entire request is aborted. Otherwise, the request is NOT
	// transactional: if it fails midway (e.g by returning internal server error),
	// some triggers may have been submitted and some may not. It is safe to retry
	// the call, supplying the same trigger IDs. Triggers with the same IDs will
	// be deduplicated. See Trigger message for more details.
	//
	// Requires TRIGGERER Job permission.
	EmitTriggers(context.Context, *EmitTriggersRequest) (*empty.Empty, error)
	mustEmbedUnimplementedSchedulerServer()
}

// UnimplementedSchedulerServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (UnimplementedSchedulerServer) GetJobs(context.Context, *JobsRequest) (*JobsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobs not implemented")
}
func (UnimplementedSchedulerServer) GetInvocations(context.Context, *InvocationsRequest) (*InvocationsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvocations not implemented")
}
func (UnimplementedSchedulerServer) GetInvocation(context.Context, *InvocationRef) (*Invocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvocation not implemented")
}
func (UnimplementedSchedulerServer) PauseJob(context.Context, *JobRef) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseJob not implemented")
}
func (UnimplementedSchedulerServer) ResumeJob(context.Context, *JobRef) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeJob not implemented")
}
func (UnimplementedSchedulerServer) AbortJob(context.Context, *JobRef) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortJob not implemented")
}
func (UnimplementedSchedulerServer) AbortInvocation(context.Context, *InvocationRef) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortInvocation not implemented")
}
func (UnimplementedSchedulerServer) EmitTriggers(context.Context, *EmitTriggersRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitTriggers not implemented")
}
func (UnimplementedSchedulerServer) mustEmbedUnimplementedSchedulerServer() {}

// UnsafeSchedulerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulerServer will
// result in compilation errors.
type UnsafeSchedulerServer interface {
	mustEmbedUnimplementedSchedulerServer()
}

func RegisterSchedulerServer(s grpc.ServiceRegistrar, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_GetJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/GetJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetJobs(ctx, req.(*JobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetInvocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetInvocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/GetInvocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetInvocations(ctx, req.(*InvocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_GetInvocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvocationRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).GetInvocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/GetInvocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).GetInvocation(ctx, req.(*InvocationRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_PauseJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).PauseJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/PauseJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).PauseJob(ctx, req.(*JobRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_ResumeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).ResumeJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/ResumeJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).ResumeJob(ctx, req.(*JobRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AbortJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AbortJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/AbortJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AbortJob(ctx, req.(*JobRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_AbortInvocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvocationRef)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).AbortInvocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/AbortInvocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).AbortInvocation(ctx, req.(*InvocationRef))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_EmitTriggers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitTriggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).EmitTriggers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/EmitTriggers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).EmitTriggers(ctx, req.(*EmitTriggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJobs",
			Handler:    _Scheduler_GetJobs_Handler,
		},
		{
			MethodName: "GetInvocations",
			Handler:    _Scheduler_GetInvocations_Handler,
		},
		{
			MethodName: "GetInvocation",
			Handler:    _Scheduler_GetInvocation_Handler,
		},
		{
			MethodName: "PauseJob",
			Handler:    _Scheduler_PauseJob_Handler,
		},
		{
			MethodName: "ResumeJob",
			Handler:    _Scheduler_ResumeJob_Handler,
		},
		{
			MethodName: "AbortJob",
			Handler:    _Scheduler_AbortJob_Handler,
		},
		{
			MethodName: "AbortInvocation",
			Handler:    _Scheduler_AbortInvocation_Handler,
		},
		{
			MethodName: "EmitTriggers",
			Handler:    _Scheduler_EmitTriggers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/scheduler/api/scheduler/v1/scheduler.proto",
}