package srlog

import (
	"github.com/happyhippyhippo/slate/senv"
	"github.com/happyhippyhippo/slate/slog"
	"github.com/happyhippyhippo/slate/srest"
	"strings"
)

const (
	// ContainerID defines the id to be used as the container
	// registration id of a logging middleware instance factory function.
	ContainerID = srest.ContainerID + ".log"

	// ContainerOkID defines the id to be used as the
	// container registration id of a standard log middleware generator
	// for a 200(OK) status code.
	ContainerOkID = ContainerID + ".ok"

	// ContainerCreatedID defines the id to be used as
	// the container registration id of a standard log middleware generator
	// for a 201(Created) status code.
	ContainerCreatedID = ContainerID + ".created"

	// ContainerNoContentID defines the id to be used as
	// the container registration id of a standard log middleware generator
	// for a 204(NoContent) status code.
	ContainerNoContentID = ContainerID + ".no_content"
)
const (
	// EnvID defines the slate.srest.log package base environment variable name.
	EnvID = srest.EnvID + "_LOG"
)

var (
	// RequestChannel defines the channel id to be used when
	// the log middleware sends the request logging signal to the logger
	// instance.
	RequestChannel = senv.String(EnvID+"_REQUEST_CHANNEL", "transport")

	// RequestLevel defines the logging level to be used when
	// the log middleware sends the request logging signal to the logger
	// instance.
	RequestLevel = envToLogLevel(EnvID+"_REQUEST_LEVEL", slog.DEBUG)

	// RequestMessage defines the request event logging message to
	// be used when the log middleware sends the logging signal to the logger
	// instance.
	RequestMessage = senv.String(EnvID+"_REQUEST_MESSAGE", "Request")

	// ResponseChannel defines the channel id to be used when the
	// log middleware sends the response logging signal to the logger instance.
	ResponseChannel = senv.String(EnvID+"_RESPONSE_CHANNEL", "transport")

	// ResponseLevel defines the logging level to be used when the
	// log middleware sends the response logging signal to the logger instance.
	ResponseLevel = envToLogLevel(EnvID+"_RESPONSE_LEVEL", slog.INFO)

	// ResponseMessage defines the response event logging message
	// to be used when the log middleware sends the logging signal to the
	// logger instance.
	ResponseMessage = senv.String(EnvID+"_RESPONSE_MESSAGE", "Response")

	// DecorateJSON flag that defines the decoration of the log entries
	// for JSON body content.
	DecorateJSON = senv.Bool(EnvID+"_DECORATE_JSON", true)

	// DecorateXML flag that defines the decoration of the log entries
	// for XML body content.
	DecorateXML = senv.Bool(EnvID+"_DECORATE_XML", false)
)

func envToLogLevel(env string, def slog.Level) slog.Level {
	v, ok := slog.LevelMap[strings.ToLower(env)]
	if !ok {
		return def
	}
	return v
}
