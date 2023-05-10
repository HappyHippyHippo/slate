package source

import (
	"sync"

	"github.com/happyhippyhippo/slate/config"
)

// BaseSource defines a config source that read a file content
// and stores its config contents to be used as a config.
type BaseSource struct {
	Mutex  sync.Locker
	Config config.Config
}

var _ config.ISource = &BaseSource{}

// Has will check if the requested path is present in the source
// configuration content.
func (s *BaseSource) Has(
	path string,
) bool {
	// lock the source for changes
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	// check if the source stored config has the requested path
	return s.Config.Has(path)
}

// Get will retrieve the value stored in the requested path present in the
// configuration content.
// If the path does not exist, then the value nil will be returned.
// This method will mostly be used by the config object to obtain the full
// content of the source to aggregate all the data into his internal storing
// config instance.
func (s *BaseSource) Get(
	path string,
	def ...interface{},
) (interface{}, error) {
	// lock the source for changes
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	// retrieve the source stored config path stored value
	return s.Config.Get(path, def...)
}
