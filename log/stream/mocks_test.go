package stream

import (
	"io"
	"reflect"

	"github.com/golang/mock/gomock"
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
