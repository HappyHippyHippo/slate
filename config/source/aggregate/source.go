package aggregate

import (
	"sync"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
)

// Source defines a config source that aggregates a list of
// config partials into a single aggregate provided source.
type Source struct {
	source.BaseSource
	configs []config.IConfig
}

var _ config.ISource = &Source{}

// NewSource will instantiate a new config source
// that aggregate a list of configs elements.
func NewSource(
	configs []config.IConfig,
) (*Source, error) {
	// check config list reference
	if configs == nil {
		return nil, errNilPointer("configs")
	}
	// instantiates the config source
	s := &Source{
		BaseSource: source.BaseSource{
			Mutex:  &sync.Mutex{},
			Config: config.Config{},
		},
		configs: configs,
	}
	// load the config information from the passed config partials
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *Source) load() error {
	// merge all the cfg partials given into the local cfg
	c := config.Config{}
	for _, cfg := range s.configs {
		partial, e := cfg.Config("", config.Config{})
		if e != nil {
			return e
		}
		c.Merge(*partial.(*config.Config))
	}
	// assign the merged configs to the source config
	s.Mutex.Lock()
	s.Config = c
	s.Mutex.Unlock()
	return nil
}
