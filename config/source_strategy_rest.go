package config

import (
	"net/http"
)

type sourceStrategyRest struct {
	clientFactory  func() HTTPClient
	decoderFactory IDecoderFactory
}

var _ ISourceStrategy = &sourceStrategyRest{}

func newSourceStrategyRest(decoderFactory IDecoderFactory) (ISourceStrategy, error) {
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}

	return &sourceStrategyRest{
		clientFactory:  func() HTTPClient { return &http.Client{} },
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source decoderFactory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyRest) Accept(sourceType string) bool {
	return sourceType == SourceTypeRest
}

// AcceptFromConfig will check if the source decoderFactory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyRest) AcceptFromConfig(cfg IConfig) bool {
	if cfg == nil {
		return false
	}

	if sourceType, e := cfg.String("type"); e == nil {
		return s.Accept(sourceType)
	}

	return false
}

// Create will instantiate the desired rest source instance.
func (s sourceStrategyRest) Create(args ...interface{}) (ISource, error) {
	if len(args) < 3 {
		return nil, errNilPointer("args")
	}

	if uri, ok := args[0].(string); !ok {
		return nil, errConversion(args[0], "string")
	} else if format, ok := args[1].(string); !ok {
		return nil, errConversion(args[1], "string")
	} else if configPath, ok := args[2].(string); !ok {
		return nil, errConversion(args[2], "string")
	} else {
		return NewSourceRest(s.clientFactory(), uri, format, s.decoderFactory, configPath)
	}
}

// CreateFromConfig will instantiate the desired rest source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyRest) CreateFromConfig(cfg IConfig) (ISource, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if uri, e := cfg.String("uri"); e != nil {
		return nil, e
	} else if format, e := cfg.String("format", DefaultRestFormat); e != nil {
		return nil, e
	} else if configPath, e := cfg.String("configPath"); e != nil {
		return nil, e
	} else {
		return s.Create(uri, format, configPath)
	}
}