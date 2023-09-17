package response

import (
	"bufio"
	"net"
	"net/http"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockWriter is a mock of Writer interface.
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterRecorder
}

// MockWriterRecorder is the mock recorder for MockWriter.
type MockWriterRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance.
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriter) EXPECT() *MockWriterRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockWriter) Body() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockWriterRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockWriter)(nil).Body))
}

// CloseNotify mocks base method.
func (m *MockWriter) CloseNotify() <-chan bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseNotify")
	ret0, _ := ret[0].(<-chan bool)
	return ret0
}

// CloseNotify indicates an expected call of CloseNotify.
func (mr *MockWriterRecorder) CloseNotify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseNotify", reflect.TypeOf((*MockWriter)(nil).CloseNotify))
}

// Flush mocks base method.
func (m *MockWriter) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush.
func (mr *MockWriterRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockWriter)(nil).Flush))
}

// Header mocks base method.
func (m *MockWriter) Header() http.Header {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(http.Header)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockWriterRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockWriter)(nil).Header))
}

// Hijack mocks base method.
func (m *MockWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hijack")
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(*bufio.ReadWriter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Hijack indicates an expected call of Hijack.
func (mr *MockWriterRecorder) Hijack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hijack", reflect.TypeOf((*MockWriter)(nil).Hijack))
}

// Pusher mocks base method.
func (m *MockWriter) Pusher() http.Pusher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pusher")
	ret0, _ := ret[0].(http.Pusher)
	return ret0
}

// Pusher indicates an expected call of Pusher.
func (mr *MockWriterRecorder) Pusher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pusher", reflect.TypeOf((*MockWriter)(nil).Pusher))
}

// Size mocks base method.
func (m *MockWriter) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockWriterRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockWriter)(nil).Size))
}

// Status mocks base method.
func (m *MockWriter) Status() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(int)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockWriterRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockWriter)(nil).Status))
}

// Write mocks base method.
func (m *MockWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockWriterRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriter)(nil).Write), arg0)
}

// WriteHeader mocks base method.
func (m *MockWriter) WriteHeader(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeader", arg0)
}

// WriteHeader indicates an expected call of WriteHeader.
func (mr *MockWriterRecorder) WriteHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeader", reflect.TypeOf((*MockWriter)(nil).WriteHeader), arg0)
}

// WriteHeaderNow mocks base method.
func (m *MockWriter) WriteHeaderNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteHeaderNow")
}

// WriteHeaderNow indicates an expected call of WriteHeaderNow.
func (mr *MockWriterRecorder) WriteHeaderNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteHeaderNow", reflect.TypeOf((*MockWriter)(nil).WriteHeaderNow))
}

// WriteString mocks base method.
func (m *MockWriter) WriteString(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteString", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteString indicates an expected call of WriteString.
func (mr *MockWriterRecorder) WriteString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockWriter)(nil).WriteString), arg0)
}

// Written mocks base method.
func (m *MockWriter) Written() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Written")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Written indicates an expected call of Written.
func (mr *MockWriterRecorder) Written() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Written", reflect.TypeOf((*MockWriter)(nil).Written))
}
