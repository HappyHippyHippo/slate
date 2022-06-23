package gconfig

import "sync"

// Source defines the base interface of a config source.
type Source interface {
	Has(path string) bool
	Get(path string, def ...interface{}) (interface{}, error)
}

// SourceObservable interface extends the Source interface with methods
// specific to sources that will be checked for updates in a regular
// periodicity defined in the config object where the source will be
// registered.
type SourceObservable interface {
	Source
	Reload() (bool, error)
}

// source defines a base code of a config source instance.
type source struct {
	mutex   sync.Locker
	partial Partial
}

var _ Source = &source{}

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
// This method will mostly be used by the config object to obtain the full
// content of the source to aggregate all the data into his internal storing
// Partial instance.
func (s *source) Get(path string, def ...interface{}) (interface{}, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.partial.Get(path, def...)
}
