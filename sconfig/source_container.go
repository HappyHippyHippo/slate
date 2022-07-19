package sconfig

import (
	"sync"
)

type sourceContainer struct {
	source
	cfgs []IConfig
}

var _ ISource = &sourceContainer{}

func newSourceContainer(cfgs []IConfig) (ISource, error) {
	if cfgs == nil {
		return nil, errNilPointer("partials")
	}

	s := &sourceContainer{
		source: source{
			mutex:   &sync.Mutex{},
			partial: Partial{},
		},
		cfgs: cfgs,
	}

	if e := s.load(); e != nil {
		return nil, e
	}

	return s, nil
}

func (s *sourceContainer) load() error {
	for _, cfg := range s.cfgs {
		partial, e := cfg.Partial("", Partial{})
		if e != nil {
			return e
		}
		s.partial.merge(partial)
	}
	return nil
}
