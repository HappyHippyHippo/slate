package log

const (
	// FormatterFormatUnknown defines the value to be used to declare an
	// unknown Log formatter format.
	FormatterFormatUnknown = "unknown"
)

// IFormatterFactory defined the interface of a log formatter
// factory instance.
type IFormatterFactory interface {
	Register(strategy IFormatterStrategy) error
	Create(format string, args ...interface{}) (IFormatter, error)
}

// FormatterFactory defines the Log formatter factory structure used to
// instantiate Log formatters, based on registered instantiation strategies.
type FormatterFactory []IFormatterStrategy

var _ IFormatterFactory = &FormatterFactory{}

// Register will register a new formatter factory strategy to be used
// on requesting to create a formatter for a defined format.
func (f *FormatterFactory) Register(
	strategy IFormatterStrategy,
) error {
	// check the strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// add the strategy to the factory strategy pool
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate and return a new content formatter.
func (f FormatterFactory) Create(
	format string,
	args ...interface{},
) (IFormatter, error) {
	// search in the factory strategy pool for one that would accept
	// to generate the requested formatter with the requested format
	for _, s := range f {
		if s.Accept(format) {
			// return the creation of the requested formatter
			return s.Create(args...)
		}
	}
	return nil, errInvalidFormat(format)
}
