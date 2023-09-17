package rest

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockEndpointRegister is a mock of EndpointRegister interface.
type MockEndpointRegister struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointRegisterRecorder
}

var _ EndpointRegister = &MockEndpointRegister{}

// MockEndpointRegisterRecorder is the mock recorder for MockEndpointRegister.
type MockEndpointRegisterRecorder struct {
	mock *MockEndpointRegister
}

// NewMockEndpointRegister creates a new mock instance.
func NewMockEndpointRegister(ctrl *gomock.Controller) *MockEndpointRegister {
	mock := &MockEndpointRegister{ctrl: ctrl}
	mock.recorder = &MockEndpointRegisterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndpointRegister) EXPECT() *MockEndpointRegisterRecorder {
	return m.recorder
}

// Reg mocks base method.
func (m *MockEndpointRegister) Reg(engine Engine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reg", engine)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reg indicates an expected call of Reg.
func (mr *MockEndpointRegisterRecorder) Reg(engine interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reg", reflect.TypeOf((*MockEndpointRegister)(nil).Reg), engine)
}
