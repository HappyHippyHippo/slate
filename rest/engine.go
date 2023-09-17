package rest

import (
	"html/template"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Engine interface for the gin-gonic engine object.
type Engine interface {
	gin.IRoutes
	gin.IRouter

	Handler() http.Handler
	Delims(left string, right string) *gin.Engine
	SecureJsonPrefix(prefix string) *gin.Engine
	LoadHTMLGlob(pattern string)
	LoadHTMLFiles(files ...string)
	SetHTMLTemplate(templ *template.Template)
	SetFuncMap(funcMap template.FuncMap)
	NoRoute(handlers ...gin.HandlerFunc)
	NoMethod(handlers ...gin.HandlerFunc)
	Use(middleware ...gin.HandlerFunc) gin.IRoutes
	Routes() (routes gin.RoutesInfo)
	Run(addr ...string) (err error)
	SetTrustedProxies(trustedProxies []string) error
	RunTLS(addr string, certFile string, keyFile string) (err error)
	RunUnix(file string) (err error)
	RunFd(fd int) (err error)
	RunListener(listener net.Listener) (err error)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	HandleContext(c *gin.Context)
}

var _ Engine = &gin.Engine{}
