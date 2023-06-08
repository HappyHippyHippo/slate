package migration

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// MockMigrator is a mock instance of Migrator interface.
type MockMigrator struct {
	ctrl     *gomock.Controller
	recorder *MockMigratorRecorder
}

var _ gorm.Migrator = &MockMigrator{}

// MockMigratorRecorder is the mock recorder for MockMigrator.
type MockMigratorRecorder struct {
	mock *MockMigrator
}

// NewMockMigrator creates a new mock instance.
func NewMockMigrator(ctrl *gomock.Controller) *MockMigrator {
	mock := &MockMigrator{ctrl: ctrl}
	mock.recorder = &MockMigratorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrator) EXPECT() *MockMigratorRecorder {
	return m.recorder
}

// AddColumn mocks base method.
func (m *MockMigrator) AddColumn(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddColumn indicates an expected call of AddColumn.
func (mr *MockMigratorRecorder) AddColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddColumn", reflect.TypeOf((*MockMigrator)(nil).AddColumn), arg0, arg1)
}

// AlterColumn mocks base method.
func (m *MockMigrator) AlterColumn(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlterColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlterColumn indicates an expected call of AlterColumn.
func (mr *MockMigratorRecorder) AlterColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlterColumn", reflect.TypeOf((*MockMigrator)(nil).AlterColumn), arg0, arg1)
}

// AutoMigrate mocks base method.
func (m *MockMigrator) AutoMigrate(arg0 ...interface{}) error {
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
func (mr *MockMigratorRecorder) AutoMigrate(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoMigrate", reflect.TypeOf((*MockMigrator)(nil).AutoMigrate), arg0...)
}

// ColumnTypes mocks base method.
func (m *MockMigrator) ColumnTypes(arg0 interface{}) ([]gorm.ColumnType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ColumnTypes", arg0)
	ret0, _ := ret[0].([]gorm.ColumnType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ColumnTypes indicates an expected call of ColumnTypes.
func (mr *MockMigratorRecorder) ColumnTypes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ColumnTypes", reflect.TypeOf((*MockMigrator)(nil).ColumnTypes), arg0)
}

// CreateConstraint mocks base method.
func (m *MockMigrator) CreateConstraint(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConstraint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateConstraint indicates an expected call of CreateConstraint.
func (mr *MockMigratorRecorder) CreateConstraint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConstraint", reflect.TypeOf((*MockMigrator)(nil).CreateConstraint), arg0, arg1)
}

// CreateIndex mocks base method.
func (m *MockMigrator) CreateIndex(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndex", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIndex indicates an expected call of CreateIndex.
func (mr *MockMigratorRecorder) CreateIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndex", reflect.TypeOf((*MockMigrator)(nil).CreateIndex), arg0, arg1)
}

// CreateTable mocks base method.
func (m *MockMigrator) CreateTable(arg0 ...interface{}) error {
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
func (mr *MockMigratorRecorder) CreateTable(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockMigrator)(nil).CreateTable), arg0...)
}

// CreateView mocks base method.
func (m *MockMigrator) CreateView(arg0 string, arg1 gorm.ViewOption) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateView", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateView indicates an expected call of CreateView.
func (mr *MockMigratorRecorder) CreateView(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateView", reflect.TypeOf((*MockMigrator)(nil).CreateView), arg0, arg1)
}

// CurrentDatabase mocks base method.
func (m *MockMigrator) CurrentDatabase() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentDatabase")
	ret0, _ := ret[0].(string)
	return ret0
}

// CurrentDatabase indicates an expected call of CurrentDatabase.
func (mr *MockMigratorRecorder) CurrentDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentDatabase", reflect.TypeOf((*MockMigrator)(nil).CurrentDatabase))
}

// DropColumn mocks base method.
func (m *MockMigrator) DropColumn(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropColumn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropColumn indicates an expected call of DropColumn.
func (mr *MockMigratorRecorder) DropColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropColumn", reflect.TypeOf((*MockMigrator)(nil).DropColumn), arg0, arg1)
}

// DropConstraint mocks base method.
func (m *MockMigrator) DropConstraint(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropConstraint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropConstraint indicates an expected call of DropConstraint.
func (mr *MockMigratorRecorder) DropConstraint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropConstraint", reflect.TypeOf((*MockMigrator)(nil).DropConstraint), arg0, arg1)
}

// DropIndex mocks base method.
func (m *MockMigrator) DropIndex(arg0 interface{}, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropIndex", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndex indicates an expected call of DropIndex.
func (mr *MockMigratorRecorder) DropIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndex", reflect.TypeOf((*MockMigrator)(nil).DropIndex), arg0, arg1)
}

