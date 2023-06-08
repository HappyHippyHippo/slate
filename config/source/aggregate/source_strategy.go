package aggregate

import (
	"github.com/happyhippyhippo/slate/config"
)

const (
	// Type defines the value to be used to declare a
	// container loading configs source type.
	Type = "aggregate"
)

// SourceStrategy defines a strategy used to instantiate
// a config aggregation config source creation strategy.
type SourceStrategy struct {
	sources []config.Source
}

var _ config.SourceStrategy = &SourceStrategy{}

// Accept will check if the source dFactory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s SourceStrategy) Accept(
	cfg config.Partial,
) bool {
	// check the config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the data from the configuration
	sc := struct{ Type string }{}
	if _, e := cfg.Populate("", &sc); e != nil {
		return false
	}
	// return acceptance for the read config type
	return sc.Type == Type
}

// Create will instantiate the desired environment source instance
// where the initialization data comes from a configuration instance.
func (s SourceStrategy) Create(
	_ config.Partial,
) (config.Source, error) {
	// create the aggregate config source
	return NewSource(s.sources)
}
