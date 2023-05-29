package rdb

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// MockConnectionFactory is a mock of connectionCreator interface.
type MockConnectionFactory struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionFactoryRecorder
}

// MockConnectionFactoryRecorder is the mock recorder for MockConnectionFactory.
type MockConnectionFactoryRecorder struct {
	mock *MockConnectionFactory
}

// NewMockConnectionFactory creates a new mock instance.
func NewMockConnectionFactory(ctrl *gomock.Controller) *MockConnectionFactory {
	mock := &MockConnectionFactory{ctrl: ctrl}
	mock.recorder = &MockConnectionFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionFactory) EXPECT() *MockConnectionFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockConnectionFactory) Create(cfg *config.Partial, gormCfg *gorm.Config) (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg, gormCfg)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockConnectionFactoryRecorder) Create(cfg, gormCfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockConnectionFactory)(nil).Create), cfg, gormCfg)
}
