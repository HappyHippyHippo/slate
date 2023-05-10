package fs

import (
	"github.com/happyhippyhippo/slate"
	"github.com/spf13/afero"
)

const (
	// ID defines the application container registration string
	// of the file system adapter.
	ID = slate.ID + ".fs"
)

// Provider defines the slate.fs module service provider to
// be used on the application initialization to register the file system
// adapter service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will add to the application container the module services.
func (Provider) Register(
	container slate.IContainer,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	return container.Service(ID, afero.NewOsFs)
}

// Boot (no-op).
func (Provider) Boot(
	container slate.IContainer,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
