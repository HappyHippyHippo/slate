package config

import (
	"sync"
)

// AggregateSource defines a config source that aggregates a list of
// config partials into a single aggregate provided source.
type AggregateSource struct {
	Source
	configs []IConfig
}

var _ ISource = &AggregateSource{}

// NewAggregateSource will instantiate a new config source
// that aggregate a list of configs elements.
func NewAggregateSource(
	configs []IConfig,
) (*AggregateSource, error) {
	// check config list reference
	if configs == nil {
		return nil, errNilPointer("configs")
	}
	// instantiates the config source
	s := &AggregateSource{
		Source: Source{
			mutex:   &sync.Mutex{},
			partial: Config{},
		},
		configs: configs,
	}
	// load the config information from the passed config partials
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *AggregateSource) load() error {
	// merge all the config partials given into the local config
	for _, config := range s.configs {
		partial, e := config.Config("", Config{})
		if e != nil {
			return e
		}
		s.partial.merge(*partial.(*Config))
	}
	return nil
}
