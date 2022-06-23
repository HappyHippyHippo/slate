package gconfig

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gtrigger"
)

// MockTrigger is a mock of Trigger interface.
type MockTrigger struct {
	ctrl     *gomock.Controller
	recorder *MockTriggerRecorder
}

var _ gtrigger.Trigger = &MockTrigger{}

// MockTriggerRecorder is the mock recorder for MockTrigger.
type MockTriggerRecorder struct {
	mock *MockTrigger
}

// NewMockTrigger creates a new mock instance.
func NewMockTrigger(ctrl *gomock.Controller) *MockTrigger {
	mock := &MockTrigger{ctrl: ctrl}
	mock.recorder = &MockTriggerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrigger) EXPECT() *MockTriggerRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTrigger) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTriggerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTrigger)(nil).Close))
}

// Delay mocks base method.
func (m *MockTrigger) Delay() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Delay indicates an expected call of Timer.
func (mr *MockTriggerRecorder) Delay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockTrigger)(nil).Delay))
}
