package config

import (
	"github.com/spf13/afero"
)

const (
	// ObservableFileSourceType defines the value to be used to
	// declare an observable file config source type.
	ObservableFileSourceType = "observable-file"
)

type observableFileSourceConfig struct {
	Path   string
	Format string
}

// ObservableFileSourceStrategy defines a strategy used to instantiate
// an observable file config source creation strategy.
type ObservableFileSourceStrategy struct {
	FileSourceStrategy
}

var _ ISourceStrategy = &ObservableFileSourceStrategy{}

// NewObservableFileSourceStrategy instantiates a new observable
// file config source creation strategy.
func NewObservableFileSourceStrategy(
	fileSystem afero.Fs,
	decoderFactory IDecoderFactory,
) (*ObservableFileSourceStrategy, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &ObservableFileSourceStrategy{
		FileSourceStrategy: FileSourceStrategy{
			fileSystem:     fileSystem,
			decoderFactory: decoderFactory,
		},
	}, nil
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration instance.
func (s ObservableFileSourceStrategy) Accept(
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
	return sc.Type == ObservableFileSourceType
}

// Create will instantiate the desired file source instance where
// the initialization data comes from a configuration instance.
func (s ObservableFileSourceStrategy) Create(
	cfg IConfig,
) (ISource, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := observableFileSourceConfig{Format: DefaultFileFormat}
	_, e := cfg.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.Path == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{"description": "missing path"})
	}
	// return acceptance for the read config type
	return NewObservableFileSource(sc.Path, sc.Format, s.fileSystem, s.decoderFactory)
}
