package watchdog

import (
	"github.com/happyhippyhippo/slate/config"
)

const (
	// UnknownLogFormatter is used to signal that the log formatter is unknown.
	UnknownLogFormatter = "unknown"
)

// LogFormatterStrategy defines a log formatter creation strategy
// instance interface.
type LogFormatterStrategy interface {
	Accept(config *config.Partial) bool
	Create(config *config.Partial) (LogFormatter, error)
}
