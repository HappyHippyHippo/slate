package watchdog

import (
	"github.com/happyhippyhippo/slate/config"
)

const (
	// UnknownLogFormatterType is used to signal that the log formatter is unknown.
	UnknownLogFormatterType = "unknown"
)

// ILogFormatterStrategy defines a log formatter creation strategy
// instance interface.
type ILogFormatterStrategy interface {
	Accept(config config.IConfig) bool
	Create(config config.IConfig) (ILogFormatter, error)
}
