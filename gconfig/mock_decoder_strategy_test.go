package gconfig

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockDecoderStrategy is a mock of DecoderStrategy interface
type MockDecoderStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderStrategyRecorder
}

var _ DecoderStrategy = &MockDecoderStrategy{}

// MockDecoderStrategyRecorder is the mock recorder for MockDecoderStrategy
type MockDecoderStrategyRecorder struct {
	mock *MockDecoderStrategy
}

// NewMockDecoderStrategy creates a new mock instance
func NewMockDecoderStrategy(ctrl *gomock.Controller) *MockDecoderStrategy {
	mock := &MockDecoderStrategy{ctrl: ctrl}
	mock.recorder = &MockDecoderStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDecoderStrategy) EXPECT() *MockDecoderStrategyRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockDecoderStrategy) Accept(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockDecoderStrategyRecorder) Accept(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockDecoderStrategy)(nil).Accept), varargs...)
}

// Create mocks base method
func (m *MockDecoderStrategy) Create(args ...interface{}) (Decoder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(Decoder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDecoderStrategyRecorder) Create(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDecoderStrategy)(nil).Create), args...)
}
