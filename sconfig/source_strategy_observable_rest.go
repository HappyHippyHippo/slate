package sconfig

import (
	"net/http"
)

type sourceStrategyObservableRest struct {
	sourceStrategyRest
}

var _ ISourceStrategy = &sourceStrategyObservableRest{}

func newSourceStrategyObservableRest(decoderFactory IDecoderFactory) (ISourceStrategy, error) {
	if decoderFactory == nil {
		return nil, errNilPointer("DecoderFactory")
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
func (sourceStrategyObservableRest) Accept(sourceType string) bool {
	return sourceType == SourceTypeObservableRest
}

// AcceptFromConfig will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyObservableRest) AcceptFromConfig(cfg IConfig) bool {
	if cfg == nil {
		return false
	}

	if sourceType, err := cfg.String("type"); err == nil {
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
		return newSourceObservableRest(s.clientFactory(), uri, format, s.decoderFactory, timestampPath, configPath)
	}
}

// CreateFromConfig will instantiate the desired rest source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyObservableRest) CreateFromConfig(cfg IConfig) (ISource, error) {
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
