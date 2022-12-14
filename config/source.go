package config

import (
	"sync"
)

const (
	// SourceUnknown defines the value to be used to declare an
	// unknown config source type.
	SourceUnknown = "unknown"

	// SourceEnv defines the value to be used to declare an
	// environment config source type.
	SourceEnv = "env"

	// SourceFile defines the value to be used to declare a
	// simple file config source type.
	SourceFile = "file"

	// SourceObservableFile defines the value to be used to
	// declare an observable file config source type.
	SourceObservableFile = "observable-file"

	// SourceDirectory defines the value to be used to declare a
	// simple dir config source type.
	SourceDirectory = "dir"

	// SourceRest defines the value to be used to declare a
	// rest config source type.
	SourceRest = "rest"

	// SourceObservableRest defines the value to be used to
	// declare an observable rest config source type.
	SourceObservableRest = "observable-rest"

	// SourceAggregate defines the value to be used to declare a
	// container loading configs source type.
	SourceAggregate = "aggregate"
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
	mutex   sync.Locker
	partial Config
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
	return s.partial.Has(path)
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
	return s.partial.Get(path, def...)
}
