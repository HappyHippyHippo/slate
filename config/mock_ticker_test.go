package config

import (
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
)

// MockTicker is a mock instance of trigger.Recurring interface.
type MockTicker struct {
	ctrl     *gomock.Controller
	recorder *MockTickerRecorder
}

var _ ticker = &MockTicker{}

// MockTickerRecorder is the mock recorder for MockTicker.
type MockTickerRecorder struct {
	mock *MockTicker
}

// NewMockTicker creates a new mock instance.
func NewMockTicker(ctrl *gomock.Controller) *MockTicker {
	mock := &MockTicker{ctrl: ctrl}
	mock.recorder = &MockTickerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTicker) EXPECT() *MockTickerRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTicker) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTickerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTicker)(nil).Close))
}

// Delay mocks base method.
func (m *MockTicker) Delay() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Delay indicates an expected call of Timer.
func (mr *MockTickerRecorder) Delay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockTicker)(nil).Delay))
}
