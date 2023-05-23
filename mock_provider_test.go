package slate

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockProvider is a mocked instance of Provider interface.
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
func (m *MockProvider) Boot(arg0 *Container) error {
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
func (m *MockProvider) Register(arg0 *Container) error {
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
