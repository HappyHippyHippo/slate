package slog

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockFormatterStrategy is a mock of LogFormatterFactoryStrategy interface
type MockFormatterStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockFormatterStrategyRecorder
}

var _ FormatterStrategy = &MockFormatterStrategy{}

// MockFormatterStrategyRecorder is the mock recorder for MockFormatterStrategy
type MockFormatterStrategyRecorder struct {
	mock *MockFormatterStrategy
}

// NewMockFormatterStrategy creates a new mock instance
func NewMockFormatterStrategy(ctrl *gomock.Controller) *MockFormatterStrategy {
	mock := &MockFormatterStrategy{ctrl: ctrl}
	mock.recorder = &MockFormatterStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFormatterStrategy) EXPECT() *MockFormatterStrategyRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockFormatterStrategy) Accept(format string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", format)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockFormatterStrategyRecorder) Accept(format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockFormatterStrategy)(nil).Accept), format)
}

// Create mocks base method
func (m *MockFormatterStrategy) Create(args ...interface{}) (Formatter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(Formatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockFormatterStrategyRecorder) Create(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFormatterStrategy)(nil).Create), args...)
}
