package rdb

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// MockDialectStrategy is a mock instance of IDialectStrategy interface.
type MockDialectStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockDialectStrategyRecorder
}

var _ DialectStrategy = &MockDialectStrategy{}

// MockDialectStrategyRecorder is the mock recorder for MockDialectStrategy.
type MockDialectStrategyRecorder struct {
	mock *MockDialectStrategy
}

// NewMockDialectStrategy creates a new mock instance.
func NewMockDialectStrategy(ctrl *gomock.Controller) *MockDialectStrategy {
	mock := &MockDialectStrategy{ctrl: ctrl}
	mock.recorder = &MockDialectStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialectStrategy) EXPECT() *MockDialectStrategyRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockDialectStrategy) Accept(cfg *config.Partial) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockDialectStrategyRecorder) Accept(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockDialectStrategy)(nil).Accept), cfg)
}

// Create mocks base method.
func (m *MockDialectStrategy) Create(cfg *config.Partial) (gorm.Dialector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(gorm.Dialector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDialectStrategyRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDialectStrategy)(nil).Create), cfg)
}
