package config

// SourceFactory defines an object responsible to instantiate a
// new partial source.
type SourceFactory []SourceStrategy

// NewSourceFactory will instantiate a new source factory instance
func NewSourceFactory() *SourceFactory {
	return &SourceFactory{}
}

// Register will register a new source factory strategy to be used
// on creation request.
func (f *SourceFactory) Register(
	strategy SourceStrategy,
) error {
	// check the strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// add the strategy to the factory strategy pool
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate and return a new partial source where the
// data used to decide the strategy to be used and also the initialization
// data comes from a configuration storing Partial instance.
func (f *SourceFactory) Create(
	partial *Partial,
) (Source, error) {
	// check the partial argument reference
	if partial == nil {
		return nil, errNilPointer("partial")
	}
	// find a strategy that accepts the requested source type
	for _, strategy := range *f {
		if strategy.Accept(partial) {
			// create the requested partial source
			return strategy.Create(partial)
		}
	}
	return nil, errInvalidSource(partial)
}
