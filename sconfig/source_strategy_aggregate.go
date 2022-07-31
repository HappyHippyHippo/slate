package sconfig

type sourceStrategyAggregate struct {
	partials []IConfig
}

var _ ISourceStrategy = &sourceStrategyAggregate{}

// Accept will check if the source dFactory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyAggregate) Accept(sourceType string) bool {
	return sourceType == SourceTypeAggregate
}

// AcceptFromConfig will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyAggregate) AcceptFromConfig(cfg IConfig) bool {
	if cfg == nil {
		return false
	}

	if sourceType, e := cfg.String("type"); e == nil {
		return s.Accept(sourceType)
	}

	return false
}

// Create will instantiate the desired environment source instance.
func (s sourceStrategyAggregate) Create(_ ...interface{}) (ISource, error) {
	return newSourceAggregate(s.partials)
}

// CreateFromConfig will instantiate the desired environment source instance
// where the initialization data comes from a configuration Partial instance.
func (s sourceStrategyAggregate) CreateFromConfig(cfg IConfig) (ISource, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	return s.Create()
}
