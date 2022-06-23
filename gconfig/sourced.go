package gconfig

import "io"

// Sourced defined an interface to an instance that holds
// configuration values from different sources
type Sourced interface {
	io.Closer

	HasSource(id string) bool
	AddSource(id string, priority int, src Source) error
	RemoveSource(id string) error
	RemoveAllSources() error
	Source(id string) (Source, error)
	SourcePriority(id string, priority int) error
}
