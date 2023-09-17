package validation

import (
	"reflect"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
	"github.com/golang/mock/gomock"
)

// MockTranslator is a mocked instance of Translator interface.
type MockTranslator struct {
	ctrl     *gomock.Controller
	recorder *MockTranslatorRecorder
}

// MockTranslatorRecorder is the mock recorder for MockTranslator.
type MockTranslatorRecorder struct {
	mock *MockTranslator
}

// NewMockTranslator creates a new mock instance.
func NewMockTranslator(ctrl *gomock.Controller) *MockTranslator {
	mock := &MockTranslator{ctrl: ctrl}
	mock.recorder = &MockTranslatorRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTranslator) EXPECT() *MockTranslatorRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTranslator) Add(arg0 interface{}, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockTranslatorRecorder) Add(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTranslator)(nil).Add), arg0, arg1, arg2)
}

// AddCardinal mocks base method.
func (m *MockTranslator) AddCardinal(arg0 interface{}, arg1 string, arg2 locales.PluralRule, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCardinal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCardinal indicates an expected call of AddCardinal.
func (mr *MockTranslatorRecorder) AddCardinal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCardinal", reflect.TypeOf((*MockTranslator)(nil).AddCardinal), arg0, arg1, arg2, arg3)
}

// AddOrdinal mocks base method.
func (m *MockTranslator) AddOrdinal(arg0 interface{}, arg1 string, arg2 locales.PluralRule, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrdinal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrdinal indicates an expected call of AddOrdinal.
func (mr *MockTranslatorRecorder) AddOrdinal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrdinal", reflect.TypeOf((*MockTranslator)(nil).AddOrdinal), arg0, arg1, arg2, arg3)
}

// AddRange mocks base method.
func (m *MockTranslator) AddRange(arg0 interface{}, arg1 string, arg2 locales.PluralRule, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRange indicates an expected call of AddRange.
func (mr *MockTranslatorRecorder) AddRange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRange", reflect.TypeOf((*MockTranslator)(nil).AddRange), arg0, arg1, arg2, arg3)
}

// C mocks base method.
func (m *MockTranslator) C(arg0 interface{}, arg1 float64, arg2 uint64, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "C", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// C indicates an expected call of C.
func (mr *MockTranslatorRecorder) C(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockTranslator)(nil).C), arg0, arg1, arg2, arg3)
}

// CardinalPluralRule mocks base method.
func (m *MockTranslator) CardinalPluralRule(arg0 float64, arg1 uint64) locales.PluralRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CardinalPluralRule", arg0, arg1)
	ret0, _ := ret[0].(locales.PluralRule)
	return ret0
}

// CardinalPluralRule indicates an expected call of CardinalPluralRule.
func (mr *MockTranslatorRecorder) CardinalPluralRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CardinalPluralRule", reflect.TypeOf((*MockTranslator)(nil).CardinalPluralRule), arg0, arg1)
}

// FmtAccounting mocks base method.
func (m *MockTranslator) FmtAccounting(arg0 float64, arg1 uint64, arg2 currency.Type) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtAccounting", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtAccounting indicates an expected call of FmtAccounting.
func (mr *MockTranslatorRecorder) FmtAccounting(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtAccounting", reflect.TypeOf((*MockTranslator)(nil).FmtAccounting), arg0, arg1, arg2)
}

// FmtCurrency mocks base method.
func (m *MockTranslator) FmtCurrency(arg0 float64, arg1 uint64, arg2 currency.Type) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtCurrency", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtCurrency indicates an expected call of FmtCurrency.
func (mr *MockTranslatorRecorder) FmtCurrency(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtCurrency", reflect.TypeOf((*MockTranslator)(nil).FmtCurrency), arg0, arg1, arg2)
}

// FmtDateFull mocks base method.
func (m *MockTranslator) FmtDateFull(arg0 time.Time) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtDateFull", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtDateFull indicates an expected call of FmtDateFull.
func (mr *MockTranslatorRecorder) FmtDateFull(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtDateFull", reflect.TypeOf((*MockTranslator)(nil).FmtDateFull), arg0)
}

// FmtDateLong mocks base method.
func (m *MockTranslator) FmtDateLong(arg0 time.Time) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtDateLong", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtDateLong indicates an expected call of FmtDateLong.
func (mr *MockTranslatorRecorder) FmtDateLong(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtDateLong", reflect.TypeOf((*MockTranslator)(nil).FmtDateLong), arg0)
}

