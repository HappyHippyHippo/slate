package sfs

import (
	"github.com/happyhippyhippo/slate"
	"github.com/spf13/afero"
)

// Provider defines the slate.sfs module service provider to be used on
// the application initialization to register the file system adapter service.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will add to the container a new file system adapter instance.
func (Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	return c.Service(ContainerID, func() (interface{}, error) {
		return afero.NewOsFs(), nil
	})
}

// Boot (no-op).
func (Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	return nil
}
