package logmw

import (
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/log"
)

// ResponseReader defines the interface methods of a response
// context reader used to compose the data to be sent to the logger on a
// response event.
type ResponseReader func(ctx *gin.Context, writer responseWriter, statusCode int) (log.Context, error)
