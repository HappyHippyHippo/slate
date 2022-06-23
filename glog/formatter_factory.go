package glog

// FormatterFactory defines the logger formatter factory structure used to
// instantiate logger formatters, based on registered instantiation strategies.
type FormatterFactory []FormatterStrategy

// Register will register a new formatter factory strategy to be used
// on requesting to create a formatter for a defined format.
func (f *FormatterFactory) Register(strategy FormatterStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	*f = append(*f, strategy)

	return nil
}

// Create will instantiate and return a new content formatter.
func (f FormatterFactory) Create(format string, args ...interface{}) (Formatter, error) {
	for _, s := range f {
		if s.Accept(format) {
			return s.Create(args...)
		}
	}
	return nil, errInvalidFormat(format)
}
