package config

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockSourceStrategy is a mock instance of SourceFactoryStrategy interface
type MockSourceStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockSourceStrategyRecorder
}

var _ SourceStrategy = &MockSourceStrategy{}

// MockSourceStrategyRecorder is the mock recorder for MockSourceStrategy
type MockSourceStrategyRecorder struct {
	mock *MockSourceStrategy
}

// NewMockSourceStrategy creates a new mock instance
func NewMockSourceStrategy(ctrl *gomock.Controller) *MockSourceStrategy {
	mock := &MockSourceStrategy{ctrl: ctrl}
	mock.recorder = &MockSourceStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceStrategy) EXPECT() *MockSourceStrategyRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockSourceStrategy) Accept(partial Partial) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", partial)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockSourceStrategyRecorder) Accept(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockSourceStrategy)(nil).Accept), cfg)
}

// Create mocks base method
func (m *MockSourceStrategy) Create(partial Partial) (Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", partial)
	ret0, _ := ret[0].(Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockSourceStrategyRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceStrategy)(nil).Create), cfg)
}
