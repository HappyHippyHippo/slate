package dir

import (
	"github.com/happyhippyhippo/slate/config"
	"github.com/spf13/afero"
)

const (
	// Type defines the value to be used to declare a
	// simple dir config source type.
	Type = "dir"
)

// SourceStrategy defines a strategy used to instantiate
// a dir config source creation strategy.
type SourceStrategy struct {
	fileSystem     afero.Fs
	decoderFactory *config.DecoderFactory
}

var _ config.SourceStrategy = &SourceStrategy{}

// NewSourceStrategy instantiates a new dir config
// source creation strategy.
func NewSourceStrategy(
	fileSystem afero.Fs,
	decoderFactory *config.DecoderFactory,
) (*SourceStrategy, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &SourceStrategy{
		fileSystem:     fileSystem,
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s SourceStrategy) Accept(
	cfg config.Partial,
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
	return sc.Type == Type
}

// Create will instantiate the desired file source instance where
// the initialization data comes from a configuration instance.
func (s SourceStrategy) Create(
	cfg config.Partial,
) (config.Source, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	sc := struct {
		Path      string
		Format    string
		Recursive bool
	}{
		Format:    config.DefaultFileFormat,
		Recursive: false,
	}
	if _, e := cfg.Populate("", &sc); e != nil {
		return nil, e
	}
	// validate configuration
	if sc.Path == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{"description": "missing path"})
	}
	// return acceptance for the read config type
	return NewSource(
		sc.Path,
		sc.Format,
		sc.Recursive,
		s.fileSystem,
		s.decoderFactory,
	)
}
