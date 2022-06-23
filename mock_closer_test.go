package slate

import (
	"io"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockCloser is a mock of MockCloser interface
type MockCloser struct {
	ctrl     *gomock.Controller
	recorder *MockCloserRecorder
}

var _ io.Closer = &MockCloser{}

// MockCloserRecorder is the mock recorder for MockCloser
type MockCloserRecorder struct {
	mock *MockCloser
}

// NewMockCloser creates a new mock instance
func NewMockCloser(ctrl *gomock.Controller) *MockCloser {
	mock := &MockCloser{ctrl: ctrl}
	mock.recorder = &MockCloserRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloser) EXPECT() *MockCloserRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockCloserRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloser)(nil).Close))
}
