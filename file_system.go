package slate

import (
	"github.com/spf13/afero"
)

// ----------------------------------------------------------------------------
// defs
// ----------------------------------------------------------------------------

const (
	// FileSystemContainerID defines the application service Provider
	// registration string of the file system adapter service.
	FileSystemContainerID = ContainerID + ".file_system"
)

// ----------------------------------------------------------------------------
// service register
// ----------------------------------------------------------------------------

// FileSystemServiceRegister defines the module service provider to
// be used on the application initialization to register the file system
// adapter service.
type FileSystemServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &FileSystemServiceRegister{}

// NewFileSystemServiceRegister will generate a new file system
// register instance
func NewFileSystemServiceRegister(
	app ...*App,
) *FileSystemServiceRegister {
	return &FileSystemServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will add to the application Provider the module services.
func (FileSystemServiceRegister) Provide(
	container *ServiceContainer,
) error {
	if container == nil {
		return errNilPointer("container")
	}
	return container.Add(FileSystemContainerID, afero.NewOsFs)
}
