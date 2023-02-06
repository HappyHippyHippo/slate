package config

import (
	"github.com/spf13/afero"
)

const (
	// SourceStrategyDirectory defines the value to be used to declare a
	// simple dir config source type.
	SourceStrategyDirectory = "dir"
)

type dirSourceConfig struct {
	Path      string
	Format    string
	Recursive bool
}

// DirSourceStrategy defines a strategy used to instantiate
// a dir config source creation strategy.
type DirSourceStrategy struct {
	fileSystem     afero.Fs
	decoderFactory IDecoderFactory
}

var _ ISourceStrategy = &DirSourceStrategy{}

// NewDirSourceStrategy instantiates a new dir config
// source creation strategy.
func NewDirSourceStrategy(
	fileSystem afero.Fs,
	decoderFactory IDecoderFactory,
) (*DirSourceStrategy, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &DirSourceStrategy{
		fileSystem:     fileSystem,
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s DirSourceStrategy) Accept(
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
		return sc.Type == SourceStrategyDirectory
	}
	return false
}

// Create will instantiate the desired file source instance where
// the initialization data comes from a configuration instance.
func (s DirSourceStrategy) Create(
	config IConfig,
) (ISource, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := dirSourceConfig{Format: DefaultFileFormat}
	_, e := config.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.Path == "" {
		return nil, errPathNotFound("path")
	}
	// return acceptance for the read config type
	return NewDirSource(
		sc.Path,
		sc.Format,
		sc.Recursive,
		s.fileSystem,
		s.decoderFactory,
	)
}
