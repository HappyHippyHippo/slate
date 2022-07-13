package slog

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/sconfig"
)

// MockStreamStrategy is a mock of LogStreamFactoryStrategy interface
type MockStreamStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockStreamStrategyRecorder
}

var _ IStreamStrategy = &MockStreamStrategy{}

// MockStreamStrategyRecorder is the mock recorder for MockStreamStrategy
type MockStreamStrategyRecorder struct {
	mock *MockStreamStrategy
}

// NewMockStreamStrategy creates a new mock instance
func NewMockStreamStrategy(ctrl *gomock.Controller) *MockStreamStrategy {
	mock := &MockStreamStrategy{ctrl: ctrl}
	mock.recorder = &MockStreamStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStreamStrategy) EXPECT() *MockStreamStrategyRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockStreamStrategy) Accept(sourceType string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", sourceType)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockStreamStrategyRecorder) Accept(sourceType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockStreamStrategy)(nil).Accept), sourceType)
}

// AcceptFromConfig mocks base method
func (m *MockStreamStrategy) AcceptFromConfig(cfg sconfig.IConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptFromConfig", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AcceptFromConfig indicates an expected call of AcceptFromConfig
func (mr *MockStreamStrategyRecorder) AcceptFromConfig(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptFromConfig", reflect.TypeOf((*MockStreamStrategy)(nil).AcceptFromConfig), cfg)
}

// Create mocks base method
func (m *MockStreamStrategy) Create(args ...interface{}) (IStream, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(IStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockStreamStrategyRecorder) Create(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStreamStrategy)(nil).Create), args...)
}

// CreateFromConfig mocks base method
func (m *MockStreamStrategy) CreateFromConfig(cfg sconfig.IConfig) (IStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFromConfig", cfg)
	ret0, _ := ret[0].(IStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFromConfig indicates an expected call of CreateConfig
func (mr *MockStreamStrategyRecorder) CreateFromConfig(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFromConfig", reflect.TypeOf((*MockStreamStrategy)(nil).CreateFromConfig), cfg)
}
