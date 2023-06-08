package migration

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockStorer is a mock instance of storer interface.
type MockStorer struct {
	ctrl     *gomock.Controller
	recorder *MockStorerRecorder
}

var _ storer = &MockStorer{}

// MockStorerRecorder is the mock recorder for MockStorer.
type MockStorerRecorder struct {
	mock *MockStorer
}

// NewMockStorer creates a new mock instance.
func NewMockStorer(ctrl *gomock.Controller) *MockStorer {
	mock := &MockStorer{ctrl: ctrl}
	mock.recorder = &MockStorerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorer) EXPECT() *MockStorerRecorder {
	return m.recorder
}

// Down mocks base method.
func (m *MockStorer) Down(last Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Down", last)
	ret0, _ := ret[0].(error)
	return ret0
}

// Down indicates an expected call of Down.
func (mr *MockStorerRecorder) Down(last interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Down", reflect.TypeOf((*MockStorer)(nil).Down), last)
}

// Last mocks base method.
func (m *MockStorer) Last() (Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Last indicates an expected call of Last.
func (mr *MockStorerRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockStorer)(nil).Last))
}

// Up mocks base method.
func (m *MockStorer) Up(version uint64) (Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Up", version)
	ret0, _ := ret[0].(Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Up indicates an expected call of Up.
func (mr *MockStorerRecorder) Up(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Up", reflect.TypeOf((*MockStorer)(nil).Up), version)
}
