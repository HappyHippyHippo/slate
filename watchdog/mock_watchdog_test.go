package watchdog

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockWatchdog is a mocked instance of IWatchdog interface.
type MockWatchdog struct {
	ctrl     *gomock.Controller
	recorder *MockWatchdogRecorder
}

// MockWatchdogRecorder is the mock recorder for MockWatchdog.
type MockWatchdogRecorder struct {
	mock *MockWatchdog
}

// NewMockWatchdog creates a new mock instance.
func NewMockWatchdog(ctrl *gomock.Controller) *MockWatchdog {
	mock := &MockWatchdog{ctrl: ctrl}
	mock.recorder = &MockWatchdogRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchdog) EXPECT() *MockWatchdogRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockWatchdog) Run(process Processor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", process)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockWatchdogRecorder) Run(process interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWatchdog)(nil).Run), process)
}
