package log

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockFormatter is a mock instance of Formatter interface
type MockFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockFormatterRecorder
}

var _ Formatter = &MockFormatter{}

// MockFormatterRecorder is the mock recorder for MockFormatter
type MockFormatterRecorder struct {
	mock *MockFormatter
}

// NewMockFormatter creates a new mock instance
func NewMockFormatter(ctrl *gomock.Controller) *MockFormatter {
	mock := &MockFormatter{ctrl: ctrl}
	mock.recorder = &MockFormatterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFormatter) EXPECT() *MockFormatterRecorder {
	return m.recorder
}

// Format mocks base method
func (m *MockFormatter) Format(level Level, message string, ctx ...Context) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{level, message}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "format", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// Format indicates an expected call of Format
func (mr *MockFormatterRecorder) Format(level, message interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{level, message}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "format", reflect.TypeOf((*MockFormatter)(nil).Format), varargs...)
}
