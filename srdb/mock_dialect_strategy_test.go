package srdb

import (
	"gorm.io/gorm"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/sconfig"
)

// MockDialectStrategy is a mock of DialectStrategy interface.
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
func (m *MockDialectStrategy) Accept(name string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockDialectStrategyRecorder) Accept(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockDialectStrategy)(nil).Accept), name)
}

// Get mocks base method.
func (m *MockDialectStrategy) Get(cfg sconfig.Config) (gorm.Dialector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", cfg)
	ret0, _ := ret[0].(gorm.Dialector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDialectStrategyRecorder) Get(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDialectStrategy)(nil).Get), cfg)
}
