package sconfig

type sourceStrategyEnv struct{}

var _ ISourceStrategy = &sourceStrategyEnv{}

// Accept will check if the source dFactory strategy can instantiate a
// new source of the requested type.
func (sourceStrategyEnv) Accept(sourceType string) bool {
	return sourceType == SourceTypeEnv
}

// AcceptFromConfig will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration Partial
// instance.
func (s sourceStrategyEnv) AcceptFromConfig(cfg IConfig) bool {
	if cfg == nil {
		return false
	}

	if sourceType, err := cfg.String("type"); err == nil {
		return s.Accept(sourceType)
	}

	return false
}

// Create will instantiate the desired environment source instance.
func (s sourceStrategyEnv) Create(args ...interface{}) (ISource, error) {
	if len(args) < 1 {
		return nil, errNilPointer("args[0]")
	}

	if mappings, ok := args[0].(map[string]string); ok {
		return newSourceEnv(mappings)
	}

	return nil, errConversion(args[0], "map[string]string")
}

// CreateFromConfig will instantiate the desired environment source instance
// where the initialization data comes from a configuration Partial instance.
func (s sourceStrategyEnv) CreateFromConfig(cfg IConfig) (ISource, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	if mappings, err := cfg.Partial("mappings", Partial{}); err == nil {
		m := map[string]string{}
		for key, val := range mappings {
			m[key.(string)] = val.(string)
		}
		return s.Create(m)
	}

	return nil, errInvalidConfigSourcePartial(cfg)
}
