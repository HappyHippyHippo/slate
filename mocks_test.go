package slate

import (
	"io"
	"reflect"

	"github.com/golang/mock/gomock"
)

//------------------------------------------------------------------------------
// Closer
//------------------------------------------------------------------------------

// MockCloser is a mock instance of MockCloser interface
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

//------------------------------------------------------------------------------
// Provider
//------------------------------------------------------------------------------

// MockProvider is a mocked instance of IProvider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderRecorder
}

// MockProviderRecorder is the mock recorder for MockProvider.
type MockProviderRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderRecorder {
	return m.recorder
}

// Boot mocks base method.
func (m *MockProvider) Boot(arg0 IContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Boot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Boot indicates an expected call of Boot.
func (mr *MockProviderRecorder) Boot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Boot", reflect.TypeOf((*MockProvider)(nil).Boot), arg0)
}

// Register mocks base method.
func (m *MockProvider) Register(arg0 IContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockProviderRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockProvider)(nil).Register), arg0)
}
