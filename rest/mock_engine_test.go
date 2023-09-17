package rest

import (
	"html/template"
	"net"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

// MockEngine is a mock instance of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineRecorder
}

var _ Engine = &MockEngine{}

// MockEngineRecorder is the mock recorder for MockEngine.
type MockEngineRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineRecorder {
	return m.recorder
}

// Any mocks base method.
func (m *MockEngine) Any(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Any", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Any indicates an expected call of Any.
func (mr *MockEngineRecorder) Any(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Any", reflect.TypeOf((*MockEngine)(nil).Any), varargs...)
}

// DELETE mocks base method.
func (m *MockEngine) DELETE(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DELETE", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// DELETE indicates an expected call of DELETE.
func (mr *MockEngineRecorder) DELETE(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DELETE", reflect.TypeOf((*MockEngine)(nil).DELETE), varargs...)
}

// Delims mocks base method.
func (m *MockEngine) Delims(left, right string) *gin.Engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delims", left, right)
	ret0, _ := ret[0].(*gin.Engine)
	return ret0
}

// Delims indicates an expected call of Delims.
func (mr *MockEngineRecorder) Delims(left, right interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delims", reflect.TypeOf((*MockEngine)(nil).Delims), left, right)
}

// GET mocks base method.
func (m *MockEngine) GET(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GET", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// GET indicates an expected call of GET.
func (mr *MockEngineRecorder) GET(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GET", reflect.TypeOf((*MockEngine)(nil).GET), varargs...)
}

// Group mocks base method.
func (m *MockEngine) Group(arg0 string, arg1 ...gin.HandlerFunc) *gin.RouterGroup {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Group", varargs...)
	ret0, _ := ret[0].(*gin.RouterGroup)
	return ret0
}

// Group indicates an expected call of Group.
func (mr *MockEngineRecorder) Group(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockEngine)(nil).Group), varargs...)
}

// HEAD mocks base method.
func (m *MockEngine) HEAD(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HEAD", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// HEAD indicates an expected call of HEAD.
func (mr *MockEngineRecorder) HEAD(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HEAD", reflect.TypeOf((*MockEngine)(nil).HEAD), varargs...)
}

// Handle mocks base method.
func (m *MockEngine) Handle(arg0, arg1 string, arg2 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Handle", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockEngineRecorder) Handle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockEngine)(nil).Handle), varargs...)
}

// HandleContext mocks base method.
func (m *MockEngine) HandleContext(c *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleContext", c)
}

// HandleContext indicates an expected call of HandleContext.
func (mr *MockEngineRecorder) HandleContext(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleContext", reflect.TypeOf((*MockEngine)(nil).HandleContext), c)
}

// Handler mocks base method.
func (m *MockEngine) Handler() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handler")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Handler indicates an expected call of Handler.
func (mr *MockEngineRecorder) Handler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handler", reflect.TypeOf((*MockEngine)(nil).Handler))
}

// LoadHTMLFiles mocks base method.
func (m *MockEngine) LoadHTMLFiles(files ...string) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range files {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "LoadHTMLFiles", varargs...)
}

// LoadHTMLFiles indicates an expected call of LoadHTMLFiles.
func (mr *MockEngineRecorder) LoadHTMLFiles(files ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadHTMLFiles", reflect.TypeOf((*MockEngine)(nil).LoadHTMLFiles), files...)
}

// LoadHTMLGlob mocks base method.
func (m *MockEngine) LoadHTMLGlob(pattern string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LoadHTMLGlob", pattern)
}

// LoadHTMLGlob indicates an expected call of LoadHTMLGlob.
func (mr *MockEngineRecorder) LoadHTMLGlob(pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadHTMLGlob", reflect.TypeOf((*MockEngine)(nil).LoadHTMLGlob), pattern)
}

// Match mocks base method.
func (m *MockEngine) Match(arg0 []string, arg1 string, arg2 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Match", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Match indicates an expected call of Match.
func (mr *MockEngineRecorder) Match(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockEngine)(nil).Match), varargs...)
}

