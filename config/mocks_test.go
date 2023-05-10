package config

import (
	"io"
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/trigger"
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
// Trigger
//------------------------------------------------------------------------------

// MockTrigger is a mock of ITrigger interface.
type MockTrigger struct {
	ctrl     *gomock.Controller
	recorder *MockTriggerRecorder
}

var _ trigger.ITrigger = &MockTrigger{}

// MockTriggerRecorder is the mock recorder for MockTrigger.
type MockTriggerRecorder struct {
	mock *MockTrigger
}

// NewMockTrigger creates a new mock instance.
func NewMockTrigger(ctrl *gomock.Controller) *MockTrigger {
	mock := &MockTrigger{ctrl: ctrl}
	mock.recorder = &MockTriggerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrigger) EXPECT() *MockTriggerRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTrigger) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTriggerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTrigger)(nil).Close))
}

// Delay mocks base method.
func (m *MockTrigger) Delay() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Delay indicates an expected call of Timer.
func (mr *MockTriggerRecorder) Delay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockTrigger)(nil).Delay))
}

//------------------------------------------------------------------------------
// Decoder
//------------------------------------------------------------------------------

// MockDecoder is a mock instance of IDecoder interface
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderRecorder
}

var _ IDecoder = &MockDecoder{}

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
func (m *MockDecoder) Decode() (IConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode")
	ret0, _ := ret[0].(IConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockDecoderRecorder) Decode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode))
}

//------------------------------------------------------------------------------
// Decoder Strategy
//------------------------------------------------------------------------------

// MockDecoderStrategy is a mock instance of IDecoderStrategy interface
type MockDecoderStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderStrategyRecorder
}

var _ IDecoderStrategy = &MockDecoderStrategy{}

// MockDecoderStrategyRecorder is the mock recorder for MockDecoderStrategy
type MockDecoderStrategyRecorder struct {
	mock *MockDecoderStrategy
}

// NewMockDecoderStrategy creates a new mock instance
func NewMockDecoderStrategy(ctrl *gomock.Controller) *MockDecoderStrategy {
	mock := &MockDecoderStrategy{ctrl: ctrl}
	mock.recorder = &MockDecoderStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDecoderStrategy) EXPECT() *MockDecoderStrategyRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockDecoderStrategy) Accept(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockDecoderStrategyRecorder) Accept(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockDecoderStrategy)(nil).Accept), varargs...)
}

// Create mocks base method
func (m *MockDecoderStrategy) Create(args ...interface{}) (IDecoder, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(IDecoder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDecoderStrategyRecorder) Create(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDecoderStrategy)(nil).Create), args...)
}

//------------------------------------------------------------------------------
// Source
//------------------------------------------------------------------------------

// MockSource is a mock instance of ISource interface
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceRecorder
}

var _ ISource = &MockSource{}

// MockSourceRecorder is the mock recorder for MockSource
type MockSourceRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSource) EXPECT() *MockSourceRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSourceRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSource)(nil).Close))
}

// Has mocks base method
func (m *MockSource) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockSourceRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockSource)(nil).Has), path)
}

// Get mocks base method
func (m *MockSource) Get(path string, def ...interface{}) (interface{}, error) {
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
func (mr *MockSourceRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSource)(nil).Get), varargs...)
}

//------------------------------------------------------------------------------
// Observable Source
//------------------------------------------------------------------------------

// MockObsSource is a mock instance of ObsSource interface
type MockObsSource struct {
	ctrl     *gomock.Controller
	recorder *MockObsSourceRecorder
}

var _ IObsSource = &MockObsSource{}

// MockObsSourceRecorder is the mock recorder for MockObsSource
type MockObsSourceRecorder struct {
	mock *MockObsSource
}

// NewMockObsSource creates a new mock instance
func NewMockObsSource(ctrl *gomock.Controller) *MockObsSource {
	mock := &MockObsSource{ctrl: ctrl}
	mock.recorder = &MockObsSourceRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObsSource) EXPECT() *MockObsSourceRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockObsSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockObsSourceRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockObsSource)(nil).Close))
}

// Has mocks base method
func (m *MockObsSource) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockObsSourceRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockObsSource)(nil).Has), path)
}

