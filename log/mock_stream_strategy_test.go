package log

import (
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"reflect"
)

// MockStreamStrategy is a mock an instance of StreamStrategy interface
type MockStreamStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockStreamStrategyRecorder
}

var _ StreamStrategy = &MockStreamStrategy{}

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
func (m *MockStreamStrategy) Accept(cfg *config.Partial) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockStreamStrategyRecorder) Accept(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockStreamStrategy)(nil).Accept), cfg)
}

// Create mocks base method
func (m *MockStreamStrategy) Create(cfg *config.Partial) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of CreateConfig
func (mr *MockStreamStrategyRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStreamStrategy)(nil).Create), cfg)
}
