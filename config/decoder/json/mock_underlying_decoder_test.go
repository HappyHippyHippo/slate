package json

import (
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate/config/decoder"
)

// MockUnderlyingDecoder is a mock instance of decoder.UnderlyingDecoder interface
type MockUnderlyingDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockUnderlyingDecoderRecorder
}

var _ decoder.UnderlyingDecoder = &MockUnderlyingDecoder{}

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