// Get mocks base method
func (m *MockObsSource) Get(path string, def ...interface{}) (interface{}, error) {
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
func (mr *MockObsSourceRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSource)(nil).Get), varargs...)
}

// Reload mocks base method
func (m *MockObsSource) Reload() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reload indicates an expected call of Reload
func (mr *MockObsSourceRecorder) Reload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockObsSource)(nil).Reload))
}

//------------------------------------------------------------------------------
// Source Strategy
//------------------------------------------------------------------------------

// MockSourceStrategy is a mock instance of SourceFactoryStrategy interface
type MockSourceStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockSourceStrategyRecorder
}

var _ ISourceStrategy = &MockSourceStrategy{}

// MockSourceStrategyRecorder is the mock recorder for MockSourceStrategy
type MockSourceStrategyRecorder struct {
	mock *MockSourceStrategy
}

// NewMockSourceStrategy creates a new mock instance
func NewMockSourceStrategy(ctrl *gomock.Controller) *MockSourceStrategy {
	mock := &MockSourceStrategy{ctrl: ctrl}
	mock.recorder = &MockSourceStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceStrategy) EXPECT() *MockSourceStrategyRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockSourceStrategy) Accept(cfg IConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockSourceStrategyRecorder) Accept(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockSourceStrategy)(nil).Accept), cfg)
}

// Create mocks base method
func (m *MockSourceStrategy) Create(cfg IConfig) (ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockSourceStrategyRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceStrategy)(nil).Create), cfg)
}

//------------------------------------------------------------------------------
// Source Factory
//------------------------------------------------------------------------------

// MockSourceFactory is a mock instance of ISourceFactory interface.
type MockSourceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSourceFactoryRecorder
}

var _ ISourceFactory = &MockSourceFactory{}

// MockSourceFactoryRecorder is the mock recorder for MockSourceFactory.
type MockSourceFactoryRecorder struct {
	mock *MockSourceFactory
}

// NewMockSourceFactory creates a new mock instance.
func NewMockSourceFactory(ctrl *gomock.Controller) *MockSourceFactory {
	mock := &MockSourceFactory{ctrl: ctrl}
	mock.recorder = &MockSourceFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceFactory) EXPECT() *MockSourceFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSourceFactory) Create(cfg IConfig) (ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSourceFactoryRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceFactory)(nil).Create), cfg)
}

// Register mocks base method.
func (m *MockSourceFactory) Register(strategy ISourceStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockSourceFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSourceFactory)(nil).Register), strategy)
}

//------------------------------------------------------------------------------
// Loader
//------------------------------------------------------------------------------

// MockLoader is a mock instance of ILoader interface.
type MockLoader struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderRecorder
}

// MockLoaderRecorder is the mock recorder for MockLoader.
type MockLoaderRecorder struct {
	mock *MockLoader
}

// NewMockLoader creates a new mock instance.
func NewMockLoader(ctrl *gomock.Controller) *MockLoader {
	mock := &MockLoader{ctrl: ctrl}
	mock.recorder = &MockLoaderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoader) EXPECT() *MockLoaderRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockLoader) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockLoaderRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockLoader)(nil).Load))
}

//------------------------------------------------------------------------------
// Manager
//------------------------------------------------------------------------------

// MockManager is a mock instance of IManager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerRecorder
}

var _ IManager = &MockManager{}

// MockManagerRecorder is the mock recorder for MockManager.
type MockManagerRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerRecorder {
	return m.recorder
}

// AddObserver mocks base method.
func (m *MockManager) AddObserver(path string, callback IObserver) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddObserver", path, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddObserver indicates an expected call of AddObserver.
func (mr *MockManagerRecorder) AddObserver(path, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddObserver", reflect.TypeOf((*MockManager)(nil).AddObserver), path, callback)
}

// AddSource mocks base method.
func (m *MockManager) AddSource(id string, priority int, src ISource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSource", id, priority, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSource indicates an expected call of AddSource.
func (mr *MockManagerRecorder) AddSource(id, priority, src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSource", reflect.TypeOf((*MockManager)(nil).AddSource), id, priority, src)
}

// Bool mocks base method.
func (m *MockManager) Bool(path string, def ...bool) (bool, error) {
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
func (mr *MockManagerRecorder) Bool(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockManager)(nil).Bool), varargs...)
}

