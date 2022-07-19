package sconfig

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockDecoderFactory is a mock of IDecoderFactory interface.
type MockDecoderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderFactoryRecorder
}

var _ IDecoderFactory = &MockDecoderFactory{}

// MockDecoderFactoryRecorder is the mock recorder for MockDecoderFactory.
type MockDecoderFactoryRecorder struct {
	mock *MockDecoderFactory
}

// NewMockDecoderFactory creates a new mock instance.
func NewMockDecoderFactory(ctrl *gomock.Controller) *MockDecoderFactory {
	mock := &MockDecoderFactory{ctrl: ctrl}
	mock.recorder = &MockDecoderFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoderFactory) EXPECT() *MockDecoderFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDecoderFactory) Create(format string, args ...interface{}) (IDecoder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(IDecoder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDecoderFactoryRecorder) Create(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDecoderFactory)(nil).Create), varargs...)
}

// Register mocks base method.
func (m *MockDecoderFactory) Register(strategy IDecoderStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockDecoderFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockDecoderFactory)(nil).Register), strategy)
}
