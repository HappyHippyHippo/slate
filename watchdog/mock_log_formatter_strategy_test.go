package watchdog

import (
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate/config"
)

// MockLogFormatterStrategy is a mock instance of LogFormatterStrategy interface.
type MockLogFormatterStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterStrategyRecorder
}

var _ LogFormatterStrategy = &MockLogFormatterStrategy{}

// MockLogFormatterStrategyRecorder is the mock recorder for MockLogFormatterStrategy.
type MockLogFormatterStrategyRecorder struct {
	mock *MockLogFormatterStrategy
}

// NewMockLogFormatterStrategy creates a new mock instance.
func NewMockLogFormatterStrategy(ctrl *gomock.Controller) *MockLogFormatterStrategy {
	mock := &MockLogFormatterStrategy{ctrl: ctrl}
	mock.recorder = &MockLogFormatterStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFormatterStrategy) EXPECT() *MockLogFormatterStrategyRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockLogFormatterStrategy) Accept(cfg *config.Partial) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockLogFormatterStrategyRecorder) Accept(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockLogFormatterStrategy)(nil).Accept), cfg)
}

// Create mocks base method.
func (m *MockLogFormatterStrategy) Create(cfg *config.Partial) (LogFormatter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(LogFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLogFormatterStrategyRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLogFormatterStrategy)(nil).Create), cfg)
}
