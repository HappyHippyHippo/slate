package logmw

import (
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/log"
)

// RequestReader defines the function used by the middleware that compose the
// logging request context object.
type RequestReader func(ctx *gin.Context) (log.Context, error)
