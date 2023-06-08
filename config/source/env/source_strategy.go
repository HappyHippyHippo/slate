package env

import (
	"github.com/happyhippyhippo/slate/config"
)

const (
	// Type defines the value to be used to
	// declare an environment config source type.
	Type = "env"
)

// SourceStrategy defines a strategy used to instantiate an
// environment variable mapped config source creation strategy.
type SourceStrategy struct{}

var _ config.SourceStrategy = &SourceStrategy{}

// NewSourceStrategy instantiates a new environment config
// source creation strategy.
func NewSourceStrategy() *SourceStrategy {
	return &SourceStrategy{}
}

// Accept will check if the source factory strategy can instantiate
// a source where the data to check comes from a configuration
// instance.
func (s SourceStrategy) Accept(
	cfg *config.Partial,
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
	return sc.Type == Type
}

// Create will instantiate the desired environment source instance
// where the initialization data comes from a configuration instance.
func (s SourceStrategy) Create(
	cfg *config.Partial,
) (config.Source, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	sc := struct{ Mappings config.Partial }{}
	if _, e := cfg.Populate("", &sc); e != nil {
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
	return NewSource(mapping)
}