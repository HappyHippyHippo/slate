package watchdog

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockLogFormatter is a mock an instance of DefaultLogFormatter interface.
type MockLogFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterRecorder
}

var _ LogFormatter = &MockLogFormatter{}

// MockLogFormatterRecorder is the mock recorder for MockLogFormatter.
type MockLogFormatterRecorder struct {
	mock *MockLogFormatter
}

// NewMockLogFormatter creates a new mock instance.
func NewMockLogFormatter(ctrl *gomock.Controller) *MockLogFormatter {
	mock := &MockLogFormatter{ctrl: ctrl}
	mock.recorder = &MockLogFormatterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFormatter) EXPECT() *MockLogFormatterRecorder {
	return m.recorder
}

// Done mocks base method.
func (m *MockLogFormatter) Done(service string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done", service)
	ret0, _ := ret[0].(string)
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockLogFormatterRecorder) Done(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockLogFormatter)(nil).Done), service)
}

// Error mocks base method.
func (m *MockLogFormatter) Error(service string, e error) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", service, e)
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockLogFormatterRecorder) Error(service, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogFormatter)(nil).Error), service, e)
}

// Start mocks base method.
func (m *MockLogFormatter) Start(service string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", service)
	ret0, _ := ret[0].(string)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockLogFormatterRecorder) Start(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockLogFormatter)(nil).Start), service)
}
