package sconfig

// SourceFactory defines a config source factory that uses a list of
// registered instantiation strategies to perform the config source
// instantiation.
type SourceFactory []SourceStrategy

// Register will register a new source factory strategy to be used
// on creation request.
func (f *SourceFactory) Register(strategy SourceStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	*f = append(*f, strategy)

	return nil
}

// Create will instantiate and return a new config source by the type requested.
func (f SourceFactory) Create(stype string, args ...interface{}) (Source, error) {
	for _, strategy := range f {
		if strategy.Accept(stype) {
			return strategy.Create(args...)
		}
	}

	return nil, errInvalidConfigSourceType(stype)
}

// CreateFromConfig will instantiate and return a new config source where the
// data used to decide the strategy to be used and also the initialization
// data comes from a configuration storing Partial instance.
func (f SourceFactory) CreateFromConfig(cfg Config) (Source, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	for _, strategy := range f {
		if strategy.AcceptFromConfig(cfg) {
			return strategy.CreateFromConfig(cfg)
		}
	}

	return nil, errInvalidConfigSourcePartial(cfg)
}
