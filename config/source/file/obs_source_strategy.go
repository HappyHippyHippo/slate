package file

import (
	"github.com/happyhippyhippo/slate/config"
	"github.com/spf13/afero"
)

const (
	// ObsType defines the value to be used to
	// declare an observable file config source type.
	ObsType = "observable-file"
)

// ObsSourceStrategy defines a strategy used to instantiate
// an observable file config source creation strategy.
type ObsSourceStrategy struct {
	SourceStrategy
}

var _ config.SourceStrategy = &ObsSourceStrategy{}

// NewObsSourceStrategy instantiates a new observable
// file config source creation strategy.
func NewObsSourceStrategy(
	fileSystem afero.Fs,
	decoderFactory *config.DecoderFactory,
) (*ObsSourceStrategy, error) {
	// check the file system argument reference
	if fileSystem == nil {
		return nil, errNilPointer("fileSystem")
	}
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &ObsSourceStrategy{
		SourceStrategy: SourceStrategy{
			fileSystem:     fileSystem,
			decoderFactory: decoderFactory,
		},
	}, nil
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration instance.
func (s ObsSourceStrategy) Accept(
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
	return sc.Type == ObsType
}

// Create will instantiate the desired file source instance where
// the initialization data comes from a configuration instance.
func (s ObsSourceStrategy) Create(
	cfg config.Partial,
) (config.Source, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("partial")
	}
	// retrieve the data from the configuration
	sc := struct {
		Path   string
		Format string
	}{
		Format: config.DefaultFileFormat,
	}
	if _, e := cfg.Populate("", &sc); e != nil {
		return nil, e
	}
	// validate configuration
	if sc.Path == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{
			"description": "missing path",
		})
	}
	// return acceptance for the read config type
	return NewObsSource(sc.Path, sc.Format, s.fileSystem, s.decoderFactory)
}
