package sconfig

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/senv"
)

const (
	// DecoderFormatUnknown defines the value to be used to declare an
	// unknown config source format.
	DecoderFormatUnknown = "unknown"

	// DecoderFormatYAML defines the value to be used to declare a YAML
	// config source format.
	DecoderFormatYAML = "yaml"

	// DecoderFormatJSON defines the value to be used to declare a JSON
	// config source format.
	DecoderFormatJSON = "json"
)

const (
	// SourceTypeUnknown defines the value to be used to declare a
	// unknown config source type.
	SourceTypeUnknown = "unknown"

	// SourceTypeFile defines the value to be used to declare a
	// simple file config source type.
	SourceTypeFile = "file"

	// SourceTypeObservableFile defines the value to be used to
	// declare an observable file config source type.
	SourceTypeObservableFile = "observable-file"

	// SourceTypeDirectory defines the value to be used to declare a
	// simple dir config source type.
	SourceTypeDirectory = "dir"

	// SourceTypeRest defines the value to be used to declare a
	// rest config source type.
	SourceTypeRest = "rest"

	// SourceTypeObservableRest defines the value to be used to
	// declare an observable rest config source type.
	SourceTypeObservableRest = "observable-rest"

	// SourceTypeEnv defines the value to be used to declare an
	// environment config source type.
	SourceTypeEnv = "env"
)

const (
	// ContainerID defines the id to be used as the container
	// registration id of a config instance, as a base id of all other config
	// package instances registered in the application container.
	ContainerID = slate.ContainerID + ".config"

	// ContainerDecoderStrategyTag defines the tag to be assigned to all
	// container decoders strategies.
	ContainerDecoderStrategyTag = ContainerID + ".decoder.strategy"

	// ContainerDecoderStrategyYAMLID defines the id to be used as the
	// container registration id of a yaml config decoder factory strategy
	// instance.
	ContainerDecoderStrategyYAMLID = ContainerID + ".decoder.strategy.yaml"

	// ContainerDecoderStrategyJSONID defines the id to be used as the
	// container registration id of a json config decoder factory strategy
	// instance.
	ContainerDecoderStrategyJSONID = ContainerID + ".decoder.strategy.json"

	// ContainerDecoderFactoryID defines the id to be used as the
	// container registration id of a config decoder factory instance.
	ContainerDecoderFactoryID = ContainerID + ".decoder.factory"

	// ContainerSourceStrategyTag defines the tag to be assigned to all
	// container source strategies.
	ContainerSourceStrategyTag = ContainerID + ".source.strategy"

	// ContainerSourceStrategyFileID defines the id to be used as the
	// container registration id of a config file source factory strategy
	// instance.
	ContainerSourceStrategyFileID = ContainerID + ".source.strategy.file"

	// ContainerSourceStrategyFileObservableID defines the id to be used
	// as the container registration id of a config observable file source
	// factory strategy instance.
	ContainerSourceStrategyFileObservableID = ContainerID + ".source.strategy.observable_file"

	// ContainerSourceStrategyDirID defines the id to be used as the
	// container registration id of a config dir source factory strategy
	// instance.
	ContainerSourceStrategyDirID = ContainerID + ".source.strategy.dir"

	// ContainerSourceStrategyRestID defines the id to be used as the
	// container registration id of a config rest source factory strategy
	// instance.
	ContainerSourceStrategyRestID = ContainerID + ".source.strategy.rest"

	// ContainerSourceStrategyRestObservableID defines the id to be used
	// as the container registration id of a config observable rest source
	// factory strategy instance.
	ContainerSourceStrategyRestObservableID = ContainerID + ".source.strategy.observable_rest"

	// ContainerSourceStrategyEnvID defines the id to be used as
	// the container registration id of a config environment source
	// factory strategy instance.
	ContainerSourceStrategyEnvID = ContainerID + ".source.strategy.env"

	// ContainerSourceFactoryID defines the id to be used as the
	// container registration id config source factory instance.
	ContainerSourceFactoryID = ContainerID + ".source.factory"

	// ContainerLoaderID defines the id to be used as the container
	// registration id of a config loader instance.
	ContainerLoaderID = ContainerID + ".loader"
)

const (
	// EnvID defines the slate.sconfig package base environment variable name.
	EnvID = slate.EnvID + "_SCONFIG"
)

var (
	// DefaultFileFormat defines the file base config source default format
	// if the format is not present in the config.
	DefaultFileFormat = senv.String(EnvID+"_DEFAULT_FILE_FORMAT", DecoderFormatYAML)

	// DefaultRestFormat defines the rest base config source default format
	// if the format is not present in the config.
	DefaultRestFormat = senv.String(EnvID+"_DEFAULT_REST_FORMAT", DecoderFormatJSON)

	// PathSeparator defines the element(s) that will be used to split
	// a config path string into path elements.
	PathSeparator = senv.String(EnvID+"_PATH_SEPARATOR", ".")

	// LoaderSourceID defines the id to be used as the default of the
	// entry config source id to be used as the loader entry.
	LoaderSourceID = senv.String(EnvID+"_LOADER_SOURCE_ID", "_sources")

	// LoaderSourcePath defines the entry config source path
	// to be used as the loader entry.
	LoaderSourcePath = senv.String(EnvID+"_LOADER_SOURCE_PATH", "config/config.yaml")

	// LoaderSourceFormat defines the entry config source format
	// to be used as the loader entry.
	LoaderSourceFormat = senv.String(EnvID+"_LOADER_SOURCE_FORMAT", DecoderFormatYAML)

	// LoaderSourceListPath defines the entry config source path of
	// loading sources.
	LoaderSourceListPath = senv.String(EnvID+"_LOADER_SOURCE_LIST_PATH", "configs")

	// ObserveFrequency defines the id to be used as the default of a
	// config observable source frequency time in seconds.
	ObserveFrequency = senv.Int(EnvID+"_OBSERVE_FREQUENCY", 0)

	// LoaderActive defines if the config loader should be executed
	// while the provider boot
	LoaderActive = senv.Bool(EnvID+"_LOADER_ACTIVE", true)
)
