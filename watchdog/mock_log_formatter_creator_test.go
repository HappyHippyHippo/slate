package watchdog

import (
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"reflect"
)

// MockLogFormatterCreator is a mocked instance of ILogFormatterFactory interface.
type MockLogFormatterCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterCreatorRecorder
}

var _ logFormatterCreator = &MockLogFormatterCreator{}

// MockLogFormatterCreatorRecorder is the mock recorder for MockLogFormatterCreator.
type MockLogFormatterCreatorRecorder struct {
	mock *MockLogFormatterCreator
}

// NewMockLogFormatterCreator creates a new mock instance.
func NewMockLogFormatterCreator(ctrl *gomock.Controller) *MockLogFormatterCreator {
	mock := &MockLogFormatterCreator{ctrl: ctrl}
	mock.recorder = &MockLogFormatterCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFormatterCreator) EXPECT() *MockLogFormatterCreatorRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLogFormatterCreator) Create(cfg *config.Partial) (LogFormatter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(LogFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLogFormatterCreatorRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLogFormatterCreator)(nil).Create), cfg)
}

// Register mocks base method.
func (m *MockLogFormatterCreator) Register(strategy LogFormatterStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockLogFormatterCreatorRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockLogFormatterCreator)(nil).Register), strategy)
}
