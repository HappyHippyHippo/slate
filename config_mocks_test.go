package slate

import (
	"net/http"
	"reflect"

	"github.com/golang/mock/gomock"
)

// ----------------------------------------------------------------------------
// ConfigParserCreator
// ----------------------------------------------------------------------------

// MockConfigParserCreator is a mock instance of DecoderStrategy interface
type MockConfigParserCreator struct {
	ctrl     *gomock.Controller
	recorder *MockConfigParserCreatorRecorder
}

var _ ConfigParserCreator = &MockConfigParserCreator{}

// MockConfigParserCreatorRecorder is the mock recorder for MockConfigParserCreator
type MockConfigParserCreatorRecorder struct {
	mock *MockConfigParserCreator
}

// NewMockConfigParserCreator creates a new mock instance
func NewMockConfigParserCreator(ctrl *gomock.Controller) *MockConfigParserCreator {
	mock := &MockConfigParserCreator{ctrl: ctrl}
	mock.recorder = &MockConfigParserCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigParserCreator) EXPECT() *MockConfigParserCreatorRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockConfigParserCreator) Accept(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockConfigParserCreatorRecorder) Accept(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockConfigParserCreator)(nil).Accept), varargs...)
}

// Create mocks base method
func (m *MockConfigParserCreator) Create(args ...interface{}) (ConfigParser, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(ConfigParser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockConfigParserCreatorRecorder) Create(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockConfigParserCreator)(nil).Create), args...)
}

// ----------------------------------------------------------------------------
// ConfigUnderlyingDecoder
// ----------------------------------------------------------------------------

// MockConfigUnderlyingDecoder is a mock instance of UnderlyingDecoder interface
type MockConfigUnderlyingDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockConfigUnderlyingDecoderRecorder
}

var _ ConfigUnderlyingDecoder = &MockConfigUnderlyingDecoder{}

// MockConfigUnderlyingDecoderRecorder is the mock recorder for MockConfigUnderlyingDecoder
type MockConfigUnderlyingDecoderRecorder struct {
	mock *MockConfigUnderlyingDecoder
}

// NewMockConfigUnderlyingDecoder creates a new mock instance
func NewMockConfigUnderlyingDecoder(ctrl *gomock.Controller) *MockConfigUnderlyingDecoder {
	mock := &MockConfigUnderlyingDecoder{ctrl: ctrl}
	mock.recorder = &MockConfigUnderlyingDecoderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigUnderlyingDecoder) EXPECT() *MockConfigUnderlyingDecoderRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockConfigUnderlyingDecoder) Decode(partial interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", partial)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockConfigUnderlyingDecoderRecorder) Decode(partial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockConfigUnderlyingDecoder)(nil).Decode), partial)
}

// ----------------------------------------------------------------------------
// ConfigParser
// ----------------------------------------------------------------------------

// MockConfigParser is a mock instance of Decoder interface
type MockConfigParser struct {
	ctrl     *gomock.Controller
	recorder *MockConfigParserRecorder
}

var _ ConfigParser = &MockConfigParser{}

// MockConfigParserRecorder is the mock recorder for MockConfigParser
type MockConfigParserRecorder struct {
	mock *MockConfigParser
}

// NewMockConfigParser creates a new mock instance
func NewMockConfigParser(ctrl *gomock.Controller) *MockConfigParser {
	mock := &MockConfigParser{ctrl: ctrl}
	mock.recorder = &MockConfigParserRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigParser) EXPECT() *MockConfigParserRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockConfigParser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConfigParserRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfigParser)(nil).Close))
}

// Parse mocks base method
func (m *MockConfigParser) Parse() (*ConfigPartial, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse")
	ret0, _ := ret[0].(*ConfigPartial)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse
func (mr *MockConfigParserRecorder) Parse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockConfigParser)(nil).Parse))
}

// ----------------------------------------------------------------------------
// ConfigSupplierCreator
// ----------------------------------------------------------------------------

// MockConfigSupplierCreator is a mock instance of SourceFactoryStrategy interface
type MockConfigSupplierCreator struct {
	ctrl     *gomock.Controller
	recorder *MockConfigSupplierCreatorRecorder
}

var _ ConfigSupplierCreator = &MockConfigSupplierCreator{}

// MockConfigSupplierCreatorRecorder is the mock recorder for MockConfigSupplierCreator
type MockConfigSupplierCreatorRecorder struct {
	mock *MockConfigSupplierCreator
}

