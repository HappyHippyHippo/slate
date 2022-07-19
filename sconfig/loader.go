package sconfig

// ILoader defines the interface of a sconfig loader instance.
type ILoader interface {
	Load() error
}

type loader struct {
	cfg      IManager
	sFactory ISourceFactory
}

var _ ILoader = &loader{}

func newLoader(cfg IManager, sFactory ISourceFactory) (ILoader, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	if sFactory == nil {
		return nil, errNilPointer("dFactory")
	}

	return &loader{
		cfg:      cfg,
		sFactory: sFactory,
	}, nil
}

// Load loads the configuration from a base sconfig file defined by a
// path and format.
func (l loader) Load() error {
	if src, e := l.sFactory.Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat); e != nil {
		return e
	} else if e := l.cfg.AddSource(LoaderSourceID, 0, src); e != nil {
		return e
	} else if sources, e := l.cfg.List(LoaderSourceListPath); e != nil {
		return nil
	} else {
		for _, src := range sources {
			if s, ok := src.(Partial); ok {
				if e := l.loadSource(&s); e != nil {
					return e
				}
			}
		}
	}

	return nil
}

func (l loader) loadSource(cfg IConfig) error {
	if id, e := cfg.String("id"); e != nil {
		return e
	} else if priority, e := cfg.Int("priority", 0); e != nil {
		return e
	} else if src, e := l.sFactory.CreateFromConfig(cfg); e != nil {
		return e
	} else {
		return l.cfg.AddSource(id, priority, src)
	}
}
