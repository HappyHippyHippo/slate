package slog

import (
	"github.com/happyhippyhippo/slate/sconfig"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockConfigManager is a mock of IManager interface.
type MockConfigManager struct {
	ctrl     *gomock.Controller
	recorder *MockConfigManagerRecorder
}

var _ sconfig.IManager = &MockConfigManager{}

// MockConfigManagerRecorder is the mock recorder for MockConfigManager.
type MockConfigManagerRecorder struct {
	mock *MockConfigManager
}

// NewMockConfigManager creates a new mock instance.
func NewMockConfigManager(ctrl *gomock.Controller) *MockConfigManager {
	mock := &MockConfigManager{ctrl: ctrl}
	mock.recorder = &MockConfigManagerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigManager) EXPECT() *MockConfigManagerRecorder {
	return m.recorder
}

// AddObserver mocks base method.
func (m *MockConfigManager) AddObserver(path string, callback sconfig.IObserver) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddObserver", path, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddObserver indicates an expected call of AddObserver.
func (mr *MockConfigManagerRecorder) AddObserver(path, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddObserver", reflect.TypeOf((*MockConfigManager)(nil).AddObserver), path, callback)
}

// AddSource mocks base method.
func (m *MockConfigManager) AddSource(id string, priority int, src sconfig.ISource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSource", id, priority, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSource indicates an expected call of AddSource.
func (mr *MockConfigManagerRecorder) AddSource(id, priority, src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSource", reflect.TypeOf((*MockConfigManager)(nil).AddSource), id, priority, src)
}

// Bool mocks base method.
func (m *MockConfigManager) Bool(path string, def ...bool) (bool, error) {
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
func (mr *MockConfigManagerRecorder) Bool(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockConfigManager)(nil).Bool), varargs...)
}

// Close mocks base method.
func (m *MockConfigManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConfigManagerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfigManager)(nil).Close))
}

// Float mocks base method.
func (m *MockConfigManager) Float(path string, def ...float64) (float64, error) {
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
func (mr *MockConfigManagerRecorder) Float(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float", reflect.TypeOf((*MockConfigManager)(nil).Float), varargs...)
}

// Get mocks base method.
func (m *MockConfigManager) Get(path string, def ...interface{}) (interface{}, error) {
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
func (mr *MockConfigManagerRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigManager)(nil).Get), varargs...)
}

// Has mocks base method.
func (m *MockConfigManager) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockConfigManagerRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockConfigManager)(nil).Has), path)
}

// HasObserver mocks base method.
func (m *MockConfigManager) HasObserver(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasObserver", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasObserver indicates an expected call of HasObserver.
func (mr *MockConfigManagerRecorder) HasObserver(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasObserver", reflect.TypeOf((*MockConfigManager)(nil).HasObserver), path)
}

// HasSource mocks base method.
func (m *MockConfigManager) HasSource(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSource", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSource indicates an expected call of HasSource.
func (mr *MockConfigManagerRecorder) HasSource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSource", reflect.TypeOf((*MockConfigManager)(nil).HasSource), id)
}

// Int mocks base method.
func (m *MockConfigManager) Int(path string, def ...int) (int, error) {
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
func (mr *MockConfigManagerRecorder) Int(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockConfigManager)(nil).Int), varargs...)
}

// List mocks base method.
func (m *MockConfigManager) List(path string, def ...[]interface{}) ([]interface{}, error) {
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
func (mr *MockConfigManagerRecorder) List(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockConfigManager)(nil).List), varargs...)
}

// Partial mocks base method.
func (m *MockConfigManager) Partial(path string, def ...sconfig.Partial) (sconfig.Partial, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Partial", varargs...)
	ret0, _ := ret[0].(sconfig.Partial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partial indicates an expected call of Partial.
func (mr *MockConfigManagerRecorder) Partial(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partial", reflect.TypeOf((*MockConfigManager)(nil).Partial), varargs...)
}

// RemoveAllSources mocks base method.
func (m *MockConfigManager) RemoveAllSources() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllSources")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllSources indicates an expected call of RemoveAllSources.
func (mr *MockConfigManagerRecorder) RemoveAllSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllSources", reflect.TypeOf((*MockConfigManager)(nil).RemoveAllSources))
}

// RemoveObserver mocks base method.
func (m *MockConfigManager) RemoveObserver(path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveObserver", path)
}

// RemoveObserver indicates an expected call of RemoveObserver.
func (mr *MockConfigManagerRecorder) RemoveObserver(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveObserver", reflect.TypeOf((*MockConfigManager)(nil).RemoveObserver), path)
}

// RemoveSource mocks base method.
func (m *MockConfigManager) RemoveSource(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSource", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSource indicates an expected call of RemoveSource.
func (mr *MockConfigManagerRecorder) RemoveSource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSource", reflect.TypeOf((*MockConfigManager)(nil).RemoveSource), id)
}

// Source mocks base method.
func (m *MockConfigManager) Source(id string) (sconfig.ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Source", id)
	ret0, _ := ret[0].(sconfig.ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Source indicates an expected call of Source.
func (mr *MockConfigManagerRecorder) Source(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Source", reflect.TypeOf((*MockConfigManager)(nil).Source), id)
}

// SourcePriority mocks base method.
func (m *MockConfigManager) SourcePriority(id string, priority int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourcePriority", id, priority)
	ret0, _ := ret[0].(error)
	return ret0
}

// SourcePriority indicates an expected call of SourcePriority.
func (mr *MockConfigManagerRecorder) SourcePriority(id, priority interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourcePriority", reflect.TypeOf((*MockConfigManager)(nil).SourcePriority), id, priority)
}

// String mocks base method.
func (m *MockConfigManager) String(path string, def ...string) (string, error) {
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
func (mr *MockConfigManagerRecorder) String(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockConfigManager)(nil).String), varargs...)
}
