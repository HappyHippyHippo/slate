package fs

import (
	"github.com/happyhippyhippo/slate"
	"github.com/spf13/afero"
)

const (
	// ID defines the application container package
	// that provides the service's id base string.
	ID = slate.ID + ".fs"
)

// Provider defines the fs module service provider to
// be used on the application initialization to register the file system
// adapter service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will add to the container a new file system adapter instance.
func (Provider) Register(
	container ...slate.IContainer,
) error {
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	return container[0].Service(ID, afero.NewOsFs)
}

// Boot (no-op).
func (Provider) Boot(
	container ...slate.IContainer,
) error {
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	return nil
}
