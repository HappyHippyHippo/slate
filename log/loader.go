package log

import (
	"github.com/happyhippyhippo/slate/config"
)

type configurer interface {
	Partial(path string, def ...config.Partial) (*config.Partial, error)
	AddObserver(path string, callback config.Observer) error
}

type logger interface {
	RemoveAllStreams()
	AddStream(id string, stream Stream) error
}

type streamFactory interface {
	Create(cfg *config.Partial) (Stream, error)
}

// Loader defines the logger instantiation and initialization of a new
// logger proxy.
type Loader struct {
	cfg           configurer
	log           logger
	streamFactory streamFactory
}

// NewLoader generates a new logger initialization instance.
func NewLoader(
	cfg *config.Config,
	log *Log,
	streamFactory *StreamFactory,
) (*Loader, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// check the logger argument reference
	if log == nil {
		return nil, errNilPointer("logger")
	}
	// check the stream factory argument reference
	if streamFactory == nil {
		return nil, errNilPointer("factory")
	}
	// instantiate the loader
	return &Loader{
		cfg:           cfg,
		log:           log,
		streamFactory: streamFactory,
	}, nil
}

// Load will parse the configuration and instantiates logging streams
// depending the data on the configuration.
func (l Loader) Load() error {
	// retrieve the logger entries from the config instance
	entries, e := l.cfg.Partial(LoaderConfigPath, config.Partial{})
	if e != nil {
		return e
	}
	// load the retrieved entries
	if e := l.load(entries); e != nil {
		return e
	}
	// check if the logger streams list should be observed for updates
	if LoaderObserveConfig {
		// add the observer to the given config
		_ = l.cfg.AddObserver(
			LoaderConfigPath,
			func(_ interface{}, newConfig interface{}) {
				// type check the new logger config with the logging streams
				cfg, ok := newConfig.(config.Partial)
				if !ok {
					return
				}
				// remove all the current registered streams
				l.log.RemoveAllStreams()
				// load the new stream entries
				_ = l.load(&cfg)
			},
		)
	}
	return nil
}

func (l Loader) load(
	cfg *config.Partial,
) error {
	// iterate through the given logger config stream list
	for _, id := range cfg.Entries() {
		// get the configuration
		entry, e := cfg.Partial(id)
		if e != nil {
			return e
		}
		// generate the new stream
		stream, e := l.streamFactory.Create(entry)
		if e != nil {
			return e
		}
		// add the stream to the logger stream pool
		e = l.log.AddStream(id, stream)
		if e != nil {
			return e
		}
	}
	return nil
}
