package sconfig

// Loader defines the config instantiation and initialization of a new
// config managing structure.
type Loader struct {
	cfg     Manager
	factory *SourceFactory
}

// NewLoader instantiate a new configuration loader.
func NewLoader(cfg Manager, factory *SourceFactory) (*Loader, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	return &Loader{
		cfg:     cfg,
		factory: factory,
	}, nil
}

// Load loads the configuration from a base config file defined by a
// path and format.
func (l Loader) Load() error {
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

func (l Loader) loadSource(cfg Config) error {
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
