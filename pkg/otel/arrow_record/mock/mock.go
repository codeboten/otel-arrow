// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/f5/otel-arrow-adapter/pkg/otel/arrow_record (interfaces: ProducerAPI,ConsumerAPI)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	v1 "github.com/f5/otel-arrow-adapter/api/collector/arrow/v1"
	gomock "github.com/golang/mock/gomock"
	plog "go.opentelemetry.io/collector/pdata/plog"
	ptrace "go.opentelemetry.io/collector/pdata/ptrace"
)

// MockProducerAPI is a mock of ProducerAPI interface.
type MockProducerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockProducerAPIMockRecorder
}

// MockProducerAPIMockRecorder is the mock recorder for MockProducerAPI.
type MockProducerAPIMockRecorder struct {
	mock *MockProducerAPI
}

// NewMockProducerAPI creates a new mock instance.
func NewMockProducerAPI(ctrl *gomock.Controller) *MockProducerAPI {
	mock := &MockProducerAPI{ctrl: ctrl}
	mock.recorder = &MockProducerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducerAPI) EXPECT() *MockProducerAPIMockRecorder {
	return m.recorder
}

// BatchArrowRecordsFromLogs mocks base method.
func (m *MockProducerAPI) BatchArrowRecordsFromLogs(arg0 plog.Logs) (*v1.BatchArrowRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchArrowRecordsFromLogs", arg0)
	ret0, _ := ret[0].(*v1.BatchArrowRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchArrowRecordsFromLogs indicates an expected call of BatchArrowRecordsFromLogs.
func (mr *MockProducerAPIMockRecorder) BatchArrowRecordsFromLogs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchArrowRecordsFromLogs", reflect.TypeOf((*MockProducerAPI)(nil).BatchArrowRecordsFromLogs), arg0)
}

// BatchArrowRecordsFromTraces mocks base method.
func (m *MockProducerAPI) BatchArrowRecordsFromTraces(arg0 ptrace.Traces) (*v1.BatchArrowRecords, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchArrowRecordsFromTraces", arg0)
	ret0, _ := ret[0].(*v1.BatchArrowRecords)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchArrowRecordsFromTraces indicates an expected call of BatchArrowRecordsFromTraces.
func (mr *MockProducerAPIMockRecorder) BatchArrowRecordsFromTraces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchArrowRecordsFromTraces", reflect.TypeOf((*MockProducerAPI)(nil).BatchArrowRecordsFromTraces), arg0)
}

// MockConsumerAPI is a mock of ConsumerAPI interface.
type MockConsumerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerAPIMockRecorder
}

// MockConsumerAPIMockRecorder is the mock recorder for MockConsumerAPI.
type MockConsumerAPIMockRecorder struct {
	mock *MockConsumerAPI
}

// NewMockConsumerAPI creates a new mock instance.
func NewMockConsumerAPI(ctrl *gomock.Controller) *MockConsumerAPI {
	mock := &MockConsumerAPI{ctrl: ctrl}
	mock.recorder = &MockConsumerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerAPI) EXPECT() *MockConsumerAPIMockRecorder {
	return m.recorder
}

// LogsFrom mocks base method.
func (m *MockConsumerAPI) LogsFrom(arg0 *v1.BatchArrowRecords) ([]plog.Logs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogsFrom", arg0)
	ret0, _ := ret[0].([]plog.Logs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogsFrom indicates an expected call of LogsFrom.
func (mr *MockConsumerAPIMockRecorder) LogsFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogsFrom", reflect.TypeOf((*MockConsumerAPI)(nil).LogsFrom), arg0)
}

// TracesFrom mocks base method.
func (m *MockConsumerAPI) TracesFrom(arg0 *v1.BatchArrowRecords) ([]ptrace.Traces, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TracesFrom", arg0)
	ret0, _ := ret[0].([]ptrace.Traces)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TracesFrom indicates an expected call of TracesFrom.
func (mr *MockConsumerAPIMockRecorder) TracesFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TracesFrom", reflect.TypeOf((*MockConsumerAPI)(nil).TracesFrom), arg0)
}