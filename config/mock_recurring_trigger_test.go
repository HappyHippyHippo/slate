package config

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
)

// MockRecurringTrigger is a mock instance of trigger.Recurring interface.
type MockRecurringTrigger struct {
	ctrl     *gomock.Controller
	recorder *MockRecurringTriggerRecorder
}

var _ ticker = &MockRecurringTrigger{}

// MockRecurringTriggerRecorder is the mock recorder for MockRecurringTrigger.
type MockRecurringTriggerRecorder struct {
	mock *MockRecurringTrigger
}

// NewMockRecurringTrigger creates a new mock instance.
func NewMockRecurringTrigger(ctrl *gomock.Controller) *MockRecurringTrigger {
	mock := &MockRecurringTrigger{ctrl: ctrl}
	mock.recorder = &MockRecurringTriggerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecurringTrigger) EXPECT() *MockRecurringTriggerRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRecurringTrigger) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRecurringTriggerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRecurringTrigger)(nil).Close))
}

// Delay mocks base method.
func (m *MockRecurringTrigger) Delay() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Delay indicates an expected call of Timer.
func (mr *MockRecurringTriggerRecorder) Delay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockRecurringTrigger)(nil).Delay))
}
