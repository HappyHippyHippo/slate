package sconfig

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockSourceObservable is a mock of ConfigSourceObservable interface
type MockSourceObservable struct {
	ctrl     *gomock.Controller
	recorder *MockSourceObservableRecorder
}

var _ ISourceObservable = &MockSourceObservable{}

// MockSourceObservableRecorder is the mock recorder for MockSourceObservable
type MockSourceObservableRecorder struct {
	mock *MockSourceObservable
}

// NewMockSourceObservable creates a new mock instance
func NewMockSourceObservable(ctrl *gomock.Controller) *MockSourceObservable {
	mock := &MockSourceObservable{ctrl: ctrl}
	mock.recorder = &MockSourceObservableRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceObservable) EXPECT() *MockSourceObservableRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSourceObservable) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSourceObservableRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSourceObservable)(nil).Close))
}

// Has mocks base method
func (m *MockSourceObservable) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockSourceObservableRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockSourceObservable)(nil).Has), path)
}

// Get mocks base method
func (m *MockSourceObservable) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSourceObservableRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := []interface{}{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSource)(nil).Get), varargs...)
}

// Reload mocks base method
func (m *MockSourceObservable) Reload() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reload indicates an expected call of Reload
func (mr *MockSourceObservableRecorder) Reload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockSourceObservable)(nil).Reload))
}