// FmtDateMedium mocks base method.
func (m *MockTranslator) FmtDateMedium(arg0 time.Time) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtDateMedium", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtDateMedium indicates an expected call of FmtDateMedium.
func (mr *MockTranslatorRecorder) FmtDateMedium(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtDateMedium", reflect.TypeOf((*MockTranslator)(nil).FmtDateMedium), arg0)
}

// FmtDateShort mocks base method.
func (m *MockTranslator) FmtDateShort(arg0 time.Time) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtDateShort", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtDateShort indicates an expected call of FmtDateShort.
func (mr *MockTranslatorRecorder) FmtDateShort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtDateShort", reflect.TypeOf((*MockTranslator)(nil).FmtDateShort), arg0)
}

// FmtNumber mocks base method.
func (m *MockTranslator) FmtNumber(arg0 float64, arg1 uint64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtNumber", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtNumber indicates an expected call of FmtNumber.
func (mr *MockTranslatorRecorder) FmtNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtNumber", reflect.TypeOf((*MockTranslator)(nil).FmtNumber), arg0, arg1)
}

// FmtPercent mocks base method.
func (m *MockTranslator) FmtPercent(arg0 float64, arg1 uint64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtPercent", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtPercent indicates an expected call of FmtPercent.
func (mr *MockTranslatorRecorder) FmtPercent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtPercent", reflect.TypeOf((*MockTranslator)(nil).FmtPercent), arg0, arg1)
}

// FmtTimeFull mocks base method.
func (m *MockTranslator) FmtTimeFull(arg0 time.Time) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtTimeFull", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtTimeFull indicates an expected call of FmtTimeFull.
func (mr *MockTranslatorRecorder) FmtTimeFull(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtTimeFull", reflect.TypeOf((*MockTranslator)(nil).FmtTimeFull), arg0)
}

// FmtTimeLong mocks base method.
func (m *MockTranslator) FmtTimeLong(arg0 time.Time) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtTimeLong", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtTimeLong indicates an expected call of FmtTimeLong.
func (mr *MockTranslatorRecorder) FmtTimeLong(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtTimeLong", reflect.TypeOf((*MockTranslator)(nil).FmtTimeLong), arg0)
}

// FmtTimeMedium mocks base method.
func (m *MockTranslator) FmtTimeMedium(arg0 time.Time) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtTimeMedium", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtTimeMedium indicates an expected call of FmtTimeMedium.
func (mr *MockTranslatorRecorder) FmtTimeMedium(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtTimeMedium", reflect.TypeOf((*MockTranslator)(nil).FmtTimeMedium), arg0)
}

// FmtTimeShort mocks base method.
func (m *MockTranslator) FmtTimeShort(arg0 time.Time) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FmtTimeShort", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FmtTimeShort indicates an expected call of FmtTimeShort.
func (mr *MockTranslatorRecorder) FmtTimeShort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FmtTimeShort", reflect.TypeOf((*MockTranslator)(nil).FmtTimeShort), arg0)
}

// Locale mocks base method.
func (m *MockTranslator) Locale() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Locale")
	ret0, _ := ret[0].(string)
	return ret0
}

// Locale indicates an expected call of Locale.
func (mr *MockTranslatorRecorder) Locale() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Locale", reflect.TypeOf((*MockTranslator)(nil).Locale))
}

// MonthAbbreviated mocks base method.
func (m *MockTranslator) MonthAbbreviated(arg0 time.Month) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonthAbbreviated", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// MonthAbbreviated indicates an expected call of MonthAbbreviated.
func (mr *MockTranslatorRecorder) MonthAbbreviated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthAbbreviated", reflect.TypeOf((*MockTranslator)(nil).MonthAbbreviated), arg0)
}

// MonthNarrow mocks base method.
func (m *MockTranslator) MonthNarrow(arg0 time.Month) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonthNarrow", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// MonthNarrow indicates an expected call of MonthNarrow.
func (mr *MockTranslatorRecorder) MonthNarrow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthNarrow", reflect.TypeOf((*MockTranslator)(nil).MonthNarrow), arg0)
}

