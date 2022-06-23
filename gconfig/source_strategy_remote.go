package gconfig

import (
	"net/http"
)

// sourceStrategyRemote defines a remote config source instantiation
// strategy to be used by the config sources factory instance.
type sourceStrategyRemote struct {
	clientFactory  func() HTTPClient
	decoderFactory *DecoderFactory
}

var _ SourceStrategy = &sourceStrategyRemote{}

// NewSourceStrategyRemote instantiate a new remote source factory
// strategy that will enable the source factory to instantiate a new
// remote configuration source.
func NewSourceStrategyRemote(decoderFactory *DecoderFactory) (SourceStrategy, error) {
	if decoderFactory == nil {
		return nil, errNilPointer("decoderFactory")
	}

	return &sourceStrategyRemote{
		clientFactory:  func() HTTPClient { return &http.Client{} },
		decoderFactory: decoderFactory,
	}, nil
}

// Accept will check if the source factory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyRemote) Accept(stype string) bool {
	return stype == SourceTypeRemote
}

// AcceptFromConfig will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyRemote) AcceptFromConfig(cfg Config) bool {
	if cfg == nil {
		return false
	}

	if stype, err := cfg.String("type"); err == nil {
		return s.Accept(stype)
	}

	return false
}

// Create will instantiate the desired remote source instance.
func (s sourceStrategyRemote) Create(args ...interface{}) (Source, error) {
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
		return NewSourceRemote(s.clientFactory(), uri, format, s.decoderFactory, configPath)
	}
}

// CreateFromConfig will instantiate the desired remote source instance where
// the initialization data comes from a configuration Partial instance.
func (s sourceStrategyRemote) CreateFromConfig(cfg Config) (Source, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if uri, err := cfg.String("uri"); err != nil {
		return nil, err
	} else if format, err := cfg.String("format", DefaultRemoteFormat); err != nil {
		return nil, err
	} else if configPath, err := cfg.String("configPath"); err != nil {
		return nil, err
	} else {
		return s.Create(uri, format, configPath)
	}
}
