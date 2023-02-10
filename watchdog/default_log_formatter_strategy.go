package watchdog

import (
	"github.com/happyhippyhippo/slate/config"
)

const (
	// DefaultLogFormatterType is used to signal that the log formatter to be
	// used is the default slate message formatter.
	DefaultLogFormatterType = "default"
)

// DefaultLogFormatterStrategy defines a log formatter generation strategy instance.
type DefaultLogFormatterStrategy struct{}

var _ ILogFormatterStrategy = &DefaultLogFormatterStrategy{}

// NewDefaultLogFormatterStrategy @todo doc
func NewDefaultLogFormatterStrategy() *DefaultLogFormatterStrategy {
	return &DefaultLogFormatterStrategy{}
}

// Accept will check if the strategy will accept the configuration
// used to create a new log formatter.
func (s DefaultLogFormatterStrategy) Accept(
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
		return fc.Type == DefaultLogFormatterType
	}
	return false
}

// Create will try to generate a log formatter based on the
// passed configuration.
func (s DefaultLogFormatterStrategy) Create(
	_ config.IConfig,
) (ILogFormatter, error) {
	// return the default log formatter instance
	return &DefaultLogFormatter{}, nil
}
