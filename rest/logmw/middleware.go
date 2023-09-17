package logmw

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/rest"
)

// MiddlewareGenerator @todo doc
type MiddlewareGenerator func(statusCode int) rest.Middleware

// NewMiddlewareGenerator @todo doc
func NewMiddlewareGenerator(
	logger *log.Log,
	requestReader RequestReader,
	responseReader ResponseReader,
) (MiddlewareGenerator, error) {
	// check logger argument reference
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check request reader argument reference
	if requestReader == nil {
		return nil, errNilPointer("requestReader")
	}
	// check response reader argument reference
	if responseReader == nil {
		return nil, errNilPointer("responseReader")
	}
	// return the middleware generator function
	return func(
		statusCode int,
	) rest.Middleware {
		// return the middleware method with the expected status code
		return func(
			next gin.HandlerFunc,
		) gin.HandlerFunc {
			// return the middleware handler function
			return func(
				ctx *gin.Context,
			) {
				// override the context writer
				w, _ := NewWriter(ctx.Writer)
				ctx.Writer = w
				// obtain and log the request content
				req, _ := requestReader(ctx)
				_ = logger.Signal(
					RequestChannel,
					RequestLevel,
					RequestMessage,
					log.Context{
						"request": req,
					},
				)
				// execute the endpoint process and calculate the elapsed
				// time of it
				startTimestamp := time.Now().UnixMilli()
				if next != nil {
					next(ctx)
				}
				duration := time.Now().UnixMilli() - startTimestamp
				// obtain and log the request, response and execution duration
				resp, _ := responseReader(ctx, w, statusCode)
				_ = logger.Signal(
					ResponseChannel,
					ResponseLevel,
					ResponseMessage,
					log.Context{
						"request":  req,
						"response": resp,
						"duration": duration,
					},
				)
			}
		}
	}, nil
}
