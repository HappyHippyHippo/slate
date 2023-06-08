package rdb

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// MockConnectionCreator is a mock instance of connectionCreator interface.
type MockConnectionCreator struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionCreatorRecorder
}

var _ connectionCreator = &MockConnectionCreator{}

// MockConnectionCreatorRecorder is the mock recorder for MockConnectionCreator.
type MockConnectionCreatorRecorder struct {
	mock *MockConnectionCreator
}

// NewMockConnectionCreator creates a new mock instance.
func NewMockConnectionCreator(ctrl *gomock.Controller) *MockConnectionCreator {
	mock := &MockConnectionCreator{ctrl: ctrl}
	mock.recorder = &MockConnectionCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionCreator) EXPECT() *MockConnectionCreatorRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockConnectionCreator) Create(cfg config.Partial, gormCfg *gorm.Config) (*gorm.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg, gormCfg)
	ret0, _ := ret[0].(*gorm.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockConnectionCreatorRecorder) Create(cfg, gormCfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockConnectionCreator)(nil).Create), cfg, gormCfg)
}
