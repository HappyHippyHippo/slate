package config

import (
	"net/http"
)

const (
	// ObservableRestSourceType defines the value to be used to
	// declare an observable rest config source type.
	ObservableRestSourceType = "observable-rest"
)

type observableRestSourceConfig struct {
	URI    string
	Format string
	Path   struct {
		Config    string
		Timestamp string
	}
}

// ObservableRestSourceStrategy defines a strategy used to instantiate
// an observable REST service config source creation strategy.
type ObservableRestSourceStrategy struct {
	RestSourceStrategy
}

var _ ISourceStrategy = &ObservableRestSourceStrategy{}

// NewObservableRestSourceStrategy instantiates a new observable REST
// service config source creation strategy.
func NewObservableRestSourceStrategy(
	decoderFactory IDecoderFactory,
) (*ObservableRestSourceStrategy, error) {
	// check the decoder factory argument reference
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}
	// instantiate the strategy
	return &ObservableRestSourceStrategy{
		RestSourceStrategy: RestSourceStrategy{
			clientFactory:  func() httpClient { return &http.Client{} },
			decoderFactory: decoderFactory,
		},
	}, nil
}

// Accept will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s ObservableRestSourceStrategy) Accept(
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
	return sc.Type == ObservableRestSourceType
}

// Create will instantiate the desired rest source instance where
// the initialization data comes from a configuration instance.
func (s ObservableRestSourceStrategy) Create(
	cfg IConfig,
) (ISource, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := observableRestSourceConfig{Format: DefaultRestFormat}
	_, e := cfg.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	if sc.URI == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{"description": "missing URI"})
	}
	if sc.Path.Config == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{"description": "missing response config path"})
	}
	if sc.Path.Timestamp == "" {
		return nil, errInvalidSource(cfg, map[string]interface{}{"description": "missing response config timestamp"})
	}
	// return acceptance for the read config type
	return NewObservableRestSource(
		s.clientFactory(),
		sc.URI,
		sc.Format,
		s.decoderFactory,
		sc.Path.Timestamp,
		sc.Path.Config,
	)
}
