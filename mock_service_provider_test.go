package slate

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockServiceProvider is a mock of Provider interface
type MockServiceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockserviceProviderRecorder
}

var _ IServiceProvider = &MockServiceProvider{}

// MockserviceProviderRecorder is the mock recorder for MockServiceProvider
type MockserviceProviderRecorder struct {
	mock *MockServiceProvider
}

// NewMockServiceProvider creates a new mock instance
func NewMockServiceProvider(ctrl *gomock.Controller) *MockServiceProvider {
	mock := &MockServiceProvider{ctrl: ctrl}
	mock.recorder = &MockserviceProviderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceProvider) EXPECT() *MockserviceProviderRecorder {
	return m.recorder
}

// Register mocks base method
func (m *MockServiceProvider) Register(arg0 ServiceContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockserviceProviderRecorder) Register(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockServiceProvider)(nil).Register), arg0)
}

// Boot mocks base method
func (m *MockServiceProvider) Boot(arg0 ServiceContainer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Boot", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Boot indicates an expected call of Boot
func (mr *MockserviceProviderRecorder) Boot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Boot", reflect.TypeOf((*MockServiceProvider)(nil).Boot), arg0)
}
