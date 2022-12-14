package log

// JSONFormatterStrategy defines a log message JSON formatter
// generation strategy.
type JSONFormatterStrategy struct{}

var _ IFormatterStrategy = &JSONFormatterStrategy{}

// Accept will check if the formatter factory strategy can instantiate a
// formatter of the requested format.
func (JSONFormatterStrategy) Accept(
	format string,
) bool {
	// only accept to create a JSON format formatter
	return format == FormatJSON
}

// Create will instantiate the desired formatter instance.
func (JSONFormatterStrategy) Create(
	_ ...interface{},
) (IFormatter, error) {
	// generate the JSON formatter
	return &JSONFormatter{}, nil
}