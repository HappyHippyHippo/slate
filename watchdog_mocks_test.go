package slate

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// ----------------------------------------------------------------------------
// WatchdogLogFormatterCreator
// ----------------------------------------------------------------------------

// MockWatchdogLogFormatterCreator is a mock instance of LogFormatterStrategy interface.
type MockWatchdogLogFormatterCreator struct {
	ctrl     *gomock.Controller
	recorder *MockWatchdogLogFormatterCreatorRecorder
}

var _ WatchdogLogFormatterCreator = &MockWatchdogLogFormatterCreator{}

// MockWatchdogLogFormatterCreatorRecorder is the mock recorder for MockWatchdogLogFormatterCreator.
type MockWatchdogLogFormatterCreatorRecorder struct {
	mock *MockWatchdogLogFormatterCreator
}

// NewMockWatchdogLogFormatterCreator creates a new mock instance.
func NewMockWatchdogLogFormatterCreator(ctrl *gomock.Controller) *MockWatchdogLogFormatterCreator {
	mock := &MockWatchdogLogFormatterCreator{ctrl: ctrl}
	mock.recorder = &MockWatchdogLogFormatterCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchdogLogFormatterCreator) EXPECT() *MockWatchdogLogFormatterCreatorRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockWatchdogLogFormatterCreator) Accept(config *ConfigPartial) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", config)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockWatchdogLogFormatterCreatorRecorder) Accept(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockWatchdogLogFormatterCreator)(nil).Accept), config)
}

// Create mocks base method.
func (m *MockWatchdogLogFormatterCreator) Create(config *ConfigPartial) (WatchdogLogFormatter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", config)
	ret0, _ := ret[0].(WatchdogLogFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWatchdogLogFormatterCreatorRecorder) Create(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWatchdogLogFormatterCreator)(nil).Create), config)
}

// ----------------------------------------------------------------------------
// WatchdogLogFormatter
// ----------------------------------------------------------------------------

// MockWatchdogLogFormatter is a mock instance of DefaultLogFormatter interface.
type MockWatchdogLogFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockWatchdogLogFormatterRecorder
}

var _ WatchdogLogFormatter = &MockWatchdogLogFormatter{}

// MockWatchdogLogFormatterRecorder is the mock recorder for MockWatchdogLogFormatter.
type MockWatchdogLogFormatterRecorder struct {
	mock *MockWatchdogLogFormatter
}

// NewMockWatchdogLogFormatter creates a new mock instance.
func NewMockWatchdogLogFormatter(ctrl *gomock.Controller) *MockWatchdogLogFormatter {
	mock := &MockWatchdogLogFormatter{ctrl: ctrl}
	mock.recorder = &MockWatchdogLogFormatterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchdogLogFormatter) EXPECT() *MockWatchdogLogFormatterRecorder {
	return m.recorder
}

// Done mocks base method.
func (m *MockWatchdogLogFormatter) Done(service string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done", service)
	ret0, _ := ret[0].(string)
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockWatchdogLogFormatterRecorder) Done(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockWatchdogLogFormatter)(nil).Done), service)
}

// Error mocks base method.
func (m *MockWatchdogLogFormatter) Error(service string, e error) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", service, e)
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockWatchdogLogFormatterRecorder) Error(service, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockWatchdogLogFormatter)(nil).Error), service, e)
}

// Start mocks base method.
func (m *MockWatchdogLogFormatter) Start(service string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", service)
	ret0, _ := ret[0].(string)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockWatchdogLogFormatterRecorder) Start(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockWatchdogLogFormatter)(nil).Start), service)
}

// ----------------------------------------------------------------------------
// watchdogLogger
// ----------------------------------------------------------------------------

// MockWatchdogLogger is a mock instance of logger interface.
type MockWatchdogLogger struct {
	ctrl     *gomock.Controller
	recorder *MockWatchdogLoggerRecorder
}

var _ watchdogLogger = &MockWatchdogLogger{}

// MockWatchdogLoggerRecorder is the mock recorder for MockWatchdogLogger.
type MockWatchdogLoggerRecorder struct {
	mock *MockWatchdogLogger
}

// NewMockWatchdogLogger creates a new mock instance.
func NewMockWatchdogLogger(ctrl *gomock.Controller) *MockWatchdogLogger {
	mock := &MockWatchdogLogger{ctrl: ctrl}
	mock.recorder = &MockWatchdogLoggerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchdogLogger) EXPECT() *MockWatchdogLoggerRecorder {
	return m.recorder
}

// Signal mocks base method.
func (m *MockWatchdogLogger) Signal(channel string, level LogLevel, msg string, ctx ...LogContext) error {
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
func (mr *MockWatchdogLoggerRecorder) Signal(channel, level, msg interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{channel, level, msg}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockWatchdogLogger)(nil).Signal), varargs...)
}
