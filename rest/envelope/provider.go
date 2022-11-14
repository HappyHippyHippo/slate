package envelope

import (
	"github.com/happyhippyhippo/slate"
	sconfig "github.com/happyhippyhippo/slate/config"
)

// Provider defines the slaterest.envelope module service provider
// to be used on the application initialization to register the
// response envelope middleware services.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will add to the container a new file system adapter instance.
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Factory(ContainerID, func() (interface{}, error) {
		cfg, err := sconfig.Get(c)
		if err != nil {
			return nil, err
		}

		return NewMiddlewareGenerator(cfg)
	})

	return nil
}

// Boot (no-op).
func (Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	return nil
}
