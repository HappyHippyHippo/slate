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
// Service Provider
//------------------------------------------------------------------------------

// MockServiceProvider is a mock instance of Provider interface
type MockServiceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProviderRecorder
}

var _ IProvider = &MockServiceProvider{}

// MockServiceProviderRecorder is the mock recorder for MockServiceProvider
type MockServiceProviderRecorder struct {
	mock *MockServiceProvider
}

// NewMockServiceProvider creates a new mock instance
func NewMockServiceProvider(ctrl *gomock.Controller) *MockServiceProvider {
	mock := &MockServiceProvider{ctrl: ctrl}
	mock.recorder = &MockServiceProviderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceProvider) EXPECT() *MockServiceProviderRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockServiceProvider) Register(arg0 IContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockServiceProviderRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockServiceProvider)(nil).Register), arg0)
}

// Boot mocks base method
func (m *MockServiceProvider) Boot(arg0 IContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Boot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Boot indicates an expected call of Boot
func (mr *MockServiceProviderRecorder) Boot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Boot", reflect.TypeOf((*MockServiceProvider)(nil).Boot), arg0)
}
