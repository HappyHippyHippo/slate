package sconfig

type sourceStrategyContainer struct {
	partials []IConfig
}

var _ ISourceStrategy = &sourceStrategyContainer{}

// Accept will check if the source dFactory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyContainer) Accept(sourceType string) bool {
	return sourceType == SourceTypeContainer
}

// AcceptFromConfig will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyContainer) AcceptFromConfig(cfg IConfig) bool {
	if cfg == nil {
		return false
	}

	if sourceType, err := cfg.String("type"); err == nil {
		return s.Accept(sourceType)
	}

	return false
}

// Create will instantiate the desired environment source instance.
func (s sourceStrategyContainer) Create(_ ...interface{}) (ISource, error) {
	return newSourceContainer(s.partials)
}

// CreateFromConfig will instantiate the desired environment source instance
// where the initialization data comes from a configuration Partial instance.
func (s sourceStrategyContainer) CreateFromConfig(cfg IConfig) (ISource, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	return s.Create()
}
