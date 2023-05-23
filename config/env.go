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
	// DefaultFileFormat defines the file base config source def
	// format if the format is not present in the config.
	DefaultFileFormat = env.String(EnvID+"_DEFAULT_FILE_FORMAT", "yaml")

	// DefaultRestFormat defines the rest base config source def
	// format if the format is not present in the config.
	DefaultRestFormat = env.String(EnvID+"_DEFAULT_REST_FORMAT", "json")

	// PathSeparator defines the element(s) that will be used to split
	// a config path string into path elements.
	PathSeparator = env.String(EnvID+"_PATH_SEPARATOR", ".")

	// LoaderActive defines if the config observer should be executed
	// while the provider boot
	LoaderActive = env.Bool(EnvID+"_LOADER_ACTIVE", true)

	// LoaderSourceID defines the id to be used as the def of the
	// entry config source id to be used as the observer entry.
	LoaderSourceID = env.String(EnvID+"_LOADER_SOURCE_ID", "_sources")

	// LoaderSourcePath defines the entry config source path
	// to be used as the observer entry.
	LoaderSourcePath = env.String(EnvID+"_LOADER_SOURCE_PATH", "config/sources.yaml")

	// LoaderSourceFormat defines the entry config source format
	// to be used as the observer entry.
	LoaderSourceFormat = env.String(EnvID+"_LOADER_SOURCE_FORMAT", "yaml")

	// LoaderSourceListPath defines the entry config source path of
	// loading sources.
	LoaderSourceListPath = env.String(EnvID+"_LOADER_SOURCE_LIST_PATH", "slate.config.sources")

	// ObserveFrequency defines the id to be used as the def of a
	// config observable source frequency time in seconds.
	ObserveFrequency = env.Int(EnvID+"_OBSERVE_FREQUENCY", 0)
)
