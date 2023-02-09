package config

import (
	"io"
	"net/http"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/happyhippyhippo/slate/trigger"
	"github.com/spf13/afero"

	"github.com/golang/mock/gomock"
)

//------------------------------------------------------------------------------
// Reader
//------------------------------------------------------------------------------

// MockReader is a mock instance of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderRecorder
}

var _ io.Reader = &MockReader{}

// MockReaderRecorder is the mock recorder for MockReader
type MockReaderRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReader) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReaderRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReader)(nil).Read), arg0)
}

// Close mocks base method
func (m *MockReader) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockReaderRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReader)(nil).Close))
}

//------------------------------------------------------------------------------
// Locker
//------------------------------------------------------------------------------

// MockLocker is a mock instance of Locker interface
type MockLocker struct {
	ctrl     *gomock.Controller
	recorder *MockLockerRecorder
}

var _ sync.Locker = &MockLocker{}

// MockLockerRecorder is the mock recorder for MockLocker
type MockLockerRecorder struct {
	mock *MockLocker
}

// NewMockLocker creates a new mock instance
func NewMockLocker(ctrl *gomock.Controller) *MockLocker {
	mock := &MockLocker{ctrl: ctrl}
	mock.recorder = &MockLockerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocker) EXPECT() *MockLockerRecorder {
	return m.recorder
}

// Lock mocks base method
func (m *MockLocker) Lock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock
func (mr *MockLockerRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockLocker)(nil).Lock))
}

// Unlock mocks base method
func (m *MockLocker) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock
func (mr *MockLockerRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockLocker)(nil).Unlock))
}

//------------------------------------------------------------------------------
// Trigger
//------------------------------------------------------------------------------

// MockTrigger is a mock of ITrigger interface.
type MockTrigger struct {
	ctrl     *gomock.Controller
	recorder *MockTriggerRecorder
}

var _ trigger.ITrigger = &MockTrigger{}

// MockTriggerRecorder is the mock recorder for MockTrigger.
type MockTriggerRecorder struct {
	mock *MockTrigger
}

// NewMockTrigger creates a new mock instance.
func NewMockTrigger(ctrl *gomock.Controller) *MockTrigger {
	mock := &MockTrigger{ctrl: ctrl}
	mock.recorder = &MockTriggerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrigger) EXPECT() *MockTriggerRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTrigger) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTriggerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTrigger)(nil).Close))
}

// Delay mocks base method.
func (m *MockTrigger) Delay() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Delay indicates an expected call of Timer.
func (mr *MockTriggerRecorder) Delay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delay", reflect.TypeOf((*MockTrigger)(nil).Delay))
}

//------------------------------------------------------------------------------
// HTTP Client
//------------------------------------------------------------------------------

// MockHTTPClient is a mock instance of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientRecorder
}

var _ httpClient = &MockHTTPClient{}

// MockHTTPClientRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientRecorder {
	return m.recorder
}

// Do mock base method.
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicate an expected call of Do.
func (mr *MockHTTPClientRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), req)
}

//------------------------------------------------------------------------------
// File System
//------------------------------------------------------------------------------

// MockFs is a mock instance of Fs interface
type MockFs struct {
	ctrl     *gomock.Controller
	recorder *MockFsRecorder
}

var _ afero.Fs = &MockFs{}

// MockFsRecorder is the mock recorder for MockFs
type MockFsRecorder struct {
	mock *MockFs
}

// NewMockFs creates a new mock instance
func NewMockFs(ctrl *gomock.Controller) *MockFs {
	mock := &MockFs{ctrl: ctrl}
	mock.recorder = &MockFsRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFs) EXPECT() *MockFsRecorder {
	return m.recorder
}

// Chmod mocks base method
func (m *MockFs) Chmod(arg0 string, arg1 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chmod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chmod indicates an expected call of Chmod
func (mr *MockFsRecorder) Chmod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chmod", reflect.TypeOf((*MockFs)(nil).Chmod), arg0, arg1)
}

// Chtimes mocks base method
func (m *MockFs) Chtimes(arg0 string, arg1, arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chtimes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chtimes indicates an expected call of Chtimes
func (mr *MockFsRecorder) Chtimes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chtimes", reflect.TypeOf((*MockFs)(nil).Chtimes), arg0, arg1, arg2)
}

