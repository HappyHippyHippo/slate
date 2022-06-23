package glog

import "github.com/happyhippyhippo/slate/gconfig"

type streamStrategyConsole struct {
	streamStrategy
	factory *FormatterFactory
}

var _ StreamStrategy = &streamStrategyConsole{}

// NewStreamStrategyConsole instantiate a new console stream factory
// strategy that will enable the stream factory to instantiate a new console
// stream.
func NewStreamStrategyConsole(factory *FormatterFactory) (StreamStrategy, error) {
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	return &streamStrategyConsole{
		factory: factory,
	}, nil
}

// Accept will check if the console stream factory strategy can instantiate a
// stream of the requested type and with the calling parameters.
func (streamStrategyConsole) Accept(stype string) bool {
	return stype == StreamConsole
}

// AcceptFromConfig will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s streamStrategyConsole) AcceptFromConfig(cfg gconfig.Config) bool {
	if cfg == nil {
		return false
	}

	if stype, err := cfg.String("type"); err == nil {
		return s.Accept(stype)
	}

	return false
}

// Create will instantiate the desired stream instance.
func (s streamStrategyConsole) Create(args ...interface{}) (Stream, error) {
	if len(args) < 3 {
		return nil, errNilPointer("args")
	}

	if format, ok := args[0].(string); !ok {
		return nil, errConversion(args[0], "string")
	} else if channels, ok := args[1].([]string); !ok {
		return nil, errConversion(args[1], "[]string")
	} else if level, ok := args[2].(Level); !ok {
		return nil, errConversion(args[2], "log.Level")
	} else if formatter, err := s.factory.Create(format); err != nil {
		return nil, err
	} else {
		return NewStreamConsole(formatter, channels, level)
	}
}

// CreateFromConfig will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s streamStrategyConsole) CreateFromConfig(cfg gconfig.Config) (Stream, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if format, err := cfg.String("format"); err != nil {
		return nil, err
	} else if channels, err := s.channels(cfg); err != nil {
		return nil, err
	} else if level, err := s.level(cfg); err != nil {
		return nil, err
	} else {
		return s.Create(format, channels, level)
	}
}
