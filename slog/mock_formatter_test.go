package slog

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockFormatter is a mock of LogFormatter interface
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
func (m *MockFormatter) Format(level Level, message string, fields map[string]interface{}) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Format", level, message, fields)
	ret0, _ := ret[0].(string)
	return ret0
}

// Format indicates an expected call of Format
func (mr *MockFormatterRecorder) Format(level, message, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Format", reflect.TypeOf((*MockFormatter)(nil).Format), level, message, fields)
}
