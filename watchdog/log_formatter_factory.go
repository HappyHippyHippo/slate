package watchdog

import (
	"github.com/happyhippyhippo/slate/config"
)

// LogFormatterFactory defines an object responsible to instantiate a
// new watchdog log formatter.
type LogFormatterFactory []LogFormatterStrategy

// NewLogFormatterFactory will instantiate a new logging formatter
// creator instance.
func NewLogFormatterFactory() *LogFormatterFactory {
	return &LogFormatterFactory{}
}

// Register will register a new watchdog log formatter creator
// strategy to be used on creation request.
func (f *LogFormatterFactory) Register(
	strategy LogFormatterStrategy,
) error {
	// check the strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// add the strategy to the creator strategy pool
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate and return a new watchdog log formatter where the
// data used to decide the strategy to be used and also the initialization
// data comes from a configuration storing Partial instance.
func (f *LogFormatterFactory) Create(
	cfg *config.Partial,
) (LogFormatter, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// find a strategy that accepts the requested DefaultLogFormatter type
	for _, strategy := range *f {
		if strategy.Accept(cfg) {
			// create the requested config DefaultLogFormatter
			return strategy.Create(cfg)
		}
	}
	return nil, errInvalidWatchdog(cfg)
}
