package rest

import (
	"html/template"
	"net"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
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

// Config mocks base method.
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

// Config indicates an expected call of Config.
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

//------------------------------------------------------------------------------
// Config Manager
//------------------------------------------------------------------------------

// MockConfigManager is a mock an instance of IManager interface.
type MockConfigManager struct {
	ctrl     *gomock.Controller
	recorder *MockConfigManagerRecorder
}

var _ config.IManager = &MockConfigManager{}

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
func (m *MockConfigManager) AddObserver(path string, callback config.IObserver) error {
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
func (m *MockConfigManager) AddSource(id string, priority int, src config.ISource) error {
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

// Config mocks base method.
func (m *MockConfigManager) Config(path string, def ...config.Config) (config.IConfig, error) {
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

// Config indicates an expected call of Config.
func (mr *MockConfigManagerRecorder) Config(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockConfigManager)(nil).Config), varargs...)
}

// Entries mocks base method.
func (m *MockConfigManager) Entries() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entries")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Entries indicates an expected call of Entries.
func (mr *MockConfigManagerRecorder) Entries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entries", reflect.TypeOf((*MockConfigManager)(nil).Entries))
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

// Populate mocks base method.
func (m *MockConfigManager) Populate(path string, target interface{}, icase ...bool) (interface{}, error) {
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
func (mr *MockConfigManagerRecorder) Populate(path, target interface{}, icase ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path, target}, icase...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Populate", reflect.TypeOf((*MockConfigManager)(nil).Populate), varargs...)
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
func (m *MockConfigManager) Source(id string) (config.ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Source", id)
	ret0, _ := ret[0].(config.ISource)
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

//------------------------------------------------------------------------------
// Log
//------------------------------------------------------------------------------

// MockLog is a mock an instance of ILogger interface.
type MockLog struct {
	ctrl     *gomock.Controller
	recorder *MockLogRecorder
}

var _ log.ILog = &MockLog{}

// MockLogRecorder is the mock recorder for MockLog.
type MockLogRecorder struct {
	mock *MockLog
}

// NewMockLog creates a new mock instance.
func NewMockLog(ctrl *gomock.Controller) *MockLog {
	mock := &MockLog{ctrl: ctrl}
	mock.recorder = &MockLogRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLog) EXPECT() *MockLogRecorder {
	return m.recorder
}

// AddStream mocks base method.
func (m *MockLog) AddStream(id string, stream log.IStream) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStream", id, stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddStream indicates an expected call of AddStream.
func (mr *MockLogRecorder) AddStream(id, stream interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStream", reflect.TypeOf((*MockLog)(nil).AddStream), id, stream)
}

// Broadcast mocks base method.
func (m *MockLog) Broadcast(level log.Level, msg string, ctx ...log.Context) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{level, msg}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Broadcast", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockLogRecorder) Broadcast(level, msg interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{level, msg}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockLog)(nil).Broadcast), varargs...)
}

// Close mocks base method.
func (m *MockLog) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockLogRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLog)(nil).Close))
}

// HasStream mocks base method.
func (m *MockLog) HasStream(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasStream", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasStream indicates an expected call of HasStream.
func (mr *MockLogRecorder) HasStream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasStream", reflect.TypeOf((*MockLog)(nil).HasStream), id)
}

// ListStreams mocks base method.
func (m *MockLog) ListStreams() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStreams")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListStreams indicates an expected call of ListStreams.
func (mr *MockLogRecorder) ListStreams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreams", reflect.TypeOf((*MockLog)(nil).ListStreams))
}

// RemoveAllStreams mocks base method.
func (m *MockLog) RemoveAllStreams() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveAllStreams")
}

// RemoveAllStreams indicates an expected call of RemoveAllStreams.
func (mr *MockLogRecorder) RemoveAllStreams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllStreams", reflect.TypeOf((*MockLog)(nil).RemoveAllStreams))
}

// RemoveStream mocks base method.
func (m *MockLog) RemoveStream(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveStream", id)
}

