package watchdog

import (
	"github.com/happyhippyhippo/slate/config"
)

const (
	// FormatterUnknown is used to signal that the log formatter is unknown.
	FormatterUnknown = "unknown"

	// FormatterDefault is used to signal that the log formatter to be
	// used is the default slate message formatter.
	FormatterDefault = "default"
)

// ILogFormatterStrategy defines a log formatter creation strategy
// instance interface.
type ILogFormatterStrategy interface {
	Accept(config config.IConfig) bool
	Create(config config.IConfig) (ILogFormatter, error)
}

// LogFormatterStrategy defines a log formatter generation strategy instance.
type LogFormatterStrategy struct{}

var _ ILogFormatterStrategy = &LogFormatterStrategy{}

// Accept will check if the strategy will accept the configuration
// used to create a new log formatter.
func (s LogFormatterStrategy) Accept(
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
		return fc.Type == FormatterDefault
	}
	return false
}

// Create will try to generate a log formatter based on the
// passed configuration.
func (s LogFormatterStrategy) Create(
	_ config.IConfig,
) (ILogFormatter, error) {
	// return the default log formatter instance
	return &LogFormatter{}, nil
}
