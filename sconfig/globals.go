package sconfig

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/senv"
)

const (
	// DecoderFormatUnknown defines the value to be used to declare an
	// unknown sconfig source format.
	DecoderFormatUnknown = "unknown"

	// DecoderFormatYAML defines the value to be used to declare a YAML
	// sconfig source format.
	DecoderFormatYAML = "yaml"

	// DecoderFormatJSON defines the value to be used to declare a JSON
	// sconfig source format.
	DecoderFormatJSON = "json"
)

const (
	// SourceTypeUnknown defines the value to be used to declare an
	// unknown sconfig source type.
	SourceTypeUnknown = "unknown"

	// SourceTypeFile defines the value to be used to declare a
	// simple file sconfig source type.
	SourceTypeFile = "file"

	// SourceTypeObservableFile defines the value to be used to
	// declare an observable file sconfig source type.
	SourceTypeObservableFile = "observable-file"

	// SourceTypeDirectory defines the value to be used to declare a
	// simple dir sconfig source type.
	SourceTypeDirectory = "dir"

	// SourceTypeRest defines the value to be used to declare a
	// rest sconfig source type.
	SourceTypeRest = "rest"

	// SourceTypeObservableRest defines the value to be used to
	// declare an observable rest sconfig source type.
	SourceTypeObservableRest = "observable-rest"

	// SourceTypeEnv defines the value to be used to declare an
	// environment sconfig source type.
	SourceTypeEnv = "senv"

	// SourceTypeContainer defines the value to be used to declare a
	// container loading configs source type.
	SourceTypeContainer = "container"
)

const (
	// ContainerID defines the id to be used as the container
	// registration id of a sconfig instance, as a base id of all other sconfig
	// package instances registered in the application container.
	ContainerID = slate.ContainerID + ".sconfig"

	// ContainerDecoderStrategyTag defines the tag to be assigned to all
	// container decoders strategies.
	ContainerDecoderStrategyTag = ContainerID + ".decoder.strategy"

	// ContainerDecoderStrategyYAMLID defines the id to be used as the
	// container registration id of a yaml sconfig decoder factory strategy
	// instance.
	ContainerDecoderStrategyYAMLID = ContainerID + ".decoder.strategy.yaml"

	// ContainerDecoderStrategyJSONID defines the id to be used as the
	// container registration id of a json sconfig decoder factory strategy
	// instance.
	ContainerDecoderStrategyJSONID = ContainerID + ".decoder.strategy.json"

	// ContainerDecoderFactoryID defines the id to be used as the
	// container registration id of a sconfig decoder factory instance.
	ContainerDecoderFactoryID = ContainerID + ".decoder.factory"

	// ContainerSourceStrategyTag defines the tag to be assigned to all
	// container source strategies.
	ContainerSourceStrategyTag = ContainerID + ".source.strategy"

	// ContainerSourceStrategyFileID defines the id to be used as the
	// container registration id of a sconfig file source factory strategy
	// instance.
	ContainerSourceStrategyFileID = ContainerID + ".source.strategy.file"

	// ContainerSourceStrategyFileObservableID defines the id to be used
	// as the container registration id of a sconfig observable file source
	// factory strategy instance.
	ContainerSourceStrategyFileObservableID = ContainerID + ".source.strategy.observable_file"

	// ContainerSourceStrategyDirID defines the id to be used as the
	// container registration id of a sconfig dir source factory strategy
	// instance.
	ContainerSourceStrategyDirID = ContainerID + ".source.strategy.dir"

	// ContainerSourceStrategyRestID defines the id to be used as the
	// container registration id of a sconfig rest source factory strategy
	// instance.
	ContainerSourceStrategyRestID = ContainerID + ".source.strategy.rest"

	// ContainerSourceStrategyRestObservableID defines the id to be used
	// as the container registration id of a sconfig observable rest source
	// factory strategy instance.
	ContainerSourceStrategyRestObservableID = ContainerID + ".source.strategy.observable_rest"

	// ContainerSourceStrategyEnvID defines the id to be used as
	// the container registration id of a sconfig environment source
	// factory strategy instance.
	ContainerSourceStrategyEnvID = ContainerID + ".source.strategy.senv"

	// ContainerSourceStrategyContainerID defines the id to be used as
	// the container registration id of a container loading sconfig source
	// factory strategy instance.
	ContainerSourceStrategyContainerID = ContainerID + ".source.strategy.container"

	// ContainerSourceFactoryID defines the id to be used as the
	// container registration id sconfig source factory instance.
	ContainerSourceFactoryID = ContainerID + ".source.factory"

	// ContainerSourceContainerPartialTag defines the tag to be assigned
	// to all container defined sconfig partials.
	ContainerSourceContainerPartialTag = ContainerID + ".source.container.tag"

	// ContainerLoaderID defines the id to be used as the container
	// registration id of a sconfig loader instance.
	ContainerLoaderID = ContainerID + ".loader"
)

const (
	// EnvID defines the slate.sconfig package base environment variable name.
	EnvID = slate.EnvID + "_CONFIG"
)

var (
	// DefaultFileFormat defines the file base sconfig source default format
	// if the format is not present in the sconfig.
	DefaultFileFormat = senv.String(EnvID+"_DEFAULT_FILE_FORMAT", DecoderFormatYAML)

	// DefaultRestFormat defines the rest base sconfig source default format
	// if the format is not present in the sconfig.
	DefaultRestFormat = senv.String(EnvID+"_DEFAULT_REST_FORMAT", DecoderFormatJSON)

	// PathSeparator defines the element(s) that will be used to split
	// a sconfig path string into path elements.
	PathSeparator = senv.String(EnvID+"_PATH_SEPARATOR", ".")

	// LoaderSourceID defines the id to be used as the default of the
	// entry sconfig source id to be used as the loader entry.
	LoaderSourceID = senv.String(EnvID+"_LOADER_SOURCE_ID", "_sources")

	// LoaderSourcePath defines the entry sconfig source path
	// to be used as the loader entry.
	LoaderSourcePath = senv.String(EnvID+"_LOADER_SOURCE_PATH", "sconfig/sconfig.yaml")

	// LoaderSourceFormat defines the entry sconfig source format
	// to be used as the loader entry.
	LoaderSourceFormat = senv.String(EnvID+"_LOADER_SOURCE_FORMAT", DecoderFormatYAML)

	// LoaderSourceListPath defines the entry sconfig source path of
	// loading sources.
	LoaderSourceListPath = senv.String(EnvID+"_LOADER_SOURCE_LIST_PATH", "configs")

	// ObserveFrequency defines the id to be used as the default of a
	// sconfig observable source frequency time in seconds.
	ObserveFrequency = senv.Int(EnvID+"_OBSERVE_FREQUENCY", 0)

	// LoaderActive defines if the sconfig loader should be executed
	// while the provider boot
	LoaderActive = senv.Bool(EnvID+"_LOADER_ACTIVE", true)
)
