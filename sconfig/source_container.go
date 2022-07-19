package sconfig

import (
	"sync"
)

type sourceContainer struct {
	source
	configs []IConfig
}

var _ ISource = &sourceContainer{}

func newSourceContainer(configs []IConfig) (ISource, error) {
	if configs == nil {
		return nil, errNilPointer("configs")
	}

	s := &sourceContainer{
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

func (s *sourceContainer) load() error {
	for _, config := range s.configs {
		partial, e := config.Partial("", Partial{})
		if e != nil {
			return e
		}
		s.partial.merge(partial)
	}
	return nil
}
