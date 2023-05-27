package log

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockLogger is a mock of logger interface.
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
func (m *MockLogger) AddStream(id string, stream Stream) error {
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
