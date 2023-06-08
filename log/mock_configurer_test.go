package log

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

// MockConfigurer is a mock instance of configurer interface.
type MockConfigurer struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurerRecorder
}

var _ configurer = &MockConfigurer{}

// MockConfigurerRecorder is the mock recorder for MockConfigurer.
type MockConfigurerRecorder struct {
	mock *MockConfigurer
}

// NewMockConfigurer creates a new mock instance.
func NewMockConfigurer(ctrl *gomock.Controller) *MockConfigurer {
	mock := &MockConfigurer{ctrl: ctrl}
	mock.recorder = &MockConfigurerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurer) EXPECT() *MockConfigurerRecorder {
	return m.recorder
}

// AddObserver mocks base method.
func (m *MockConfigurer) AddObserver(path string, callback config.Observer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddObserver", path, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddObserver indicates an expected call of AddObserver.
func (mr *MockConfigurerRecorder) AddObserver(path, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddObserver", reflect.TypeOf((*MockConfigurer)(nil).AddObserver), path, callback)
}

// Partial mocks base method.
func (m *MockConfigurer) Partial(path string, def ...config.Partial) (*config.Partial, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Partial", varargs...)
	ret0, _ := ret[0].(*config.Partial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partial indicates an expected call of Partial.
func (mr *MockConfigurerRecorder) Partial(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partial", reflect.TypeOf((*MockConfigurer)(nil).Partial), varargs...)
}
