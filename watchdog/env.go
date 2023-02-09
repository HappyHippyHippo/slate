package watchdog

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/env"
)

const (
	// EnvID defines the watchdog package base environment variable name.
	EnvID = slate.EnvID + "_WATCHDOG"
)

var (
	// ConfigPathPrefix defines the configuration path of the watchdog
	// entries of the application.
	ConfigPathPrefix = env.String(EnvID+"_CONFIG_PATH", "slate.watchdog.services")

	// LogChannel defines the logging signal channel of the watchdogs.
	LogChannel = env.String(EnvID+"_LOG_CHANNEL", "watchdog")

	// LogStartLevel defines the watchdog starting logging signal
	// message level.
	LogStartLevel = env.String(EnvID+"_LOG_START_LEVEL", "notice")

	// LogStartMessage defines the watchdog starting logging signal message.
	LogStartMessage = env.String(EnvID+"_LOG_START_MESSAGE", "[watchdog:%s] start execution")

	// LogErrorLevel defines the watchdog error logging signal
	// message level.
	LogErrorLevel = env.String(EnvID+"_LOG_ERROR_LEVEL", "error")

	// LogErrorMessage defines the watchdog error logging signal message.
	LogErrorMessage = env.String(EnvID+"_LOG_ERROR_MESSAGE", "[watchdog:%s] execution error (%v)")

	// LogDoneLevel defines the watchdog termination logging
	// signal message level.
	LogDoneLevel = env.String(EnvID+"_LOG_DONE_LEVEL", "notice")

	// LogDoneMessage defines the watchdog termination logging
	// signal message.
	LogDoneMessage = env.String(EnvID+"_LOG_DONE_MESSAGE", "[watchdog:%s] execution terminated")
)
