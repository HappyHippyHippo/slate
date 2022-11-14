package rest

import (
	"github.com/happyhippyhippo/slate"
	senv "github.com/happyhippyhippo/slate/env"
)

const (
	// ContainerID defines the slate.rest package container entry id base string.
	ContainerID = slate.ContainerID + ".rest"

	// ContainerEngineID defines the default id used to register the
	// application gin engine instance in the application container.
	ContainerEngineID = ContainerID + ".engine"
)

const (
	// EnvID defines the rest package base environment variable name.
	EnvID = slate.EnvID + "_REST"
)

var (
	// ConfigPortPath contains the config path of the server port to be used.
	ConfigPortPath = senv.String(EnvID+"_CONFIG_PORT_PATH", "slate.rest.server.port")

	// LogChannel contains the name of the logging channel to be used on
	// the rest application messages.
	LogChannel = senv.String(EnvID+"_LOG_CHANNEL", "exec")
)