// Chown mocks base method
func (m *MockFs) Chown(arg0 string, arg1, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Chown", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Chown indicates an expected call of Chtimes
func (mr *MockFsRecorder) Chown(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Chown", reflect.TypeOf((*MockFs)(nil).Chown), arg0, arg1, arg2)
}

// Create mocks base method
func (m *MockFs) Create(arg0 string) (afero.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(afero.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockFsRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFs)(nil).Create), arg0)
}

// Mkdir mocks base method
func (m *MockFs) Mkdir(arg0 string, arg1 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir
func (mr *MockFsRecorder) Mkdir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockFs)(nil).Mkdir), arg0, arg1)
}

// MkdirAll mocks base method
func (m *MockFs) MkdirAll(arg0 string, arg1 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MkdirAll", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MkdirAll indicates an expected call of MkdirAll
func (mr *MockFsRecorder) MkdirAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MkdirAll", reflect.TypeOf((*MockFs)(nil).MkdirAll), arg0, arg1)
}

// Name mocks base method
func (m *MockFs) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockFsRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockFs)(nil).Name))
}

// Open mocks base method
func (m *MockFs) Open(arg0 string) (afero.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(afero.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockFsRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockFs)(nil).Open), arg0)
}

// OpenFile mocks base method
func (m *MockFs) OpenFile(arg0 string, arg1 int, arg2 os.FileMode) (afero.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(afero.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile
func (mr *MockFsRecorder) OpenFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockFs)(nil).OpenFile), arg0, arg1, arg2)
}

// Remove mocks base method
func (m *MockFs) Remove(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockFsRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockFs)(nil).Remove), arg0)
}

// RemoveAll mocks base method
func (m *MockFs) RemoveAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *MockFsRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockFs)(nil).RemoveAll), arg0)
}

// Rename mocks base method
func (m *MockFs) Rename(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rename indicates an expected call of Rename
func (mr *MockFsRecorder) Rename(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockFs)(nil).Rename), arg0, arg1)
}

// Stat mocks base method
func (m *MockFs) Stat(arg0 string) (os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", arg0)
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat
func (mr *MockFsRecorder) Stat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockFs)(nil).Stat), arg0)
}

//------------------------------------------------------------------------------
// File
//------------------------------------------------------------------------------

// MockFile is a mock instance of File interface
type MockFile struct {
	ctrl     *gomock.Controller
	recorder *MockFileRecorder
}

var _ afero.File = &MockFile{}

// MockFileRecorder is the mock recorder for MockFile
type MockFileRecorder struct {
	mock *MockFile
}

// NewMockFile creates a new mock instance
func NewMockFile(ctrl *gomock.Controller) *MockFile {
	mock := &MockFile{ctrl: ctrl}
	mock.recorder = &MockFileRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFile) EXPECT() *MockFileRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockFile) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockFileRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFile)(nil).Close))
}

// Name mocks base method
func (m *MockFile) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockFileRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockFile)(nil).Name))
}

// Read mocks base method
func (m *MockFile) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockFileRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFile)(nil).Read), arg0)
}

// ReadAt mocks base method
func (m *MockFile) ReadAt(arg0 []byte, arg1 int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAt", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAt indicates an expected call of ReadAt
func (mr *MockFileRecorder) ReadAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAt", reflect.TypeOf((*MockFile)(nil).ReadAt), arg0, arg1)
}

// Readdir mocks base method
func (m *MockFile) Readdir(arg0 int) ([]os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Readdir", arg0)
	ret0, _ := ret[0].([]os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Readdir indicates an expected call of Readdir
func (mr *MockFileRecorder) Readdir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readdir", reflect.TypeOf((*MockFile)(nil).Readdir), arg0)
}

// Readdirnames mocks base method
func (m *MockFile) Readdirnames(arg0 int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Readdirnames", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Readdirnames indicates an expected call of Readdirnames
func (mr *MockFileRecorder) Readdirnames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readdirnames", reflect.TypeOf((*MockFile)(nil).Readdirnames), arg0)
}

// Seek mocks base method
func (m *MockFile) Seek(arg0 int64, arg1 int) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seek", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seek indicates an expected call of Seek
func (mr *MockFileRecorder) Seek(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seek", reflect.TypeOf((*MockFile)(nil).Seek), arg0, arg1)
}

// Stat mocks base method
func (m *MockFile) Stat() (os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat")
	ret0, _ := ret[0].(os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat
func (mr *MockFileRecorder) Stat() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockFile)(nil).Stat))
}

