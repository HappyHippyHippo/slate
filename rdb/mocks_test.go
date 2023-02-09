package rdb

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

//------------------------------------------------------------------------------
// Config Source
//------------------------------------------------------------------------------

// MockConfigSource is a mock an instance of ConfigSource interface
type MockConfigSource struct {
	ctrl     *gomock.Controller
	recorder *MockConfigSourceRecorder
}

var _ config.ISource = &MockConfigSource{}

// MockConfigSourceRecorder is the mock recorder for MockConfigSource
type MockConfigSourceRecorder struct {
	mock *MockConfigSource
}

// NewMockConfigSource creates a new mock instance
func NewMockConfigSource(ctrl *gomock.Controller) *MockConfigSource {
	mock := &MockConfigSource{ctrl: ctrl}
	mock.recorder = &MockConfigSourceRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigSource) EXPECT() *MockConfigSourceRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockConfigSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConfigSourceRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfigSource)(nil).Close))
}

// Has mocks base method
func (m *MockConfigSource) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockConfigSourceRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockConfigSource)(nil).Has), path)
}

// Get mocks base method
func (m *MockConfigSource) Get(path string, def ...interface{}) (interface{}, error) {
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
func (mr *MockConfigSourceRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigSource)(nil).Get), varargs...)
}

//------------------------------------------------------------------------------
// Config Manager
//------------------------------------------------------------------------------

// MockConfigManager is a mock an instance of IManager interface.
type MockConfigManager struct {
	ctrl     *gomock.Controller
	recorder *MockConfigManagerRecorder
}

var _ config.IManager = &MockConfigManager{}

// MockConfigManagerRecorder is the mock recorder for MockConfigManager.
type MockConfigManagerRecorder struct {
	mock *MockConfigManager
}

// NewMockConfigManager creates a new mock instance.
func NewMockConfigManager(ctrl *gomock.Controller) *MockConfigManager {
	mock := &MockConfigManager{ctrl: ctrl}
	mock.recorder = &MockConfigManagerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigManager) EXPECT() *MockConfigManagerRecorder {
	return m.recorder
}

// AddObserver mocks base method.
func (m *MockConfigManager) AddObserver(path string, callback config.IObserver) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddObserver", path, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddObserver indicates an expected call of AddObserver.
func (mr *MockConfigManagerRecorder) AddObserver(path, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddObserver", reflect.TypeOf((*MockConfigManager)(nil).AddObserver), path, callback)
}

// AddSource mocks base method.
func (m *MockConfigManager) AddSource(id string, priority int, src config.ISource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSource", id, priority, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSource indicates an expected call of AddSource.
func (mr *MockConfigManagerRecorder) AddSource(id, priority, src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSource", reflect.TypeOf((*MockConfigManager)(nil).AddSource), id, priority, src)
}

// Bool mocks base method.
func (m *MockConfigManager) Bool(path string, def ...bool) (bool, error) {
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
func (mr *MockConfigManagerRecorder) Bool(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockConfigManager)(nil).Bool), varargs...)
}

// Close mocks base method.
func (m *MockConfigManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockConfigManagerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConfigManager)(nil).Close))
}

// Config mocks base method.
func (m *MockConfigManager) Config(path string, def ...config.Config) (config.IConfig, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Config", varargs...)
	ret0, _ := ret[0].(config.IConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Config.
func (mr *MockConfigManagerRecorder) Config(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockConfigManager)(nil).Config), varargs...)
}

// Entries mocks base method.
func (m *MockConfigManager) Entries() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entries")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Entries indicates an expected call of Entries.
func (mr *MockConfigManagerRecorder) Entries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entries", reflect.TypeOf((*MockConfigManager)(nil).Entries))
}

// Float mocks base method.
func (m *MockConfigManager) Float(path string, def ...float64) (float64, error) {
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

// Float indicates an expected call of Float.
func (mr *MockConfigManagerRecorder) Float(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float", reflect.TypeOf((*MockConfigManager)(nil).Float), varargs...)
}

// Get mocks base method.
func (m *MockConfigManager) Get(path string, def ...interface{}) (interface{}, error) {
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
func (mr *MockConfigManagerRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigManager)(nil).Get), varargs...)
}

// Has mocks base method.
func (m *MockConfigManager) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockConfigManagerRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockConfigManager)(nil).Has), path)
}

