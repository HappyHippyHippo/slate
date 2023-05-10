package aggregate

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

//------------------------------------------------------------------------------
// Config
//------------------------------------------------------------------------------

// MockConfig is a mock instance of IConfig interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigRecorder
}

var _ config.IConfig = &MockConfig{}

// MockConfigRecorder is the mock recorder for MockConfig.
type MockConfigRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigRecorder {
	return m.recorder
}

// Bool mocks base method.
func (m *MockConfig) Bool(path string, def ...bool) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Bool", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bool indicates an expected call of Bool.
func (mr *MockConfigRecorder) Bool(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockConfig)(nil).Bool), varargs...)
}

// Partial mocks base method.
func (m *MockConfig) Config(path string, def ...config.Config) (config.IConfig, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Config", varargs...)
	ret0, _ := ret[0].(config.IConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partial indicates an expected call of Partial.
func (mr *MockConfigRecorder) Config(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockConfig)(nil).Config), varargs...)
}

// Entries mocks base method.
func (m *MockConfig) Entries() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entries")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Entries indicates an expected call of Entries.
func (mr *MockConfigRecorder) Entries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entries", reflect.TypeOf((*MockConfig)(nil).Entries))
}

// Float mocks base method.
func (m *MockConfig) Float(path string, def ...float64) (float64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float", varargs...)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float indicates an expected call of Float.
func (mr *MockConfigRecorder) Float(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float", reflect.TypeOf((*MockConfig)(nil).Float), varargs...)
}

// Get mocks base method.
func (m *MockConfig) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockConfigRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfig)(nil).Get), varargs...)
}

// Has mocks base method.
func (m *MockConfig) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockConfigRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockConfig)(nil).Has), path)
}

// Int mocks base method.
func (m *MockConfig) Int(path string, def ...int) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int indicates an expected call of Int.
func (mr *MockConfigRecorder) Int(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockConfig)(nil).Int), varargs...)
}

// List mocks base method.
func (m *MockConfig) List(path string, def ...[]interface{}) ([]interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockConfigRecorder) List(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockConfig)(nil).List), varargs...)
}

// Populate mocks base method.
func (m *MockConfig) Populate(path string, target interface{}, icase ...bool) (interface{}, error) {
	m.ctrl.T.Helper()
	m.ctrl.T.Helper()
	varargs := []interface{}{path, target}
	for _, a := range icase {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Populate", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Populate indicates an expected call of Partial.
func (mr *MockConfigRecorder) Populate(path, target interface{}, icase ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path, target}, icase...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Populate", reflect.TypeOf((*MockConfig)(nil).Populate), varargs...)
}

// String mocks base method.
func (m *MockConfig) String(path string, def ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "String", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// String indicates an expected call of String.
func (mr *MockConfigRecorder) String(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockConfig)(nil).String), varargs...)
}
