package config

import (
	"net/http"
)

// RestSourceStrategy defines a strategy used to instantiate
// a REST service config source creation strategy.
type RestSourceStrategy struct {
	clientFactory  func() httpClient
	decoderFactory IDecoderFactory
}

var _ ISourceStrategy = &RestSourceStrategy{}

type restSourceConfig struct {
	URI    string
	Format string
	Path   struct {
		Config string
	}
}

// NewRestSourceStrategy instantiates a new REST service config
// source creation strategy.
func NewRestSourceStrategy(
	decoderFactory IDecoderFactory,
) (*RestSourceStrategy, error) {
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &RestSourceStrategy{
		clientFactory:  func() httpClient { return &http.Client{} },
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s RestSourceStrategy) Accept(
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
		return sc.Type == SourceRest
	}
	return false
}

// Create will instantiate the desired rest source instance where
// the initialization data comes from a configuration instance.
func (s RestSourceStrategy) Create(
	config IConfig,
) (ISource, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := restSourceConfig{Format: DefaultRestFormat}
	_, e := config.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.URI == "" {
		return nil, errPathNotFound("uri")
	}
	if sc.Path.Config == "" {
		return nil, errPathNotFound("path.config")
	}
	// return acceptance for the read config type
	return NewRestSource(
		s.clientFactory(),
		sc.URI,
		sc.Format,
		s.decoderFactory,
		sc.Path.Config,
	)
}
