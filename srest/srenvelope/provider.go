package srenvelope

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
)

// Provider defines the default envelope provider to be used on
// the application initialization to register the file system adapter service.
type Provider struct{}

var _ slate.ServiceProvider = &Provider{}

// Register will add to the container a new file system adapter instance.
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Factory(ContainerID, func() (interface{}, error) {
		cfg, err := sconfig.GetConfig(c)
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
