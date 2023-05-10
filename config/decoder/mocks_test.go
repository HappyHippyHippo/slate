package decoder

import (
	"io"
	"reflect"

	"github.com/golang/mock/gomock"
)

//------------------------------------------------------------------------------
// Reader
//------------------------------------------------------------------------------

// MockReader is a mock instance of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderRecorder
}

var _ io.Reader = &MockReader{}

// MockReaderRecorder is the mock recorder for MockDecoder
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
// Underlying Decoder
//------------------------------------------------------------------------------

// MockUnderlyingDecoder is a mock instance of IUnderlyingDecoder interface
type MockUnderlyingDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockUnderlyingDecoderRecorder
}

var _ IUnderlyingDecoder = &MockUnderlyingDecoder{}

// MockUnderlyingDecoderRecorder is the mock recorder for MockUnderlyingDecoder
type MockUnderlyingDecoderRecorder struct {
	mock *MockUnderlyingDecoder
}

// NewMockUnderlyingDecoder creates a new mock instance
func NewMockUnderlyingDecoder(ctrl *gomock.Controller) *MockUnderlyingDecoder {
	mock := &MockUnderlyingDecoder{ctrl: ctrl}
	mock.recorder = &MockUnderlyingDecoderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnderlyingDecoder) EXPECT() *MockUnderlyingDecoderRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockUnderlyingDecoder) Decode(partial interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", partial)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockUnderlyingDecoderRecorder) Decode(partial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockUnderlyingDecoder)(nil).Decode), partial)
}
