package log

import (
	"github.com/happyhippyhippo/slate/config"
)

// ConsoleStreamStrategy defines a console log stream generation strategy.
type ConsoleStreamStrategy struct {
	StreamStrategy
	formatterFactory IFormatterFactory
}

var _ IStreamStrategy = &ConsoleStreamStrategy{}

type consoleStreamConfig struct {
	Format   string
	Channels []interface{}
	Level    string
}

// NewConsoleStreamStrategy generates a new console log stream
// generation strategy instance.
func NewConsoleStreamStrategy(
	formatterFactory IFormatterFactory,
) (*ConsoleStreamStrategy, error) {
	// check formatter factory argument reference
	if formatterFactory == nil {
		return nil, errNilPointer("formatterFactory")
	}
	// instantiates the console stream strategy
	return &ConsoleStreamStrategy{
		formatterFactory: formatterFactory,
	}, nil
}

// Accept will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s ConsoleStreamStrategy) Accept(
	cfg config.IConfig,
) bool {
	// check config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	_, e := cfg.Populate("", &sc)
	if e == nil {
		// return acceptance for the read config type
		return sc.Type == StreamConsole
	}
	return false
}

// Create will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s ConsoleStreamStrategy) Create(
	cfg config.IConfig,
) (IStream, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := consoleStreamConfig{}
	_, e := cfg.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// validate configuration
	level, e := s.level(sc.Level)
	if e != nil {
		return nil, e
	}
	// create the stream formatter to be given to the console stream
	formatter, e := s.formatterFactory.Create(sc.Format)
	if e != nil {
		return nil, e
	}
	// instantiate the console stream
	return NewConsoleStream(formatter, s.channels(sc.Channels), level)
}
