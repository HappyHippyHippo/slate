package gconfig

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/genv"
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

	// SourceTypeRemote defines the value to be used to declare a
	// remote config source type.
	SourceTypeRemote = "remote"

	// SourceTypeObservableRemote defines the value to be used to
	// declare an observable remote config source type.
	SourceTypeObservableRemote = "observable-remote"

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
	ContainerSourceStrategyFileObservableID = ContainerID + ".source.strategy.file_observable"

	// ContainerSourceStrategyDirID defines the id to be used as the
	// container registration id of a config dir source factory strategy
	// instance.
	ContainerSourceStrategyDirID = ContainerID + ".source.strategy.dir"

	// ContainerSourceStrategyRemoteID defines the id to be used as the
	// container registration id of a config remote source factory strategy
	// instance.
	ContainerSourceStrategyRemoteID = ContainerID + ".source.strategy.remote"

	// ContainerSourceStrategyRemoteObservableID defines the id to be used
	// as the container registration id of a config observable remote source
	// factory strategy instance.
	ContainerSourceStrategyRemoteObservableID = ContainerID + ".source.strategy.remote_observable"

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

var (
	// DefaultFileFormat defines the file base config source default format
	// if the format is not present in the config.
	DefaultFileFormat = genv.String(
		"SLATE_GCONFIG_DEFAULT_FILE_FORMAT",
		DecoderFormatYAML)

	// DefaultRemoteFormat defines the remote base config source default format
	// if the format is not present in the config.
	DefaultRemoteFormat = genv.String(
		"SLATE_GCONFIG_DEFAULT_REMOTE_FORMAT",
		DecoderFormatJSON)

	// PathSeparator defines the element(s) that will be used to split
	// a config path string into path elements.
	PathSeparator = genv.String(
		"SLATE_GCONFIG_PATH_SEPARATOR",
		".")

	// LoaderSourceID defines the id to be used as the default of the
	// entry config source id to be used as the loader entry.
	LoaderSourceID = genv.String(
		"SLATE_GCONFIG_LOADER_SOURCE_ID",
		"_sources")

	// LoaderSourcePath defines the entry config source path
	// to be used as the loader entry.
	LoaderSourcePath = genv.String(
		"SLATE_GCONFIG_LOADER_SOURCE_PATH",
		"config/config.yaml")

	// LoaderSourceFormat defines the entry config source format
	// to be used as the loader entry.
	LoaderSourceFormat = genv.String(
		"SLATE_GCONFIG_LOADER_SOURCE_FORMAT",
		DecoderFormatYAML)

	// LoaderSourceListPath defines the entry config source path of
	// loading sources.
	LoaderSourceListPath = genv.String(
		"SLATE_GCONFIG_LOADER_SOURCE_LIST_PATH",
		"configs")

	// ObserveFrequency defines the id to be used as the default of a
	// config observable source frequency time in seconds.
	ObserveFrequency = genv.Int(
		"SLATE_GCONFIG_OBSERVE_FREQUENCY",
		0)

	// LoaderActive defines if the config loader should be executed
	// while the provider boot
	LoaderActive = genv.Bool(
		"SLATE_GCONFIG_LOADER_ACTIVE",
		true)
)
