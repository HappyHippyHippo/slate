package gfs

import (
	"github.com/happyhippyhippo/slate"
	"github.com/spf13/afero"
)

// GetFileSystem wii try to retrieve the registered file
// system adapter from the application service container.
func GetFileSystem(c slate.ServiceContainer) (afero.Fs, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(afero.Fs)
	if !ok {
		return nil, errConversion(instance, "afero.Fs")
	}
	return i, nil
}
