package sconfig

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockSourceStrategy is a mock of SourceFactoryStrategy interface
type MockSourceStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockSourceStrategyRecorder
}

var _ ISourceStrategy = &MockSourceStrategy{}

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
func (m *MockSourceStrategy) Accept(sourceType string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", sourceType)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockSourceStrategyRecorder) Accept(sourceType interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sourceType}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockSourceStrategy)(nil).Accept), varargs...)
}

// AcceptFromConfig mocks base method
func (m *MockSourceStrategy) AcceptFromConfig(cfg IConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptFromConfig", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AcceptFromConfig indicates an expected call of AcceptFromConfig
func (mr *MockSourceStrategyRecorder) AcceptFromConfig(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptFromConfig", reflect.TypeOf((*MockSourceStrategy)(nil).AcceptFromConfig), cfg)
}

// Create mocks base method
func (m *MockSourceStrategy) Create(args ...interface{}) (ISource, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockSourceStrategyRecorder) Create(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceStrategy)(nil).Create), args...)
}

// CreateFromConfig mocks base method
func (m *MockSourceStrategy) CreateFromConfig(cfg IConfig) (ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFromConfig", cfg)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFromConfig indicates an expected call of CreateFromConfig
func (mr *MockSourceStrategyRecorder) CreateFromConfig(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFromConfig", reflect.TypeOf((*MockSourceStrategy)(nil).CreateFromConfig), cfg)
}
