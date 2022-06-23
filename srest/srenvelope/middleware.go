package srenvelope

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/srest"
	"net/http"
)

// NewMiddlewareGenerator returns a middleware generator function
// based on the application configuration. This middleware generator function
// should be called with the corresponding endpoint name, so it can generate
// the appropriate middleware function.
func NewMiddlewareGenerator(cfg sconfig.Manager) (func(string) (srest.Middleware, error), error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	service, err := cfg.Int(ConfigPathServiceID, 0)
	if err != nil {
		return nil, err
	}

	_ = cfg.AddObserver(ConfigPathServiceID, func(old interface{}, new interface{}) {
		service = new.(int)
	})

	acceptedList, err := cfg.List(ConfigPathTransportAcceptList)
	if err != nil {
		return nil, err
	}

	var accepted []string
	for _, v := range acceptedList {
		if tv, ok := v.(string); ok {
			accepted = append(accepted, tv)
		}
	}

	_ = cfg.AddObserver(ConfigPathTransportAcceptList, func(old interface{}, new interface{}) {
		accepted = []string{}
		for _, v := range new.([]interface{}) {
			accepted = append(accepted, v.(string))
		}
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

					ctx.Negotiate(
						response.GetStatusCode(),
						gin.Negotiate{
							Offered: accepted,
							Data:    response.SetService(service).SetEndpoint(endpoint),
						},
					)
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
