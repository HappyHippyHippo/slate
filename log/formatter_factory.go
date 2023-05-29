package log

// FormatterFactory defines the Log Formatter factory structure used to
// instantiate Log formatters, based on registered instantiation strategies.
type FormatterFactory []FormatterStrategy

// NewFormatterFactory will instantiate a new formatter factory instance
func NewFormatterFactory() *FormatterFactory {
	return &FormatterFactory{}
}

// Register will register a new Formatter factory strategy to be used
// on requesting to create a Formatter for a defined format.
func (f *FormatterFactory) Register(
	strategy FormatterStrategy,
) error {
	// check the strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// add the strategy to the factory strategy pool
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate and return a new content Formatter.
func (f *FormatterFactory) Create(
	format string,
	args ...interface{},
) (Formatter, error) {
	// search in the factory strategy pool for one that would accept
	// to generate the requested Formatter with the requested format
	for _, s := range *f {
		if s.Accept(format) {
			// return the creation of the requested Formatter
			return s.Create(args...)
		}
	}
	return nil, errInvalidFormat(format)
}
