package log

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockFormatterFactory is a mock of IFormatterFactory interface.
type MockFormatterFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFormatterFactoryRecorder
}

var _ IFormatterFactory = &MockFormatterFactory{}

// MockFormatterFactoryRecorder is the mock recorder for MockFormatterFactory.
type MockFormatterFactoryRecorder struct {
	mock *MockFormatterFactory
}

// NewMockFormatterFactory creates a new mock instance.
func NewMockFormatterFactory(ctrl *gomock.Controller) *MockFormatterFactory {
	mock := &MockFormatterFactory{ctrl: ctrl}
	mock.recorder = &MockFormatterFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFormatterFactory) EXPECT() *MockFormatterFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFormatterFactory) Create(format string, args ...interface{}) (IFormatter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(IFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFormatterFactoryRecorder) Create(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFormatterFactory)(nil).Create), varargs...)
}

// Register mocks base method.
func (m *MockFormatterFactory) Register(strategy IFormatterStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockFormatterFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockFormatterFactory)(nil).Register), strategy)
}
