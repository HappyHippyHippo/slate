package config

import (
	"net/http"
)

type sourceStrategyObservableRest struct {
	sourceStrategyRest
}

var _ ISourceStrategy = &sourceStrategyObservableRest{}

func newSourceStrategyObservableRest(dFactory IDecoderFactory) (ISourceStrategy, error) {
	if dFactory == nil {
		return nil, errNilPointer("dFactory")
	}

	return &sourceStrategyObservableRest{
		sourceStrategyRest: sourceStrategyRest{
			cFactory: func() HTTPClient { return &http.Client{} },
			dFactory: dFactory,
		},
	}, nil
}

// Accept will check if the source dFactory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyObservableRest) Accept(sourceType string) bool {
	return sourceType == SourceTypeObservableRest
}

// AcceptFromConfig will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyObservableRest) AcceptFromConfig(cfg IConfig) bool {
	if cfg == nil {
		return false
	}

	if sourceType, e := cfg.String("type"); e == nil {
		return s.Accept(sourceType)
	}

	return false
}

// Create will instantiate the desired observable rest source instance.
func (s sourceStrategyObservableRest) Create(args ...interface{}) (ISource, error) {
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
		return newSourceObservableRest(s.cFactory(), uri, format, s.dFactory, timestampPath, configPath)
	}
}

// CreateFromConfig will instantiate the desired rest source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyObservableRest) CreateFromConfig(cfg IConfig) (ISource, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if uri, e := cfg.String("uri"); e != nil {
		return nil, e
	} else if format, e := cfg.String("format", DefaultRestFormat); e != nil {
		return nil, e
	} else if timestampPath, e := cfg.String("timestampPath"); e != nil {
		return nil, e
	} else if configPath, e := cfg.String("configPath"); e != nil {
		return nil, e
	} else {
		return s.Create(uri, format, timestampPath, configPath)
	}
}
