package log

import (
	"github.com/happyhippyhippo/slate/config"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockStreamFactory is a mock of streamCreator interface.
type MockStreamFactory struct {
	ctrl     *gomock.Controller
	recorder *MockStreamFactoryRecorder
}

// MockStreamFactoryRecorder is the mock recorder for MockStreamFactory.
type MockStreamFactoryRecorder struct {
	mock *MockStreamFactory
}

// NewMockStreamFactory creates a new mock instance.
func NewMockStreamFactory(ctrl *gomock.Controller) *MockStreamFactory {
	mock := &MockStreamFactory{ctrl: ctrl}
	mock.recorder = &MockStreamFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamFactory) EXPECT() *MockStreamFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStreamFactory) Create(cfg *config.Partial) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStreamFactoryRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStreamFactory)(nil).Create), cfg)
}
