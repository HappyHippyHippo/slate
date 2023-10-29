package slate

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// ----------------------------------------------------------------------------
// gorm.Migrator
// ----------------------------------------------------------------------------

// MockGormMigrator is a mock instance of Migrator interface.
type MockGormMigrator struct {
	ctrl     *gomock.Controller
	recorder *MockGormMigratorRecorder
}

var _ gorm.Migrator = &MockGormMigrator{}

// MockGormMigratorRecorder is the mock recorder for MockGormMigrator.
type MockGormMigratorRecorder struct {
	mock *MockGormMigrator
}

// NewMockGormMigrator creates a new mock instance.
func NewMockGormMigrator(ctrl *gomock.Controller) *MockGormMigrator {
	mock := &MockGormMigrator{ctrl: ctrl}
	mock.recorder = &MockGormMigratorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGormMigrator) EXPECT() *MockGormMigratorRecorder {
	return m.recorder
}

// AddColumn mocks base method.
func (m *MockGormMigrator) AddColumn(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddColumn indicates an expected call of AddColumn.
func (mr *MockGormMigratorRecorder) AddColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddColumn", reflect.TypeOf((*MockGormMigrator)(nil).AddColumn), arg0, arg1)
}

// AlterColumn mocks base method.
func (m *MockGormMigrator) AlterColumn(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlterColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlterColumn indicates an expected call of AlterColumn.
func (mr *MockGormMigratorRecorder) AlterColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlterColumn", reflect.TypeOf((*MockGormMigrator)(nil).AlterColumn), arg0, arg1)
}

// AutoMigrate mocks base method.
func (m *MockGormMigrator) AutoMigrate(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AutoMigrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoMigrate indicates an expected call of AutoMigrate.
func (mr *MockGormMigratorRecorder) AutoMigrate(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockGormMigrator)(nil).AutoMigrate), arg0...)
}

// ColumnTypes mocks base method.
func (m *MockGormMigrator) ColumnTypes(arg0 interface{}) ([]gorm.ColumnType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ColumnTypes", arg0)
	ret0, _ := ret[0].([]gorm.ColumnType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ColumnTypes indicates an expected call of ColumnTypes.
func (mr *MockGormMigratorRecorder) ColumnTypes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColumnTypes", reflect.TypeOf((*MockGormMigrator)(nil).ColumnTypes), arg0)
}

// CreateConstraint mocks base method.
func (m *MockGormMigrator) CreateConstraint(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConstraint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateConstraint indicates an expected call of CreateConstraint.
func (mr *MockGormMigratorRecorder) CreateConstraint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConstraint", reflect.TypeOf((*MockGormMigrator)(nil).CreateConstraint), arg0, arg1)
}

// CreateIndex mocks base method.
func (m *MockGormMigrator) CreateIndex(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndex", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIndex indicates an expected call of CreateIndex.
func (mr *MockGormMigratorRecorder) CreateIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndex", reflect.TypeOf((*MockGormMigrator)(nil).CreateIndex), arg0, arg1)
}

// CreateTable mocks base method.
func (m *MockGormMigrator) CreateTable(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTable", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTable indicates an expected call of CreateTable.
func (mr *MockGormMigratorRecorder) CreateTable(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockGormMigrator)(nil).CreateTable), arg0...)
}

// CreateView mocks base method.
func (m *MockGormMigrator) CreateView(arg0 string, arg1 gorm.ViewOption) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateView", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateView indicates an expected call of CreateView.
func (mr *MockGormMigratorRecorder) CreateView(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateView", reflect.TypeOf((*MockGormMigrator)(nil).CreateView), arg0, arg1)
}

// CurrentDatabase mocks base method.
func (m *MockGormMigrator) CurrentDatabase() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentDatabase")
	ret0, _ := ret[0].(string)
	return ret0
}

// CurrentDatabase indicates an expected call of CurrentDatabase.
func (mr *MockGormMigratorRecorder) CurrentDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentDatabase", reflect.TypeOf((*MockGormMigrator)(nil).CurrentDatabase))
}

// DropColumn mocks base method.
func (m *MockGormMigrator) DropColumn(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropColumn indicates an expected call of DropColumn.
func (mr *MockGormMigratorRecorder) DropColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropColumn", reflect.TypeOf((*MockGormMigrator)(nil).DropColumn), arg0, arg1)
}

// DropConstraint mocks base method.
func (m *MockGormMigrator) DropConstraint(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropConstraint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropConstraint indicates an expected call of DropConstraint.
func (mr *MockGormMigratorRecorder) DropConstraint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropConstraint", reflect.TypeOf((*MockGormMigrator)(nil).DropConstraint), arg0, arg1)
}

// DropIndex mocks base method.
func (m *MockGormMigrator) DropIndex(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropIndex", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndex indicates an expected call of DropIndex.
func (mr *MockGormMigratorRecorder) DropIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndex", reflect.TypeOf((*MockGormMigrator)(nil).DropIndex), arg0, arg1)
}

