package json

import (
	"github.com/happyhippyhippo/slate/log"
)

const (
	// Format defines the value to be used to declare a
	// JSON Log formatter format.
	Format = "json"
)

// FormatterStrategy defines a log message JSON formatter
// generation strategy.
type FormatterStrategy struct{}

var _ log.IFormatterStrategy = &FormatterStrategy{}

// NewFormatterStrategy generates a new JSON formatter
// generation strategy instance.
func NewFormatterStrategy() *FormatterStrategy {
	return &FormatterStrategy{}
}

// Accept will check if the formatter factory strategy can instantiate a
// formatter of the requested format.
func (FormatterStrategy) Accept(
	format string,
) bool {
	// only accept to create a JSON format formatter
	return format == Format
}

// Create will instantiate the desired formatter instance.
func (FormatterStrategy) Create(
	_ ...interface{},
) (log.IFormatter, error) {
	// generate the JSON formatter
	return &Formatter{}, nil
}
