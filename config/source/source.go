package source

import (
	"sync"

	"github.com/happyhippyhippo/slate/config"
)

// Source defines a partial source that read a file content
// and stores its config contents to be used as a partial.
type Source struct {
	Mutex   sync.Locker
	Partial config.Partial
}

var _ config.Source = &Source{}

// Has will check if the requested path is present in the source
// configuration content.
func (s *Source) Has(
	path string,
) bool {
	// lock the source for changes
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	// check if the source stored config has the requested path
	return s.Partial.Has(path)
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
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	// retrieve the source stored config path stored value
	return s.Partial.Get(path, def...)
}
