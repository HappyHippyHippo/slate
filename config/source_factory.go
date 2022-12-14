package config

// ISourceFactory defined the interface of a
// config source factory instance.
type ISourceFactory interface {
	Register(strategy ISourceStrategy) error
	Create(cfg IConfig) (ISource, error)
}

// SourceFactory defines an object responsible to instantiate a
// new config source.
type SourceFactory []ISourceStrategy

var _ ISourceFactory = &SourceFactory{}

// Register will register a new source factory strategy to be used
// on creation request.
func (f *SourceFactory) Register(
	strategy ISourceStrategy,
) error {
	// check the strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// add the strategy to the factory strategy pool
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate and return a new config source where the
// data used to decide the strategy to be used and also the initialization
// data comes from a configuration storing Partial instance.
func (f SourceFactory) Create(
	config IConfig,
) (ISource, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// find a strategy that accepts the requested source type
	for _, strategy := range f {
		if strategy.Accept(config) {
			// create the requested config source
			return strategy.Create(config)
		}
	}
	return nil, errInvalidSourceData(config)
}
