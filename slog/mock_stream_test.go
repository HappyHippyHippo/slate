package slog

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockStreamer is a mock of LogStream interface
type MockStreamer struct {
	ctrl     *gomock.Controller
	recorder *MockStreamerRecorder
}

var _ Stream = &MockStreamer{}

// MockStreamerRecorder is the mock recorder for MockStreamer
type MockStreamerRecorder struct {
	mock *MockStreamer
}

// NewMockStream creates a new mock instance
func NewMockStream(ctrl *gomock.Controller) *MockStreamer {
	mock := &MockStreamer{ctrl: ctrl}
	mock.recorder = &MockStreamerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamer) EXPECT() *MockStreamerRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockStreamer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockStreamerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStreamer)(nil).Close))
}

// Level mocks base method
func (m *MockStreamer) Level() Level {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Level")
	ret0, _ := ret[0].(Level)
	return ret0
}

// Level indicates an expected call of Level
func (mr *MockStreamerRecorder) Level() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Level", reflect.TypeOf((*MockStreamer)(nil).Level))
}

// Signal mocks base method
func (m *MockStreamer) Signal(channel string, level Level, message string, fields map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signal", channel, level, message, fields)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal
func (mr *MockStreamerRecorder) Signal(channel, level, message, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockStreamer)(nil).Signal), channel, level, message, fields)
}

// Broadcast mocks base method
func (m *MockStreamer) Broadcast(level Level, message string, fields map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", level, message, fields)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockStreamerRecorder) Broadcast(level, message, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockStreamer)(nil).Broadcast), level, message, fields)
}

// HasChannel mocks base method
func (m *MockStreamer) HasChannel(channel string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasChannel", channel)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasChannel indicates an expected call of HasChannel
func (mr *MockStreamerRecorder) HasChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChannel", reflect.TypeOf((*MockStreamer)(nil).HasChannel), channel)
}

// ListChannels mocks base method
func (m *MockStreamer) ListChannels() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChannels")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListChannels indicates an expected call of ListChannels
func (mr *MockStreamerRecorder) ListChannels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChannels", reflect.TypeOf((*MockStreamer)(nil).ListChannels))
}

// AddChannel mocks base method
func (m *MockStreamer) AddChannel(channel string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChannel", channel)
}

// AddChannel indicates an expected call of AddChannel
func (mr *MockStreamerRecorder) AddChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChannel", reflect.TypeOf((*MockStreamer)(nil).AddChannel), channel)
}

// RemoveChannel mocks base method
func (m *MockStreamer) RemoveChannel(channel string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveChannel", channel)
}

// RemoveChannel indicates an expected call of RemoveChannel
func (mr *MockStreamerRecorder) RemoveChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChannel", reflect.TypeOf((*MockStreamer)(nil).RemoveChannel), channel)
}
