package rest

import (
	"net/http"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

//------------------------------------------------------------------------------
// HTTP Client
//------------------------------------------------------------------------------

// MockHTTPClient is a mock instance of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientRecorder
}

var _ httpClient = &MockHTTPClient{}

// MockHTTPClientRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientRecorder {
	return m.recorder
}

// Do mock base method.
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicate an expected call of Do.
func (mr *MockHTTPClientRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), req)
}

//------------------------------------------------------------------------------
// Decoder
//------------------------------------------------------------------------------

// MockDecoder is a mock instance of IDecoder interface
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderRecorder
}

var _ config.IDecoder = &MockDecoder{}

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
func (m *MockDecoder) Decode() (config.IConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode")
	ret0, _ := ret[0].(config.IConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockDecoderRecorder) Decode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode))
}

//------------------------------------------------------------------------------
// Decoder Factory
//------------------------------------------------------------------------------

// MockDecoderFactory is a mock instance of IDecoderFactory interface.
type MockDecoderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderFactoryRecorder
}

var _ config.IDecoderFactory = &MockDecoderFactory{}

// MockDecoderFactoryRecorder is the mock recorder for MockDecoderFactory.
type MockDecoderFactoryRecorder struct {
	mock *MockDecoderFactory
}

// NewMockDecoderFactory creates a new mock instance.
func NewMockDecoderFactory(ctrl *gomock.Controller) *MockDecoderFactory {
	mock := &MockDecoderFactory{ctrl: ctrl}
	mock.recorder = &MockDecoderFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoderFactory) EXPECT() *MockDecoderFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDecoderFactory) Create(format string, args ...interface{}) (config.IDecoder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(config.IDecoder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDecoderFactoryRecorder) Create(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDecoderFactory)(nil).Create), varargs...)
}

// Register mocks base method.
func (m *MockDecoderFactory) Register(strategy config.IDecoderStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockDecoderFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockDecoderFactory)(nil).Register), strategy)
}
