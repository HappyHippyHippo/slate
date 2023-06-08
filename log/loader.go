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

type streamCreator interface {
	Create(cfg *config.Partial) (Stream, error)
}

// Loader defines the logger instantiation and initialization of a new
// logger proxy.
type Loader struct {
	configurer    configurer
	logger        logger
	streamCreator streamCreator
}

// NewLoader generates a new logger initialization instance.
func NewLoader(
	configurer *config.Config,
	logger *Log,
	streamCreator *StreamFactory,
) (*Loader, error) {
	// check the config argument reference
	if configurer == nil {
		return nil, errNilPointer("configurer")
	}
	// check the logger argument reference
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check the stream factory argument reference
	if streamCreator == nil {
		return nil, errNilPointer("streamCreator")
	}
	// instantiate the loader
	return &Loader{
		configurer:    configurer,
		logger:        logger,
		streamCreator: streamCreator,
	}, nil
}

// Load will parse the configuration and instantiates logging streams
// depending the data on the configuration.
func (l Loader) Load() error {
	// retrieve the logger entries from the config instance
	entries, e := l.configurer.Partial(LoaderConfigPath, config.Partial{})
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
		_ = l.configurer.AddObserver(
			LoaderConfigPath,
			func(_ interface{}, newConfig interface{}) {
				// type check the new logger config with the logging streams
				cfg, ok := newConfig.(config.Partial)
				if !ok {
					return
				}
				// remove all the current registered streams
				l.logger.RemoveAllStreams()
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
		stream, e := l.streamCreator.Create(entry)
		if e != nil {
			return e
		}
		// add the stream to the logger stream pool
		if e := l.logger.AddStream(id, stream); e != nil {
			return e
		}
	}
	return nil
}
