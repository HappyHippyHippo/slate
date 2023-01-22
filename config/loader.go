package config

// ILoader defines the interface of a config loader instance.
type ILoader interface {
	Load() error
}

// Loader defines an object responsible to initialize a
// configuration manager.
type Loader struct {
	manager       IManager
	sourceFactory ISourceFactory
}

var _ ILoader = &Loader{}

type sourceConfig struct {
	ID       string
	Priority int
}

// NewLoader instantiate a new configuration loader instance.
func NewLoader(
	manager IManager,
	sourceFactory ISourceFactory,
) (*Loader, error) {
	// check manager argument reference
	if manager == nil {
		return nil, errNilPointer("manager")
	}
	// check source factory argument reference
	if sourceFactory == nil {
		return nil, errNilPointer("sourceFactory")
	}
	// instantiate the config loader
	return &Loader{
		manager:       manager,
		sourceFactory: sourceFactory,
	}, nil
}

// Load loads the configuration from a base config file defined by a
// path and format.
func (l Loader) Load() error {
	// retrieve the loader entry file config content
	sc := &Config{"type": SourceFile, "path": LoaderSourcePath, "format": LoaderSourceFormat}
	src, e := l.sourceFactory.Create(sc)
	if e != nil {
		return e
	}
	// add the loaded entry file content into the manager
	if e := l.manager.AddSource(LoaderSourceID, 0, src); e != nil {
		return e
	}
	// retrieve from the loaded info the config entries list
	sources, e := l.manager.List(LoaderSourceListPath)
	if e != nil {
		return nil
	}
	// iterate through the sources list
	for _, src := range sources {
		if s, ok := src.(Config); ok {
			// load the source
			if e := l.loadSource(&s); e != nil {
				return e
			}
		}
	}
	return nil
}

func (l Loader) loadSource(
	cfg IConfig,
) error {
	// parse the configuration
	sc := sourceConfig{}
	_, e := cfg.Populate("", &sc)
	if e != nil {
		return e
	}
	// validate configuration
	if sc.ID == "" {
		return errPathNotFound("id")
	}
	// create the config source
	src, e := l.sourceFactory.Create(cfg)
	if e != nil {
		return e
	}
	// add the loaded source to the manager
	return l.manager.AddSource(sc.ID, sc.Priority, src)
}
