package log

import (
	"fmt"
	sconfig "github.com/happyhippyhippo/slate/config"
	"time"

	"github.com/gin-gonic/gin"
	slog "github.com/happyhippyhippo/slate/log"
	srest "github.com/happyhippyhippo/slate/rest"
)

// MiddlewareGenerator @todo doc
type MiddlewareGenerator func(string) (srest.Middleware, error)

// NewMiddlewareGenerator will instantiate a new middleware function generator
// that will create an endpoint function that will emit logging signals on
// a request event and on a response event.
func NewMiddlewareGenerator(cfg sconfig.IManager, logger slog.ILogger) (MiddlewareGenerator, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	if logger == nil {
		return nil, errNilPointer("logger")
	}

	return func(id string) (srest.Middleware, error) {
		statusCode, e := cfg.Int(fmt.Sprintf(ConfigPathEndpointStatusFormat, id))
		if e != nil {
			return nil, e
		}

		requestReader := RequestReaderDefault
		responseReader := NewResponseReaderDefault(statusCode)

		if DecorateJSON {
			requestReader, _ = NewRequestReaderDecoratorJSON(requestReader, nil)
			responseReader, _ = NewResponseReaderDecoratorJSON(responseReader, nil)
		}

		if DecorateXML {
			requestReader, _ = NewRequestReaderDecoratorXML(requestReader, nil)
			responseReader, _ = NewResponseReaderDecoratorXML(responseReader, nil)
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
	}, nil
}
