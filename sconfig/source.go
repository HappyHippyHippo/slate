package sconfig

import "sync"

// ISource defines the base interface of a sconfig source.
type ISource interface {
	Has(path string) bool
	Get(path string, def ...interface{}) (interface{}, error)
}

// ISourceObservable interface extends the ISource interface with methods
// specific to sources that will be checked for updates in a regular
// periodicity defined in the sconfig object where the source will be
// registered.
type ISourceObservable interface {
	ISource
	Reload() (bool, error)
}

type source struct {
	mutex   sync.Locker
	partial Partial
}

var _ ISource = &source{}

// Has will check if the requested path is present in the source
// configuration content.
func (s *source) Has(path string) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.partial.Has(path)
}

// Get will retrieve the value stored in the requested path present in the
// configuration content.
// If the path does not exist, then the value nil will be returned.
// This method will mostly be used by the sconfig object to obtain the full
// content of the source to aggregate all the data into his internal storing
// Partial instance.
func (s *source) Get(path string, def ...interface{}) (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.partial.Get(path, def...)
}
