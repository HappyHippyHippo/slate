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

type streamConfig struct {
	ID string
}

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
	entries, e := l.entries()
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
			func(_ interface{}, newEntries interface{}) {
				defer func() {
					if e := recover(); e != nil {
						// log the error
						_ = l.log.Signal(
							LoaderErrorChannel,
							ERROR,
							"reloading log streams error",
							Context{"error": e},
						)
					}
				}()
				// generate the stream entry list from the
				// observer new value parameter
				var entries []config.Config
				for _, entry := range newEntries.([]interface{}) {
					entries = append(entries, entry.(config.Config))
				}
				// remove all the current registered streams
				l.log.RemoveAllStreams()
				// load the new stream entries
				if e := l.load(entries); e != nil {
					_ = l.log.Signal(
						LoaderErrorChannel,
						ERROR,
						"reloading log streams error",
						Context{"error": e},
					)
				}
			},
		)
	}
	return nil
}

func (l Loader) entries() ([]config.Config, error) {
	// retrieve the stream entry list from the config
	list, e := l.cfg.List(LoaderConfigPath, []interface{}{})
	if e != nil {
		return nil, e
	}
	// iterate through the obtained list
	var entries []config.Config
	for _, item := range list {
		// validate the entry type
		entry, ok := item.(config.Config)
		if !ok {
			return nil, errConversion(item, "config.Config")
		}
		// add the iterated entry to the final list result
		entries = append(entries, entry)
	}
	return entries, nil
}

func (l Loader) load(
	entries []config.Config,
) error {
	// iterate through the given stream config list
	for _, entry := range entries {
		// parse the configuration
		sc := streamConfig{}
		_, e := entry.Populate("", &sc)
		if e != nil {
			return e
		}
		// validate configuration
		if sc.ID == "" {
			return errInvalidConfig(&entry)
		}
		// generate the new stream
		stream, e := l.streamFactory.Create(&entry)
		if e != nil {
			return e
		}
		// add the stream to the log stream pool
		e = l.log.AddStream(sc.ID, stream)
		if e != nil {
			return e
		}
	}
	return nil
}