// HasObserver mocks base method.
func (m *MockConfigManager) HasObserver(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasObserver", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasObserver indicates an expected call of HasObserver.
func (mr *MockConfigManagerRecorder) HasObserver(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasObserver", reflect.TypeOf((*MockConfigManager)(nil).HasObserver), path)
}

// HasSource mocks base method.
func (m *MockConfigManager) HasSource(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSource", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSource indicates an expected call of HasSource.
func (mr *MockConfigManagerRecorder) HasSource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSource", reflect.TypeOf((*MockConfigManager)(nil).HasSource), id)
}

// Int mocks base method.
func (m *MockConfigManager) Int(path string, def ...int) (int, error) {
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
func (mr *MockConfigManagerRecorder) Int(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockConfigManager)(nil).Int), varargs...)
}

// List mocks base method.
func (m *MockConfigManager) List(path string, def ...[]interface{}) ([]interface{}, error) {
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
func (mr *MockConfigManagerRecorder) List(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockConfigManager)(nil).List), varargs...)
}

// Populate mocks base method.
func (m *MockConfigManager) Populate(path string, target interface{}, icase ...bool) (interface{}, error) {
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
func (mr *MockConfigManagerRecorder) Populate(path, target interface{}, icase ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path, target}, icase...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Populate", reflect.TypeOf((*MockConfigManager)(nil).Populate), varargs...)
}

// RemoveAllSources mocks base method.
func (m *MockConfigManager) RemoveAllSources() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllSources")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllSources indicates an expected call of RemoveAllSources.
func (mr *MockConfigManagerRecorder) RemoveAllSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllSources", reflect.TypeOf((*MockConfigManager)(nil).RemoveAllSources))
}

// RemoveObserver mocks base method.
func (m *MockConfigManager) RemoveObserver(path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveObserver", path)
}

// RemoveObserver indicates an expected call of RemoveObserver.
func (mr *MockConfigManagerRecorder) RemoveObserver(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveObserver", reflect.TypeOf((*MockConfigManager)(nil).RemoveObserver), path)
}

// RemoveSource mocks base method.
func (m *MockConfigManager) RemoveSource(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSource", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSource indicates an expected call of RemoveSource.
func (mr *MockConfigManagerRecorder) RemoveSource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSource", reflect.TypeOf((*MockConfigManager)(nil).RemoveSource), id)
}

// Source mocks base method.
func (m *MockConfigManager) Source(id string) (config.ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Source", id)
	ret0, _ := ret[0].(config.ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Source indicates an expected call of Source.
func (mr *MockConfigManagerRecorder) Source(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Source", reflect.TypeOf((*MockConfigManager)(nil).Source), id)
}

// SourcePriority mocks base method.
func (m *MockConfigManager) SourcePriority(id string, priority int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourcePriority", id, priority)
	ret0, _ := ret[0].(error)
	return ret0
}

// SourcePriority indicates an expected call of SourcePriority.
func (mr *MockConfigManagerRecorder) SourcePriority(id, priority interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourcePriority", reflect.TypeOf((*MockConfigManager)(nil).SourcePriority), id, priority)
}

// String mocks base method.
func (m *MockConfigManager) String(path string, def ...string) (string, error) {
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
func (mr *MockConfigManagerRecorder) String(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockConfigManager)(nil).String), varargs...)
}

//------------------------------------------------------------------------------
// Dialect
//------------------------------------------------------------------------------

// MockDialect is a mock an instance of Dialect interface.
type MockDialect struct {
	ctrl     *gomock.Controller
	recorder *MockDialectRecorder
}

var _ gorm.Dialector = &MockDialect{}

// MockDialectRecorder is the mock recorder for MockDialect.
type MockDialectRecorder struct {
	mock *MockDialect
}

// NewMockDialect creates a new mock instance.
func NewMockDialect(ctrl *gomock.Controller) *MockDialect {
	mock := &MockDialect{ctrl: ctrl}
	mock.recorder = &MockDialectRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialect) EXPECT() *MockDialectRecorder {
	return m.recorder
}

// BindVarTo mocks base method.
func (m *MockDialect) BindVarTo(arg0 clause.Writer, arg1 *gorm.Statement, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BindVarTo", arg0, arg1, arg2)
}

// BindVarTo indicates an expected call of BindVarTo.
func (mr *MockDialectRecorder) BindVarTo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindVarTo", reflect.TypeOf((*MockDialect)(nil).BindVarTo), arg0, arg1, arg2)
}

