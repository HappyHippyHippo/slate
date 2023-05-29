package rest

import (
	"net/http"

	"github.com/happyhippyhippo/slate/config"
)

const (
	// Type defines the value to be used to declare a
	// rest config source type.
	Type = "rest"
)

// SourceStrategy defines a strategy used to instantiate
// a REST service config source creation strategy.
type SourceStrategy struct {
	clientFactory  func() requester
	decoderFactory *config.DecoderFactory
}

var _ config.SourceStrategy = &SourceStrategy{}

// NewSourceStrategy instantiates a new REST service config
// source creation strategy.
func NewSourceStrategy(
	decoderFactory *config.DecoderFactory,
) (*SourceStrategy, error) {
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderCreator")
	}
	// instantiate the strategy
	return &SourceStrategy{
		clientFactory:  func() requester { return &http.Client{} },
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s SourceStrategy) Accept(
	cfg *config.Partial,
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

// Create will instantiate the desired rest source instance where
// the initialization data comes from a configuration instance.
func (s SourceStrategy) Create(
	cfg *config.Partial,
) (config.Source, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	sc := struct {
		URI    string
		Format string
		Path   struct {
			Config string
		}
	}{
		Format: config.DefaultRestFormat,
	}
	if _, e := cfg.Populate("", &sc); e != nil {
		return nil, e
	}
	// validate configuration
	if sc.URI == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{"description": "missing URI"})
	}
	if sc.Path.Config == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{"description": "missing response config path"})
	}
	// return acceptance for the read config type
	return NewSource(
		s.clientFactory(),
		sc.URI,
		sc.Format,
		s.decoderFactory,
		sc.Path.Config,
	)
}
