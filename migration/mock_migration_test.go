package migration

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockMigration is a mocked instance of Migration interface.
type MockMigration struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationRecorder
}

var _ Migration = &MockMigration{}

// MockMigrationRecorder is the mock recorder for MockMigration.
type MockMigrationRecorder struct {
	mock *MockMigration
}

// NewMockMigration creates a new mock instance.
func NewMockMigration(ctrl *gomock.Controller) *MockMigration {
	mock := &MockMigration{ctrl: ctrl}
	mock.recorder = &MockMigrationRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigration) EXPECT() *MockMigrationRecorder {
	return m.recorder
}

// Version mocks base method.
func (m *MockMigration) Version() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockMigrationRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockMigration)(nil).Version))
}

// Up mocks base method.
func (m *MockMigration) Up() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Up")
	ret0, _ := ret[0].(error)
	return ret0
}

// Up indicates an expected call of Up.
func (mr *MockMigrationRecorder) Up() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Up", reflect.TypeOf((*MockMigration)(nil).Up))
}

// Down mocks base method.
func (m *MockMigration) Down() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Down")
	ret0, _ := ret[0].(error)
	return ret0
}

// Down indicates an expected call of Down.
func (mr *MockMigrationRecorder) Down() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Down", reflect.TypeOf((*MockMigration)(nil).Down))
}
