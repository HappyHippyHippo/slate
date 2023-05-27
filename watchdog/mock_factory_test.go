package watchdog

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockFactory is a mocked instance of factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryRecorder
}

// MockFactoryRecorder is the mock recorder for MockFactory.
type MockFactoryRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFactory) Create(service string) (*Watchdog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", service)
	ret0, _ := ret[0].(*Watchdog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFactoryRecorder) Create(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFactory)(nil).Create), service)
}
