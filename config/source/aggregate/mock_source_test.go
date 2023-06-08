package aggregate

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockSource is a mock instance of IConfig interface.
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceRecorder
}

// MockSourceRecorder is the mock recorder for MockSource.
type MockSourceRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance.
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSource) EXPECT() *MockSourceRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSource) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSourceRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSource)(nil).Get), varargs...)
}

// Has mocks base method.
func (m *MockSource) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockSourceRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockSource)(nil).Has), path)
}