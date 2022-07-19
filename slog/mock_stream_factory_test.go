package slog

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/sconfig"
)

// MockStreamFactory is a mock of IStreamFactory interface.
type MockStreamFactory struct {
	ctrl     *gomock.Controller
	recorder *MockStreamFactoryRecorder
}

var _ IStreamFactory = &MockStreamFactory{}

// MockStreamFactoryRecorder is the mock recorder for MockStreamFactory.
type MockStreamFactoryRecorder struct {
	mock *MockStreamFactory
}

// NewMockStreamFactory creates a new mock instance.
func NewMockStreamFactory(ctrl *gomock.Controller) *MockStreamFactory {
	mock := &MockStreamFactory{ctrl: ctrl}
	mock.recorder = &MockStreamFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamFactory) EXPECT() *MockStreamFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStreamFactory) Create(streamType string, args ...interface{}) (IStream, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{streamType}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(IStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStreamFactoryRecorder) Create(streamType interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{streamType}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStreamFactory)(nil).Create), varargs...)
}

// CreateFromConfig mocks base method.
func (m *MockStreamFactory) CreateFromConfig(cfg sconfig.IConfig) (IStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFromConfig", cfg)
	ret0, _ := ret[0].(IStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFromConfig indicates an expected call of CreateFromConfig.
func (mr *MockStreamFactoryRecorder) CreateFromConfig(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFromConfig", reflect.TypeOf((*MockStreamFactory)(nil).CreateFromConfig), cfg)
}

// Register mocks base method.
func (m *MockStreamFactory) Register(strategy IStreamStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockStreamFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockStreamFactory)(nil).Register), strategy)
}