// DropTable mocks base method.
func (m *MockMigrator) DropTable(arg0 ...interface{}) error {
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
func (mr *MockMigratorRecorder) DropTable(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropTable", reflect.TypeOf((*MockMigrator)(nil).DropTable), arg0...)
}

// DropView mocks base method.
func (m *MockMigrator) DropView(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropView", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropView indicates an expected call of DropView.
func (mr *MockMigratorRecorder) DropView(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropView", reflect.TypeOf((*MockMigrator)(nil).DropView), arg0)
}

// FullDataTypeOf mocks base method.
func (m *MockMigrator) FullDataTypeOf(arg0 *schema.Field) clause.Expr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullDataTypeOf", arg0)
	ret0, _ := ret[0].(clause.Expr)
	return ret0
}

// FullDataTypeOf indicates an expected call of FullDataTypeOf.
func (mr *MockMigratorRecorder) FullDataTypeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullDataTypeOf", reflect.TypeOf((*MockMigrator)(nil).FullDataTypeOf), arg0)
}

// GetIndexes mocks base method.
func (m *MockMigrator) GetIndexes(arg0 interface{}) ([]gorm.Index, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIndexes", arg0)
	ret0, _ := ret[0].([]gorm.Index)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndexes indicates an expected call of GetIndexes.
func (mr *MockMigratorRecorder) GetIndexes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexes", reflect.TypeOf((*MockMigrator)(nil).GetIndexes), arg0)
}

// GetTables mocks base method.
func (m *MockMigrator) GetTables() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTables")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTables indicates an expected call of GetTables.
func (mr *MockMigratorRecorder) GetTables() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTables", reflect.TypeOf((*MockMigrator)(nil).GetTables))
}

// GetTypeAliases mocks base method.
func (m *MockMigrator) GetTypeAliases(arg0 string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeAliases", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetTypeAliases indicates an expected call of GetTypeAliases.
func (mr *MockMigratorRecorder) GetTypeAliases(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeAliases", reflect.TypeOf((*MockMigrator)(nil).GetTypeAliases), arg0)
}

// HasColumn mocks base method.
func (m *MockMigrator) HasColumn(arg0 interface{}, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasColumn", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasColumn indicates an expected call of HasColumn.
func (mr *MockMigratorRecorder) HasColumn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasColumn", reflect.TypeOf((*MockMigrator)(nil).HasColumn), arg0, arg1)
}

// HasConstraint mocks base method.
func (m *MockMigrator) HasConstraint(arg0 interface{}, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasConstraint", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasConstraint indicates an expected call of HasConstraint.
func (mr *MockMigratorRecorder) HasConstraint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasConstraint", reflect.TypeOf((*MockMigrator)(nil).HasConstraint), arg0, arg1)
}

// HasIndex mocks base method.
func (m *MockMigrator) HasIndex(arg0 interface{}, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasIndex", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasIndex indicates an expected call of HasIndex.
func (mr *MockMigratorRecorder) HasIndex(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasIndex", reflect.TypeOf((*MockMigrator)(nil).HasIndex), arg0, arg1)
}

// HasTable mocks base method.
func (m *MockMigrator) HasTable(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasTable", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasTable indicates an expected call of HasTable.
func (mr *MockMigratorRecorder) HasTable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasTable", reflect.TypeOf((*MockMigrator)(nil).HasTable), arg0)
}

// MigrateColumn mocks base method.
func (m *MockMigrator) MigrateColumn(arg0 interface{}, arg1 *schema.Field, arg2 gorm.ColumnType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MigrateColumn", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MigrateColumn indicates an expected call of MigrateColumn.
func (mr *MockMigratorRecorder) MigrateColumn(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MigrateColumn", reflect.TypeOf((*MockMigrator)(nil).MigrateColumn), arg0, arg1, arg2)
}

// RenameColumn mocks base method.
func (m *MockMigrator) RenameColumn(arg0 interface{}, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameColumn", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameColumn indicates an expected call of RenameColumn.
func (mr *MockMigratorRecorder) RenameColumn(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameColumn", reflect.TypeOf((*MockMigrator)(nil).RenameColumn), arg0, arg1, arg2)
}

// RenameIndex mocks base method.
func (m *MockMigrator) RenameIndex(arg0 interface{}, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameIndex", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameIndex indicates an expected call of RenameIndex.
func (mr *MockMigratorRecorder) RenameIndex(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameIndex", reflect.TypeOf((*MockMigrator)(nil).RenameIndex), arg0, arg1, arg2)
}

// RenameTable mocks base method.
func (m *MockMigrator) RenameTable(arg0, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameTable", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameTable indicates an expected call of RenameTable.
func (mr *MockMigratorRecorder) RenameTable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameTable", reflect.TypeOf((*MockMigrator)(nil).RenameTable), arg0, arg1)
}
