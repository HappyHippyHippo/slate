package sconfig

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockSourceFactory is a mock of ISourceFactory interface.
type MockSourceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSourceFactoryRecorder
}

var _ ISourceFactory = &MockSourceFactory{}

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
func (m *MockSourceFactory) Create(sourceType string, args ...interface{}) (ISource, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{sourceType}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSourceFactoryRecorder) Create(sourceType interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{sourceType}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceFactory)(nil).Create), varargs...)
}

// CreateFromConfig mocks base method.
func (m *MockSourceFactory) CreateFromConfig(cfg IConfig) (ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFromConfig", cfg)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFromConfig indicates an expected call of CreateFromConfig.
func (mr *MockSourceFactoryRecorder) CreateFromConfig(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFromConfig", reflect.TypeOf((*MockSourceFactory)(nil).CreateFromConfig), cfg)
}

// Register mocks base method.
func (m *MockSourceFactory) Register(strategy ISourceStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockSourceFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSourceFactory)(nil).Register), strategy)
}
