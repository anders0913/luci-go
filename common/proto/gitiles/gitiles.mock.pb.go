// Code generated by MockGen. DO NOT EDIT.
// Source: gitiles.pb.go

// Package gitiles is a generated GoMock package.
package gitiles

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockGitilesClient is a mock of GitilesClient interface.
type MockGitilesClient struct {
	ctrl     *gomock.Controller
	recorder *MockGitilesClientMockRecorder
}

// MockGitilesClientMockRecorder is the mock recorder for MockGitilesClient.
type MockGitilesClientMockRecorder struct {
	mock *MockGitilesClient
}

// NewMockGitilesClient creates a new mock instance.
func NewMockGitilesClient(ctrl *gomock.Controller) *MockGitilesClient {
	mock := &MockGitilesClient{ctrl: ctrl}
	mock.recorder = &MockGitilesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitilesClient) EXPECT() *MockGitilesClientMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockGitilesClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Log", varargs...)
	ret0, _ := ret[0].(*LogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Log indicates an expected call of Log.
func (mr *MockGitilesClientMockRecorder) Log(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockGitilesClient)(nil).Log), varargs...)
}

// Refs mocks base method.
func (m *MockGitilesClient) Refs(ctx context.Context, in *RefsRequest, opts ...grpc.CallOption) (*RefsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Refs", varargs...)
	ret0, _ := ret[0].(*RefsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refs indicates an expected call of Refs.
func (mr *MockGitilesClientMockRecorder) Refs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refs", reflect.TypeOf((*MockGitilesClient)(nil).Refs), varargs...)
}

// Archive mocks base method.
func (m *MockGitilesClient) Archive(ctx context.Context, in *ArchiveRequest, opts ...grpc.CallOption) (*ArchiveResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Archive", varargs...)
	ret0, _ := ret[0].(*ArchiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Archive indicates an expected call of Archive.
func (mr *MockGitilesClientMockRecorder) Archive(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Archive", reflect.TypeOf((*MockGitilesClient)(nil).Archive), varargs...)
}

// DownloadFile mocks base method.
func (m *MockGitilesClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadFile", varargs...)
	ret0, _ := ret[0].(*DownloadFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFile indicates an expected call of DownloadFile.
func (mr *MockGitilesClientMockRecorder) DownloadFile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockGitilesClient)(nil).DownloadFile), varargs...)
}

// MockGitilesServer is a mock of GitilesServer interface.
type MockGitilesServer struct {
	ctrl     *gomock.Controller
	recorder *MockGitilesServerMockRecorder
}

// MockGitilesServerMockRecorder is the mock recorder for MockGitilesServer.
type MockGitilesServerMockRecorder struct {
	mock *MockGitilesServer
}

// NewMockGitilesServer creates a new mock instance.
func NewMockGitilesServer(ctrl *gomock.Controller) *MockGitilesServer {
	mock := &MockGitilesServer{ctrl: ctrl}
	mock.recorder = &MockGitilesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitilesServer) EXPECT() *MockGitilesServerMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockGitilesServer) Log(arg0 context.Context, arg1 *LogRequest) (*LogResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log", arg0, arg1)
	ret0, _ := ret[0].(*LogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Log indicates an expected call of Log.
func (mr *MockGitilesServerMockRecorder) Log(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockGitilesServer)(nil).Log), arg0, arg1)
}

// Refs mocks base method.
func (m *MockGitilesServer) Refs(arg0 context.Context, arg1 *RefsRequest) (*RefsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refs", arg0, arg1)
	ret0, _ := ret[0].(*RefsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refs indicates an expected call of Refs.
func (mr *MockGitilesServerMockRecorder) Refs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refs", reflect.TypeOf((*MockGitilesServer)(nil).Refs), arg0, arg1)
}

// Archive mocks base method.
func (m *MockGitilesServer) Archive(arg0 context.Context, arg1 *ArchiveRequest) (*ArchiveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Archive", arg0, arg1)
	ret0, _ := ret[0].(*ArchiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Archive indicates an expected call of Archive.
func (mr *MockGitilesServerMockRecorder) Archive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Archive", reflect.TypeOf((*MockGitilesServer)(nil).Archive), arg0, arg1)
}

// DownloadFile mocks base method.
func (m *MockGitilesServer) DownloadFile(arg0 context.Context, arg1 *DownloadFileRequest) (*DownloadFileResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", arg0, arg1)
	ret0, _ := ret[0].(*DownloadFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFile indicates an expected call of DownloadFile.
func (mr *MockGitilesServerMockRecorder) DownloadFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockGitilesServer)(nil).DownloadFile), arg0, arg1)
}
