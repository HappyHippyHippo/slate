package sconfig

// ISourceFactory defined the interface of a config source dFactory instance.
type ISourceFactory interface {
	Register(strategy ISourceStrategy) error
	Create(sourceType string, args ...interface{}) (ISource, error)
	CreateFromConfig(cfg IConfig) (ISource, error)
}

type sourceFactory []ISourceStrategy

var _ ISourceFactory = &sourceFactory{}

// Register will register a new source dFactory strategy to be used
// on creation request.
func (f *sourceFactory) Register(strategy ISourceStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	*f = append(*f, strategy)

	return nil
}

// Create will instantiate and return a new config source by the type requested.
func (f sourceFactory) Create(sourceType string, args ...interface{}) (ISource, error) {
	for _, strategy := range f {
		if strategy.Accept(sourceType) {
			return strategy.Create(args...)
		}
	}

	return nil, errInvalidConfigSourceType(sourceType)
}

// CreateFromConfig will instantiate and return a new config source where the
// data used to decide the strategy to be used and also the initialization
// data comes from a configuration storing Partial instance.
func (f sourceFactory) CreateFromConfig(cfg IConfig) (ISource, error) {
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
