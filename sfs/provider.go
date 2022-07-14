package sfs

import (
	"github.com/happyhippyhippo/slate"
	"github.com/spf13/afero"
)

// Provider defines the slate.fs module service provider to be used on
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

// GetFileSystem wii try to retrieve the registered file
// system adapter from the application service container.
func GetFileSystem(c slate.ServiceContainer) (afero.Fs, error) {
	instance, e := c.Get(ContainerID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(afero.Fs)
	if !ok {
		return nil, errConversion(instance, "afero.Fs")
	}
	return i, nil
}
