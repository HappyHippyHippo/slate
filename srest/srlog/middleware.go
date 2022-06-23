package srlog

import (
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/slog"
	"time"
)

// NewMiddleware will instantiate a new middleware that will emit
// logging signals on a request event and on a response event.
func NewMiddleware(
	logger *slog.Logger,
	requestReader RequestReader,
	responseReader ResponseReader,
) (func(gin.HandlerFunc) gin.HandlerFunc, error) {
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	if requestReader == nil {
		return nil, errNilPointer("requestReader")
	}
	if responseReader == nil {
		return nil, errNilPointer("responseReader")
	}

	return func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			w, _ := newResponseWriter(ctx.Writer)
			ctx.Writer = w

			request, _ := requestReader(ctx)
			_ = logger.Signal(
				RequestChannel,
				RequestLevel,
				RequestMessage,
				map[string]interface{}{
					"request": request,
				},
			)

			startTimestamp := time.Now().UnixMilli()
			if next != nil {
				next(ctx)
			}
			duration := time.Now().UnixMilli() - startTimestamp

			response, _ := responseReader(ctx, w)
			_ = logger.Signal(
				ResponseChannel,
				ResponseLevel,
				ResponseMessage,
				map[string]interface{}{
					"request":  request,
					"response": response,
					"duration": duration,
				},
			)
		}
	}, nil
}
