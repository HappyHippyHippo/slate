package watchdog

import (
	"reflect"

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
func (m *MockLog) Broadcast(level log.Level, msg string, ctx map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", level, msg, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockLogRecorder) Broadcast(level, msg, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockLog)(nil).Broadcast), level, msg, ctx)
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
func (m *MockLog) Signal(channel string, level log.Level, msg string, ctx map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signal", channel, level, msg, ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal.
func (mr *MockLogRecorder) Signal(channel, level, msg, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockLog)(nil).Signal), channel, level, msg, ctx)
}

// Stream mocks base method.
func (m *MockLog) Stream(id string) log.IStream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", id)
	ret0, _ := ret[0].(log.IStream)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockLogRecorder) Stream(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockLog)(nil).Stream), id)
}

//------------------------------------------------------------------------------
// Log Formatter
//------------------------------------------------------------------------------

// MockLogFormatter is a mock an instance of LogFormatter interface.
type MockLogFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterRecorder
}

var _ ILogFormatter = &MockLogFormatter{}

// MockLogFormatterRecorder is the mock recorder for MockLogFormatter.
type MockLogFormatterRecorder struct {
	mock *MockLogFormatter
}

// NewMockLogFormatter creates a new mock instance.
func NewMockLogFormatter(ctrl *gomock.Controller) *MockLogFormatter {
	mock := &MockLogFormatter{ctrl: ctrl}
	mock.recorder = &MockLogFormatterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFormatter) EXPECT() *MockLogFormatterRecorder {
	return m.recorder
}

// Done mocks base method.
func (m *MockLogFormatter) Done(service string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done", service)
	ret0, _ := ret[0].(string)
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockLogFormatterRecorder) Done(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockLogFormatter)(nil).Done), service)
}

// Error mocks base method.
func (m *MockLogFormatter) Error(service string, e error) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", service, e)
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockLogFormatterRecorder) Error(service, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogFormatter)(nil).Error), service, e)
}

// Start mocks base method.
func (m *MockLogFormatter) Start(service string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", service)
	ret0, _ := ret[0].(string)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockLogFormatterRecorder) Start(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockLogFormatter)(nil).Start), service)
}

//------------------------------------------------------------------------------
// Log Formatter Strategy
//------------------------------------------------------------------------------

// MockLogFormatterStrategy is a mock an instance of LogFormatterStrategy interface.
type MockLogFormatterStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterStrategyRecorder
}

var _ ILogFormatterStrategy = &MockLogFormatterStrategy{}

// MockLogFormatterStrategyRecorder is the mock recorder for MockLogFormatterStrategy.
type MockLogFormatterStrategyRecorder struct {
	mock *MockLogFormatterStrategy
}

// NewMockLogFormatterStrategy creates a new mock instance.
func NewMockLogFormatterStrategy(ctrl *gomock.Controller) *MockLogFormatterStrategy {
	mock := &MockLogFormatterStrategy{ctrl: ctrl}
	mock.recorder = &MockLogFormatterStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFormatterStrategy) EXPECT() *MockLogFormatterStrategyRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockLogFormatterStrategy) Accept(cfg config.IConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockLogFormatterStrategyRecorder) Accept(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockLogFormatterStrategy)(nil).Accept), cfg)
}

// Create mocks base method.
func (m *MockLogFormatterStrategy) Create(cfg config.IConfig) (ILogFormatter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(ILogFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLogFormatterStrategyRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLogFormatterStrategy)(nil).Create), cfg)
}

//------------------------------------------------------------------------------
// Log Formatter Factory
//------------------------------------------------------------------------------

// MockLogFormatterFactory is a mock of ILogFormatterFactory interface.
type MockLogFormatterFactory struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterFactoryRecorder
}

// MockLogFormatterFactoryRecorder is the mock recorder for MockLogFormatterFactory.
type MockLogFormatterFactoryRecorder struct {
	mock *MockLogFormatterFactory
}

// NewMockLogFormatterFactory creates a new mock instance.
func NewMockLogFormatterFactory(ctrl *gomock.Controller) *MockLogFormatterFactory {
	mock := &MockLogFormatterFactory{ctrl: ctrl}
	mock.recorder = &MockLogFormatterFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogFormatterFactory) EXPECT() *MockLogFormatterFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockLogFormatterFactory) Create(cfg config.IConfig) (ILogFormatter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(ILogFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockLogFormatterFactoryRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLogFormatterFactory)(nil).Create), cfg)
}

// Register mocks base method.
func (m *MockLogFormatterFactory) Register(strategy ILogFormatterStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockLogFormatterFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockLogFormatterFactory)(nil).Register), strategy)
}

//------------------------------------------------------------------------------
// Log Adapter
//------------------------------------------------------------------------------

// MockLogAdapter is a mock an instance of LogAdapter interface.
type MockLogAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockLogAdapterRecorder
}

// MockLogAdapterRecorder is the mock recorder for MockLogAdapter.
type MockLogAdapterRecorder struct {
	mock *MockLogAdapter
}

// NewMockLogAdapter creates a new mock instance.
func NewMockLogAdapter(ctrl *gomock.Controller) *MockLogAdapter {
	mock := &MockLogAdapter{ctrl: ctrl}
	mock.recorder = &MockLogAdapterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogAdapter) EXPECT() *MockLogAdapterRecorder {
	return m.recorder
}

// Done mocks base method.
func (m *MockLogAdapter) Done() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Done")
	ret0, _ := ret[0].(error)
	return ret0
}

// Done indicates an expected call of Done.
func (mr *MockLogAdapterRecorder) Done() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Done", reflect.TypeOf((*MockLogAdapter)(nil).Done))
}

// Error mocks base method.
func (m *MockLogAdapter) Error(e error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockLogAdapterRecorder) Error(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogAdapter)(nil).Error), e)
}

// Start mocks base method.
func (m *MockLogAdapter) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockLogAdapterRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockLogAdapter)(nil).Start))
}

//------------------------------------------------------------------------------
// Watchdog
//------------------------------------------------------------------------------

// MockWatchdog is a mock of IWatchdog interface.
type MockWatchdog struct {
	ctrl     *gomock.Controller
	recorder *MockWatchdogRecorder
}

// MockWatchdogRecorder is the mock recorder for MockWatchdog.
type MockWatchdogRecorder struct {
	mock *MockWatchdog
}

// NewMockWatchdog creates a new mock instance.
func NewMockWatchdog(ctrl *gomock.Controller) *MockWatchdog {
	mock := &MockWatchdog{ctrl: ctrl}
	mock.recorder = &MockWatchdogRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchdog) EXPECT() *MockWatchdogRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockWatchdog) Run(process IProcess) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", process)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockWatchdogRecorder) Run(process interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWatchdog)(nil).Run), process)
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// MockFactory is a mock of IFactory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryRecorder
}

// MockFactoryRecorder is the mock recorder for MockFactory.
type MockFactoryRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFactory) Create(service string) (IWatchdog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", service)
	ret0, _ := ret[0].(IWatchdog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFactoryRecorder) Create(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFactory)(nil).Create), service)
}