// MonthWide mocks base method.
func (m *MockTranslator) MonthWide(arg0 time.Month) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonthWide", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// MonthWide indicates an expected call of MonthWide.
func (mr *MockTranslatorRecorder) MonthWide(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthWide", reflect.TypeOf((*MockTranslator)(nil).MonthWide), arg0)
}

// MonthsAbbreviated mocks base method.
func (m *MockTranslator) MonthsAbbreviated() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonthsAbbreviated")
	ret0, _ := ret[0].([]string)
	return ret0
}

// MonthsAbbreviated indicates an expected call of MonthsAbbreviated.
func (mr *MockTranslatorRecorder) MonthsAbbreviated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthsAbbreviated", reflect.TypeOf((*MockTranslator)(nil).MonthsAbbreviated))
}

// MonthsNarrow mocks base method.
func (m *MockTranslator) MonthsNarrow() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonthsNarrow")
	ret0, _ := ret[0].([]string)
	return ret0
}

// MonthsNarrow indicates an expected call of MonthsNarrow.
func (mr *MockTranslatorRecorder) MonthsNarrow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthsNarrow", reflect.TypeOf((*MockTranslator)(nil).MonthsNarrow))
}

// MonthsWide mocks base method.
func (m *MockTranslator) MonthsWide() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MonthsWide")
	ret0, _ := ret[0].([]string)
	return ret0
}

// MonthsWide indicates an expected call of MonthsWide.
func (mr *MockTranslatorRecorder) MonthsWide() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MonthsWide", reflect.TypeOf((*MockTranslator)(nil).MonthsWide))
}

// O mocks base method.
func (m *MockTranslator) O(arg0 interface{}, arg1 float64, arg2 uint64, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "O", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// O indicates an expected call of O.
func (mr *MockTranslatorRecorder) O(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "O", reflect.TypeOf((*MockTranslator)(nil).O), arg0, arg1, arg2, arg3)
}

// OrdinalPluralRule mocks base method.
func (m *MockTranslator) OrdinalPluralRule(arg0 float64, arg1 uint64) locales.PluralRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrdinalPluralRule", arg0, arg1)
	ret0, _ := ret[0].(locales.PluralRule)
	return ret0
}

// OrdinalPluralRule indicates an expected call of OrdinalPluralRule.
func (mr *MockTranslatorRecorder) OrdinalPluralRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrdinalPluralRule", reflect.TypeOf((*MockTranslator)(nil).OrdinalPluralRule), arg0, arg1)
}

// PluralsCardinal mocks base method.
func (m *MockTranslator) PluralsCardinal() []locales.PluralRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PluralsCardinal")
	ret0, _ := ret[0].([]locales.PluralRule)
	return ret0
}

// PluralsCardinal indicates an expected call of PluralsCardinal.
func (mr *MockTranslatorRecorder) PluralsCardinal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PluralsCardinal", reflect.TypeOf((*MockTranslator)(nil).PluralsCardinal))
}

// PluralsOrdinal mocks base method.
func (m *MockTranslator) PluralsOrdinal() []locales.PluralRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PluralsOrdinal")
	ret0, _ := ret[0].([]locales.PluralRule)
	return ret0
}

// PluralsOrdinal indicates an expected call of PluralsOrdinal.
func (mr *MockTranslatorRecorder) PluralsOrdinal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PluralsOrdinal", reflect.TypeOf((*MockTranslator)(nil).PluralsOrdinal))
}

// PluralsRange mocks base method.
func (m *MockTranslator) PluralsRange() []locales.PluralRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PluralsRange")
	ret0, _ := ret[0].([]locales.PluralRule)
	return ret0
}

// PluralsRange indicates an expected call of PluralsRange.
func (mr *MockTranslatorRecorder) PluralsRange() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PluralsRange", reflect.TypeOf((*MockTranslator)(nil).PluralsRange))
}

// R mocks base method.
func (m *MockTranslator) R(arg0 interface{}, arg1 float64, arg2 uint64, arg3 float64, arg4 uint64, arg5, arg6 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "R", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// R indicates an expected call of R.
func (mr *MockTranslatorRecorder) R(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "R", reflect.TypeOf((*MockTranslator)(nil).R), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// RangePluralRule mocks base method.
func (m *MockTranslator) RangePluralRule(arg0 float64, arg1 uint64, arg2 float64, arg3 uint64) locales.PluralRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RangePluralRule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(locales.PluralRule)
	return ret0
}

// RangePluralRule indicates an expected call of RangePluralRule.
func (mr *MockTranslatorRecorder) RangePluralRule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangePluralRule", reflect.TypeOf((*MockTranslator)(nil).RangePluralRule), arg0, arg1, arg2, arg3)
}

