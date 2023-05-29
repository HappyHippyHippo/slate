package rest

import (
	"net/http"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockHRequester is a mock instance of HTTPClient interface.
type MockHRequester struct {
	ctrl     *gomock.Controller
	recorder *MockRequesterRecorder
}

var _ requester = &MockHRequester{}

// MockRequesterRecorder is the mock recorder for MockHRequester.
type MockRequesterRecorder struct {
	mock *MockHRequester
}

// NewMockRequester creates a new mock instance.
func NewMockRequester(ctrl *gomock.Controller) *MockHRequester {
	mock := &MockHRequester{ctrl: ctrl}
	mock.recorder = &MockRequesterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHRequester) EXPECT() *MockRequesterRecorder {
	return m.recorder
}

// Do mock base method.
func (m *MockHRequester) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicate an expected call of Do.
func (mr *MockRequesterRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHRequester)(nil).Do), req)
}
