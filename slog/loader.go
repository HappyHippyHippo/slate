package slog

import "github.com/happyhippyhippo/slate/sconfig"

// ILoader defines the interface of a log loader instance.
type ILoader interface {
	Load() error
}

// Loader defines the logger instantiation and initialization of a new
// logger proxy.
type Loader struct {
	config  sconfig.IManager
	logger  ILogger
	factory IStreamFactory
}

var _ ILoader = &Loader{}

func newLoader(config sconfig.IManager, logger ILogger, factory IStreamFactory) (ILoader, error) {
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

			var entries []sconfig.Partial
			for _, e := range newEntries.([]interface{}) {
				entries = append(entries, e.(sconfig.Partial))
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

func (l Loader) entries() ([]sconfig.Partial, error) {
	list, err := l.config.List(LoaderConfigPath, []interface{}{})
	if err != nil {
		return nil, err
	}

	var entries []sconfig.Partial
	for _, item := range list {
		entry, ok := item.(sconfig.Partial)
		if !ok {
			return nil, errConversion(item, "sconfig.Partial")
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (l Loader) load(entries []sconfig.Partial) error {
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
