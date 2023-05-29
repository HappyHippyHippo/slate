package watchdog

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockLogAdapter is a mock an instance of LogAdapter interface.
type MockLogAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockLogAdapterRecorder
}

// MockLogAdapterRecorder is the mock recorder for MockLogAdapter.
type MockLogAdapterRecorder struct {
	mock *MockLogAdapter
}

// NewMockLogAdapter creates a new mock instance.
func NewMockLogAdapter(ctrl *gomock.Controller) *MockLogAdapter {
	mock := &MockLogAdapter{ctrl: ctrl}
	mock.recorder = &MockLogAdapterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogAdapter) EXPECT() *MockLogAdapterRecorder {
	return m.recorder
}

// Done mocks base method.
func (m *MockLogAdapter) Done() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(error)
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockLogAdapterRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockLogAdapter)(nil).Done))
}

// Error mocks base method.
func (m *MockLogAdapter) Error(e error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockLogAdapterRecorder) Error(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogAdapter)(nil).Error), e)
}

// Start mocks base method.
func (m *MockLogAdapter) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockLogAdapterRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockLogAdapter)(nil).Start))
}
