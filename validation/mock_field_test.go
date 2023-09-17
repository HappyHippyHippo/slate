package validation

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
)

// MockFieldError is a mocked instance of FieldError interface.
type MockFieldError struct {
	ctrl     *gomock.Controller
	recorder *MockFieldErrorRecorder
}

var _ validator.FieldError = &MockFieldError{}

// MockFieldErrorRecorder is the mock recorder for MockFieldError.
type MockFieldErrorRecorder struct {
	mock *MockFieldError
}

// NewMockFieldError creates a new mock instance.
func NewMockFieldError(ctrl *gomock.Controller) *MockFieldError {
	mock := &MockFieldError{ctrl: ctrl}
	mock.recorder = &MockFieldErrorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFieldError) EXPECT() *MockFieldErrorRecorder {
	return m.recorder
}

// ActualTag mocks base method.
func (m *MockFieldError) ActualTag() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActualTag")
	ret0, _ := ret[0].(string)
	return ret0
}

// ActualTag indicates an expected call of ActualTag.
func (mr *MockFieldErrorRecorder) ActualTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActualTag", reflect.TypeOf((*MockFieldError)(nil).ActualTag))
}

// Error mocks base method.
func (m *MockFieldError) Error() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(string)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockFieldErrorRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockFieldError)(nil).Error))
}

// Field mocks base method.
func (m *MockFieldError) Field() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Field")
	ret0, _ := ret[0].(string)
	return ret0
}

// Field indicates an expected call of Field.
func (mr *MockFieldErrorRecorder) Field() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Field", reflect.TypeOf((*MockFieldError)(nil).Field))
}

// Kind mocks base method.
func (m *MockFieldError) Kind() reflect.Kind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(reflect.Kind)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockFieldErrorRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockFieldError)(nil).Kind))
}

// Namespace mocks base method.
func (m *MockFieldError) Namespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Namespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// Namespace indicates an expected call of Namespace.
func (mr *MockFieldErrorRecorder) Namespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockFieldError)(nil).Namespace))
}

// Param mocks base method.
func (m *MockFieldError) Param() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param")
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param.
func (mr *MockFieldErrorRecorder) Param() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockFieldError)(nil).Param))
}

// StructField mocks base method.
func (m *MockFieldError) StructField() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StructField")
	ret0, _ := ret[0].(string)
	return ret0
}

// StructField indicates an expected call of StructField.
func (mr *MockFieldErrorRecorder) StructField() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StructField", reflect.TypeOf((*MockFieldError)(nil).StructField))
}

// StructNamespace mocks base method.
func (m *MockFieldError) StructNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StructNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// StructNamespace indicates an expected call of StructNamespace.
func (mr *MockFieldErrorRecorder) StructNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StructNamespace", reflect.TypeOf((*MockFieldError)(nil).StructNamespace))
}

// Tag mocks base method.
func (m *MockFieldError) Tag() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(string)
	return ret0
}

// Tag indicates an expected call of Tag.
func (mr *MockFieldErrorRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockFieldError)(nil).Tag))
}

// Translate mocks base method.
func (m *MockFieldError) Translate(arg0 ut.Translator) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Translate", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Translate indicates an expected call of Translate.
func (mr *MockFieldErrorRecorder) Translate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockFieldError)(nil).Translate), arg0)
}

// Type mocks base method.
func (m *MockFieldError) Type() reflect.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(reflect.Type)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockFieldErrorRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockFieldError)(nil).Type))
}

// Value mocks base method.
func (m *MockFieldError) Value() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockFieldErrorRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockFieldError)(nil).Value))
}
