package slate

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// ----------------------------------------------------------------------------
// LogFormatterCreator
// ----------------------------------------------------------------------------

// MockLogFormatterCreator is a mock instance of FormatterStrategy interface
type MockLogFormatterCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterCreatorRecorder
}

var _ LogFormatterCreator = &MockLogFormatterCreator{}

// MockLogFormatterCreatorRecorder is the mock recorder for MockLogFormatterCreator
type MockLogFormatterCreatorRecorder struct {
	mock *MockLogFormatterCreator
}

// NewMockLogFormatterCreator creates a new mock instance
func NewMockLogFormatterCreator(ctrl *gomock.Controller) *MockLogFormatterCreator {
	mock := &MockLogFormatterCreator{ctrl: ctrl}
	mock.recorder = &MockLogFormatterCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogFormatterCreator) EXPECT() *MockLogFormatterCreatorRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockLogFormatterCreator) Accept(format string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", format)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockLogFormatterCreatorRecorder) Accept(format interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockLogFormatterCreator)(nil).Accept), format)
}

// Create mocks base method
func (m *MockLogFormatterCreator) Create(args ...interface{}) (LogFormatter, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(LogFormatter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockLogFormatterCreatorRecorder) Create(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLogFormatterCreator)(nil).Create), args...)
}

// ----------------------------------------------------------------------------
// LogFormatter
// ----------------------------------------------------------------------------

// MockLogFormatter is a mock instance of Formatter interface
type MockLogFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockLogFormatterRecorder
}

var _ LogFormatter = &MockLogFormatter{}

// MockLogFormatterRecorder is the mock recorder for MockLogFormatter
type MockLogFormatterRecorder struct {
	mock *MockLogFormatter
}

// NewMockLogFormatter creates a new mock instance
func NewMockLogFormatter(ctrl *gomock.Controller) *MockLogFormatter {
	mock := &MockLogFormatter{ctrl: ctrl}
	mock.recorder = &MockLogFormatterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogFormatter) EXPECT() *MockLogFormatterRecorder {
	return m.recorder
}

// Format mocks base method
func (m *MockLogFormatter) Format(level LogLevel, message string, ctx ...LogContext) string {
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
func (mr *MockLogFormatterRecorder) Format(level, message interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{level, message}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "format", reflect.TypeOf((*MockLogFormatter)(nil).Format), varargs...)
}

// ----------------------------------------------------------------------------
// LogWriterCreator
// ----------------------------------------------------------------------------

// MockLogWriterCreator is a mock instance of StreamStrategy interface
type MockLogWriterCreator struct {
	ctrl     *gomock.Controller
	recorder *MockLogWriterCreatorRecorder
}

var _ LogWriterCreator = &MockLogWriterCreator{}

// MockLogWriterCreatorRecorder is the mock recorder for MockLogWriterCreator
type MockLogWriterCreatorRecorder struct {
	mock *MockLogWriterCreator
}

// NewMockLogWriterCreator creates a new mock instance
func NewMockLogWriterCreator(ctrl *gomock.Controller) *MockLogWriterCreator {
	mock := &MockLogWriterCreator{ctrl: ctrl}
	mock.recorder = &MockLogWriterCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogWriterCreator) EXPECT() *MockLogWriterCreatorRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockLogWriterCreator) Accept(config *ConfigPartial) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", config)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockLogWriterCreatorRecorder) Accept(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockLogWriterCreator)(nil).Accept), config)
}

// Create mocks base method
func (m *MockLogWriterCreator) Create(config *ConfigPartial) (LogWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", config)
	ret0, _ := ret[0].(LogWriter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of CreateConfig
func (mr *MockLogWriterCreatorRecorder) Create(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLogWriterCreator)(nil).Create), config)
}

// ----------------------------------------------------------------------------
// LogWriter
// ----------------------------------------------------------------------------

// MockLogWriter is a mock instance of Stream interface
type MockLogWriter struct {
	ctrl     *gomock.Controller
	recorder *MockLogWriterRecorder
}

var _ LogWriter = &MockLogWriter{}

// MockLogWriterRecorder is the mock recorder for MockLogWriter
type MockLogWriterRecorder struct {
	mock *MockLogWriter
}

// NewMockLogWriter creates a new mock instance
func NewMockLogWriter(ctrl *gomock.Controller) *MockLogWriter {
	mock := &MockLogWriter{ctrl: ctrl}
	mock.recorder = &MockLogWriterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogWriter) EXPECT() *MockLogWriterRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockLogWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockLogWriterRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLogWriter)(nil).Close))
}

// Level mocks base method
func (m *MockLogWriter) Level() LogLevel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Level")
	ret0, _ := ret[0].(LogLevel)
	return ret0
}

// Level indicates an expected call of Level
func (mr *MockLogWriterRecorder) Level() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Level", reflect.TypeOf((*MockLogWriter)(nil).Level))
}

// Signal mocks base method
func (m *MockLogWriter) Signal(channel string, level LogLevel, message string, ctx ...LogContext) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{channel, level, message}
	for _, a := range ctx {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Signal", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Signal indicates an expected call of Signal
func (mr *MockLogWriterRecorder) Signal(channel, level, message interface{}, ctx ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{channel, level, message}, ctx...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signal", reflect.TypeOf((*MockLogWriter)(nil).Signal), varargs...)
}

// Broadcast mocks base method
func (m *MockLogWriter) Broadcast(level LogLevel, message string, fields ...LogContext) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{level, message}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Broadcast", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast
func (mr *MockLogWriterRecorder) Broadcast(level, message interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{level, message}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockLogWriter)(nil).Broadcast), varargs...)
}

// Flush mocks base method
func (m *MockLogWriter) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockLogWriterRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockLogWriter)(nil).Flush))
}

// HasChannel mocks base method
func (m *MockLogWriter) HasChannel(channel string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasChannel", channel)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasChannel indicates an expected call of HasChannel
func (mr *MockLogWriterRecorder) HasChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChannel", reflect.TypeOf((*MockLogWriter)(nil).HasChannel), channel)
}

// ListChannels mocks base method
func (m *MockLogWriter) ListChannels() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChannels")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ListChannels indicates an expected call of ListChannels
func (mr *MockLogWriterRecorder) ListChannels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChannels", reflect.TypeOf((*MockLogWriter)(nil).ListChannels))
}

// AddChannel mocks base method
func (m *MockLogWriter) AddChannel(channel string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChannel", channel)
}

// AddChannel indicates an expected call of AddChannel
func (mr *MockLogWriterRecorder) AddChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChannel", reflect.TypeOf((*MockLogWriter)(nil).AddChannel), channel)
}

// RemoveChannel mocks base method
func (m *MockLogWriter) RemoveChannel(channel string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveChannel", channel)
}

// RemoveChannel indicates an expected call of RemoveChannel
func (mr *MockLogWriterRecorder) RemoveChannel(channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveChannel", reflect.TypeOf((*MockLogWriter)(nil).RemoveChannel), channel)
}
