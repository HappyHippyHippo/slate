package sconfig

import (
	"net/http"
)

type sourceStrategyRest struct {
	clientFactory  func() HTTPClient
	decoderFactory *DecoderFactory
}

var _ SourceStrategy = &sourceStrategyRest{}

func newSourceStrategyRest(decoderFactory *DecoderFactory) (SourceStrategy, error) {
	if decoderFactory == nil {
		return nil, errNilPointer("DecoderFactory")
	}

	return &sourceStrategyRest{
		clientFactory:  func() HTTPClient { return &http.Client{} },
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source factory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyRest) Accept(stype string) bool {
	return stype == SourceTypeRest
}

// AcceptFromConfig will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyRest) AcceptFromConfig(cfg Config) bool {
	if cfg == nil {
		return false
	}

	if stype, err := cfg.String("type"); err == nil {
		return s.Accept(stype)
	}

	return false
}

// Create will instantiate the desired rest source instance.
func (s sourceStrategyRest) Create(args ...interface{}) (Source, error) {
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
		return newSourceRest(s.clientFactory(), uri, format, s.decoderFactory, configPath)
	}
}

// CreateFromConfig will instantiate the desired rest source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyRest) CreateFromConfig(cfg Config) (Source, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if uri, err := cfg.String("uri"); err != nil {
		return nil, err
	} else if format, err := cfg.String("format", DefaultRestFormat); err != nil {
		return nil, err
	} else if configPath, err := cfg.String("configPath"); err != nil {
		return nil, err
	} else {
		return s.Create(uri, format, configPath)
	}
}
