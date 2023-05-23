package rest

import (
	"net/http"

	"github.com/happyhippyhippo/slate/config"
)

const (
	// ObsType defines the value to be used to
	// declare an observable rest config source type.
	ObsType = "observable-rest"
)

type obsSourceConfig struct {
	URI    string
	Format string
	Path   struct {
		Config    string
		Timestamp string
	}
}

// ObsSourceStrategy defines a strategy used to instantiate
// an observable REST service config source creation strategy.
type ObsSourceStrategy struct {
	SourceStrategy
}

var _ config.SourceStrategy = &ObsSourceStrategy{}

// NewObsSourceStrategy instantiates a new observable REST
// service config source creation strategy.
func NewObsSourceStrategy(
	decoderFactory *config.DecoderFactory,
) (*ObsSourceStrategy, error) {
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &ObsSourceStrategy{
		SourceStrategy: SourceStrategy{
			clientFactory:  func() httpClient { return &http.Client{} },
			decoderFactory: decoderFactory,
		},
	}, nil
}

// Accept will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s ObsSourceStrategy) Accept(
	partial *config.Partial,
) bool {
	// check the config argument reference
	if partial == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	if _, e := partial.Populate("", &sc); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sc.Type == ObsType
}

// Create will instantiate the desired rest source instance where
// the initialization data comes from a configuration instance.
func (s ObsSourceStrategy) Create(
	partial *config.Partial,
) (config.Source, error) {
	// check the config argument reference
	if partial == nil {
		return nil, errNilPointer("partial")
	}
	// retrieve the data from the configuration
	sc := obsSourceConfig{Format: config.DefaultRestFormat}
	_, e := partial.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.URI == "" {
		return nil, errInvalidSource(partial, map[string]interface{}{"description": "missing URI"})
	}
	if sc.Path.Config == "" {
		return nil, errInvalidSource(partial, map[string]interface{}{"description": "missing response config path"})
	}
	if sc.Path.Timestamp == "" {
		return nil, errInvalidSource(partial, map[string]interface{}{"description": "missing response config timestamp"})
	}
	// return acceptance for the read config type
	return NewObsSource(
		s.clientFactory(),
		sc.URI,
		sc.Format,
		s.decoderFactory,
		sc.Path.Timestamp,
		sc.Path.Config,
	)
}
