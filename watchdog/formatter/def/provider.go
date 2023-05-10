package def

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/watchdog"
	"github.com/happyhippyhippo/slate/watchdog/formatter"
)

const (
	// ID defines the application container registration string for the
	// default logging formatter strategy.
	ID = formatter.ID + ".default"
)

// Provider defines the slate.config module service provider to be used
// on the application initialization to register the config service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the configuration section instances in the
// application container.
func (Provider) Register(
	container slate.IContainer,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(ID, NewFormatterStrategy, watchdog.LogFormatterStrategyTag)
	return nil
}

// Boot will start the configuration config instance by calling the
// configuration loader with the defined provider base entry information.
func (p Provider) Boot(
	container slate.IContainer,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
