package rdb

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// MockDialectFactory is a mock of IDialectFactory interface.
type MockDialectFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDialectFactoryRecorder
}

var _ IDialectFactory = &MockDialectFactory{}

// MockDialectFactoryRecorder is the mock recorder for MockDialectFactory.
type MockDialectFactoryRecorder struct {
	mock *MockDialectFactory
}

// NewMockDialectFactory creates a new mock instance.
func NewMockDialectFactory(ctrl *gomock.Controller) *MockDialectFactory {
	mock := &MockDialectFactory{ctrl: ctrl}
	mock.recorder = &MockDialectFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialectFactory) EXPECT() *MockDialectFactoryRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDialectFactory) Get(cfg config.IConfig) (gorm.Dialector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", cfg)
	ret0, _ := ret[0].(gorm.Dialector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDialectFactoryRecorder) Get(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDialectFactory)(nil).Get), cfg)
}

// Register mocks base method.
func (m *MockDialectFactory) Register(strategy IDialectStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockDialectFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockDialectFactory)(nil).Register), strategy)
}
