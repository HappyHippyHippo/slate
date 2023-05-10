package watchdog

import (
	"github.com/happyhippyhippo/slate/config"
)

// ILogFormatterFactory defined the interface of a
// watchdog log formatter factory instance.
type ILogFormatterFactory interface {
	Register(strategy ILogFormatterStrategy) error
	Create(cfg config.IConfig) (ILogFormatter, error)
}

// LogFormatterFactory defines an object responsible to instantiate a
// new watchdog log formatter.
type LogFormatterFactory []ILogFormatterStrategy

var _ ILogFormatterFactory = &LogFormatterFactory{}

// NewLogFormatterFactory will instantiate a new logging formatter
// factory instance.
func NewLogFormatterFactory() ILogFormatterFactory {
	return &LogFormatterFactory{}
}

// Register will register a new watchdog log formatter factory
// strategy to be used on creation request.
func (f *LogFormatterFactory) Register(
	strategy ILogFormatterStrategy,
) error {
	// check the strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// add the strategy to the factory strategy pool
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate and return a new watchdog log formatter where the
// data used to decide the strategy to be used and also the initialization
// data comes from a configuration storing Partial instance.
func (f LogFormatterFactory) Create(
	cfg config.IConfig,
) (ILogFormatter, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// find a strategy that accepts the requested DefaultLogFormatter type
	for _, strategy := range f {
		if strategy.Accept(cfg) {
			// create the requested config DefaultLogFormatter
			return strategy.Create(cfg)
		}
	}
	return nil, errInvalidWatchdog(cfg)
}
