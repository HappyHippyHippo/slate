package sconfig

// ILoader defines the interface of a config loader instance.
type ILoader interface {
	Load() error
}

type loader struct {
	cfg     IManager
	factory ISourceFactory
}

var _ ILoader = &loader{}

func newLoader(cfg IManager, factory ISourceFactory) (ILoader, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	return &loader{
		cfg:     cfg,
		factory: factory,
	}, nil
}

// Load loads the configuration from a base config file defined by a
// path and format.
func (l loader) Load() error {
	if src, err := l.factory.Create(SourceTypeFile, LoaderSourcePath, LoaderSourceFormat); err != nil {
		return err
	} else if err := l.cfg.AddSource(LoaderSourceID, 0, src); err != nil {
		return err
	} else if sources, err := l.cfg.List(LoaderSourceListPath); err != nil {
		return nil
	} else {
		for _, source := range sources {
			if s, ok := source.(Partial); ok {
				if err := l.loadSource(&s); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (l loader) loadSource(cfg IConfig) error {
	if id, err := cfg.String("id"); err != nil {
		return err
	} else if priority, err := cfg.Int("priority", 0); err != nil {
		return err
	} else if src, err := l.factory.CreateFromConfig(cfg); err != nil {
		return err
	} else {
		return l.cfg.AddSource(id, priority, src)
	}
}
