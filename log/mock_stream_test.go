package log

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockStream is a mock an instance of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamRecorder
}

var _ Stream = &MockStream{}

// MockStreamRecorder is the mock recorder for MockStream
type MockStreamRecorder struct {
	mock *MockStream
}

// NewMockStream creates a new mock instance
func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &MockStreamRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStream) EXPECT() *MockStreamRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStream) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStreamRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStream)(nil).Close))
}

// Level mocks base method
func (m *MockStream) Level() Level {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Level")
	ret0, _ := ret[0].(Level)
	return ret0
}

// Level indicates an expected call of Level
func (mr *MockStreamRecorder) Level() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Level", reflect.TypeOf((*MockStream)(nil).Level))
}

// Signal mocks base method
func (m *MockStream) Signal(channel string, level Level, message string, ctx ...Context) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{channel, level, message}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Signal", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal
func (mr *MockStreamRecorder) Signal(channel, level, message interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{channel, level, message}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockStream)(nil).Signal), varargs...)
}

// Broadcast mocks base method
func (m *MockStream) Broadcast(level Level, message string, fields ...Context) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{level, message}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Broadcast", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockStreamRecorder) Broadcast(level, message interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{level, message}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockStream)(nil).Broadcast), varargs...)
}

// HasChannel mocks base method
func (m *MockStream) HasChannel(channel string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasChannel", channel)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasChannel indicates an expected call of HasChannel
func (mr *MockStreamRecorder) HasChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChannel", reflect.TypeOf((*MockStream)(nil).HasChannel), channel)
}

// ListChannels mocks base method
func (m *MockStream) ListChannels() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChannels")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListChannels indicates an expected call of ListChannels
func (mr *MockStreamRecorder) ListChannels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChannels", reflect.TypeOf((*MockStream)(nil).ListChannels))
}

// AddChannel mocks base method
func (m *MockStream) AddChannel(channel string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChannel", channel)
}

// AddChannel indicates an expected call of AddChannel
func (mr *MockStreamRecorder) AddChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChannel", reflect.TypeOf((*MockStream)(nil).AddChannel), channel)
}

// RemoveChannel mocks base method
func (m *MockStream) RemoveChannel(channel string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveChannel", channel)
}

// RemoveChannel indicates an expected call of RemoveChannel
func (mr *MockStreamRecorder) RemoveChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChannel", reflect.TypeOf((*MockStream)(nil).RemoveChannel), channel)
}
