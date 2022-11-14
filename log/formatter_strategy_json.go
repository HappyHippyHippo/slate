package log

type formatterStrategyJSON struct{}

var _ IFormatterStrategy = &formatterStrategyJSON{}

// Accept will check if the formatter factory strategy can instantiate a
// formatter of the requested format.
func (formatterStrategyJSON) Accept(format string) bool {
	return format == FormatJSON
}

// Create will instantiate the desired formatter instance.
func (formatterStrategyJSON) Create(_ ...interface{}) (IFormatter, error) {
	return &formatterJSON{}, nil
}
