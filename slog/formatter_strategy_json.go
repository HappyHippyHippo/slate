package slog

// FormatterStrategyJSON defines the logger formatter instantiation
// strategy to be registered in the factory so a Json based logger formatter
// could be instantiated.
type FormatterStrategyJSON struct{}

var _ FormatterStrategy = &FormatterStrategyJSON{}

// Accept will check if the formatter factory strategy can instantiate a
// formatter of the requested format.
func (FormatterStrategyJSON) Accept(format string) bool {
	return format == FormatJSON
}

// Create will instantiate the desired formatter instance.
func (FormatterStrategyJSON) Create(_ ...interface{}) (Formatter, error) {
	return &FormatterJSON{}, nil
}
