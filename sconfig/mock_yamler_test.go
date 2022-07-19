package sconfig

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockYamler is a mock of yamller interface
type MockYamler struct {
	ctrl     *gomock.Controller
	recorder *MockYamlerRecorder
}

var _ yamler = &MockYamler{}

// MockYamlerRecorder is the mock recorder for MockYamler
type MockYamlerRecorder struct {
	mock *MockYamler
}

// NewMockYamler creates a new mock instance
func NewMockYamler(ctrl *gomock.Controller) *MockYamler {
	mock := &MockYamler{ctrl: ctrl}
	mock.recorder = &MockYamlerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockYamler) EXPECT() *MockYamlerRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockYamler) Decode(partial interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", partial)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockYamlerRecorder) Decode(partial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockYamler)(nil).Decode), partial)
}
