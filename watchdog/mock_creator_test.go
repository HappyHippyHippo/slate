package watchdog

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockCreator is a mocked instance of creator interface.
type MockCreator struct {
	ctrl     *gomock.Controller
	recorder *MockCreatorRecorder
}

var _ creator = &MockCreator{}

// MockCreatorRecorder is the mock recorder for MockCreator.
type MockCreatorRecorder struct {
	mock *MockCreator
}

// NewMockCreator creates a new mock instance.
func NewMockCreator(ctrl *gomock.Controller) *MockCreator {
	mock := &MockCreator{ctrl: ctrl}
	mock.recorder = &MockCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreator) EXPECT() *MockCreatorRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCreator) Create(service string) (*Watchdog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", service)
	ret0, _ := ret[0].(*Watchdog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCreatorRecorder) Create(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCreator)(nil).Create), service)
}
