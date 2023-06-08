package rdb

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// MockDialectCreator is a mock instance of dialectCreator interface.
type MockDialectCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDialectCreatorRecorder
}

var _ dialectCreator = &MockDialectCreator{}

// MockDialectCreatorRecorder is the mock recorder for MockDialectCreator.
type MockDialectCreatorRecorder struct {
	mock *MockDialectCreator
}

// NewMockDialectCreator creates a new mock instance.
func NewMockDialectCreator(ctrl *gomock.Controller) *MockDialectCreator {
	mock := &MockDialectCreator{ctrl: ctrl}
	mock.recorder = &MockDialectCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialectCreator) EXPECT() *MockDialectCreatorRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDialectCreator) Create(cfg config.Partial) (gorm.Dialector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(gorm.Dialector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDialectCreatorRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDialectCreator)(nil).Create), cfg)
}
