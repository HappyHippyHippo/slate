package slog

type formatterStrategyJSON struct{}

var _ FormatterStrategy = &formatterStrategyJSON{}

// Accept will check if the formatter factory strategy can instantiate a
// formatter of the requested format.
func (formatterStrategyJSON) Accept(format string) bool {
	return format == FormatJSON
}

// Create will instantiate the desired formatter instance.
func (formatterStrategyJSON) Create(_ ...interface{}) (Formatter, error) {
	return &formatterJSON{}, nil
}
