package log

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockLogger is a mock of ILogger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerRecorder
}

// MockLoggerRecorder is the mock recorder for MockLogger.
type MockLoggerRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerRecorder {
	return m.recorder
}

// AddStream mocks base method.
func (m *MockLogger) AddStream(id string, stream IStream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStream", id, stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddStream indicates an expected call of AddStream.
func (mr *MockLoggerRecorder) AddStream(id, stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStream", reflect.TypeOf((*MockLogger)(nil).AddStream), id, stream)
}

// Broadcast mocks base method.
func (m *MockLogger) Broadcast(level Level, msg string, ctx map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", level, msg, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockLoggerRecorder) Broadcast(level, msg, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockLogger)(nil).Broadcast), level, msg, ctx)
}

// Close mocks base method.
func (m *MockLogger) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockLoggerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLogger)(nil).Close))
}

// HasStream mocks base method.
func (m *MockLogger) HasStream(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasStream", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasStream indicates an expected call of HasStream.
func (mr *MockLoggerRecorder) HasStream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasStream", reflect.TypeOf((*MockLogger)(nil).HasStream), id)
}

// ListStreams mocks base method.
func (m *MockLogger) ListStreams() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStreams")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListStreams indicates an expected call of ListStreams.
func (mr *MockLoggerRecorder) ListStreams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreams", reflect.TypeOf((*MockLogger)(nil).ListStreams))
}

// RemoveAllStreams mocks base method.
func (m *MockLogger) RemoveAllStreams() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveAllStreams")
}

// RemoveAllStreams indicates an expected call of RemoveAllStreams.
func (mr *MockLoggerRecorder) RemoveAllStreams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllStreams", reflect.TypeOf((*MockLogger)(nil).RemoveAllStreams))
}

// RemoveStream mocks base method.
func (m *MockLogger) RemoveStream(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveStream", id)
}

// RemoveStream indicates an expected call of RemoveStream.
func (mr *MockLoggerRecorder) RemoveStream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStream", reflect.TypeOf((*MockLogger)(nil).RemoveStream), id)
}

// Signal mocks base method.
func (m *MockLogger) Signal(channel string, level Level, msg string, ctx map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signal", channel, level, msg, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal.
func (mr *MockLoggerRecorder) Signal(channel, level, msg, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockLogger)(nil).Signal), channel, level, msg, ctx)
}

// Stream mocks base method.
func (m *MockLogger) Stream(id string) IStream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", id)
	ret0, _ := ret[0].(IStream)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockLoggerRecorder) Stream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockLogger)(nil).Stream), id)
}
