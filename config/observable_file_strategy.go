package config

import (
	"github.com/spf13/afero"
)

// ObservableFileSourceStrategy defines a strategy used to instantiate
// an observable file config source creation strategy.
type ObservableFileSourceStrategy struct {
	FileSourceStrategy
}

var _ ISourceStrategy = &ObservableFileSourceStrategy{}

type observableFileSourceConfig struct {
	Path   string
	Format string
}

// NewObservableFileSourceStrategy instantiates a new observable
// file config source creation strategy.
func NewObservableFileSourceStrategy(
	fs afero.Fs,
	decoderFactory IDecoderFactory,
) (*ObservableFileSourceStrategy, error) {
	// check the file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &ObservableFileSourceStrategy{
		FileSourceStrategy: FileSourceStrategy{
			fs:             fs,
			decoderFactory: decoderFactory,
		},
	}, nil
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration instance.
func (s ObservableFileSourceStrategy) Accept(
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
		return sc.Type == SourceObservableFile
	}
	return false
}

// Create will instantiate the desired file source instance where
// the initialization data comes from a configuration instance.
func (s ObservableFileSourceStrategy) Create(
	config IConfig,
) (ISource, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := observableFileSourceConfig{Format: DefaultFileFormat}
	_, e := config.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.Path == "" {
		return nil, errPathNotFound("path")
	}
	// return acceptance for the read config type
	return NewObservableFileSource(sc.Path, sc.Format, s.fs, s.decoderFactory)
}