// Sync mocks base method
func (m *MockFile) Sync() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockFileRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockFile)(nil).Sync))
}

// Truncate mocks base method
func (m *MockFile) Truncate(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Truncate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Truncate indicates an expected call of Truncate
func (mr *MockFileRecorder) Truncate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Truncate", reflect.TypeOf((*MockFile)(nil).Truncate), arg0)
}

// Write mocks base method
func (m *MockFile) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockFileRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFile)(nil).Write), arg0)
}

// WriteAt mocks base method
func (m *MockFile) WriteAt(arg0 []byte, arg1 int64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteAt", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteAt indicates an expected call of WriteAt
func (mr *MockFileRecorder) WriteAt(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteAt", reflect.TypeOf((*MockFile)(nil).WriteAt), arg0, arg1)
}

// WriteString mocks base method
func (m *MockFile) WriteString(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteString", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteString indicates an expected call of WriteString
func (mr *MockFileRecorder) WriteString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteString", reflect.TypeOf((*MockFile)(nil).WriteString), arg0)
}

//------------------------------------------------------------------------------
// File Info
//------------------------------------------------------------------------------

// MockFileInfo is a mock instance of FileInfo interface
type MockFileInfo struct {
	ctrl     *gomock.Controller
	recorder *MockFileInfoRecorder
}

var _ os.FileInfo = &MockFileInfo{}

// MockFileInfoRecorder is the mock recorder for MockFileInfo
type MockFileInfoRecorder struct {
	mock *MockFileInfo
}

// NewMockFileInfo creates a new mock instance
func NewMockFileInfo(ctrl *gomock.Controller) *MockFileInfo {
	mock := &MockFileInfo{ctrl: ctrl}
	mock.recorder = &MockFileInfoRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFileInfo) EXPECT() *MockFileInfoRecorder {
	return m.recorder
}

// IsDir mocks base method
func (m *MockFileInfo) IsDir() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDir")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDir indicates an expected call of IsDir
func (mr *MockFileInfoRecorder) IsDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDir", reflect.TypeOf((*MockFileInfo)(nil).IsDir))
}

// ModTime mocks base method
func (m *MockFileInfo) ModTime() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// ModTime indicates an expected call of ModTime
func (mr *MockFileInfoRecorder) ModTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModTime", reflect.TypeOf((*MockFileInfo)(nil).ModTime))
}

// Mode mocks base method
func (m *MockFileInfo) Mode() os.FileMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mode")
	ret0, _ := ret[0].(os.FileMode)
	return ret0
}

// Mode indicates an expected call of Mode
func (mr *MockFileInfoRecorder) Mode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mode", reflect.TypeOf((*MockFileInfo)(nil).Mode))
}

// Name mocks base method
func (m *MockFileInfo) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockFileInfoRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockFileInfo)(nil).Name))
}

// Size mocks base method
func (m *MockFileInfo) Size() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockFileInfoRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockFileInfo)(nil).Size))
}

// Sys mocks base method
func (m *MockFileInfo) Sys() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sys")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Sys indicates an expected call of Sys
func (mr *MockFileInfoRecorder) Sys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sys", reflect.TypeOf((*MockFileInfo)(nil).Sys))
}

//------------------------------------------------------------------------------
// YAML Reader
//------------------------------------------------------------------------------

// MockYAMLReader is a mock instance of configYAMLReader interface
type MockYAMLReader struct {
	ctrl     *gomock.Controller
	recorder *MockYAMLReaderRecorder
}

var _ yamlReader = &MockYAMLReader{}

// MockYAMLReaderRecorder is the mock recorder for MockYAMLReader
type MockYAMLReaderRecorder struct {
	mock *MockYAMLReader
}

// NewMockYAMLReader creates a new mock instance
func NewMockYAMLReader(ctrl *gomock.Controller) *MockYAMLReader {
	mock := &MockYAMLReader{ctrl: ctrl}
	mock.recorder = &MockYAMLReaderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockYAMLReader) EXPECT() *MockYAMLReaderRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockYAMLReader) Decode(partial interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", partial)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockYAMLReaderRecorder) Decode(partial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockYAMLReader)(nil).Decode), partial)
}

