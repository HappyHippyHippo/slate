package yaml

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/decoder"
)

const (
	// ID defines the application container registration string for the
	// YAML format decoder strategy.
	ID = decoder.ID + ".yaml"
)

// Provider defines the slate.config module service provider to be used
// on the application initialization to register the config service.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the configuration section instances in the
// application container.
func (Provider) Register(
	container *slate.Container,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(ID, NewDecoderStrategy, config.DecoderStrategyTag)
	return nil
}

// Boot will start the configuration config instance by calling the
// configuration loader with the defined provider base entry information.
func (p Provider) Boot(
	container *slate.Container,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