// RemoveStream indicates an expected call of RemoveStream.
func (mr *MockLogRecorder) RemoveStream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStream", reflect.TypeOf((*MockLog)(nil).RemoveStream), id)
}

// Signal mocks base method.
func (m *MockLog) Signal(channel string, level log.Level, msg string, ctx ...log.Context) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{channel, level, msg}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Signal", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal.
func (mr *MockLogRecorder) Signal(channel, level, msg interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{channel, level, msg}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockLog)(nil).Signal), varargs...)
}

// Stream mocks base method.
func (m *MockLog) Stream(id string) (log.IStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", id)
	ret0, _ := ret[0].(log.IStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stream indicates an expected call of Stream.
func (mr *MockLogRecorder) Stream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockLog)(nil).Stream), id)
}

//------------------------------------------------------------------------------
// Engine
//------------------------------------------------------------------------------

// MockEngine is a mock instance of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineRecorder
}

// MockEngineRecorder is the mock recorder for MockEngine.
type MockEngineRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineRecorder {
	return m.recorder
}

// Any mocks base method.
func (m *MockEngine) Any(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Any", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Any indicates an expected call of Any.
func (mr *MockEngineRecorder) Any(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Any", reflect.TypeOf((*MockEngine)(nil).Any), varargs...)
}

// DELETE mocks base method.
func (m *MockEngine) DELETE(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DELETE", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// DELETE indicates an expected call of DELETE.
func (mr *MockEngineRecorder) DELETE(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DELETE", reflect.TypeOf((*MockEngine)(nil).DELETE), varargs...)
}

// Delims mocks base method.
func (m *MockEngine) Delims(left, right string) *gin.Engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delims", left, right)
	ret0, _ := ret[0].(*gin.Engine)
	return ret0
}

// Delims indicates an expected call of Delims.
func (mr *MockEngineRecorder) Delims(left, right interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delims", reflect.TypeOf((*MockEngine)(nil).Delims), left, right)
}

// GET mocks base method.
func (m *MockEngine) GET(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GET", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// GET indicates an expected call of GET.
func (mr *MockEngineRecorder) GET(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GET", reflect.TypeOf((*MockEngine)(nil).GET), varargs...)
}

// HEAD mocks base method.
func (m *MockEngine) HEAD(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HEAD", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// HEAD indicates an expected call of HEAD.
func (mr *MockEngineRecorder) HEAD(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HEAD", reflect.TypeOf((*MockEngine)(nil).HEAD), varargs...)
}

// Handle mocks base method.
func (m *MockEngine) Handle(arg0, arg1 string, arg2 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Handle", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockEngineRecorder) Handle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockEngine)(nil).Handle), varargs...)
}

// HandleContext mocks base method.
func (m *MockEngine) HandleContext(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleContext", c)
}

// HandleContext indicates an expected call of HandleContext.
func (mr *MockEngineRecorder) HandleContext(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleContext", reflect.TypeOf((*MockEngine)(nil).HandleContext), c)
}

// LoadHTMLFiles mocks base method.
func (m *MockEngine) LoadHTMLFiles(files ...string) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range files {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LoadHTMLFiles", varargs...)
}

// LoadHTMLFiles indicates an expected call of LoadHTMLFiles.
func (mr *MockEngineRecorder) LoadHTMLFiles(files ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadHTMLFiles", reflect.TypeOf((*MockEngine)(nil).LoadHTMLFiles), files...)
}

// LoadHTMLGlob mocks base method.
func (m *MockEngine) LoadHTMLGlob(pattern string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LoadHTMLGlob", pattern)
}

// LoadHTMLGlob indicates an expected call of LoadHTMLGlob.
func (mr *MockEngineRecorder) LoadHTMLGlob(pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadHTMLGlob", reflect.TypeOf((*MockEngine)(nil).LoadHTMLGlob), pattern)
}

// NoMethod mocks base method.
func (m *MockEngine) NoMethod(handlers ...gin.HandlerFunc) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range handlers {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NoMethod", varargs...)
}

// NoMethod indicates an expected call of NoMethod.
func (mr *MockEngineRecorder) NoMethod(handlers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoMethod", reflect.TypeOf((*MockEngine)(nil).NoMethod), handlers...)
}

// NoRoute mocks base method.
func (m *MockEngine) NoRoute(handlers ...gin.HandlerFunc) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range handlers {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NoRoute", varargs...)
}

// NoRoute indicates an expected call of NoRoute.
func (mr *MockEngineRecorder) NoRoute(handlers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoRoute", reflect.TypeOf((*MockEngine)(nil).NoRoute), handlers...)
}

// OPTIONS mocks base method.
func (m *MockEngine) OPTIONS(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OPTIONS", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// OPTIONS indicates an expected call of OPTIONS.
func (mr *MockEngineRecorder) OPTIONS(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OPTIONS", reflect.TypeOf((*MockEngine)(nil).OPTIONS), varargs...)
}

// PATCH mocks base method.
func (m *MockEngine) PATCH(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PATCH", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// PATCH indicates an expected call of PATCH.
func (mr *MockEngineRecorder) PATCH(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PATCH", reflect.TypeOf((*MockEngine)(nil).PATCH), varargs...)
}

// POST mocks base method.
func (m *MockEngine) POST(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "POST", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// POST indicates an expected call of POST.
func (mr *MockEngineRecorder) POST(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "POST", reflect.TypeOf((*MockEngine)(nil).POST), varargs...)
}

// PUT mocks base method.
func (m *MockEngine) PUT(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PUT", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// PUT indicates an expected call of PUT.
func (mr *MockEngineRecorder) PUT(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PUT", reflect.TypeOf((*MockEngine)(nil).PUT), varargs...)
}

// Routes mocks base method.
func (m *MockEngine) Routes() gin.RoutesInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Routes")
	ret0, _ := ret[0].(gin.RoutesInfo)
	return ret0
}

// Routes indicates an expected call of Routes.
func (mr *MockEngineRecorder) Routes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routes", reflect.TypeOf((*MockEngine)(nil).Routes))
}

// Run mocks base method.
func (m *MockEngine) Run(addr ...string) error {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range addr {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockEngineRecorder) Run(addr ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockEngine)(nil).Run), addr...)
}

// RunFd mocks base method.
func (m *MockEngine) RunFd(fd int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunFd", fd)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFd indicates an expected call of RunFd.
func (mr *MockEngineRecorder) RunFd(fd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFd", reflect.TypeOf((*MockEngine)(nil).RunFd), fd)
}

// RunListener mocks base method.
func (m *MockEngine) RunListener(listener net.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunListener", listener)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunListener indicates an expected call of RunListener.
func (mr *MockEngineRecorder) RunListener(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunListener", reflect.TypeOf((*MockEngine)(nil).RunListener), listener)
}

// RunTLS mocks base method.
func (m *MockEngine) RunTLS(addr, certFile, keyFile string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTLS", addr, certFile, keyFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTLS indicates an expected call of RunTLS.
func (mr *MockEngineRecorder) RunTLS(addr, certFile, keyFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTLS", reflect.TypeOf((*MockEngine)(nil).RunTLS), addr, certFile, keyFile)
}

// RunUnix mocks base method.
func (m *MockEngine) RunUnix(file string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunUnix", file)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunUnix indicates an expected call of RunUnix.
func (mr *MockEngineRecorder) RunUnix(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunUnix", reflect.TypeOf((*MockEngine)(nil).RunUnix), file)
}

//revive:disable

// SecureJsonPrefix mocks base method.
func (m *MockEngine) SecureJsonPrefix(prefix string) *gin.Engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecureJsonPrefix", prefix)
	ret0, _ := ret[0].(*gin.Engine)
	return ret0
}

// SecureJsonPrefix indicates an expected call of SecureJsonPrefix.
func (mr *MockEngineRecorder) SecureJsonPrefix(prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecureJsonPrefix", reflect.TypeOf((*MockEngine)(nil).SecureJsonPrefix), prefix)
}

//revive:enable

// ServeHTTP mocks base method.
func (m *MockEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeHTTP", w, req)
}

// ServeHTTP indicates an expected call of ServeHTTP.
func (mr *MockEngineRecorder) ServeHTTP(w, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeHTTP", reflect.TypeOf((*MockEngine)(nil).ServeHTTP), w, req)
}

// SetFuncMap mocks base method.
func (m *MockEngine) SetFuncMap(funcMap template.FuncMap) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFuncMap", funcMap)
}

// SetFuncMap indicates an expected call of SetFuncMap.
func (mr *MockEngineRecorder) SetFuncMap(funcMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFuncMap", reflect.TypeOf((*MockEngine)(nil).SetFuncMap), funcMap)
}

// SetHTMLTemplate mocks base method.
func (m *MockEngine) SetHTMLTemplate(tpl *template.Template) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHTMLTemplate", tpl)
}

// SetHTMLTemplate indicates an expected call of SetHTMLTemplate.
func (mr *MockEngineRecorder) SetHTMLTemplate(tpl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHTMLTemplate", reflect.TypeOf((*MockEngine)(nil).SetHTMLTemplate), tpl)
}

// Static mocks base method.
func (m *MockEngine) Static(arg0, arg1 string) gin.IRoutes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Static", arg0, arg1)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Static indicates an expected call of Static.
func (mr *MockEngineRecorder) Static(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Static", reflect.TypeOf((*MockEngine)(nil).Static), arg0, arg1)
}

// StaticFS mocks base method.
func (m *MockEngine) StaticFS(arg0 string, arg1 http.FileSystem) gin.IRoutes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StaticFS", arg0, arg1)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// StaticFS indicates an expected call of StaticFS.
func (mr *MockEngineRecorder) StaticFS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaticFS", reflect.TypeOf((*MockEngine)(nil).StaticFS), arg0, arg1)
}

// StaticFile mocks base method.
func (m *MockEngine) StaticFile(arg0, arg1 string) gin.IRoutes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StaticFile", arg0, arg1)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// StaticFile indicates an expected call of StaticFile.
func (mr *MockEngineRecorder) StaticFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaticFile", reflect.TypeOf((*MockEngine)(nil).StaticFile), arg0, arg1)
}

// StaticFileFS mocks base method.
func (m *MockEngine) StaticFileFS(arg0, arg1 string, arg2 http.FileSystem) gin.IRoutes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StaticFileFS", arg0, arg1, arg2)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// StaticFileFS indicates an expected call of StaticFileFS.
func (mr *MockEngineRecorder) StaticFileFS(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaticFileFS", reflect.TypeOf((*MockEngine)(nil).StaticFileFS), arg0, arg1, arg2)
}

// Use mocks base method.
func (m *MockEngine) Use(arg0 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Use", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *MockEngineRecorder) Use(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockEngine)(nil).Use), arg0...)
}

//------------------------------------------------------------------------------
// Register
//------------------------------------------------------------------------------

// MockRegister is a mock of IRegister interface.
type MockRegister struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterRecorder
}

// MockRegisterRecorder is the mock recorder for MockRegister.
type MockRegisterRecorder struct {
	mock *MockRegister
}

// NewMockRegister creates a new mock instance.
func NewMockRegister(ctrl *gomock.Controller) *MockRegister {
	mock := &MockRegister{ctrl: ctrl}
	mock.recorder = &MockRegisterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegister) EXPECT() *MockRegisterRecorder {
	return m.recorder
}

// Reg mocks base method.
func (m *MockRegister) Reg(engine Engine) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reg", engine)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reg indicates an expected call of Reg.
func (mr *MockRegisterRecorder) Reg(engine interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reg", reflect.TypeOf((*MockRegister)(nil).Reg), engine)
}
