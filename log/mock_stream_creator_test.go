package log

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

// MockStreamCreator is a mock instance of streamCreator interface.
type MockStreamCreator struct {
	ctrl     *gomock.Controller
	recorder *MockStreamCreatorRecorder
}

var _ streamCreator = &MockStreamCreator{}

// MockStreamCreatorRecorder is the mock recorder for MockStreamCreator.
type MockStreamCreatorRecorder struct {
	mock *MockStreamCreator
}

// NewMockStreamCreator creates a new mock instance.
func NewMockStreamCreator(ctrl *gomock.Controller) *MockStreamCreator {
	mock := &MockStreamCreator{ctrl: ctrl}
	mock.recorder = &MockStreamCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStreamCreator) EXPECT() *MockStreamCreatorRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStreamCreator) Create(cfg *config.Partial) (Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockStreamCreatorRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStreamCreator)(nil).Create), cfg)
}
