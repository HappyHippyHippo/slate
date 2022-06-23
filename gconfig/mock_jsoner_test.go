package gconfig

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockJsoner is a mock of jsoner interface
type MockJsoner struct {
	ctrl     *gomock.Controller
	recorder *MockJsonerRecorder
}

var _ jsoner = &MockJsoner{}

// MockJsonerRecorder is the mock recorder for MockJsoner
type MockJsonerRecorder struct {
	mock *MockJsoner
}

// NewMockJsoner creates a new mock instance
func NewMockJsoner(ctrl *gomock.Controller) *MockJsoner {
	mock := &MockJsoner{ctrl: ctrl}
	mock.recorder = &MockJsonerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJsoner) EXPECT() *MockJsonerRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockJsoner) Decode(partial interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", partial)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockJsonerRecorder) Decode(partial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockJsoner)(nil).Decode), partial)
}
