package rdb

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// MockDialectFactory is a mock of dialectCreator interface.
type MockDialectFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDialectFactoryRecorder
}

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

// Create mocks base method.
func (m *MockDialectFactory) Create(cfg *config.Partial) (gorm.Dialector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(gorm.Dialector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDialectFactoryRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDialectFactory)(nil).Create), cfg)
}
