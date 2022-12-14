package log

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/env"
)

const (
	// EnvID defines the log package base environment variable name.
	EnvID = slate.EnvID + "_LOG"
)

var (
	// LoaderActive defines the entry config source active flag
	// used to signal the config loader to load the streams or not
	LoaderActive = env.Bool(EnvID+"_LOADER_ACTIVE", true)

	// LoaderConfigPath defines the entry config source path
	// to be used as the loader entry.
	LoaderConfigPath = env.String(EnvID+"_LOADER_CONFIG_PATH", "slate.log.streams")

	// LoaderObserveConfig defines the loader config observing flag
	// used to register in the config object an observer of the log
	// config entries list, so it can reload the Log streams.
	LoaderObserveConfig = env.Bool(EnvID+"_LOADER_OBSERVE_CONFIG", true)

	// LoaderErrorChannel defines the loader err logging channel.
	LoaderErrorChannel = env.String(EnvID+"_LOADER_ERROR_CHANNEL", "exec")
)
