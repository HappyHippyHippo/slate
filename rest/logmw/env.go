package logmw

import (
	"strings"

	"github.com/happyhippyhippo/slate/env"
	"github.com/happyhippyhippo/slate/log"
	"github.com/happyhippyhippo/slate/rest"
)

const (
	// EnvID defines the slate.rest.log package base environment variable name.
	EnvID = rest.EnvID + "_LOG"
)

var (
	// RequestChannel defines the channel id to be used when
	// the log middleware sends the request logging signal to the logger
	// instance.
	RequestChannel = env.String(EnvID+"_REQUEST_CHANNEL", "rest")

	// RequestLevel defines the logging level to be used when
	// the log middleware sends the request logging signal to the logger
	// instance.
	RequestLevel = envToLogLevel(EnvID+"_REQUEST_LEVEL", log.DEBUG)

	// RequestMessage defines the request event logging message to
	// be used when the log middleware sends the logging signal to the logger
	// instance.
	RequestMessage = env.String(EnvID+"_REQUEST_MESSAGE", "Request")

	// ResponseChannel defines the channel id to be used when the
	// log middleware sends the response logging signal to the logger instance.
	ResponseChannel = env.String(EnvID+"_RESPONSE_CHANNEL", "rest")

	// ResponseLevel defines the logging level to be used when the
	// log middleware sends the response logging signal to the logger instance.
	ResponseLevel = envToLogLevel(EnvID+"_RESPONSE_LEVEL", log.INFO)

	// ResponseMessage defines the response event logging message
	// to be used when the log middleware sends the logging signal to the
	// logger instance.
	ResponseMessage = env.String(EnvID+"_RESPONSE_MESSAGE", "Response")

	// DecorateJSON flag that defines the decoration of the log entries
	// for JSON body content.
	DecorateJSON = env.Bool(EnvID+"_DECORATE_JSON", true)

	// DecorateXML flag that defines the decoration of the log entries
	// for XML body content.
	DecorateXML = env.Bool(EnvID+"_DECORATE_XML", false)
)

func envToLogLevel(ev string, def log.Level) log.Level {
	v, ok := log.LevelMap[strings.ToLower(ev)]
	if !ok {
		return def
	}
	return v
}
