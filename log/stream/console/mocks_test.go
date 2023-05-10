package console

import (
	"io"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
)

//------------------------------------------------------------------------------
// Writer
//------------------------------------------------------------------------------

// MockWriter is a mock instance of Writer interface
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterRecorder
}

var _ io.Writer = &MockWriter{}

// MockWriterRecorder is the mock recorder for MockWriter
type MockWriterRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriter) EXPECT() *MockWriterRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockWriterRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriter)(nil).Write), arg0)
}

// Close mocks base method
func (m *MockWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockWriterRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWriter)(nil).Close))
}

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
// Formatter
//------------------------------------------------------------------------------

// MockFormatter is a mock instance of LogFormatter interface
type MockFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockFormatterRecorder
}

var _ log.IFormatter = &MockFormatter{}

// MockFormatterRecorder is the mock recorder for MockFormatter
type MockFormatterRecorder struct {
	mock *MockFormatter
}

// NewMockFormatter creates a new mock instance
func NewMockFormatter(ctrl *gomock.Controller) *MockFormatter {
	mock := &MockFormatter{ctrl: ctrl}
	mock.recorder = &MockFormatterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFormatter) EXPECT() *MockFormatterRecorder {
	return m.recorder
}

// Format mocks base method
func (m *MockFormatter) Format(level log.Level, message string, ctx ...log.Context) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{level, message}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "format", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// Format indicates an expected call of Format
func (mr *MockFormatterRecorder) Format(level, message interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{level, message}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "format", reflect.TypeOf((*MockFormatter)(nil).Format), varargs...)
}

//------------------------------------------------------------------------------
// Formatter Factory
//------------------------------------------------------------------------------

// MockFormatterFactory is a mock instance of IFormatterFactory interface.
type MockFormatterFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFormatterFactoryRecorder
}

var _ log.IFormatterFactory = &MockFormatterFactory{}

// MockFormatterFactoryRecorder is the mock recorder for MockFormatterFactory.
type MockFormatterFactoryRecorder struct {
	mock *MockFormatterFactory
}

// NewMockFormatterFactory creates a new mock instance.
func NewMockFormatterFactory(ctrl *gomock.Controller) *MockFormatterFactory {
	mock := &MockFormatterFactory{ctrl: ctrl}
	mock.recorder = &MockFormatterFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFormatterFactory) EXPECT() *MockFormatterFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFormatterFactory) Create(format string, args ...interface{}) (log.IFormatter, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(log.IFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFormatterFactoryRecorder) Create(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFormatterFactory)(nil).Create), varargs...)
}

// Register mocks base method.
func (m *MockFormatterFactory) Register(strategy log.IFormatterStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockFormatterFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockFormatterFactory)(nil).Register), strategy)
}
