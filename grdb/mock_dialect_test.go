package grdb

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// MockDialect is a mock of Dialect interface.
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
