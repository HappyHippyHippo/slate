package config

type configurer interface {
	Partial(path string, def ...Partial) (*Partial, error)
	AddSource(id string, priority int, src Source) error
}

var _ configurer = &Config{}

type sourceCreator interface {
	Create(cfg *Partial) (Source, error)
}

var _ sourceCreator = &SourceFactory{}

// Loader defines an object responsible to initialize a
// configuration manager.
type Loader struct {
	config        configurer
	sourceCreator sourceCreator
}

// NewLoader instantiate a new configuration loader instance.
func NewLoader(
	cfg *Config,
	sourceCreator *SourceFactory,
) (*Loader, error) {
	// check manager argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// check source factory argument reference
	if sourceCreator == nil {
		return nil, errNilPointer("sourceCreator")
	}
	// instantiate the loader
	return &Loader{
		config:        cfg,
		sourceCreator: sourceCreator,
	}, nil
}

// Load loads the configuration from a base partial file defined by a
// path and format.
func (l Loader) Load() error {
	// retrieve the loader entry file partial content
	src, e := l.sourceCreator.Create(&Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat})
	if e != nil {
		return e
	}
	// add the loaded entry file content into the manager
	if e := l.config.AddSource(LoaderSourceID, 0, src); e != nil {
		return e
	}
	// retrieve from the loaded info the partial entries list
	sources, e := l.config.Partial(LoaderSourceListPath)
	if e != nil {
		return nil
	}
	// iterate through the sources list
	for _, id := range sources.Entries() {
		// retrieve the source list entry
		if partial, e := sources.Partial(id); e == nil {
			// load the source
			if e := l.loadSource(id, partial); e != nil {
				return e
			}
		}
	}
	return nil
}

func (l Loader) loadSource(
	id string,
	cfg *Partial,
) error {
	// parse the configuration
	sc := struct{ Priority int }{}
	if _, e := cfg.Populate("", &sc); e != nil {
		return e
	}
	// create the partial source
	src, e := l.sourceCreator.Create(cfg)
	if e != nil {
		return e
	}
	// add the loaded source to the manager
	return l.config.AddSource(id, sc.Priority, src)
}
