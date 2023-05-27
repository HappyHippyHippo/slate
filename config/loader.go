package config

type configurer interface {
	Partial(path string, def ...Partial) (*Partial, error)
	AddSource(id string, priority int, src Source) error
}

type sourceFactory interface {
	Create(partial *Partial) (Source, error)
}

// Loader defines an object responsible to initialize a
// configuration manager.
type Loader struct {
	config        configurer
	sourceFactory sourceFactory
}

type sourceConfig struct {
	Priority int
}

// NewLoader instantiate a new configuration loader instance.
func NewLoader(
	config *Config,
	sourceFactory *SourceFactory,
) (*Loader, error) {
	// check manager argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// check source factory argument reference
	if sourceFactory == nil {
		return nil, errNilPointer("sourceFactory")
	}
	// instantiate the loader
	return &Loader{
		config:        config,
		sourceFactory: sourceFactory,
	}, nil
}

// Load loads the configuration from a base partial file defined by a
// path and format.
func (l Loader) Load() error {
	// retrieve the loader entry file partial content
	src, e := l.sourceFactory.Create(&Partial{"type": "file", "path": LoaderSourcePath, "format": LoaderSourceFormat})
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
	partial *Partial,
) error {
	// parse the configuration
	sc := sourceConfig{}
	_, e := partial.Populate("", &sc)
	if e != nil {
		return e
	}
	// create the partial source
	src, e := l.sourceFactory.Create(partial)
	if e != nil {
		return e
	}
	// add the loaded source to the manager
	return l.config.AddSource(id, sc.Priority, src)
}
