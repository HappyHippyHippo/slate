package sconfig

import (
	"sync"
)

type sourceAggregate struct {
	source
	configs []IConfig
}

var _ ISource = &sourceAggregate{}

// NewSourceAggregate will instantiate a new config source
// that aggregate a list of configs elements.
func NewSourceAggregate(configs []IConfig) (ISource, error) {
	if configs == nil {
		return nil, errNilPointer("configs")
	}

	s := &sourceAggregate{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		configs: configs,
	}

	if e := s.load(); e != nil {
		return nil, e
	}

	return s, nil
}

func (s *sourceAggregate) load() error {
	for _, config := range s.configs {
		partial, e := config.Partial("", Partial{})
		if e != nil {
			return e
		}
		s.partial.merge(partial)
	}
	return nil
}
