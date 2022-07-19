package slog

import (
	"github.com/happyhippyhippo/slate/sconfig"
)

// ILoader defines the interface of a slog loader instance.
type ILoader interface {
	Load() error
}

// Loader defines the logger instantiation and initialization of a new
// logger proxy.
type Loader struct {
	cfg      sconfig.IManager
	logger   ILogger
	sFactory IStreamFactory
}

var _ ILoader = &Loader{}

func newLoader(cfg sconfig.IManager, logger ILogger, sFactory IStreamFactory) (ILoader, error) {
	if cfg == nil {
		return nil, errNilPointer("sconfig")
	}
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	if sFactory == nil {
		return nil, errNilPointer("factory")
	}

	return &Loader{
		cfg:      cfg,
		logger:   logger,
		sFactory: sFactory,
	}, nil
}

// Load will parse the configuration and instantiates logging streams
// depending the data on the configuration.
func (l Loader) Load() error {
	if entries, e := l.entries(); e != nil {
		return e
	} else if e := l.load(entries); e != nil {
		return e
	}

	if LoaderObserveConfig {
		_ = l.cfg.AddObserver(LoaderConfigPath, func(_ interface{}, newEntries interface{}) {
			defer func() {
				if e := recover(); e != nil {
					_ = l.logger.Signal(
						LoaderErrorChannel,
						ERROR,
						"reloading slog streams error",
						map[string]interface{}{"error": e},
					)
				}
			}()

			var entries []sconfig.Partial
			for _, entry := range newEntries.([]interface{}) {
				entries = append(entries, entry.(sconfig.Partial))
			}

			l.logger.RemoveAllStreams()

			if e := l.load(entries); e != nil {
				_ = l.logger.Signal(
					LoaderErrorChannel,
					ERROR,
					"reloading slog streams error",
					map[string]interface{}{"error": e},
				)
			}
		})
	}

	return nil
}

func (l Loader) entries() ([]sconfig.Partial, error) {
	list, e := l.cfg.List(LoaderConfigPath, []interface{}{})
	if e != nil {
		return nil, e
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
		if id, e := entry.String("id"); e != nil {
			return e
		} else if s, e := l.sFactory.CreateFromConfig(&entry); e != nil {
			return e
		} else if e := l.logger.AddStream(id, s); e != nil {
			return e
		}
	}

	return nil
}
