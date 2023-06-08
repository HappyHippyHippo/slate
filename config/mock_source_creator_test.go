package config

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockSourceCreator is a mock instance of sourceCreator interface.
type MockSourceCreator struct {
	ctrl     *gomock.Controller
	recorder *MockSourceCreatorRecorder
}

var _ sourceCreator = &MockSourceCreator{}

// MockSourceCreatorRecorder is the mock recorder for MockSourceCreator.
type MockSourceCreatorRecorder struct {
	mock *MockSourceCreator
}

// NewMockSourceCreator creates a new mock instance.
func NewMockSourceCreator(ctrl *gomock.Controller) *MockSourceCreator {
	mock := &MockSourceCreator{ctrl: ctrl}
	mock.recorder = &MockSourceCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceCreator) EXPECT() *MockSourceCreatorRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSourceCreator) Create(partial Partial) (Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", partial)
	ret0, _ := ret[0].(Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSourceCreatorRecorder) Create(partial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceCreator)(nil).Create), partial)
}
