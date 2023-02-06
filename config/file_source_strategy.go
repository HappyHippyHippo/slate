package config

import (
	"github.com/spf13/afero"
)

const (
	// SourceStrategyFile defines the value to be used to declare a
	// simple file config source type.
	SourceStrategyFile = "file"
)

type fileSourceConfig struct {
	Path   string
	Format string
}

// FileSourceStrategy defines a strategy used to instantiate a
// file config source creation strategy.
type FileSourceStrategy struct {
	fileSystem     afero.Fs
	decoderFactory IDecoderFactory
}

var _ ISourceStrategy = &FileSourceStrategy{}

// NewFileSourceStrategy instantiates a new file config source
// creation strategy.
func NewFileSourceStrategy(
	fileSystem afero.Fs,
	decoderFactory IDecoderFactory,
) (*FileSourceStrategy, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &FileSourceStrategy{
		fileSystem:     fileSystem,
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s FileSourceStrategy) Accept(
	config IConfig,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	_, e := config.Populate("", &sc)
	if e == nil {
		// return acceptance for the read config type
		return sc.Type == SourceStrategyFile
	}
	return false
}

// Create will instantiate the desired file source instance where
// the initialization data comes from a configuration instance.
func (s FileSourceStrategy) Create(
	config IConfig,
) (ISource, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := fileSourceConfig{Format: DefaultFileFormat}
	_, e := config.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.Path == "" {
		return nil, errPathNotFound("path")
	}
	// return acceptance for the read config type
	return NewFileSource(sc.Path, sc.Format, s.fileSystem, s.decoderFactory)
}