// Close mocks base method.
func (m *MockManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockManagerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockManager)(nil).Close))
}

// Config mocks base method.
func (m *MockManager) Config(path string, def ...Config) (IConfig, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Config", varargs...)
	ret0, _ := ret[0].(IConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Partial.
func (mr *MockManagerRecorder) Config(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockManager)(nil).Config), varargs...)
}

// Entries mocks base method.
func (m *MockManager) Entries() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entries")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Float mocks base method.
func (m *MockManager) Float(path string, def ...float64) (float64, error) {
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

// Entries indicates an expected call of Entries.
func (mr *MockManagerRecorder) Entries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entries", reflect.TypeOf((*MockManager)(nil).Entries))
}

// Float indicates an expected call of Float.
func (mr *MockManagerRecorder) Float(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float", reflect.TypeOf((*MockManager)(nil).Float), varargs...)
}

// Get mocks base method.
func (m *MockManager) Get(path string, def ...interface{}) (interface{}, error) {
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
func (mr *MockManagerRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockManager)(nil).Get), varargs...)
}

// Has mocks base method.
func (m *MockManager) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockManagerRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockManager)(nil).Has), path)
}

// HasObserver mocks base method.
func (m *MockManager) HasObserver(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasObserver", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasObserver indicates an expected call of HasObserver.
func (mr *MockManagerRecorder) HasObserver(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasObserver", reflect.TypeOf((*MockManager)(nil).HasObserver), path)
}

// HasSource mocks base method.
func (m *MockManager) HasSource(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSource", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSource indicates an expected call of HasSource.
func (mr *MockManagerRecorder) HasSource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSource", reflect.TypeOf((*MockManager)(nil).HasSource), id)
}

// Int mocks base method.
func (m *MockManager) Int(path string, def ...int) (int, error) {
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
func (mr *MockManagerRecorder) Int(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockManager)(nil).Int), varargs...)
}

// List mocks base method.
func (m *MockManager) List(path string, def ...[]interface{}) ([]interface{}, error) {
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
func (mr *MockManagerRecorder) List(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManager)(nil).List), varargs...)
}

// Populate mocks base method.
func (m *MockManager) Populate(path string, target interface{}, icase ...bool) (interface{}, error) {
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
func (mr *MockManagerRecorder) Populate(path, target interface{}, icase ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path, target}, icase...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Populate", reflect.TypeOf((*MockManager)(nil).Populate), varargs...)
}

// RemoveAllSources mocks base method.
func (m *MockManager) RemoveAllSources() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllSources")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllSources indicates an expected call of RemoveAllSources.
func (mr *MockManagerRecorder) RemoveAllSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllSources", reflect.TypeOf((*MockManager)(nil).RemoveAllSources))
}

// RemoveObserver mocks base method.
func (m *MockManager) RemoveObserver(path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveObserver", path)
}

// RemoveObserver indicates an expected call of RemoveObserver.
func (mr *MockManagerRecorder) RemoveObserver(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveObserver", reflect.TypeOf((*MockManager)(nil).RemoveObserver), path)
}

// RemoveSource mocks base method.
func (m *MockManager) RemoveSource(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSource", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSource indicates an expected call of RemoveSource.
func (mr *MockManagerRecorder) RemoveSource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSource", reflect.TypeOf((*MockManager)(nil).RemoveSource), id)
}

// Source mocks base method.
func (m *MockManager) Source(id string) (ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseSource", id)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Source indicates an expected call of Source.
func (mr *MockManagerRecorder) Source(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseSource", reflect.TypeOf((*MockManager)(nil).Source), id)
}

// SourcePriority mocks base method.
func (m *MockManager) SourcePriority(id string, priority int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourcePriority", id, priority)
	ret0, _ := ret[0].(error)
	return ret0
}

// SourcePriority indicates an expected call of SourcePriority.
func (mr *MockManagerRecorder) SourcePriority(id, priority interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourcePriority", reflect.TypeOf((*MockManager)(nil).SourcePriority), id, priority)
}

// String mocks base method.
func (m *MockManager) String(path string, def ...string) (string, error) {
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
func (mr *MockManagerRecorder) String(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockManager)(nil).String), varargs...)
}
