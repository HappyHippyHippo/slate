package config

import (
	"github.com/spf13/afero"
)

const (
	// DirectorySourceType defines the value to be used to declare a
	// simple dir config source type.
	DirectorySourceType = "dir"
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
	cfg IConfig,
) bool {
	// check the config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	if _, e := cfg.Populate("", &sc); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sc.Type == DirectorySourceType
}

// Create will instantiate the desired file source instance where
// the initialization data comes from a configuration instance.
func (s DirSourceStrategy) Create(
	cfg IConfig,
) (ISource, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := dirSourceConfig{Format: DefaultFileFormat}
	_, e := cfg.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.Path == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{"description": "missing path"})
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
