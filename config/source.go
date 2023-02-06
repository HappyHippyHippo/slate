package config

import (
	"sync"
)

// ISource defines the base interface of a config source.
type ISource interface {
	Has(path string) bool
	Get(path string, def ...interface{}) (interface{}, error)
}

// IObservableSource interface extends the ISource interface with methods
// specific to sources that will be checked for updates in a regular
// periodicity defined in the config object where the source will be
// registered.
type IObservableSource interface {
	ISource
	Reload() (bool, error)
}

// Source defines a base structure and functionalities of a
// configuration source.
type Source struct {
	mutex  sync.Locker
	config Config
}

var _ ISource = &Source{}

// Has will check if the requested path is present in the source
// configuration content.
func (s *Source) Has(
	path string,
) bool {
	// lock the source for changes
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// check if the source stored config has the requested path
	return s.config.Has(path)
}

// Get will retrieve the value stored in the requested path present in the
// configuration content.
// If the path does not exist, then the value nil will be returned.
// This method will mostly be used by the config object to obtain the full
// content of the source to aggregate all the data into his internal storing
// config instance.
func (s *Source) Get(
	path string,
	def ...interface{},
) (interface{}, error) {
	// lock the source for changes
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// retrieve the source stored config path stored value
	return s.config.Get(path, def...)
}
