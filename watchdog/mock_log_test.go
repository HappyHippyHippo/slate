package watchdog

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/log"
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

// Signal mocks base method.
func (m *MockLogger) Signal(channel string, level log.Level, msg string, ctx ...log.Context) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{channel, level, msg}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Signal", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal.
func (mr *MockLoggerRecorder) Signal(channel, level, msg interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{channel, level, msg}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockLogger)(nil).Signal), varargs...)
}
