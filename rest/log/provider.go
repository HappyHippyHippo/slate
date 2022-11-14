package log

import (
	"github.com/happyhippyhippo/slate"
	sconfig "github.com/happyhippyhippo/slate/config"
	slog "github.com/happyhippyhippo/slate/log"
)

// Provider defines the slaterest.log module service provider to be used on
// the application initialization to register the logging middleware services.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will register the log middleware package instances in the
// application container
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Factory(ContainerID, func() (interface{}, error) {
		cfg, err := sconfig.Get(c)
		if err != nil {
			return nil, err
		}

		logger, err := slog.Get(c)
		if err != nil {
			return nil, err
		}

		return NewMiddlewareGenerator(cfg, logger)
	})

	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (p Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	return nil
}