// DropTable mocks base method.
func (m *MockGormMigrator) DropTable(arg0 ...interface{}) error {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropTable", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropTable indicates an expected call of DropTable.
func (mr *MockGormMigratorRecorder) DropTable(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropTable", reflect.TypeOf((*MockGormMigrator)(nil).DropTable), arg0...)
}

// DropView mocks base method.
func (m *MockGormMigrator) DropView(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropView", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropView indicates an expected call of DropView.
func (mr *MockGormMigratorRecorder) DropView(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropView", reflect.TypeOf((*MockGormMigrator)(nil).DropView), arg0)
}

// FullDataTypeOf mocks base method.
func (m *MockGormMigrator) FullDataTypeOf(arg0 *schema.Field) clause.Expr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullDataTypeOf", arg0)
	ret0, _ := ret[0].(clause.Expr)
	return ret0
}

// FullDataTypeOf indicates an expected call of FullDataTypeOf.
func (mr *MockGormMigratorRecorder) FullDataTypeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullDataTypeOf", reflect.TypeOf((*MockGormMigrator)(nil).FullDataTypeOf), arg0)
}

// GetIndexes mocks base method.
func (m *MockGormMigrator) GetIndexes(arg0 interface{}) ([]gorm.Index, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexes", arg0)
	ret0, _ := ret[0].([]gorm.Index)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndexes indicates an expected call of GetIndexes.
func (mr *MockGormMigratorRecorder) GetIndexes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexes", reflect.TypeOf((*MockGormMigrator)(nil).GetIndexes), arg0)
}

// GetTables mocks base method.
func (m *MockGormMigrator) GetTables() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTables")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTables indicates an expected call of GetTables.
func (mr *MockGormMigratorRecorder) GetTables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTables", reflect.TypeOf((*MockGormMigrator)(nil).GetTables))
}

// GetTypeAliases mocks base method.
func (m *MockGormMigrator) GetTypeAliases(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeAliases", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetTypeAliases indicates an expected call of GetTypeAliases.
func (mr *MockGormMigratorRecorder) GetTypeAliases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeAliases", reflect.TypeOf((*MockGormMigrator)(nil).GetTypeAliases), arg0)
}

// HasColumn mocks base method.
func (m *MockGormMigrator) HasColumn(arg0 interface{}, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasColumn", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasColumn indicates an expected call of HasColumn.
func (mr *MockGormMigratorRecorder) HasColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasColumn", reflect.TypeOf((*MockGormMigrator)(nil).HasColumn), arg0, arg1)
}

// HasConstraint mocks base method.
func (m *MockGormMigrator) HasConstraint(arg0 interface{}, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasConstraint", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasConstraint indicates an expected call of HasConstraint.
func (mr *MockGormMigratorRecorder) HasConstraint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasConstraint", reflect.TypeOf((*MockGormMigrator)(nil).HasConstraint), arg0, arg1)
}

// HasIndex mocks base method.
func (m *MockGormMigrator) HasIndex(arg0 interface{}, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasIndex", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasIndex indicates an expected call of HasIndex.
func (mr *MockGormMigratorRecorder) HasIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasIndex", reflect.TypeOf((*MockGormMigrator)(nil).HasIndex), arg0, arg1)
}

// HasTable mocks base method.
func (m *MockGormMigrator) HasTable(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasTable", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasTable indicates an expected call of HasTable.
func (mr *MockGormMigratorRecorder) HasTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTable", reflect.TypeOf((*MockGormMigrator)(nil).HasTable), arg0)
}

