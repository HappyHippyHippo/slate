package gconfig

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockDecoder is a mock of Decoder interface
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderRecorder
}

var _ Decoder = &MockDecoder{}

// MockDecoderRecorder is the mock recorder for MockDecoder
type MockDecoderRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDecoder) EXPECT() *MockDecoderRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDecoder) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockDecoderRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDecoder)(nil).Close))
}

// Decode mocks base method
func (m *MockDecoder) Decode() (Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode")
	ret0, _ := ret[0].(Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockDecoderRecorder) Decode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode))
}
