package watchdog

import (
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"reflect"
)

// MockLogFormatterFactory is a mocked instance of ILogFormatterFactory interface.
type MockLogFormatterFactory struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterFactoryRecorder
}

// MockLogFormatterFactoryRecorder is the mock recorder for MockLogFormatterFactory.
type MockLogFormatterFactoryRecorder struct {
	mock *MockLogFormatterFactory
}

// NewMockLogFormatterFactory creates a new mock instance.
func NewMockLogFormatterFactory(ctrl *gomock.Controller) *MockLogFormatterFactory {
	mock := &MockLogFormatterFactory{ctrl: ctrl}
	mock.recorder = &MockLogFormatterFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFormatterFactory) EXPECT() *MockLogFormatterFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLogFormatterFactory) Create(cfg *config.Partial) (LogFormatter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(LogFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLogFormatterFactoryRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLogFormatterFactory)(nil).Create), cfg)
}

// Register mocks base method.
func (m *MockLogFormatterFactory) Register(strategy LogFormatterStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockLogFormatterFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockLogFormatterFactory)(nil).Register), strategy)
}
