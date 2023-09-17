package rest

import (
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate/log"
)

// MockLogStream is a mock instance of Stream interface.
type MockLogStream struct {
	ctrl     *gomock.Controller
	recorder *MockLogStreamRecorder
}

var _ log.Stream = &MockLogStream{}

// MockLogStreamRecorder is the mock recorder for MockLogStream.
type MockLogStreamRecorder struct {
	mock *MockLogStream
}

// NewMockLogStream creates a new mock instance.
func NewMockLogStream(ctrl *gomock.Controller) *MockLogStream {
	mock := &MockLogStream{ctrl: ctrl}
	mock.recorder = &MockLogStreamRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogStream) EXPECT() *MockLogStreamRecorder {
	return m.recorder
}

// AddChannel mocks base method.
func (m *MockLogStream) AddChannel(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChannel", arg0)
}

// AddChannel indicates an expected call of AddChannel.
func (mr *MockLogStreamRecorder) AddChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChannel", reflect.TypeOf((*MockLogStream)(nil).AddChannel), arg0)
}

// Broadcast mocks base method.
func (m *MockLogStream) Broadcast(arg0 log.Level, arg1 string, arg2 ...log.Context) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Broadcast", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockLogStreamRecorder) Broadcast(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockLogStream)(nil).Broadcast), varargs...)
}

// HasChannel mocks base method.
func (m *MockLogStream) HasChannel(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasChannel", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasChannel indicates an expected call of HasChannel.
func (mr *MockLogStreamRecorder) HasChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChannel", reflect.TypeOf((*MockLogStream)(nil).HasChannel), arg0)
}

// ListChannels mocks base method.
func (m *MockLogStream) ListChannels() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChannels")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListChannels indicates an expected call of ListChannels.
func (mr *MockLogStreamRecorder) ListChannels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChannels", reflect.TypeOf((*MockLogStream)(nil).ListChannels))
}

// RemoveChannel mocks base method.
func (m *MockLogStream) RemoveChannel(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveChannel", arg0)
}

// RemoveChannel indicates an expected call of RemoveChannel.
func (mr *MockLogStreamRecorder) RemoveChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChannel", reflect.TypeOf((*MockLogStream)(nil).RemoveChannel), arg0)
}

// Signal mocks base method.
func (m *MockLogStream) Signal(arg0 string, arg1 log.Level, arg2 string, arg3 ...log.Context) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Signal", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal.
func (mr *MockLogStreamRecorder) Signal(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockLogStream)(nil).Signal), varargs...)
}
