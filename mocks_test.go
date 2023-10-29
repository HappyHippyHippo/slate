package slate

import (
	"io"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
)

// ----------------------------------------------------------------------------
// io.Closer
// ----------------------------------------------------------------------------

// MockCloser is a mock instance of MockCloser interface
type MockCloser struct {
	ctrl     *gomock.Controller
	recorder *MockCloserRecorder
}

var _ io.Closer = &MockCloser{}

// MockCloserRecorder is the mock recorder for MockCloser
type MockCloserRecorder struct {
	mock *MockCloser
}

// NewMockCloser creates a new mock instance
func NewMockCloser(ctrl *gomock.Controller) *MockCloser {
	mock := &MockCloser{ctrl: ctrl}
	mock.recorder = &MockCloserRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloser) EXPECT() *MockCloserRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockCloserRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloser)(nil).Close))
}

// ----------------------------------------------------------------------------
// io.Reader
// ----------------------------------------------------------------------------

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

// ----------------------------------------------------------------------------
// io.Writer
// ----------------------------------------------------------------------------

// MockWriter is a mock instance of Writer interface
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterRecorder
}

var _ io.Writer = &MockWriter{}

// MockWriterRecorder is the mock recorder for MockWriter
type MockWriterRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriter) EXPECT() *MockWriterRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockWriterRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriter)(nil).Write), arg0)
}

// Close mocks base method
func (m *MockWriter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockWriterRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockWriter)(nil).Close))
}

// ----------------------------------------------------------------------------
// sync.Locker
// ----------------------------------------------------------------------------

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

// ----------------------------------------------------------------------------
// afero.Fs
// ----------------------------------------------------------------------------

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

// ----------------------------------------------------------------------------
// afero.File
// ----------------------------------------------------------------------------

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

// ----------------------------------------------------------------------------
// os.FileInfo
// ----------------------------------------------------------------------------

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
