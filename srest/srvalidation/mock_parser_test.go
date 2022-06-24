package srvalidation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/srest/srenvelope"
)

// MockParser is a mock of Parser interface.
type MockParser struct {
	ctrl     *gomock.Controller
	recorder *MockParserRecorder
}

// MockParserRecorder is the mock recorder for MockParser.
type MockParserRecorder struct {
	mock *MockParser
}

// NewMockParser creates a new mock instance.
func NewMockParser(ctrl *gomock.Controller) *MockParser {
	mock := &MockParser{ctrl: ctrl}
	mock.recorder = &MockParserRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParser) EXPECT() *MockParserRecorder {
	return m.recorder
}

// AddError mocks base method.
func (m *MockParser) AddError(err string, code int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddError", err, code)
}

// AddError indicates an expected call of AddError.
func (mr *MockParserRecorder) AddError(err, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddError", reflect.TypeOf((*MockParser)(nil).AddError), err, code)
}

// Parse mocks base method.
func (m *MockParser) Parse(val interface{}, errs validator.ValidationErrors) (*srenvelope.Envelope, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", val, errs)
	ret0, _ := ret[0].(*srenvelope.Envelope)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockParserRecorder) Parse(val, errs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockParser)(nil).Parse), val, errs)
}
