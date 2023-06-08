package migration

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

// MockConfigSource is a mock instance of config.Source interface
type MockConfigSource struct {
	ctrl     *gomock.Controller
	recorder *MockConfigSourceRecorder
}

var _ config.Source = &MockConfigSource{}

// MockConfigSourceRecorder is the mock recorder for MockConfigSource
type MockConfigSourceRecorder struct {
	mock *MockConfigSource
}

// NewMockConfigSource creates a new mock instance
func NewMockConfigSource(ctrl *gomock.Controller) *MockConfigSource {
	mock := &MockConfigSource{ctrl: ctrl}
	mock.recorder = &MockConfigSourceRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigSource) EXPECT() *MockConfigSourceRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockConfigSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConfigSourceRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfigSource)(nil).Close))
}

// Has mocks base method
func (m *MockConfigSource) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockConfigSourceRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockConfigSource)(nil).Has), path)
}

// Get mocks base method
func (m *MockConfigSource) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
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
func (mr *MockConfigSourceRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigSource)(nil).Get), varargs...)
}
