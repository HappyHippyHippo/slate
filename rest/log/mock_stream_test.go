package log

import (
	"reflect"

	"github.com/golang/mock/gomock"
	slog "github.com/happyhippyhippo/slate/log"
)

// MockStream is a mock of LogStream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *MockStreamRecorder
}

var _ slog.IStream = &MockStream{}

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

// Signal mocks base method
func (m *MockStream) Signal(channel string, level slog.Level, message string, fields map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signal", channel, level, message, fields)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal
func (mr *MockStreamRecorder) Signal(channel, level, message, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockStream)(nil).Signal), channel, level, message, fields)
}

// Broadcast mocks base method
func (m *MockStream) Broadcast(level slog.Level, message string, fields map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", level, message, fields)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockStreamRecorder) Broadcast(level, message, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockStream)(nil).Broadcast), level, message, fields)
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

// Level mocks base method
func (m *MockStream) Level() slog.Level {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Level")
	ret0, _ := ret[0].(slog.Level)
	return ret0
}

// Level indicates an expected call of Level
func (mr *MockStreamRecorder) Level() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Level", reflect.TypeOf((*MockStream)(nil).Level))
}
