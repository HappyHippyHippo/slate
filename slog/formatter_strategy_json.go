package slog

// formatterStrategyJSON defines the logger formatter instantiation
// strategy to be registered in the factory so a Json based logger formatter
// could be instantiated.
type formatterStrategyJSON struct{}

var _ FormatterStrategy = &formatterStrategyJSON{}

// Accept will check if the formatter factory strategy can instantiate a
// formatter of the requested format.
func (formatterStrategyJSON) Accept(format string) bool {
	return format == FormatJSON
}

// Create will instantiate the desired formatter instance.
func (formatterStrategyJSON) Create(_ ...interface{}) (Formatter, error) {
	return &FormatterJSON{}, nil
}
