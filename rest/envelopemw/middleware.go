package envelopemw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/rest"
	"github.com/happyhippyhippo/slate/rest/envelope"
)

// MiddlewareGenerator @todo doc
type MiddlewareGenerator func(string) (rest.Middleware, error)

// NewMiddlewareGenerator returns a middleware generator function
// based on the application configuration. This middleware generator function
// should be called with the corresponding endpoint name, so it can generate
// the appropriate middleware function.
func NewMiddlewareGenerator(
	cfg config.IManager,
	logger log.ILog,
) (MiddlewareGenerator, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// check the logger argument reference
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// validate log level
	logLevel, ok := log.LevelMap[LogLevel]
	if !ok {
		logLevel = log.ERROR
	}
	// retrieve the service id from the configuration
	service, e := cfg.Int(ServiceIDConfigPath, 0)
	if e != nil {
		_ = logger.Signal(LogChannel, logLevel, LogServiceErrorMessage, log.Context{"error": e})
		return nil, e
	}
	// add a config observer for the service ID
	_ = cfg.AddObserver(ServiceIDConfigPath, func(old interface{}, new interface{}) {
		// new value type check for integer
		tnew, ok := new.(int)
		if !ok {
			_ = logger.Signal(LogChannel, logLevel, LogServiceErrorMessage, log.Context{"value": new})
			return
		}
		service = tnew
	})
	// retrieve the service REST accepted format list
	acceptedList, e := cfg.List(FormatAcceptListConfigPath)
	if e != nil {
		_ = logger.Signal(LogChannel, logLevel, LogAcceptListErrorMessage, log.Context{"error": e})
		return nil, e
	}
	// parse the list retrieved from the configuration
	var accepted []string
	for _, v := range acceptedList {
		if tv, ok := v.(string); ok {
			accepted = append(accepted, tv)
		}
	}
	// add a config observer for the REST accepted format list
	_ = cfg.AddObserver(FormatAcceptListConfigPath, func(old interface{}, new interface{}) {
		accepted = []string{}
		// new value type check for an array
		tnew, ok := new.([]interface{})
		if !ok {
			_ = logger.Signal(LogChannel, logLevel, LogAcceptListErrorMessage, log.Context{"list": new})
			return
		}
		// iterate through all the array elements
		for _, v := range tnew {
			// type check for a string
			if tv, ok := v.(string); !ok {
				_ = logger.Signal(LogChannel, logLevel, LogAcceptListErrorMessage, log.Context{"value": v})
			} else {
				// add the iterated element to the accepted format list
				accepted = append(accepted, tv)
			}
		}
	})
	// return the middleware generator
	return func(
		id string,
	) (rest.Middleware, error) {
		// retrieve the endpoint id integer value from the configuration
		endpointIDConfigPath := fmt.Sprintf(EndpointIDConfigPathFormat, id)
		endpoint, e := cfg.Int(endpointIDConfigPath, 0)
		if e != nil {
			_ = logger.Signal(LogChannel, logLevel, LogEndpointErrorMessage, log.Context{"error": e})
			return nil, e
		}
		// add a config observer for the endpoint id integer value
		_ = cfg.AddObserver(endpointIDConfigPath, func(old interface{}, new interface{}) {
			// new value type check for integer
			tnew, ok := new.(int)
			if !ok {
				_ = logger.Signal(LogChannel, logLevel, LogEndpointErrorMessage, log.Context{"value": new})
				return
			}
			endpoint = tnew
		})
		// return the generated middleware function
		return func(
			next gin.HandlerFunc,
		) gin.HandlerFunc {
			// return the middleware handler function
			return func(
				ctx *gin.Context,
			) {
				// declare the result parsing method
				parse := func(val interface{}) {
					var response *envelope.Envelope
					// type check the value to be enveloped
					switch v := val.(type) {
					case *envelope.Envelope:
						// just set the result as the envelope reference
						response = v
					case error:
						// set the result as a new envelope with an
						// internal server error with the given error as the
						// error message
						response =
							envelope.NewEnvelope(http.StatusInternalServerError, nil).
								AddError(envelope.NewStatusError(0, v.Error()))
					default:
						// set the result as a new envelope with an
						// internal server error with a generic error message
						response =
							envelope.NewEnvelope(http.StatusInternalServerError, nil).
								AddError(envelope.NewStatusError(0, "internal server error"))
					}
					// try to negotiate the response format with the defined
					// accepted format mime types giving the response envelope
					// as the content data of the response
					ctx.Negotiate(
						response.GetStatusCode(),
						gin.Negotiate{
							Offered: accepted,
							Data:    response.SetService(service).SetEndpoint(endpoint),
						},
					)
				}
				// always try to fallback retrieve any error to be parsed
				// and result in a proper envelope
				defer func() {
					if e := recover(); e != nil {
						parse(e)
					}
				}()
				// execute the middleware stored execution method
				next(ctx)
				// check if the response as been stored in the context to be
				// correctly parsed
				if response, exists := ctx.Get("response"); exists {
					parse(response)
				}
			}
		}, nil
	}, nil
}
