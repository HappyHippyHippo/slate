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
	Delims(left, right string) *gin.Engine
	HandleContext(c *gin.Context)
	LoadHTMLFiles(files ...string)
	LoadHTMLGlob(pattern string)
	NoMethod(handlers ...gin.HandlerFunc)
	NoRoute(handlers ...gin.HandlerFunc)
	Routes() (routes gin.RoutesInfo)
	Run(addr ...string) (err error)
	RunFd(fd int) (err error)
	RunListener(listener net.Listener) (err error)
	RunTLS(addr, certFile, keyFile string) (err error)
	RunUnix(file string) (err error)
	SecureJsonPrefix(prefix string) *gin.Engine
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	SetFuncMap(funcMap template.FuncMap)
	SetHTMLTemplate(tpl *template.Template)
}
