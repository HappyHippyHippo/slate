package config

import (
	"io"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/happyhippyhippo/slate/trigger"
)

type sourceRef struct {
	id       string
	priority int
	source   ISource
}

type sourceRefSorter []sourceRef

func (sources sourceRefSorter) Len() int {
	return len(sources)
}

func (sources sourceRefSorter) Swap(i, j int) {
	sources[i], sources[j] = sources[j], sources[i]
}

func (sources sourceRefSorter) Less(i, j int) bool {
	return sources[i].priority < sources[j].priority
}

// IObserver callback function used to be called when an observed
// configuration path has changed.
type IObserver func(interface{}, interface{})

type observerRef struct {
	path     string
	current  interface{}
	callback IObserver
}

// IManager defined an interface to an instance that
// manages configuration
type IManager interface {
	io.Closer
	IConfig

	HasSource(id string) bool
	AddSource(id string, priority int, src ISource) error
	RemoveSource(id string) error
	RemoveAllSources() error
	Source(id string) (ISource, error)
	SourcePriority(id string, priority int) error
	HasObserver(path string) bool
	AddObserver(path string, callback IObserver) error
	RemoveObserver(path string)
}

// Manager defines an object responsible to manage the application
// several defines sources and config values observers.
type Manager struct {
	mutex     sync.Locker
	sources   []sourceRef
	observers []observerRef
	config    *Config
	loader    trigger.ITrigger
}

var _ IConfig = &Manager{}
var _ IManager = &Manager{}

// NewManager instantiate a new configuration object.
// This object will manage a series of sources, alongside of the ability of
// registration of configuration path/values observer callbacks that will be
// called whenever the value has changed.
func NewManager(
	period time.Duration,
) *Manager {
	// instantiate the config manager
	c := &Manager{
		mutex:     &sync.Mutex{},
		sources:   []sourceRef{},
		observers: []observerRef{},
		config:    &Config{},
		loader:    nil,
	}
	// check if there is a need to create the observable sources
	// recurring trigger
	if period != 0 {
		// create the recurring trigger used to poll the
		// observable sources
		c.loader, _ = trigger.NewRecurring(period, func() error {
			return c.reload()
		})
	}
	return c
}

// Close terminates the config instance.
// This will stop the observer trigger and call close on
// all registered sources.
func (m *Manager) Close() error {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// iterate through all the manager sources checking if we can close then
	for _, ref := range m.sources {
		// check if the iterated source implements the closer interface
		if source, ok := ref.source.(io.Closer); ok {
			// close the source
			if e := source.Close(); e != nil {
				return e
			}
		}
	}
	// check if a recurring trigger was generated on creation
	// for observable sources polling
	if m.loader != nil {
		// terminate the recurring trigger
		if e := m.loader.Close(); e != nil {
			return e
		}
		m.loader = nil
	}
	return nil
}

// Entries will retrieve the list of stored entries any registered source.
func (m *Manager) Entries() []string {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// retrieve the stored entries list
	return m.config.Entries()
}

// Has will check if a path has been loaded.
// This means that if the values has been loaded by any registered source.
func (m *Manager) Has(
	path string,
) bool {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// check if the requested path exists in the stored config
	return m.config.Has(path)
}

// Get will retrieve a configuration value loaded from a source.
func (m *Manager) Get(
	path string,
	def ...interface{},
) (interface{}, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to retrieve the requested value
	return m.config.Get(path, def...)
}

// Bool will retrieve a bool configuration value loaded from a source.
func (m *Manager) Bool(
	path string,
	def ...bool,
) (bool, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to retrieve a boolean value from the local config
	return m.config.Bool(path, def...)
}

// Int will retrieve an integer configuration value loaded from a source.
func (m *Manager) Int(
	path string,
	def ...int,
) (int, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to retrieve an integer value from the local config
	return m.config.Int(path, def...)
}

// Float will retrieve a floating point configuration value loaded from a source.
func (m *Manager) Float(
	path string,
	def ...float64,
) (float64, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to retrieve a float value from the local config
	return m.config.Float(path, def...)
}

// String will retrieve a string configuration value loaded from a source.
func (m *Manager) String(
	path string,
	def ...string,
) (string, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to retrieve a string value from the local config
	return m.config.String(path, def...)
}

// List will retrieve a list configuration value loaded from a source.
func (m *Manager) List(
	path string,
	def ...[]interface{},
) ([]interface{}, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to retrieve a list value from the local config
	return m.config.List(path, def...)
}

// Config will retrieve config values loaded from a source.
func (m *Manager) Config(
	path string,
	def ...Config,
) (IConfig, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to retrieve a config value from the local config
	return m.config.Config(path, def...)
}

// Populate will retrieve a config value loaded from a source.
func (m *Manager) Populate(
	path string,
	data interface{},
	icase ...bool,
) (interface{}, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to populate a value from the local config
	return m.config.Populate(path, data, icase...)
}

// HasSource check if a source with a specific id has been registered.
func (m *Manager) HasSource(
	id string,
) bool {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to find a source with the requested id
	for _, ref := range m.sources {
		if ref.id == id {
			return true
		}
	}
	return false
}

