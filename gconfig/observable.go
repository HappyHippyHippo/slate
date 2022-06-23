package gconfig

import "io"

// Observable defined an interface to an instance that can
// observe configuration changes
type Observable interface {
	io.Closer

	HasObserver(path string) bool
	AddObserver(path string, callback Observer) error
	RemoveObserver(path string)
}
