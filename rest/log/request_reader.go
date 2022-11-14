package log

import (
	"github.com/gin-gonic/gin"
)

// RequestReader defines the function used by the middleware that compose the
// logging request context object.
type RequestReader func(ctx *gin.Context) (map[string]interface{}, error)
