package config

const (
	// SourceStrategyAggregate defines the value to be used to declare a
	// container loading configs source type.
	SourceStrategyAggregate = "aggregate"
)

// AggregateSourceStrategy defines a strategy used to instantiate
// a config aggregation config source creation strategy.
type AggregateSourceStrategy struct {
	configs []IConfig
}

var _ ISourceStrategy = &AggregateSourceStrategy{}

// Accept will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s AggregateSourceStrategy) Accept(
	config IConfig,
) bool {
	// check the config argument reference
	if config == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	_, e := config.Populate("", &sc)
	if e == nil {
		// return acceptance for the read config type
		return sc.Type == SourceStrategyAggregate
	}
	return false
}

// Create will instantiate the desired environment source instance
// where the initialization data comes from a configuration instance.
func (s AggregateSourceStrategy) Create(
	_ IConfig,
) (ISource, error) {
	// create the aggregate config source
	return NewAggregateSource(s.configs)
}
