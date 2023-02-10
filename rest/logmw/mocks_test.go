package logmw

import (
	"bufio"
	"net"
	"net/http"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/log"
)

//------------------------------------------------------------------------------
// Response Writer
//------------------------------------------------------------------------------

// MockResponseWriter is a mocked instance of responseWriter interface
type MockResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockResponseWriterRecorder
}

var _ responseWriter = &MockResponseWriter{}

// MockResponseWriterRecorder is the mock recorder for MockResponseWriter
type MockResponseWriterRecorder struct {
	mock *MockResponseWriter
}

// NewMockResponseWriter creates a new mock instance
func NewMockResponseWriter(ctrl *gomock.Controller) *MockResponseWriter {
	mock := &MockResponseWriter{ctrl: ctrl}
	mock.recorder = &MockResponseWriterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResponseWriter) EXPECT() *MockResponseWriterRecorder {
	return m.recorder
}

// Body mocks base method
func (m *MockResponseWriter) Body() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Body indicates an expected call of Body
func (mr *MockResponseWriterRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockResponseWriter)(nil).Body))
}

// CloseNotify mocks base method
func (m *MockResponseWriter) CloseNotify() <-chan bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseNotify")
	ret0, _ := ret[0].(<-chan bool)
	return ret0
}

// CloseNotify indicates an expected call of CloseNotify
func (mr *MockResponseWriterRecorder) CloseNotify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseNotify", reflect.TypeOf((*MockResponseWriter)(nil).CloseNotify))
}

// Flush mocks base method
func (m *MockResponseWriter) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush
func (mr *MockResponseWriterRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockResponseWriter)(nil).Flush))
}

// Header mocks base method
func (m *MockResponseWriter) Header() http.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(http.Header)
	return ret0
}

// Header indicates an expected call of Header
func (mr *MockResponseWriterRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockResponseWriter)(nil).Header))
}

// Hijack mocks base method
func (m *MockResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hijack")
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(*bufio.ReadWriter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Hijack indicates an expected call of Hijack
func (mr *MockResponseWriterRecorder) Hijack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hijack", reflect.TypeOf((*MockResponseWriter)(nil).Hijack))
}

// Pusher mocks base method
func (m *MockResponseWriter) Pusher() http.Pusher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pusher")
	ret0, _ := ret[0].(http.Pusher)
	return ret0
}

// Pusher indicates an expected call of Pusher
func (mr *MockResponseWriterRecorder) Pusher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pusher", reflect.TypeOf((*MockResponseWriter)(nil).Pusher))
}

// Size mocks base method
func (m *MockResponseWriter) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockResponseWriterRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockResponseWriter)(nil).Size))
}

// Status mocks base method
func (m *MockResponseWriter) Status() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(int)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockResponseWriterRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockResponseWriter)(nil).Status))
}

// Write mocks base method
func (m *MockResponseWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockResponseWriterRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockResponseWriter)(nil).Write), arg0)
}

// WriteHeader mocks base method
func (m *MockResponseWriter) WriteHeader(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeader", arg0)
}

// WriteHeader indicates an expected call of WriteHeader
func (mr *MockResponseWriterRecorder) WriteHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeader", reflect.TypeOf((*MockResponseWriter)(nil).WriteHeader), arg0)
}

// WriteHeaderNow mocks base method
func (m *MockResponseWriter) WriteHeaderNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeaderNow")
}

// WriteHeaderNow indicates an expected call of WriteHeaderNow
func (mr *MockResponseWriterRecorder) WriteHeaderNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeaderNow", reflect.TypeOf((*MockResponseWriter)(nil).WriteHeaderNow))
}

// WriteString mocks base method
func (m *MockResponseWriter) WriteString(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteString", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteString indicates an expected call of WriteString
func (mr *MockResponseWriterRecorder) WriteString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockResponseWriter)(nil).WriteString), arg0)
}

// Written mocks base method
func (m *MockResponseWriter) Written() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Written")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Written indicates an expected call of Written
func (mr *MockResponseWriterRecorder) Written() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Written", reflect.TypeOf((*MockResponseWriter)(nil).Written))
}

//------------------------------------------------------------------------------
// Reader
//------------------------------------------------------------------------------

// MockReader is a mocked instance of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderRecorder
}

// MockReaderRecorder is the mock recorder for MockReader
type MockReaderRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReader) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReaderRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReader)(nil).Read), arg0)
}

// Close mocks base method
func (m *MockReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockReaderRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReader)(nil).Close))
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
