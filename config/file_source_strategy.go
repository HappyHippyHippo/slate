package config

import (
	"github.com/spf13/afero"
)

// FileSourceStrategy defines a strategy used to instantiate a
// file config source creation strategy.
type FileSourceStrategy struct {
	fs             afero.Fs
	decoderFactory IDecoderFactory
}

var _ ISourceStrategy = &FileSourceStrategy{}

// NewFileSourceStrategy instantiates a new file config source
// creation strategy.
func NewFileSourceStrategy(
	fs afero.Fs,
	decoderFactory IDecoderFactory,
) (*FileSourceStrategy, error) {
	// check the file system argument reference
	if fs == nil {
		return nil, errNilPointer("fs")
	}
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &FileSourceStrategy{
		fs:             fs,
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
		return sc.Type == SourceFile
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
	sc := struct {
		Path   string
		Format string
	}{Format: DefaultFileFormat}
	_, e := config.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.Path == "" {
		return nil, errPathNotFound("path")
	}
	// return acceptance for the read config type
	return NewFileSource(sc.Path, sc.Format, s.fs, s.decoderFactory)
}