// NewMockConfigSupplierCreator creates a new mock instance
func NewMockConfigSupplierCreator(ctrl *gomock.Controller) *MockConfigSupplierCreator {
	mock := &MockConfigSupplierCreator{ctrl: ctrl}
	mock.recorder = &MockConfigSupplierCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigSupplierCreator) EXPECT() *MockConfigSupplierCreatorRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockConfigSupplierCreator) Accept(partial *ConfigPartial) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", partial)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockConfigSupplierCreatorRecorder) Accept(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockConfigSupplierCreator)(nil).Accept), config)
}

// Create mocks base method
func (m *MockConfigSupplierCreator) Create(partial *ConfigPartial) (ConfigSupplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", partial)
	ret0, _ := ret[0].(ConfigSupplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockConfigSupplierCreatorRecorder) Create(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockConfigSupplierCreator)(nil).Create), config)
}

// ----------------------------------------------------------------------------
// ConfigSupplier
// ----------------------------------------------------------------------------

// MockConfigSupplier is a mock instance of Source interface
type MockConfigSupplier struct {
	ctrl     *gomock.Controller
	recorder *MockConfigSupplierRecorder
}

var _ ConfigSupplier = &MockConfigSupplier{}

// MockConfigSupplierRecorder is the mock recorder for MockConfigSupplier
type MockConfigSupplierRecorder struct {
	mock *MockConfigSupplier
}

// NewMockConfigSupplier creates a new mock instance
func NewMockConfigSupplier(ctrl *gomock.Controller) *MockConfigSupplier {
	mock := &MockConfigSupplier{ctrl: ctrl}
	mock.recorder = &MockConfigSupplierRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigSupplier) EXPECT() *MockConfigSupplierRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockConfigSupplier) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConfigSupplierRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfigSupplier)(nil).Close))
}

// Has mocks base method
func (m *MockConfigSupplier) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockConfigSupplierRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockConfigSupplier)(nil).Has), path)
}

// Get mocks base method
func (m *MockConfigSupplier) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockConfigSupplierRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigSupplier)(nil).Get), varargs...)
}

// ----------------------------------------------------------------------------
// ConfigObsSupplier
// ----------------------------------------------------------------------------

// MockConfigObsSupplier is a mock instance of ObsSource interface
type MockConfigObsSupplier struct {
	ctrl     *gomock.Controller
	recorder *MockConfigObsSupplierRecorder
}

var _ ConfigObsSupplier = &MockConfigObsSupplier{}

// MockConfigObsSupplierRecorder is the mock recorder for MockConfigObsSupplier
type MockConfigObsSupplierRecorder struct {
	mock *MockConfigObsSupplier
}

// NewMockConfigObsSupplier creates a new mock instance
func NewMockConfigObsSupplier(ctrl *gomock.Controller) *MockConfigObsSupplier {
	mock := &MockConfigObsSupplier{ctrl: ctrl}
	mock.recorder = &MockConfigObsSupplierRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigObsSupplier) EXPECT() *MockConfigObsSupplierRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockConfigObsSupplier) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConfigObsSupplierRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfigObsSupplier)(nil).Close))
}

// Has mocks base method
func (m *MockConfigObsSupplier) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockConfigObsSupplierRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockConfigObsSupplier)(nil).Has), path)
}

// Get mocks base method
func (m *MockConfigObsSupplier) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockConfigObsSupplierRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigObsSupplier)(nil).Get), varargs...)
}

// Reload mocks base method
func (m *MockConfigObsSupplier) Reload() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reload indicates an expected call of Reload
func (mr *MockConfigObsSupplierRecorder) Reload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockConfigObsSupplier)(nil).Reload))
}

// ----------------------------------------------------------------------------
// Requester
// ----------------------------------------------------------------------------

// MockConfigRestRequester is a mock instance of HTTPClient interface.
type MockConfigRestRequester struct {
	ctrl     *gomock.Controller
	recorder *MockConfigRestRequesterRecorder
}

var _ configRestRequester = &MockConfigRestRequester{}

// MockConfigRestRequesterRecorder is the mock recorder for MockConfigRestRequester.
type MockConfigRestRequesterRecorder struct {
	mock *MockConfigRestRequester
}

// NewMockConfigRestRequester creates a new mock instance.
func NewMockConfigRestRequester(ctrl *gomock.Controller) *MockConfigRestRequester {
	mock := &MockConfigRestRequester{ctrl: ctrl}
	mock.recorder = &MockConfigRestRequesterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigRestRequester) EXPECT() *MockConfigRestRequesterRecorder {
	return m.recorder
}

// Do mock base method.
func (m *MockConfigRestRequester) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicate an expected call of Do.
func (mr *MockConfigRestRequesterRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockConfigRestRequester)(nil).Do), req)
}
