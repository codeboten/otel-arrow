// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/f5/otel-arrow-adapter/api/experimental/arrow/v1 (interfaces: ArrowStreamServiceClient,ArrowStreamService_ArrowStreamClient,ArrowStreamServiceServer,ArrowStreamService_ArrowStreamServer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	v1 "github.com/f5/otel-arrow-adapter/api/experimental/arrow/v1"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
)

// MockArrowStreamServiceClient is a mock of ArrowStreamServiceClient interface.
type MockArrowStreamServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockArrowStreamServiceClientMockRecorder
}

// MockArrowStreamServiceClientMockRecorder is the mock recorder for MockArrowStreamServiceClient.
type MockArrowStreamServiceClientMockRecorder struct {
	mock *MockArrowStreamServiceClient
}

// NewMockArrowStreamServiceClient creates a new mock instance.
func NewMockArrowStreamServiceClient(ctrl *gomock.Controller) *MockArrowStreamServiceClient {
	mock := &MockArrowStreamServiceClient{ctrl: ctrl}
	mock.recorder = &MockArrowStreamServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArrowStreamServiceClient) EXPECT() *MockArrowStreamServiceClientMockRecorder {
	return m.recorder
}

// ArrowStream mocks base method.
func (m *MockArrowStreamServiceClient) ArrowStream(arg0 context.Context, arg1 ...grpc.CallOption) (v1.ArrowStreamService_ArrowStreamClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ArrowStream", varargs...)
	ret0, _ := ret[0].(v1.ArrowStreamService_ArrowStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArrowStream indicates an expected call of ArrowStream.
func (mr *MockArrowStreamServiceClientMockRecorder) ArrowStream(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArrowStream", reflect.TypeOf((*MockArrowStreamServiceClient)(nil).ArrowStream), varargs...)
}

// MockArrowStreamService_ArrowStreamClient is a mock of ArrowStreamService_ArrowStreamClient interface.
type MockArrowStreamService_ArrowStreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockArrowStreamService_ArrowStreamClientMockRecorder
}

// MockArrowStreamService_ArrowStreamClientMockRecorder is the mock recorder for MockArrowStreamService_ArrowStreamClient.
type MockArrowStreamService_ArrowStreamClientMockRecorder struct {
	mock *MockArrowStreamService_ArrowStreamClient
}

// NewMockArrowStreamService_ArrowStreamClient creates a new mock instance.
func NewMockArrowStreamService_ArrowStreamClient(ctrl *gomock.Controller) *MockArrowStreamService_ArrowStreamClient {
	mock := &MockArrowStreamService_ArrowStreamClient{ctrl: ctrl}
	mock.recorder = &MockArrowStreamService_ArrowStreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArrowStreamService_ArrowStreamClient) EXPECT() *MockArrowStreamService_ArrowStreamClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockArrowStreamService_ArrowStreamClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockArrowStreamService_ArrowStreamClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockArrowStreamService_ArrowStreamClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockArrowStreamService_ArrowStreamClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockArrowStreamService_ArrowStreamClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockArrowStreamService_ArrowStreamClient)(nil).Context))
}

// Header mocks base method.
func (m *MockArrowStreamService_ArrowStreamClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockArrowStreamService_ArrowStreamClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockArrowStreamService_ArrowStreamClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockArrowStreamService_ArrowStreamClient) Recv() (*v1.BatchStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.BatchStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockArrowStreamService_ArrowStreamClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockArrowStreamService_ArrowStreamClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockArrowStreamService_ArrowStreamClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockArrowStreamService_ArrowStreamClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockArrowStreamService_ArrowStreamClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockArrowStreamService_ArrowStreamClient) Send(arg0 *v1.BatchArrowRecords) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockArrowStreamService_ArrowStreamClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockArrowStreamService_ArrowStreamClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockArrowStreamService_ArrowStreamClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockArrowStreamService_ArrowStreamClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockArrowStreamService_ArrowStreamClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockArrowStreamService_ArrowStreamClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockArrowStreamService_ArrowStreamClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockArrowStreamService_ArrowStreamClient)(nil).Trailer))
}

