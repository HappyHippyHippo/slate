package glog

import "github.com/happyhippyhippo/slate/gconfig"

// Loader defines the logger instantiation and initialization of a new
// logger proxy.
type Loader struct {
	config  gconfig.Manager
	logger  *Logger
	factory *StreamFactory
}

// NewLoader create a new logging configuration loader instance.
func NewLoader(config gconfig.Manager, logger *Logger, factory *StreamFactory) (*Loader, error) {
	if config == nil {
		return nil, errNilPointer("config")
	}
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	if factory == nil {
		return nil, errNilPointer("factory")
	}

	return &Loader{
		config:  config,
		logger:  logger,
		factory: factory,
	}, nil
}

// Load will parse the configuration and instantiates logging streams
// depending the data on the configuration.
func (l Loader) Load() error {
	if entries, err := l.entries(); err != nil {
		return err
	} else if err := l.load(entries); err != nil {
		return err
	}

	if LoaderObserveConfig {
		_ = l.config.AddObserver(LoaderConfigPath, func(_ interface{}, newEntries interface{}) {
			defer func() {
				if err := recover(); err != nil {
					_ = l.logger.Signal(
						LoaderErrorChannel,
						ERROR,
						"reloading log streams error",
						map[string]interface{}{"error": err},
					)
				}
			}()

			var entries []gconfig.Partial
			for _, e := range newEntries.([]interface{}) {
				entries = append(entries, e.(gconfig.Partial))
			}

			l.logger.RemoveAllStreams()

			if err := l.load(entries); err != nil {
				_ = l.logger.Signal(
					LoaderErrorChannel,
					ERROR,
					"reloading log streams error",
					map[string]interface{}{"error": err},
				)
			}
		})
	}

	return nil
}

func (l Loader) entries() ([]gconfig.Partial, error) {
	list, err := l.config.List(LoaderConfigPath, []interface{}{})
	if err != nil {
		return nil, err
	}

	var entries []gconfig.Partial
	for _, item := range list {
		entry, ok := item.(gconfig.Partial)
		if !ok {
			return nil, errConversion(item, "gconfig.Partial")
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (l Loader) load(entries []gconfig.Partial) error {
	for _, entry := range entries {
		if id, err := entry.String("id"); err != nil {
			return err
		} else if stream, err := l.factory.CreateFromConfig(&entry); err != nil {
			return err
		} else if err := l.logger.AddStream(id, stream); err != nil {
			return err
		}
	}

	return nil
}
