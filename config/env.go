package config

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/env"
)

const (
	// EnvID defines the slate.config package base environment variable name.
	EnvID = slate.EnvID + "_CONFIG"
)

var (
	// DefaultFileFormat defines the file base config source default
	// format if the format is not present in the config.
	DefaultFileFormat = env.String(EnvID+"_DEFAULT_FILE_FORMAT", FormatYAML)

	// DefaultRestFormat defines the rest base config source default
	// format if the format is not present in the config.
	DefaultRestFormat = env.String(EnvID+"_DEFAULT_REST_FORMAT", FormatJSON)

	// PathSeparator defines the element(s) that will be used to split
	// a config path string into path elements.
	PathSeparator = env.String(EnvID+"_PATH_SEPARATOR", ".")

	// LoaderActive defines if the config loader should be executed
	// while the provider boot
	LoaderActive = env.Bool(EnvID+"_LOADER_ACTIVE", true)

	// LoaderSourceID defines the id to be used as the default of the
	// entry config source id to be used as the loader entry.
	LoaderSourceID = env.String(EnvID+"_LOADER_SOURCE_ID", "_sources")

	// LoaderSourcePath defines the entry config source path
	// to be used as the loader entry.
	LoaderSourcePath = env.String(EnvID+"_LOADER_SOURCE_PATH", "config/config.yaml")

	// LoaderSourceFormat defines the entry config source format
	// to be used as the loader entry.
	LoaderSourceFormat = env.String(EnvID+"_LOADER_SOURCE_FORMAT", FormatYAML)

	// LoaderSourceListPath defines the entry config source path of
	// loading sources.
	LoaderSourceListPath = env.String(EnvID+"_LOADER_SOURCE_LIST_PATH", "slate.config.list")

	// ObserveFrequency defines the id to be used as the default of a
	// config observable source frequency time in seconds.
	ObserveFrequency = env.Int(EnvID+"_OBSERVE_FREQUENCY", 0)
)