// DataTypeOf mocks base method.
func (m *MockDialect) DataTypeOf(arg0 *schema.Field) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataTypeOf", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DataTypeOf indicates an expected call of DataTypeOf.
func (mr *MockDialectRecorder) DataTypeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataTypeOf", reflect.TypeOf((*MockDialect)(nil).DataTypeOf), arg0)
}

// DefaultValueOf mocks base method.
func (m *MockDialect) DefaultValueOf(arg0 *schema.Field) clause.Expression {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultValueOf", arg0)
	ret0, _ := ret[0].(clause.Expression)
	return ret0
}

// DefaultValueOf indicates an expected call of DefaultValueOf.
func (mr *MockDialectRecorder) DefaultValueOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultValueOf", reflect.TypeOf((*MockDialect)(nil).DefaultValueOf), arg0)
}

// Explain mocks base method.
func (m *MockDialect) Explain(arg0 string, arg1 ...interface{}) string {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Explain", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// Explain indicates an expected call of Explain.
func (mr *MockDialectRecorder) Explain(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Explain", reflect.TypeOf((*MockDialect)(nil).Explain), varargs...)
}

// Initialize mocks base method.
func (m *MockDialect) Initialize(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockDialectRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockDialect)(nil).Initialize), arg0)
}

// Migrator mocks base method.
func (m *MockDialect) Migrator(arg0 *gorm.DB) gorm.Migrator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrator", arg0)
	ret0, _ := ret[0].(gorm.Migrator)
	return ret0
}

// Migrator indicates an expected call of Migrator.
func (mr *MockDialectRecorder) Migrator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrator", reflect.TypeOf((*MockDialect)(nil).Migrator), arg0)
}

// Name mocks base method.
func (m *MockDialect) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockDialectRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDialect)(nil).Name))
}

// QuoteTo mocks base method.
func (m *MockDialect) QuoteTo(arg0 clause.Writer, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "QuoteTo", arg0, arg1)
}

// QuoteTo indicates an expected call of QuoteTo.
func (mr *MockDialectRecorder) QuoteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuoteTo", reflect.TypeOf((*MockDialect)(nil).QuoteTo), arg0, arg1)
}

//------------------------------------------------------------------------------
// Dialect Strategy
//------------------------------------------------------------------------------

// MockDialectStrategy is a mock an instance of IDialectStrategy interface.
type MockDialectStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockDialectStrategyRecorder
}

var _ IDialectStrategy = &MockDialectStrategy{}

// MockDialectStrategyRecorder is the mock recorder for MockDialectStrategy.
type MockDialectStrategyRecorder struct {
	mock *MockDialectStrategy
}

// NewMockDialectStrategy creates a new mock instance.
func NewMockDialectStrategy(ctrl *gomock.Controller) *MockDialectStrategy {
	mock := &MockDialectStrategy{ctrl: ctrl}
	mock.recorder = &MockDialectStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialectStrategy) EXPECT() *MockDialectStrategyRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockDialectStrategy) Accept(cfg config.IConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockDialectStrategyRecorder) Accept(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockDialectStrategy)(nil).Accept), cfg)
}

// Get mocks base method.
func (m *MockDialectStrategy) Get(cfg config.IConfig) (gorm.Dialector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", cfg)
	ret0, _ := ret[0].(gorm.Dialector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDialectStrategyRecorder) Get(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDialectStrategy)(nil).Get), cfg)
}

//------------------------------------------------------------------------------
// Dialect Factory
//------------------------------------------------------------------------------

// MockDialectFactory is a mock an instance of IDialectFactory interface.
type MockDialectFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDialectFactoryRecorder
}

var _ IDialectFactory = &MockDialectFactory{}

// MockDialectFactoryRecorder is the mock recorder for MockDialectFactory.
type MockDialectFactoryRecorder struct {
	mock *MockDialectFactory
}

// NewMockDialectFactory creates a new mock instance.
func NewMockDialectFactory(ctrl *gomock.Controller) *MockDialectFactory {
	mock := &MockDialectFactory{ctrl: ctrl}
	mock.recorder = &MockDialectFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialectFactory) EXPECT() *MockDialectFactoryRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDialectFactory) Get(cfg config.IConfig) (gorm.Dialector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", cfg)
	ret0, _ := ret[0].(gorm.Dialector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDialectFactoryRecorder) Get(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDialectFactory)(nil).Get), cfg)
}

// Register mocks base method.
func (m *MockDialectFactory) Register(strategy IDialectStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockDialectFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockDialectFactory)(nil).Register), strategy)
}
