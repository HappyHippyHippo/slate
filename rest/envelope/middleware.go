package envelope

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	sconfig "github.com/happyhippyhippo/slate/config"
	srest "github.com/happyhippyhippo/slate/rest"
)

// MiddlewareGenerator @todo doc
type MiddlewareGenerator func(string) (srest.Middleware, error)

// NewMiddlewareGenerator returns a middleware generator function
// based on the application configuration. This middleware generator function
// should be called with the corresponding endpoint name, so it can generate
// the appropriate middleware function.
func NewMiddlewareGenerator(cfg sconfig.IManager) (MiddlewareGenerator, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	mutex := sync.RWMutex{}

	service, err := cfg.Int(ConfigPathServiceID, 0)
	if err != nil {
		return nil, err
	}

	_ = cfg.AddObserver(ConfigPathServiceID, func(old interface{}, new interface{}) {
		mutex.Lock()
		service = new.(int)
		mutex.Unlock()
	})

	acceptedList, err := cfg.List(ConfigPathTransportAcceptList)
	if err != nil {
		return nil, err
	}

	getAccepted := func(list []interface{}) []string {
		var accepted []string
		for _, entry := range list {
			if sEntry, ok := entry.(string); ok {
				accepted = append(accepted, sEntry)
			}
		}
		return accepted
	}
	accepted := getAccepted(acceptedList)

	_ = cfg.AddObserver(ConfigPathTransportAcceptList, func(old interface{}, new interface{}) {
		mutex.Lock()
		accepted = getAccepted(new.([]interface{}))
		mutex.Unlock()
	})

	return func(id string) (srest.Middleware, error) {
		endpoint, err := cfg.Int(fmt.Sprintf(ConfigPathEndpointIDFormat, id), 0)
		if err != nil {
			return nil, err
		}

		return func(next gin.HandlerFunc) gin.HandlerFunc {
			return func(ctx *gin.Context) {
				parse := func(val interface{}) {
					var response *Envelope

					switch v := val.(type) {
					case *Envelope:
						response = v
					case error:
						response =
							NewEnvelope(http.StatusInternalServerError, nil, nil).
								AddError(NewStatusError(0, v.Error()))
					default:
						response =
							NewEnvelope(http.StatusInternalServerError, nil, nil).
								AddError(NewStatusError(0, "internal server error"))
					}

					mutex.Lock()
					ctx.Negotiate(
						response.GetStatusCode(),
						gin.Negotiate{
							Offered: accepted,
							Data:    response.SetService(service).SetEndpoint(endpoint),
						},
					)
					mutex.Unlock()
				}

				defer func() {
					if err := recover(); err != nil {
						parse(err)
					}
				}()

				next(ctx)

				if response, exists := ctx.Get("response"); exists {
					parse(response)
				}
			}
		}, nil
	}, nil
}