// MigrateColumn mocks base method.
func (m *MockGormMigrator) MigrateColumn(arg0 interface{}, arg1 *schema.Field, arg2 gorm.ColumnType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateColumn", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateColumn indicates an expected call of MigrateColumn.
func (mr *MockGormMigratorRecorder) MigrateColumn(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateColumn", reflect.TypeOf((*MockGormMigrator)(nil).MigrateColumn), arg0, arg1, arg2)
}

// RenameColumn mocks base method.
func (m *MockGormMigrator) RenameColumn(arg0 interface{}, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameColumn", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameColumn indicates an expected call of RenameColumn.
func (mr *MockGormMigratorRecorder) RenameColumn(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameColumn", reflect.TypeOf((*MockGormMigrator)(nil).RenameColumn), arg0, arg1, arg2)
}

// RenameIndex mocks base method.
func (m *MockGormMigrator) RenameIndex(arg0 interface{}, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameIndex indicates an expected call of RenameIndex.
func (mr *MockGormMigratorRecorder) RenameIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameIndex", reflect.TypeOf((*MockGormMigrator)(nil).RenameIndex), arg0, arg1, arg2)
}

// RenameTable mocks base method.
func (m *MockGormMigrator) RenameTable(arg0, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameTable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameTable indicates an expected call of RenameTable.
func (mr *MockGormMigratorRecorder) RenameTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameTable", reflect.TypeOf((*MockGormMigrator)(nil).RenameTable), arg0, arg1)
}

// TableType mocks base method.
func (m *MockGormMigrator) TableType(arg0 interface{}) (gorm.TableType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TableType", arg0)
	ret0, _ := ret[0].(gorm.TableType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TableType indicates an expected call of TableType.
func (mr *MockGormMigratorRecorder) TableType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TableType", reflect.TypeOf((*MockGormMigrator)(nil).TableType), arg0)
}

// ----------------------------------------------------------------------------
// gorm.Dialector
// ----------------------------------------------------------------------------

// MockGormDialector is a mock instance of Dialect interface.
type MockGormDialector struct {
	ctrl     *gomock.Controller
	recorder *MockGormDialectorRecorder
}

var _ gorm.Dialector = &MockGormDialector{}

// MockGormDialectorRecorder is the mock recorder for MockGormDialector.
type MockGormDialectorRecorder struct {
	mock *MockGormDialector
}

// NewMockGormDialector creates a new mock instance.
func NewMockGormDialector(ctrl *gomock.Controller) *MockGormDialector {
	mock := &MockGormDialector{ctrl: ctrl}
	mock.recorder = &MockGormDialectorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGormDialector) EXPECT() *MockGormDialectorRecorder {
	return m.recorder
}

// BindVarTo mocks base method.
func (m *MockGormDialector) BindVarTo(arg0 clause.Writer, arg1 *gorm.Statement, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BindVarTo", arg0, arg1, arg2)
}

// BindVarTo indicates an expected call of BindVarTo.
func (mr *MockGormDialectorRecorder) BindVarTo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindVarTo", reflect.TypeOf((*MockGormDialector)(nil).BindVarTo), arg0, arg1, arg2)
}

// DataTypeOf mocks base method.
func (m *MockGormDialector) DataTypeOf(arg0 *schema.Field) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataTypeOf", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DataTypeOf indicates an expected call of DataTypeOf.
func (mr *MockGormDialectorRecorder) DataTypeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataTypeOf", reflect.TypeOf((*MockGormDialector)(nil).DataTypeOf), arg0)
}

// DefaultValueOf mocks base method.
func (m *MockGormDialector) DefaultValueOf(arg0 *schema.Field) clause.Expression {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultValueOf", arg0)
	ret0, _ := ret[0].(clause.Expression)
	return ret0
}

// DefaultValueOf indicates an expected call of DefaultValueOf.
func (mr *MockGormDialectorRecorder) DefaultValueOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultValueOf", reflect.TypeOf((*MockGormDialector)(nil).DefaultValueOf), arg0)
}

// Explain mocks base method.
func (m *MockGormDialector) Explain(arg0 string, arg1 ...interface{}) string {
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
func (mr *MockGormDialectorRecorder) Explain(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Explain", reflect.TypeOf((*MockGormDialector)(nil).Explain), varargs...)
}

// Initialize mocks base method.
func (m *MockGormDialector) Initialize(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockGormDialectorRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockGormDialector)(nil).Initialize), arg0)
}

// Migrator mocks base method.
func (m *MockGormDialector) Migrator(arg0 *gorm.DB) gorm.Migrator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrator", arg0)
	ret0, _ := ret[0].(gorm.Migrator)
	return ret0
}

// Migrator indicates an expected call of Migrator.
func (mr *MockGormDialectorRecorder) Migrator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrator", reflect.TypeOf((*MockGormDialector)(nil).Migrator), arg0)
}

// Name mocks base method.
func (m *MockGormDialector) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockGormDialectorRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockGormDialector)(nil).Name))
}

// QuoteTo mocks base method.
func (m *MockGormDialector) QuoteTo(arg0 clause.Writer, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "QuoteTo", arg0, arg1)
}

// QuoteTo indicates an expected call of QuoteTo.
func (mr *MockGormDialectorRecorder) QuoteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuoteTo", reflect.TypeOf((*MockGormDialector)(nil).QuoteTo), arg0, arg1)
}

// ----------------------------------------------------------------------------
// RdbDialectCreator
// ----------------------------------------------------------------------------

// MockRdbDialectCreator is a mock instance of IDialectStrategy interface.
type MockRdbDialectCreator struct {
	ctrl     *gomock.Controller
	recorder *MockRdbDialectCreatorRecorder
}

var _ RdbDialectCreator = &MockRdbDialectCreator{}

// MockRdbDialectCreatorRecorder is the mock recorder for MockRdbDialectCreator.
type MockRdbDialectCreatorRecorder struct {
	mock *MockRdbDialectCreator
}

// NewMockRdbDialectCreator creates a new mock instance.
func NewMockRdbDialectCreator(ctrl *gomock.Controller) *MockRdbDialectCreator {
	mock := &MockRdbDialectCreator{ctrl: ctrl}
	mock.recorder = &MockRdbDialectCreatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRdbDialectCreator) EXPECT() *MockRdbDialectCreatorRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockRdbDialectCreator) Accept(config *ConfigPartial) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", config)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockRdbDialectCreatorRecorder) Accept(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockRdbDialectCreator)(nil).Accept), config)
}

// Create mocks base method.
func (m *MockRdbDialectCreator) Create(config *ConfigPartial) (gorm.Dialector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", config)
	ret0, _ := ret[0].(gorm.Dialector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRdbDialectCreatorRecorder) Create(config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRdbDialectCreator)(nil).Create), config)
}
