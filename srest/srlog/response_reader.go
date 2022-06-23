package srlog

import "github.com/gin-gonic/gin"

// ResponseReader defines the interface methods of a response
// context reader used to compose the data to be sent to the logger on a
// response event.
type ResponseReader func(ctx *gin.Context, writer responseWriter) (map[string]interface{}, error)
