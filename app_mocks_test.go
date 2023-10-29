package slate

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// ----------------------------------------------------------------------------
// ServiceProvider
// ----------------------------------------------------------------------------

// MockServiceProvider is a mocked instance of Register interface.
type MockServiceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockServiceProviderRecorder
}

var _ ServiceProvider = &MockServiceProvider{}

// MockServiceProviderRecorder is the mock recorder for MockServiceProvider.
type MockServiceProviderRecorder struct {
	mock *MockServiceProvider
}

// NewMockServiceProvider creates a new mock instance.
func NewMockServiceProvider(ctrl *gomock.Controller) *MockServiceProvider {
	mock := &MockServiceProvider{ctrl: ctrl}
	mock.recorder = &MockServiceProviderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceProvider) EXPECT() *MockServiceProviderRecorder {
	return m.recorder
}

// Boot mocks base method.
func (m *MockServiceProvider) Boot(arg0 *ServiceContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Boot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Boot indicates an expected call of Boot.
func (mr *MockServiceProviderRecorder) Boot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Boot", reflect.TypeOf((*MockServiceProvider)(nil).Boot), arg0)
}

// Provide mocks base method.
func (m *MockServiceProvider) Provide(arg0 *ServiceContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provide", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Provide indicates an expected call of Provide.
func (mr *MockServiceProviderRecorder) Provide(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provide", reflect.TypeOf((*MockServiceProvider)(nil).Provide), arg0)
}
