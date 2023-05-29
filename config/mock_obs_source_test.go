package config

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockObsSource is a mock instance of ObsSource interface
type MockObsSource struct {
	ctrl     *gomock.Controller
	recorder *MockObsSourceRecorder
}

var _ ObsSource = &MockObsSource{}

// MockObsSourceRecorder is the mock recorder for MockObsSource
type MockObsSourceRecorder struct {
	mock *MockObsSource
}

// NewMockObsSource creates a new mock instance
func NewMockObsSource(ctrl *gomock.Controller) *MockObsSource {
	mock := &MockObsSource{ctrl: ctrl}
	mock.recorder = &MockObsSourceRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObsSource) EXPECT() *MockObsSourceRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockObsSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockObsSourceRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockObsSource)(nil).Close))
}

// Has mocks base method
func (m *MockObsSource) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockObsSourceRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockObsSource)(nil).Has), path)
}

// Get mocks base method
func (m *MockObsSource) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockObsSourceRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSource)(nil).Get), varargs...)
}

// Reload mocks base method
func (m *MockObsSource) Reload() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reload indicates an expected call of Reload
func (mr *MockObsSourceRecorder) Reload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockObsSource)(nil).Reload))
}
