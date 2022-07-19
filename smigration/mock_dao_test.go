package smigration

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockDao is a mock of IDao interface.
type MockDao struct {
	ctrl     *gomock.Controller
	recorder *MockDaoRecorder
}

// MockDaoRecorder is the mock recorder for MockDao.
type MockDaoRecorder struct {
	mock *MockDao
}

// NewMockDao creates a new mock instance.
func NewMockDao(ctrl *gomock.Controller) *MockDao {
	mock := &MockDao{ctrl: ctrl}
	mock.recorder = &MockDaoRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDao) EXPECT() *MockDaoRecorder {
	return m.recorder
}

// Down mocks base method.
func (m *MockDao) Down(last Record) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Down", last)
	ret0, _ := ret[0].(error)
	return ret0
}

// Down indicates an expected call of Down.
func (mr *MockDaoRecorder) Down(last interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Down", reflect.TypeOf((*MockDao)(nil).Down), last)
}

// Last mocks base method.
func (m *MockDao) Last() (Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Last indicates an expected call of Last.
func (mr *MockDaoRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockDao)(nil).Last))
}

// Up mocks base method.
func (m *MockDao) Up(version uint64) (Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Up", version)
	ret0, _ := ret[0].(Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Up indicates an expected call of Up.
func (mr *MockDaoRecorder) Up(version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Up", reflect.TypeOf((*MockDao)(nil).Up), version)
}
