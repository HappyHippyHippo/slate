package config

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockSourceFactory is a mock of sourceFactory interface.
type MockSourceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSourceFactoryRecorder
}

// MockSourceFactoryRecorder is the mock recorder for MockSourceFactory.
type MockSourceFactoryRecorder struct {
	mock *MockSourceFactory
}

// NewMockSourceFactory creates a new mock instance.
func NewMockSourceFactory(ctrl *gomock.Controller) *MockSourceFactory {
	mock := &MockSourceFactory{ctrl: ctrl}
	mock.recorder = &MockSourceFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceFactory) EXPECT() *MockSourceFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSourceFactory) Create(partial *Partial) (Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", partial)
	ret0, _ := ret[0].(Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSourceFactoryRecorder) Create(partial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceFactory)(nil).Create), partial)
}