//------------------------------------------------------------------------------
// Config
//------------------------------------------------------------------------------

// MockConfig is a mock instance of IConfig interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigRecorder
}

var _ IConfig = &MockConfig{}

// MockConfigRecorder is the mock recorder for MockConfig.
type MockConfigRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigRecorder {
	return m.recorder
}

// Bool mocks base method.
func (m *MockConfig) Bool(path string, def ...bool) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Bool", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bool indicates an expected call of Bool.
func (mr *MockConfigRecorder) Bool(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockConfig)(nil).Bool), varargs...)
}

// Partial mocks base method.
func (m *MockConfig) Config(path string, def ...Config) (IConfig, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Config", varargs...)
	ret0, _ := ret[0].(IConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Partial indicates an expected call of Partial.
func (mr *MockConfigRecorder) Config(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockConfig)(nil).Config), varargs...)
}

// Entries mocks base method.
func (m *MockConfig) Entries() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entries")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Entries indicates an expected call of Entries.
func (mr *MockConfigRecorder) Entries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entries", reflect.TypeOf((*MockConfig)(nil).Entries))
}

// Float mocks base method.
func (m *MockConfig) Float(path string, def ...float64) (float64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float", varargs...)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Float indicates an expected call of Float.
func (mr *MockConfigRecorder) Float(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float", reflect.TypeOf((*MockConfig)(nil).Float), varargs...)
}

// Get mocks base method.
func (m *MockConfig) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockConfigRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfig)(nil).Get), varargs...)
}

// Has mocks base method.
func (m *MockConfig) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockConfigRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockConfig)(nil).Has), path)
}

// Int mocks base method.
func (m *MockConfig) Int(path string, def ...int) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int indicates an expected call of Int.
func (mr *MockConfigRecorder) Int(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockConfig)(nil).Int), varargs...)
}

// List mocks base method.
func (m *MockConfig) List(path string, def ...[]interface{}) ([]interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockConfigRecorder) List(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockConfig)(nil).List), varargs...)
}

// Populate mocks base method.
func (m *MockConfig) Populate(path string, target interface{}, icase ...bool) (interface{}, error) {
	m.ctrl.T.Helper()
	m.ctrl.T.Helper()
	varargs := []interface{}{path, target}
	for _, a := range icase {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Populate", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Populate indicates an expected call of Partial.
func (mr *MockConfigRecorder) Populate(path, target interface{}, icase ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path, target}, icase...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Populate", reflect.TypeOf((*MockConfig)(nil).Populate), varargs...)
}

// String mocks base method.
func (m *MockConfig) String(path string, def ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "String", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// String indicates an expected call of String.
func (mr *MockConfigRecorder) String(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockConfig)(nil).String), varargs...)
}

//------------------------------------------------------------------------------
// JSON Reader
//------------------------------------------------------------------------------

// MockJSONReader is a mock instance of configJSONReader interface
type MockJSONReader struct {
	ctrl     *gomock.Controller
	recorder *MockJSONReaderRecorder
}

var _ jsonReader = &MockJSONReader{}

// MockJSONReaderRecorder is the mock recorder for MockJSONReader
type MockJSONReaderRecorder struct {
	mock *MockJSONReader
}

// NewMockJSONReader creates a new mock instance
func NewMockJSONReader(ctrl *gomock.Controller) *MockJSONReader {
	mock := &MockJSONReader{ctrl: ctrl}
	mock.recorder = &MockJSONReaderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJSONReader) EXPECT() *MockJSONReaderRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockJSONReader) Decode(partial interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", partial)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockJSONReaderRecorder) Decode(partial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockJSONReader)(nil).Decode), partial)
}

//------------------------------------------------------------------------------
// Decoder
//------------------------------------------------------------------------------

// MockDecoder is a mock instance of IDecoder interface
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderRecorder
}

var _ IDecoder = &MockDecoder{}

// MockDecoderRecorder is the mock recorder for MockDecoder
type MockDecoderRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDecoder) EXPECT() *MockDecoderRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockDecoder) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockDecoderRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDecoder)(nil).Close))
}

