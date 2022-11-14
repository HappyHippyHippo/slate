package log

import sconfig "github.com/happyhippyhippo/slate/config"

type streamStrategyConsole struct {
	streamStrategy
	formatterFactory IFormatterFactory
}

var _ IStreamStrategy = &streamStrategyConsole{}

func newStreamStrategyConsole(formatterFactory IFormatterFactory) (IStreamStrategy, error) {
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}

	return &streamStrategyConsole{
		formatterFactory: formatterFactory,
	}, nil
}

// Accept will check if the console stream factory strategy can instantiate a
// stream of the requested type and with the calling parameters.
func (streamStrategyConsole) Accept(streamType string) bool {
	return streamType == StreamConsole
}

// AcceptFromConfig will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s streamStrategyConsole) AcceptFromConfig(cfg sconfig.IConfig) bool {
	if cfg == nil {
		return false
	}

	if streamType, e := cfg.String("type"); e == nil {
		return s.Accept(streamType)
	}

	return false
}

// Create will instantiate the desired stream instance.
func (s streamStrategyConsole) Create(args ...interface{}) (IStream, error) {
	if len(args) < 3 {
		return nil, errNilPointer("args")
	}

	if format, ok := args[0].(string); !ok {
		return nil, errConversion(args[0], "string")
	} else if channels, ok := args[1].([]string); !ok {
		return nil, errConversion(args[1], "[]string")
	} else if level, ok := args[2].(Level); !ok {
		return nil, errConversion(args[2], "log.Level")
	} else if formatter, e := s.formatterFactory.Create(format); e != nil {
		return nil, e
	} else {
		return newStreamConsole(formatter, channels, level)
	}
}

// CreateFromConfig will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s streamStrategyConsole) CreateFromConfig(cfg sconfig.IConfig) (IStream, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if format, e := cfg.String("format"); e != nil {
		return nil, e
	} else if channels, e := s.channels(cfg); e != nil {
		return nil, e
	} else if level, e := s.level(cfg); e != nil {
		return nil, e
	} else {
		return s.Create(format, channels, level)
	}
}