// AddSource register a new source with a specific id with a given priority.
func (m *Manager) AddSource(
	id string,
	priority int,
	src ISource,
) error {
	// check the config argument reference
	if src == nil {
		return errNilPointer("src")
	}
	// check if there is already a registered config source with the given id
	if m.HasSource(id) {
		return errDuplicateSource(id)
	}
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// add the config source to the manager and sort them so that the
	// config data can be correctly merged
	m.sources = append(m.sources, sourceRef{id, priority, src})
	sort.Sort(sourceRefSorter(m.sources))
	// rebuild the local config with the source's config information
	m.rebuild()
	return nil
}

// RemoveSource remove a source from the registration list
// of the configuration. This will also update the configuration content and
// re-validate the observed paths.
func (m *Manager) RemoveSource(
	id string,
) error {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to find the requested source to be removed
	for i, ref := range m.sources {
		if ref.id == id {
			// check if the source implements the closer interface
			if src, ok := ref.source.(io.Closer); ok {
				// close the removing source
				if e := src.Close(); e != nil {
					return e
				}
			}
			// remove the source from the managed sources
			m.sources = append(m.sources[:i], m.sources[i+1:]...)
			// rebuild the local config
			m.rebuild()
			return nil
		}
	}
	return nil
}

// RemoveAllSources remove all the registered sources from the registration
// list of the configuration. This will also update the configuration content and
// re-validate the observed paths.
func (m *Manager) RemoveAllSources() error {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// iterate through all the stored sources
	for _, ref := range m.sources {
		// check if the iterated source implements the close interface
		if src, ok := ref.source.(io.Closer); ok {
			// close the source
			if e := src.Close(); e != nil {
				return e
			}
		}
	}
	// recreate the sources array and rebuild the local config
	m.sources = []sourceRef{}
	m.rebuild()
	return nil
}

// Source retrieve a previously registered source with a requested id.
func (m *Manager) Source(
	id string,
) (ISource, error) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to find the requested source
	for _, ref := range m.sources {
		if ref.id == id {
			return ref.source, nil
		}
	}
	return nil, errSourceNotFound(id)
}

// SourcePriority set a priority value of a previously registered
// source with the specified id. This may change the defined values if there
// was an override process of the configuration paths of the changing source.
func (m *Manager) SourcePriority(
	id string,
	priority int,
) error {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to find the requested source to be updated
	for i, ref := range m.sources {
		if ref.id == id {
			// redefine the stored source priority
			m.sources[i] = sourceRef{
				id:       ref.id,
				priority: priority,
				source:   ref.source,
			}
			// sort the sources and rebuild the local config
			sort.Sort(sourceRefSorter(m.sources))
			m.rebuild()
			return nil
		}
	}
	return errSourceNotFound(id)
}

// HasObserver check if there is an observer to a configuration value path.
func (m *Manager) HasObserver(
	path string,
) bool {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// check if the requested observer is registered
	for _, oreg := range m.observers {
		if oreg.path == path {
			return true
		}
	}
	return false
}

// AddObserver register a new observer to a configuration path.
func (m *Manager) AddObserver(
	path string,
	callback IObserver,
) error {
	// validate the callback argument reference
	if callback == nil {
		return errNilPointer("callback")
	}
	// check if the requested path is present
	val, e := m.Get(path)
	if e != nil {
		return e
	}
	// if the founded value is a config, clone it, so
	// it can be used for update checks
	if v, ok := val.(Config); ok {
		val = v.Clone()
	}
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// register the requested observer with the current path value
	m.observers = append(m.observers, observerRef{path, val, callback})
	return nil
}

// RemoveObserver remove an observer to a configuration path.
func (m *Manager) RemoveObserver(
	path string,
) {
	// lock the manager for handling
	m.mutex.Lock()
	defer m.mutex.Unlock()
	// try to find the observer to be removed
	for i, oreg := range m.observers {
		if oreg.path == path {
			// remove the found observer
			m.observers = append(m.observers[:i], m.observers[i+1:]...)
			return
		}
	}
}

func (m *Manager) reload() error {
	// iterate through all stores sources
	reloaded := false
	for _, ref := range m.sources {
		// check if the iterated source is an observable source
		if s, ok := ref.source.(IObservableSource); ok {
			// reload the source and update the reloaded flag if the request
			// resulted in a source info update
			updated, _ := s.Reload()
			reloaded = reloaded || updated
		}
	}
	// check if the iteration resulted in an update of any info
	if reloaded {
		// lock the manager for handling
		m.mutex.Lock()
		defer m.mutex.Unlock()
		// rebuild the local config with the new source info
		m.rebuild()
	}
	return nil
}

func (m *Manager) rebuild() {
	// iterate through all the stored sources
	updated := Config{}
	for _, ref := range m.sources {
		// retrieve the source stored config information
		// and merge it with all parsed sources
		cfg, _ := ref.source.Get("")
		updated.merge(cfg.(Config))
	}
	// store locally the resulting config
	m.config = &updated
	// iterate through all observers
	for id, observer := range m.observers {
		// retrieve the observer path value
		// and check if the current value differs from the previous one
		val, e := m.config.Get(observer.path)
		if e == nil && !reflect.DeepEqual(observer.current, val) {
			// store the new value in the observer registry
			// and call the observer callback
			old := observer.current
			m.observers[id].current = val
			observer.callback(old, val)
		}
	}
}
