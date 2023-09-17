package rest

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/env"
)

const (
	// EnvID defines the slate.api.rest package base environment variable name.
	EnvID = slate.EnvID + "_REST"
)

var (
	// ConfigPath defines the configuration location where is
	// defined the REST service configuration.
	ConfigPath = env.String(EnvID+"_CONFIG_PATH", "slate.api.rest")

	// WatchdogName defines the default REST service watchdog name.
	WatchdogName = env.String(EnvID+"_WATCHDOG_NAME", "rest")

	// Port defines the default rest service port.
	Port = env.Int(EnvID+"_PORT", 80)

	// LogChannel defines the default logging channel.
	LogChannel = env.String(EnvID+"_LOG_CHANNEL", "rest")

	// LogLevel defines the default logging level.
	LogLevel = env.String(EnvID+"_LOG_LEVEL", "info")

	// LogStartMessage defines the default service start logging message.
	LogStartMessage = env.String(EnvID+"_LOG_START_MESSAGE", "[service:rest] service starting ...")

	// LogErrorMessage defines the default service error logging message.
	LogErrorMessage = env.String(EnvID+"_LOG_ERROR_MESSAGE", "[service:rest] service error")

	// LogEndMessage defines the default service end logging message.
	LogEndMessage = env.String(EnvID+"_LOG_END_MESSAGE", "[service:rest] service terminated")
)
