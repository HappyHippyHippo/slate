package migration

import (
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

// MockDialector is a mock instance of gorm.Dialector interface.
type MockDialector struct {
	ctrl     *gomock.Controller
	recorder *MockDialectorRecorder
}

var _ gorm.Dialector = &MockDialector{}

// MockDialectorRecorder is the mock recorder for MockDialector.
type MockDialectorRecorder struct {
	mock *MockDialector
}

// NewMockDialector creates a new mock instance.
func NewMockDialector(ctrl *gomock.Controller) *MockDialector {
	mock := &MockDialector{ctrl: ctrl}
	mock.recorder = &MockDialectorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDialector) EXPECT() *MockDialectorRecorder {
	return m.recorder
}

// BindVarTo mocks base method.
func (m *MockDialector) BindVarTo(arg0 clause.Writer, arg1 *gorm.Statement, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BindVarTo", arg0, arg1, arg2)
}

// BindVarTo indicates an expected call of BindVarTo.
func (mr *MockDialectorRecorder) BindVarTo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindVarTo", reflect.TypeOf((*MockDialector)(nil).BindVarTo), arg0, arg1, arg2)
}

// DataTypeOf mocks base method.
func (m *MockDialector) DataTypeOf(arg0 *schema.Field) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataTypeOf", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DataTypeOf indicates an expected call of DataTypeOf.
func (mr *MockDialectorRecorder) DataTypeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataTypeOf", reflect.TypeOf((*MockDialector)(nil).DataTypeOf), arg0)
}

// DefaultValueOf mocks base method.
func (m *MockDialector) DefaultValueOf(arg0 *schema.Field) clause.Expression {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultValueOf", arg0)
	ret0, _ := ret[0].(clause.Expression)
	return ret0
}

// DefaultValueOf indicates an expected call of DefaultValueOf.
func (mr *MockDialectorRecorder) DefaultValueOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultValueOf", reflect.TypeOf((*MockDialector)(nil).DefaultValueOf), arg0)
}

// Explain mocks base method.
func (m *MockDialector) Explain(arg0 string, arg1 ...interface{}) string {
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
func (mr *MockDialectorRecorder) Explain(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Explain", reflect.TypeOf((*MockDialector)(nil).Explain), varargs...)
}

// Initialize mocks base method.
func (m *MockDialector) Initialize(arg0 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockDialectorRecorder) Initialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockDialector)(nil).Initialize), arg0)
}

// Migrator mocks base method.
func (m *MockDialector) Migrator(arg0 *gorm.DB) gorm.Migrator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Migrator", arg0)
	ret0, _ := ret[0].(gorm.Migrator)
	return ret0
}

// Migrator indicates an expected call of Migrator.
func (mr *MockDialectorRecorder) Migrator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrator", reflect.TypeOf((*MockDialector)(nil).Migrator), arg0)
}

// Name mocks base method.
func (m *MockDialector) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockDialectorRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockDialector)(nil).Name))
}

// QuoteTo mocks base method.
func (m *MockDialector) QuoteTo(arg0 clause.Writer, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "QuoteTo", arg0, arg1)
}

// QuoteTo indicates an expected call of QuoteTo.
func (mr *MockDialectorRecorder) QuoteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuoteTo", reflect.TypeOf((*MockDialector)(nil).QuoteTo), arg0, arg1)
}
