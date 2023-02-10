package config

const (
	// EnvironmentSourceType defines the value to be used to
	// declare an environment config source type.
	EnvironmentSourceType = "env"
)

type envSourceConfig struct {
	Mappings Config
}

// EnvSourceStrategy defines a strategy used to instantiate an
// environment variable mapped config source creation strategy.
type EnvSourceStrategy struct{}

var _ ISourceStrategy = &EnvSourceStrategy{}

// NewEnvSourceStrategy @todo doc
func NewEnvSourceStrategy() *EnvSourceStrategy {
	return &EnvSourceStrategy{}
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s EnvSourceStrategy) Accept(
	cfg IConfig,
) bool {
	// check config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	if _, e := cfg.Populate("", &sc); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sc.Type == EnvironmentSourceType
}

// Create will instantiate the desired environment source instance
// where the initialization data comes from a configuration instance.
func (s EnvSourceStrategy) Create(
	cfg IConfig,
) (ISource, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sc := envSourceConfig{}
	_, e := cfg.Populate("", &sc)
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