// MockArrowStreamServiceServer is a mock of ArrowStreamServiceServer interface.
type MockArrowStreamServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockArrowStreamServiceServerMockRecorder
}

// MockArrowStreamServiceServerMockRecorder is the mock recorder for MockArrowStreamServiceServer.
type MockArrowStreamServiceServerMockRecorder struct {
	mock *MockArrowStreamServiceServer
}

// NewMockArrowStreamServiceServer creates a new mock instance.
func NewMockArrowStreamServiceServer(ctrl *gomock.Controller) *MockArrowStreamServiceServer {
	mock := &MockArrowStreamServiceServer{ctrl: ctrl}
	mock.recorder = &MockArrowStreamServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArrowStreamServiceServer) EXPECT() *MockArrowStreamServiceServerMockRecorder {
	return m.recorder
}

// ArrowStream mocks base method.
func (m *MockArrowStreamServiceServer) ArrowStream(arg0 v1.ArrowStreamService_ArrowStreamServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArrowStream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ArrowStream indicates an expected call of ArrowStream.
func (mr *MockArrowStreamServiceServerMockRecorder) ArrowStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArrowStream", reflect.TypeOf((*MockArrowStreamServiceServer)(nil).ArrowStream), arg0)
}

// mustEmbedUnimplementedArrowStreamServiceServer mocks base method.
func (m *MockArrowStreamServiceServer) mustEmbedUnimplementedArrowStreamServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedArrowStreamServiceServer")
}

// mustEmbedUnimplementedArrowStreamServiceServer indicates an expected call of mustEmbedUnimplementedArrowStreamServiceServer.
func (mr *MockArrowStreamServiceServerMockRecorder) mustEmbedUnimplementedArrowStreamServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedArrowStreamServiceServer", reflect.TypeOf((*MockArrowStreamServiceServer)(nil).mustEmbedUnimplementedArrowStreamServiceServer))
}

// MockArrowStreamService_ArrowStreamServer is a mock of ArrowStreamService_ArrowStreamServer interface.
type MockArrowStreamService_ArrowStreamServer struct {
	ctrl     *gomock.Controller
	recorder *MockArrowStreamService_ArrowStreamServerMockRecorder
}

// MockArrowStreamService_ArrowStreamServerMockRecorder is the mock recorder for MockArrowStreamService_ArrowStreamServer.
type MockArrowStreamService_ArrowStreamServerMockRecorder struct {
	mock *MockArrowStreamService_ArrowStreamServer
}

// NewMockArrowStreamService_ArrowStreamServer creates a new mock instance.
func NewMockArrowStreamService_ArrowStreamServer(ctrl *gomock.Controller) *MockArrowStreamService_ArrowStreamServer {
	mock := &MockArrowStreamService_ArrowStreamServer{ctrl: ctrl}
	mock.recorder = &MockArrowStreamService_ArrowStreamServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArrowStreamService_ArrowStreamServer) EXPECT() *MockArrowStreamService_ArrowStreamServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockArrowStreamService_ArrowStreamServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockArrowStreamService_ArrowStreamServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockArrowStreamService_ArrowStreamServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockArrowStreamService_ArrowStreamServer) Recv() (*v1.BatchArrowRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.BatchArrowRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockArrowStreamService_ArrowStreamServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockArrowStreamService_ArrowStreamServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockArrowStreamService_ArrowStreamServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockArrowStreamService_ArrowStreamServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockArrowStreamService_ArrowStreamServer)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockArrowStreamService_ArrowStreamServer) Send(arg0 *v1.BatchStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockArrowStreamService_ArrowStreamServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockArrowStreamService_ArrowStreamServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockArrowStreamService_ArrowStreamServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockArrowStreamService_ArrowStreamServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockArrowStreamService_ArrowStreamServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m *MockArrowStreamService_ArrowStreamServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockArrowStreamService_ArrowStreamServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockArrowStreamService_ArrowStreamServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method.
func (m *MockArrowStreamService_ArrowStreamServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockArrowStreamService_ArrowStreamServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockArrowStreamService_ArrowStreamServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockArrowStreamService_ArrowStreamServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockArrowStreamService_ArrowStreamServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockArrowStreamService_ArrowStreamServer)(nil).SetTrailer), arg0)
}