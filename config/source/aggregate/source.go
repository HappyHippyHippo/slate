package aggregate

import (
	"sync"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/source"
)

// Source defines a config source that aggregates a list of
// config partials into a single aggregate provided source.
type Source struct {
	source.Source
	sources []config.Source
}

var _ config.Source = &Source{}

// NewSource will instantiate a new config source
// that aggregate a list of sources elements.
func NewSource(
	sources []config.Source,
) (*Source, error) {
	// check config list reference
	if sources == nil {
		return nil, errNilPointer("sources")
	}
	// instantiates the config source
	s := &Source{
		Source: source.Source{
			Mutex:   &sync.Mutex{},
			Partial: config.Partial{},
		},
		sources: sources,
	}
	// load the config information from the passed config partials
	if e := s.load(); e != nil {
		return nil, e
	}
	return s, nil
}

func (s *Source) load() error {
	// merge all the cfg partials given into the local cfg
	c := config.Partial{}
	for _, s := range s.sources {
		p, e := s.Get("", config.Partial{})
		if e != nil {
			return e
		}
		c.Merge(p.(config.Partial))
	}
	// assign the merged sources to the source config
	s.Mutex.Lock()
	s.Partial = c
	s.Mutex.Unlock()
	return nil
}
