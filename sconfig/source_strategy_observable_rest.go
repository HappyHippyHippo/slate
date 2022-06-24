package sconfig

import (
	"net/http"
)

// sourceStrategyObservableRest defines am observable rest config
// source instantiation strategy to be used by the config sources factory
// instance.
type sourceStrategyObservableRest struct {
	sourceStrategyRest
}

var _ SourceStrategy = &sourceStrategyObservableRest{}

// NewSourceStrategyObservableRest instantiate a new observable rest
// source factory strategy that will enable the source factory to instantiate
// a new rest configuration source.
func NewSourceStrategyObservableRest(decoderFactory *DecoderFactory) (SourceStrategy, error) {
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}

	return &sourceStrategyObservableRest{
		sourceStrategyRest: sourceStrategyRest{
			clientFactory:  func() HTTPClient { return &http.Client{} },
			decoderFactory: decoderFactory,
		},
	}, nil
}

// Accept will check if the source factory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyObservableRest) Accept(stype string) bool {
	return stype == SourceTypeObservableRest
}

// AcceptFromConfig will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyObservableRest) AcceptFromConfig(cfg Config) bool {
	if cfg == nil {
		return false
	}

	if stype, err := cfg.String("type"); err == nil {
		return s.Accept(stype)
	}

	return false
}

// Create will instantiate the desired observable rest source instance.
func (s sourceStrategyObservableRest) Create(args ...interface{}) (Source, error) {
	if len(args) < 4 {
		return nil, errNilPointer("args")
	}

	if uri, ok := args[0].(string); !ok {
		return nil, errConversion(args[0], "string")
	} else if format, ok := args[1].(string); !ok {
		return nil, errConversion(args[1], "string")
	} else if timestampPath, ok := args[2].(string); !ok {
		return nil, errConversion(args[2], "string")
	} else if configPath, ok := args[3].(string); !ok {
		return nil, errConversion(args[3], "string")
	} else {
		return NewSourceObservableRest(s.clientFactory(), uri, format, s.decoderFactory, timestampPath, configPath)
	}
}

// CreateFromConfig will instantiate the desired rest source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyObservableRest) CreateFromConfig(cfg Config) (Source, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if uri, err := cfg.String("uri"); err != nil {
		return nil, err
	} else if format, err := cfg.String("format", DefaultRestFormat); err != nil {
		return nil, err
	} else if timestampPath, err := cfg.String("timestampPath"); err != nil {
		return nil, err
	} else if configPath, err := cfg.String("configPath"); err != nil {
		return nil, err
	} else {
		return s.Create(uri, format, timestampPath, configPath)
	}
}