// NoMethod mocks base method.
func (m *MockEngine) NoMethod(handlers ...gin.HandlerFunc) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range handlers {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NoMethod", varargs...)
}

// NoMethod indicates an expected call of NoMethod.
func (mr *MockEngineRecorder) NoMethod(handlers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoMethod", reflect.TypeOf((*MockEngine)(nil).NoMethod), handlers...)
}

// NoRoute mocks base method.
func (m *MockEngine) NoRoute(handlers ...gin.HandlerFunc) {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range handlers {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "NoRoute", varargs...)
}

// NoRoute indicates an expected call of NoRoute.
func (mr *MockEngineRecorder) NoRoute(handlers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoRoute", reflect.TypeOf((*MockEngine)(nil).NoRoute), handlers...)
}

// OPTIONS mocks base method.
func (m *MockEngine) OPTIONS(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OPTIONS", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// OPTIONS indicates an expected call of OPTIONS.
func (mr *MockEngineRecorder) OPTIONS(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OPTIONS", reflect.TypeOf((*MockEngine)(nil).OPTIONS), varargs...)
}

// PATCH mocks base method.
func (m *MockEngine) PATCH(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PATCH", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// PATCH indicates an expected call of PATCH.
func (mr *MockEngineRecorder) PATCH(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PATCH", reflect.TypeOf((*MockEngine)(nil).PATCH), varargs...)
}

// POST mocks base method.
func (m *MockEngine) POST(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "POST", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// POST indicates an expected call of POST.
func (mr *MockEngineRecorder) POST(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "POST", reflect.TypeOf((*MockEngine)(nil).POST), varargs...)
}

// PUT mocks base method.
func (m *MockEngine) PUT(arg0 string, arg1 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PUT", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// PUT indicates an expected call of PUT.
func (mr *MockEngineRecorder) PUT(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PUT", reflect.TypeOf((*MockEngine)(nil).PUT), varargs...)
}

// Routes mocks base method.
func (m *MockEngine) Routes() gin.RoutesInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Routes")
	ret0, _ := ret[0].(gin.RoutesInfo)
	return ret0
}

// Routes indicates an expected call of Routes.
func (mr *MockEngineRecorder) Routes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routes", reflect.TypeOf((*MockEngine)(nil).Routes))
}

// Run mocks base method.
func (m *MockEngine) Run(addr ...string) error {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range addr {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockEngineRecorder) Run(addr ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockEngine)(nil).Run), addr...)
}