// Decode mocks base method
func (m *MockDecoder) Decode() (IConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode")
	ret0, _ := ret[0].(IConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode
func (mr *MockDecoderRecorder) Decode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode))
}

//------------------------------------------------------------------------------
// Decoder Strategy
//------------------------------------------------------------------------------

// MockDecoderStrategy is a mock instance of IDecoderStrategy interface
type MockDecoderStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderStrategyRecorder
}

var _ IDecoderStrategy = &MockDecoderStrategy{}

// MockDecoderStrategyRecorder is the mock recorder for MockDecoderStrategy
type MockDecoderStrategyRecorder struct {
	mock *MockDecoderStrategy
}

// NewMockDecoderStrategy creates a new mock instance
func NewMockDecoderStrategy(ctrl *gomock.Controller) *MockDecoderStrategy {
	mock := &MockDecoderStrategy{ctrl: ctrl}
	mock.recorder = &MockDecoderStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDecoderStrategy) EXPECT() *MockDecoderStrategyRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockDecoderStrategy) Accept(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockDecoderStrategyRecorder) Accept(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockDecoderStrategy)(nil).Accept), varargs...)
}

// Create mocks base method
func (m *MockDecoderStrategy) Create(args ...interface{}) (IDecoder, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(IDecoder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockDecoderStrategyRecorder) Create(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDecoderStrategy)(nil).Create), args...)
}

//------------------------------------------------------------------------------
// Decoder Factory
//------------------------------------------------------------------------------

// MockDecoderFactory is a mock instance of IDecoderFactory interface.
type MockDecoderFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderFactoryRecorder
}

var _ IDecoderFactory = &MockDecoderFactory{}

// MockDecoderFactoryRecorder is the mock recorder for MockDecoderFactory.
type MockDecoderFactoryRecorder struct {
	mock *MockDecoderFactory
}

// NewMockDecoderFactory creates a new mock instance.
func NewMockDecoderFactory(ctrl *gomock.Controller) *MockDecoderFactory {
	mock := &MockDecoderFactory{ctrl: ctrl}
	mock.recorder = &MockDecoderFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoderFactory) EXPECT() *MockDecoderFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDecoderFactory) Create(format string, args ...interface{}) (IDecoder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(IDecoder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDecoderFactoryRecorder) Create(format interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDecoderFactory)(nil).Create), varargs...)
}

// Register mocks base method.
func (m *MockDecoderFactory) Register(strategy IDecoderStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockDecoderFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockDecoderFactory)(nil).Register), strategy)
}

//------------------------------------------------------------------------------
// Source
//------------------------------------------------------------------------------

// MockSource is a mock instance of ISource interface
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceRecorder
}

var _ ISource = &MockSource{}

// MockSourceRecorder is the mock recorder for MockSource
type MockSourceRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSource) EXPECT() *MockSourceRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockSourceRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSource)(nil).Close))
}

// Has mocks base method
func (m *MockSource) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockSourceRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockSource)(nil).Has), path)
}

// Get mocks base method
func (m *MockSource) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockSourceRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSource)(nil).Get), varargs...)
}

//------------------------------------------------------------------------------
// Observable Source
//------------------------------------------------------------------------------

// MockObservableSource is a mock instance of ObservableSource interface
type MockObservableSource struct {
	ctrl     *gomock.Controller
	recorder *MockObservableSourceRecorder
}

var _ IObservableSource = &MockObservableSource{}

// MockObservableSourceRecorder is the mock recorder for MockObservableSource
type MockObservableSourceRecorder struct {
	mock *MockObservableSource
}

// NewMockObservableSource creates a new mock instance
func NewMockObservableSource(ctrl *gomock.Controller) *MockObservableSource {
	mock := &MockObservableSource{ctrl: ctrl}
	mock.recorder = &MockObservableSourceRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObservableSource) EXPECT() *MockObservableSourceRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockObservableSource) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockObservableSourceRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockObservableSource)(nil).Close))
}

// Has mocks base method
func (m *MockObservableSource) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has
func (mr *MockObservableSourceRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockObservableSource)(nil).Has), path)
}

