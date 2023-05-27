package config

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockConfigurer is a mock of configurer interface.
type MockConfigurer struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurerRecorder
}

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

// AddSource mocks base method.
func (m *MockConfigurer) AddSource(id string, priority int, src Source) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSource", id, priority, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSource indicates an expected call of AddSource.
func (mr *MockConfigurerRecorder) AddSource(id, priority, src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSource", reflect.TypeOf((*MockConfigurer)(nil).AddSource), id, priority, src)
}

// Partial mocks base method.
func (m *MockConfigurer) Partial(path string, def ...Partial) (*Partial, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Partial", varargs...)
	ret0, _ := ret[0].(*Partial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partial indicates an expected call of Partial.
func (mr *MockConfigurerRecorder) Partial(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partial", reflect.TypeOf((*MockConfigurer)(nil).Partial), varargs...)
}