// RunFd mocks base method.
func (m *MockEngine) RunFd(fd int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunFd", fd)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunFd indicates an expected call of RunFd.
func (mr *MockEngineRecorder) RunFd(fd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunFd", reflect.TypeOf((*MockEngine)(nil).RunFd), fd)
}

// RunListener mocks base method.
func (m *MockEngine) RunListener(listener net.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunListener", listener)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunListener indicates an expected call of RunListener.
func (mr *MockEngineRecorder) RunListener(listener interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunListener", reflect.TypeOf((*MockEngine)(nil).RunListener), listener)
}

// RunTLS mocks base method.
func (m *MockEngine) RunTLS(addr, certFile, keyFile string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTLS", addr, certFile, keyFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTLS indicates an expected call of RunTLS.
func (mr *MockEngineRecorder) RunTLS(addr, certFile, keyFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTLS", reflect.TypeOf((*MockEngine)(nil).RunTLS), addr, certFile, keyFile)
}

// RunUnix mocks base method.
func (m *MockEngine) RunUnix(file string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunUnix", file)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunUnix indicates an expected call of RunUnix.
func (mr *MockEngineRecorder) RunUnix(file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunUnix", reflect.TypeOf((*MockEngine)(nil).RunUnix), file)
}

// SecureJsonPrefix mocks base method.
//
//revive:disable
func (m *MockEngine) SecureJsonPrefix(prefix string) *gin.Engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecureJsonPrefix", prefix)
	ret0, _ := ret[0].(*gin.Engine)
	return ret0
}

// SecureJsonPrefix indicates an expected call of SecureJsonPrefix.
func (mr *MockEngineRecorder) SecureJsonPrefix(prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecureJsonPrefix", reflect.TypeOf((*MockEngine)(nil).SecureJsonPrefix), prefix)
}

//revive:enable

// ServeHTTP mocks base method.
func (m *MockEngine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeHTTP", w, req)
}

// ServeHTTP indicates an expected call of ServeHTTP.
func (mr *MockEngineRecorder) ServeHTTP(w, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeHTTP", reflect.TypeOf((*MockEngine)(nil).ServeHTTP), w, req)
}

// SetFuncMap mocks base method.
func (m *MockEngine) SetFuncMap(funcMap template.FuncMap) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFuncMap", funcMap)
}

// SetFuncMap indicates an expected call of SetFuncMap.
func (mr *MockEngineRecorder) SetFuncMap(funcMap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFuncMap", reflect.TypeOf((*MockEngine)(nil).SetFuncMap), funcMap)
}

// SetHTMLTemplate mocks base method.
func (m *MockEngine) SetHTMLTemplate(templ *template.Template) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHTMLTemplate", templ)
}

// SetHTMLTemplate indicates an expected call of SetHTMLTemplate.
func (mr *MockEngineRecorder) SetHTMLTemplate(templ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHTMLTemplate", reflect.TypeOf((*MockEngine)(nil).SetHTMLTemplate), templ)
}

// SetTrustedProxies mocks base method.
func (m *MockEngine) SetTrustedProxies(trustedProxies []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTrustedProxies", trustedProxies)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTrustedProxies indicates an expected call of SetTrustedProxies.
func (mr *MockEngineRecorder) SetTrustedProxies(trustedProxies interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrustedProxies", reflect.TypeOf((*MockEngine)(nil).SetTrustedProxies), trustedProxies)
}

// Static mocks base method.
func (m *MockEngine) Static(arg0, arg1 string) gin.IRoutes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Static", arg0, arg1)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Static indicates an expected call of Static.
func (mr *MockEngineRecorder) Static(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Static", reflect.TypeOf((*MockEngine)(nil).Static), arg0, arg1)
}

// StaticFS mocks base method.
func (m *MockEngine) StaticFS(arg0 string, arg1 http.FileSystem) gin.IRoutes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StaticFS", arg0, arg1)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// StaticFS indicates an expected call of StaticFS.
func (mr *MockEngineRecorder) StaticFS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaticFS", reflect.TypeOf((*MockEngine)(nil).StaticFS), arg0, arg1)
}

// StaticFile mocks base method.
func (m *MockEngine) StaticFile(arg0, arg1 string) gin.IRoutes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StaticFile", arg0, arg1)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// StaticFile indicates an expected call of StaticFile.
func (mr *MockEngineRecorder) StaticFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaticFile", reflect.TypeOf((*MockEngine)(nil).StaticFile), arg0, arg1)
}

// StaticFileFS mocks base method.
func (m *MockEngine) StaticFileFS(arg0, arg1 string, arg2 http.FileSystem) gin.IRoutes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StaticFileFS", arg0, arg1, arg2)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// StaticFileFS indicates an expected call of StaticFileFS.
func (mr *MockEngineRecorder) StaticFileFS(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaticFileFS", reflect.TypeOf((*MockEngine)(nil).StaticFileFS), arg0, arg1, arg2)
}

// Use mocks base method.
func (m *MockEngine) Use(arg0 ...gin.HandlerFunc) gin.IRoutes {
	m.ctrl.T.Helper()
	var varargs []interface{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Use", varargs...)
	ret0, _ := ret[0].(gin.IRoutes)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *MockEngineRecorder) Use(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockEngine)(nil).Use), arg0...)
}