// Get mocks base method
func (m *MockObservableSource) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockObservableSourceRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	var varargs []interface{}
	varargs = append(varargs, path)
	for _, a := range def {
		varargs = append(varargs, a)
	}
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSource)(nil).Get), varargs...)
}

// Reload mocks base method
func (m *MockObservableSource) Reload() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reload")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reload indicates an expected call of Reload
func (mr *MockObservableSourceRecorder) Reload() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockObservableSource)(nil).Reload))
}

//------------------------------------------------------------------------------
// Source Strategy
//------------------------------------------------------------------------------

// MockSourceStrategy is a mock instance of SourceFactoryStrategy interface
type MockSourceStrategy struct {
	ctrl     *gomock.Controller
	recorder *MockSourceStrategyRecorder
}

var _ ISourceStrategy = &MockSourceStrategy{}

// MockSourceStrategyRecorder is the mock recorder for MockSourceStrategy
type MockSourceStrategyRecorder struct {
	mock *MockSourceStrategy
}

// NewMockSourceStrategy creates a new mock instance
func NewMockSourceStrategy(ctrl *gomock.Controller) *MockSourceStrategy {
	mock := &MockSourceStrategy{ctrl: ctrl}
	mock.recorder = &MockSourceStrategyRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceStrategy) EXPECT() *MockSourceStrategyRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockSourceStrategy) Accept(cfg IConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", cfg)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockSourceStrategyRecorder) Accept(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockSourceStrategy)(nil).Accept), cfg)
}

// Create mocks base method
func (m *MockSourceStrategy) Create(cfg IConfig) (ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockSourceStrategyRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceStrategy)(nil).Create), cfg)
}

//------------------------------------------------------------------------------
// Source Factory
//------------------------------------------------------------------------------

// MockSourceFactory is a mock instance of ISourceFactory interface.
type MockSourceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockSourceFactoryRecorder
}

var _ ISourceFactory = &MockSourceFactory{}

// MockSourceFactoryRecorder is the mock recorder for MockSourceFactory.
type MockSourceFactoryRecorder struct {
	mock *MockSourceFactory
}

// NewMockSourceFactory creates a new mock instance.
func NewMockSourceFactory(ctrl *gomock.Controller) *MockSourceFactory {
	mock := &MockSourceFactory{ctrl: ctrl}
	mock.recorder = &MockSourceFactoryRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourceFactory) EXPECT() *MockSourceFactoryRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSourceFactory) Create(cfg IConfig) (ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cfg)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockSourceFactoryRecorder) Create(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSourceFactory)(nil).Create), cfg)
}

// Register mocks base method.
func (m *MockSourceFactory) Register(strategy ISourceStrategy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", strategy)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockSourceFactoryRecorder) Register(strategy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockSourceFactory)(nil).Register), strategy)
}

//------------------------------------------------------------------------------
// Manager
//------------------------------------------------------------------------------

// MockManager is a mock instance of IManager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerRecorder
}

var _ IManager = &MockManager{}

// MockManagerRecorder is the mock recorder for MockManager.
type MockManagerRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerRecorder {
	return m.recorder
}

// AddObserver mocks base method.
func (m *MockManager) AddObserver(path string, callback IObserver) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddObserver", path, callback)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddObserver indicates an expected call of AddObserver.
func (mr *MockManagerRecorder) AddObserver(path, callback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddObserver", reflect.TypeOf((*MockManager)(nil).AddObserver), path, callback)
}

// AddSource mocks base method.
func (m *MockManager) AddSource(id string, priority int, src ISource) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSource", id, priority, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSource indicates an expected call of AddSource.
func (mr *MockManagerRecorder) AddSource(id, priority, src interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSource", reflect.TypeOf((*MockManager)(nil).AddSource), id, priority, src)
}

// Bool mocks base method.
func (m *MockManager) Bool(path string, def ...bool) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Bool", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bool indicates an expected call of Bool.
func (mr *MockManagerRecorder) Bool(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockManager)(nil).Bool), varargs...)
}

// Close mocks base method.
func (m *MockManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockManagerRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockManager)(nil).Close))
}

// Config mocks base method.
func (m *MockManager) Config(path string, def ...Config) (IConfig, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Config", varargs...)
	ret0, _ := ret[0].(IConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Config indicates an expected call of Partial.
func (mr *MockManagerRecorder) Config(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockManager)(nil).Config), varargs...)
}

