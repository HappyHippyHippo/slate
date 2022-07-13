package sconfig

import "io"

// IObservable defined an interface to an instance that can
// observe configuration changes
type IObservable interface {
	io.Closer

	HasObserver(path string) bool
	AddObserver(path string, callback Observer) error
	RemoveObserver(path string)
}
