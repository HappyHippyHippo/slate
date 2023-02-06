package config

const (
	// SourceStrategyEnvironment defines the value to be used to declare an
	// environment config source type.
	SourceStrategyEnvironment = "env"
)

type envSourceConfig struct {
	Mappings Config
}

// EnvSourceStrategy defines a strategy used to instantiate an
// environment variable mapped config source creation strategy.
type EnvSourceStrategy struct{}

var _ ISourceStrategy = &EnvSourceStrategy{}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s EnvSourceStrategy) Accept(
	config IConfig,
) bool {
	// check config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	_, e := config.Populate("", &sc)
	if e == nil {
		// return acceptance for the read config type
		return sc.Type == SourceStrategyEnvironment
	}
	return false
}

// Create will instantiate the desired environment source instance
// where the initialization data comes from a configuration instance.
func (s EnvSourceStrategy) Create(
	config IConfig,
) (ISource, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := envSourceConfig{}
	_, e := config.Populate("", &sc)
	if e != nil {
		return nil, e
	}
	// create the mappings map
	mapping := make(map[string]string)
	for k, v := range sc.Mappings {
		tk, ok := k.(string)
		if !ok {
			return nil, errConversion(k, "string")
		}
		tv, ok := v.(string)
		if !ok {
			return nil, errConversion(v, "string")
		}
		mapping[tk] = tv
	}
	// create the config source
	return NewEnvSource(mapping)
}