// T mocks base method.
func (m *MockTranslator) T(arg0 interface{}, arg1 ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "T", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// T indicates an expected call of T.
func (mr *MockTranslatorRecorder) T(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "T", reflect.TypeOf((*MockTranslator)(nil).T), varargs...)
}

// VerifyTranslations mocks base method.
func (m *MockTranslator) VerifyTranslations() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyTranslations")
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyTranslations indicates an expected call of VerifyTranslations.
func (mr *MockTranslatorRecorder) VerifyTranslations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyTranslations", reflect.TypeOf((*MockTranslator)(nil).VerifyTranslations))
}

// WeekdayAbbreviated mocks base method.
func (m *MockTranslator) WeekdayAbbreviated(arg0 time.Weekday) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeekdayAbbreviated", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// WeekdayAbbreviated indicates an expected call of WeekdayAbbreviated.
func (mr *MockTranslatorRecorder) WeekdayAbbreviated(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeekdayAbbreviated", reflect.TypeOf((*MockTranslator)(nil).WeekdayAbbreviated), arg0)
}

// WeekdayNarrow mocks base method.
func (m *MockTranslator) WeekdayNarrow(arg0 time.Weekday) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeekdayNarrow", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// WeekdayNarrow indicates an expected call of WeekdayNarrow.
func (mr *MockTranslatorRecorder) WeekdayNarrow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeekdayNarrow", reflect.TypeOf((*MockTranslator)(nil).WeekdayNarrow), arg0)
}

// WeekdayShort mocks base method.
func (m *MockTranslator) WeekdayShort(arg0 time.Weekday) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeekdayShort", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// WeekdayShort indicates an expected call of WeekdayShort.
func (mr *MockTranslatorRecorder) WeekdayShort(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeekdayShort", reflect.TypeOf((*MockTranslator)(nil).WeekdayShort), arg0)
}

// WeekdayWide mocks base method.
func (m *MockTranslator) WeekdayWide(arg0 time.Weekday) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeekdayWide", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// WeekdayWide indicates an expected call of WeekdayWide.
func (mr *MockTranslatorRecorder) WeekdayWide(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeekdayWide", reflect.TypeOf((*MockTranslator)(nil).WeekdayWide), arg0)
}

// WeekdaysAbbreviated mocks base method.
func (m *MockTranslator) WeekdaysAbbreviated() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeekdaysAbbreviated")
	ret0, _ := ret[0].([]string)
	return ret0
}

// WeekdaysAbbreviated indicates an expected call of WeekdaysAbbreviated.
func (mr *MockTranslatorRecorder) WeekdaysAbbreviated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeekdaysAbbreviated", reflect.TypeOf((*MockTranslator)(nil).WeekdaysAbbreviated))
}

// WeekdaysNarrow mocks base method.
func (m *MockTranslator) WeekdaysNarrow() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeekdaysNarrow")
	ret0, _ := ret[0].([]string)
	return ret0
}

// WeekdaysNarrow indicates an expected call of WeekdaysNarrow.
func (mr *MockTranslatorRecorder) WeekdaysNarrow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeekdaysNarrow", reflect.TypeOf((*MockTranslator)(nil).WeekdaysNarrow))
}

// WeekdaysShort mocks base method.
func (m *MockTranslator) WeekdaysShort() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeekdaysShort")
	ret0, _ := ret[0].([]string)
	return ret0
}

// WeekdaysShort indicates an expected call of WeekdaysShort.
func (mr *MockTranslatorRecorder) WeekdaysShort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeekdaysShort", reflect.TypeOf((*MockTranslator)(nil).WeekdaysShort))
}

// WeekdaysWide mocks base method.
func (m *MockTranslator) WeekdaysWide() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WeekdaysWide")
	ret0, _ := ret[0].([]string)
	return ret0
}

// WeekdaysWide indicates an expected call of WeekdaysWide.
func (mr *MockTranslatorRecorder) WeekdaysWide() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WeekdaysWide", reflect.TypeOf((*MockTranslator)(nil).WeekdaysWide))
}
