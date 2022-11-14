package fs

import (
	"github.com/happyhippyhippo/slate"
	"github.com/spf13/afero"
)

// Get wii try to retrieve the registered file
// system adapter from the application service container.
func Get(c slate.ServiceContainer) (afero.Fs, error) {
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
