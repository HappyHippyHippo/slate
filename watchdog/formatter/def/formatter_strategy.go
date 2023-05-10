package def

import (
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/watchdog"
)

const (
	// Type is used to signal that the log formatter to be
	// used is the def slate message formatter.
	Type = "def"
)

// FormatterStrategy defines a log formatter generation strategy instance.
type FormatterStrategy struct{}

var _ watchdog.ILogFormatterStrategy = &FormatterStrategy{}

// NewFormatterStrategy will instantiate a new default logging formatter
// strategy instance.
func NewFormatterStrategy() *FormatterStrategy {
	return &FormatterStrategy{}
}

// Accept will check if the strategy will accept the configuration
// used to create a new log formatter.
func (s FormatterStrategy) Accept(
	cfg config.IConfig,
) bool {
	// check the config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the data from the configuration
	fc := struct{ Type string }{}
	_, e := cfg.Populate("", &fc)
	if e == nil {
		// return acceptance for the read config type
		return fc.Type == Type
	}
	return false
}

// Create will try to generate a log formatter based on the
// passed configuration.
func (s FormatterStrategy) Create(
	_ config.IConfig,
) (watchdog.ILogFormatter, error) {
	// return the def log formatter instance
	return &Formatter{}, nil
}
