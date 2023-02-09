package log

import (
	"github.com/happyhippyhippo/slate/config"
)

// ILoader defines the interface of a log loader instance.
type ILoader interface {
	Load() error
}

// Loader defines the log instantiation and initialization of a new
// log proxy.
type Loader struct {
	cfg           config.IManager
	log           ILog
	streamFactory IStreamFactory
}

var _ ILoader = &Loader{}

// NewLoader generates a new log initialization instance.
func NewLoader(
	cfg config.IManager,
	log ILog,
	streamFactory IStreamFactory,
) (ILoader, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// check the log argument reference
	if log == nil {
		return nil, errNilPointer("log")
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
	// retrieve the log entries from the config instance
	entries, e := l.cfg.Config(LoaderConfigPath, config.Config{})
	if e != nil {
		return e
	}
	// load the retrieved entries
	if e := l.load(entries); e != nil {
		return e
	}
	// check if the log streams list should be observed for updates
	if LoaderObserveConfig {
		// add the observer to the given config
		_ = l.cfg.AddObserver(
			LoaderConfigPath,
			func(_ interface{}, newConfig interface{}) {
				// type check the new log config with the logging streams
				cfg, ok := newConfig.(config.Config)
				if !ok {
					// log the input value error
					_ = l.log.Signal(LoaderErrorChannel, ERROR, "reloading log streams error", Context{"error": e})
					return
				}
				// remove all the current registered streams
				l.log.RemoveAllStreams()
				// load the new stream entries
				if e := l.load(&cfg); e != nil {
					_ = l.log.Signal(LoaderErrorChannel, ERROR, "reloading log streams error", Context{"error": e})
				}
			},
		)
	}
	return nil
}

func (l Loader) load(
	cfg config.IConfig,
) error {
	// iterate through the given log config stream list
	for _, id := range cfg.Entries() {
		// get the configuration
		entry, e := cfg.Config(id)
		if e != nil {
			return e
		}
		// generate the new stream
		stream, e := l.streamFactory.Create(entry)
		if e != nil {
			return e
		}
		// add the stream to the log stream pool
		e = l.log.AddStream(id, stream)
		if e != nil {
			return e
		}
	}
	return nil
}