// Entries mocks base method.
func (m *MockManager) Entries() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Entries")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Float mocks base method.
func (m *MockManager) Float(path string, def ...float64) (float64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Float", varargs...)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Entries indicates an expected call of Entries.
func (mr *MockManagerRecorder) Entries() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Entries", reflect.TypeOf((*MockManager)(nil).Entries))
}

// Float indicates an expected call of Float.
func (mr *MockManagerRecorder) Float(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float", reflect.TypeOf((*MockManager)(nil).Float), varargs...)
}

// Get mocks base method.
func (m *MockManager) Get(path string, def ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockManagerRecorder) Get(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockManager)(nil).Get), varargs...)
}

// Has mocks base method.
func (m *MockManager) Has(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockManagerRecorder) Has(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockManager)(nil).Has), path)
}

// HasObserver mocks base method.
func (m *MockManager) HasObserver(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasObserver", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasObserver indicates an expected call of HasObserver.
func (mr *MockManagerRecorder) HasObserver(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasObserver", reflect.TypeOf((*MockManager)(nil).HasObserver), path)
}

// HasSource mocks base method.
func (m *MockManager) HasSource(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasSource", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasSource indicates an expected call of HasSource.
func (mr *MockManagerRecorder) HasSource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasSource", reflect.TypeOf((*MockManager)(nil).HasSource), id)
}

// Int mocks base method.
func (m *MockManager) Int(path string, def ...int) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Int", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Int indicates an expected call of Int.
func (mr *MockManagerRecorder) Int(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockManager)(nil).Int), varargs...)
}

// List mocks base method.
func (m *MockManager) List(path string, def ...[]interface{}) ([]interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockManagerRecorder) List(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManager)(nil).List), varargs...)
}

// Populate mocks base method.
func (m *MockManager) Populate(path string, target interface{}, icase ...bool) (interface{}, error) {
	m.ctrl.T.Helper()
	m.ctrl.T.Helper()
	varargs := []interface{}{path, target}
	for _, a := range icase {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Populate", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Populate indicates an expected call of Partial.
func (mr *MockManagerRecorder) Populate(path, target interface{}, icase ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path, target}, icase...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Populate", reflect.TypeOf((*MockManager)(nil).Populate), varargs...)
}

// RemoveAllSources mocks base method.
func (m *MockManager) RemoveAllSources() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllSources")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllSources indicates an expected call of RemoveAllSources.
func (mr *MockManagerRecorder) RemoveAllSources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllSources", reflect.TypeOf((*MockManager)(nil).RemoveAllSources))
}

// RemoveObserver mocks base method.
func (m *MockManager) RemoveObserver(path string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveObserver", path)
}

// RemoveObserver indicates an expected call of RemoveObserver.
func (mr *MockManagerRecorder) RemoveObserver(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveObserver", reflect.TypeOf((*MockManager)(nil).RemoveObserver), path)
}

// RemoveSource mocks base method.
func (m *MockManager) RemoveSource(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSource", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSource indicates an expected call of RemoveSource.
func (mr *MockManagerRecorder) RemoveSource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSource", reflect.TypeOf((*MockManager)(nil).RemoveSource), id)
}

// Source mocks base method.
func (m *MockManager) Source(id string) (ISource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Source", id)
	ret0, _ := ret[0].(ISource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Source indicates an expected call of Source.
func (mr *MockManagerRecorder) Source(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Source", reflect.TypeOf((*MockManager)(nil).Source), id)
}

// SourcePriority mocks base method.
func (m *MockManager) SourcePriority(id string, priority int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourcePriority", id, priority)
	ret0, _ := ret[0].(error)
	return ret0
}

// SourcePriority indicates an expected call of SourcePriority.
func (mr *MockManagerRecorder) SourcePriority(id, priority interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourcePriority", reflect.TypeOf((*MockManager)(nil).SourcePriority), id, priority)
}

// String mocks base method.
func (m *MockManager) String(path string, def ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{path}
	for _, a := range def {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "String", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// String indicates an expected call of String.
func (mr *MockManagerRecorder) String(path interface{}, def ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{path}, def...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockManager)(nil).String), varargs...)
}
