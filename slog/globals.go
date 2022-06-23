package slog

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/senv"
)

const (
	// FormatUnknown defines the value to be used to declare an unknown
	// logger formatter format.
	FormatUnknown = "unknown"

	// FormatJSON defines the value to be used to declare a JSON
	// logger formatter format.
	FormatJSON = "json"
)

const (
	// StreamUnknown defines the value to be used to declare an unknown
	// logger stream type.
	StreamUnknown = "unknown"

	// StreamConsole defines the value to be used to declare a console
	// logger stream type.
	StreamConsole = "console"

	// StreamFile defines the value to be used to declare a file
	// logger stream type.
	StreamFile = "file"

	// StreamRotatingFile defines the value to be used to declare a file
	// logger stream type that rotates regarding the current date.
	StreamRotatingFile = "rotating-file"
)

const (
	// ContainerID defines the id to be used as the container
	// registration id of a logger instance, as a base id of all other log
	// package instances registered in the application container.
	ContainerID = slate.ContainerID + ".log"

	// ContainerFormatterStrategyTag defines the tag to be assigned to all
	// container formatter strategies.
	ContainerFormatterStrategyTag = ContainerID + ".formatter.strategy"

	// ContainerFormatterStrategyJSONID defines the id to be used as
	// the container registration id of a logger json formatter factory
	// strategy instance.
	ContainerFormatterStrategyJSONID = ContainerID + ".formatter.strategy.json"

	// ContainerFormatterFactoryID defines the id to be used as the
	// container registration id of a logger formatter factory instance.
	ContainerFormatterFactoryID = ContainerID + ".formatter.factory"

	// ContainerStreamStrategyTag defines the tag to be assigned to all
	// container stream strategies.
	ContainerStreamStrategyTag = ContainerID + ".stream.strategy"

	// ContainerStreamStrategyConsoleID defines the id to be used as the
	// container registration id of a logger console stream factory strategy
	// instance.
	ContainerStreamStrategyConsoleID = ContainerID + ".stream.strategy.console"

	// ContainerStreamStrategyFileID defines the id to be used as the
	// container registration id of a logger file stream factory strategy
	// instance.
	ContainerStreamStrategyFileID = ContainerID + ".stream.strategy.file"

	// ContainerStreamStrategRotatingFileID defines the id to be used as the
	// container registration id of a logger rotating file stream factory
	// strategy instance.
	ContainerStreamStrategRotatingFileID = ContainerID + ".stream.strategy.rotating_file"

	// ContainerStreamFactoryID defines the id to be used as the container
	// registration id of a logger stream factory instance.
	ContainerStreamFactoryID = ContainerID + ".stream.factory"

	// ContainerLoaderID defines the id to be used as the container
	// registration id of a logger loader instance.
	ContainerLoaderID = ContainerID + ".loader"
)

const (
	// EnvID defines the slate.slog package base environment variable name.
	EnvID = slate.EnvID + "_SLOG"
)

var (
	// LoaderActive defines the entry config source active flag
	// used to signal the config loader to load the streams or not
	LoaderActive = senv.Bool(EnvID+"_LOADER_ACTIVE", true)

	// LoaderConfigPath defines the entry config source path
	// to be used as the loader entry.
	LoaderConfigPath = senv.String(EnvID+"_LOADER_CONFIG_PATH", "log.streams")

	// LoaderObserveConfig defines the loader config observing flag
	// used to register in the config object an observer of the log
	// config entries list, so it can reload the logger streams.
	LoaderObserveConfig = senv.Bool(EnvID+"_LOADER_OBSERVE_CONFIG", true)

	// LoaderErrorChannel defines the loader error logging channel.
	LoaderErrorChannel = senv.String(EnvID+"_LOADER_ERROR_CHANNEL", "exec")
)
