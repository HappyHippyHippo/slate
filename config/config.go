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
	source   Source
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

// Observer callback function used to be called when an observed
// configuration path has changed.
type Observer func(interface{}, interface{})

type observerRef struct {
	path     string
	current  interface{}
	callback Observer
}

type ticker interface {
	io.Closer
	Delay() time.Duration
}

// Config defines an object responsible to config the application
// several defines sources and partial values observers.
type Config struct {
	mutex     sync.Locker
	sources   []sourceRef
	observers []observerRef
	partial   *Partial
	observer  ticker
}

// NewConfig instantiate a new configuration object.
// This object will config a series of sources, alongside of the ability of
// registration of configuration path/values observer callbacks that will be
// called whenever the value has changed.
func NewConfig() *Config {
	// instantiate the config
	c := &Config{
		mutex:     &sync.Mutex{},
		sources:   []sourceRef{},
		observers: []observerRef{},
		partial:   &Partial{},
		observer:  nil,
	}
	// check if there is a need to create the observable sources
	// ticker trigger
	if period := time.Duration(ObserveFrequency) * time.Millisecond; period != 0 {
		// create the ticker trigger used to poll the
		// observable sources
		c.observer, _ = trigger.NewRecurring(period, func() error {
			return c.reload()
		})
	}
	return c
}

// Close terminates the config instance.
// This will stop the observer trigger and call close on
// all registered sources.
func (c *Config) Close() error {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// iterate through all the config sources checking if we can close then
	for _, ref := range c.sources {
		// check if the iterated source implements the closer interface
		if source, ok := ref.source.(io.Closer); ok {
			// close the source
			if e := source.Close(); e != nil {
				return e
			}
		}
	}
	// check if a ticker trigger was generated on creation
	// for observable sources polling
	if c.observer != nil {
		// terminate the ticker trigger
		if e := c.observer.Close(); e != nil {
			return e
		}
		c.observer = nil
	}
	return nil
}

// Entries will retrieve the list of stored entries any registered source.
func (c *Config) Entries() []string {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// retrieve the stored entries list
	return c.partial.Entries()
}

// Has will check if a path has been loaded.
// This means that if the values has been loaded by any registered source.
func (c *Config) Has(
	path string,
) bool {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// check if the requested path exists in the stored partial
	return c.partial.Has(path)
}

// Get will retrieve a configuration value loaded from a source.
func (c *Config) Get(
	path string,
	def ...interface{},
) (interface{}, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve the requested value
	return c.partial.Get(path, def...)
}

// Bool will retrieve a bool configuration value loaded from a source.
func (c *Config) Bool(
	path string,
	def ...bool,
) (bool, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a boolean value from the local partial
	return c.partial.Bool(path, def...)
}

// Int will retrieve an integer configuration value loaded from a source.
func (c *Config) Int(
	path string,
	def ...int,
) (int, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve an integer value from the local partial
	return c.partial.Int(path, def...)
}

// Float will retrieve a floating point configuration value loaded from a source.
func (c *Config) Float(
	path string,
	def ...float64,
) (float64, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a float value from the local partial
	return c.partial.Float(path, def...)
}

// String will retrieve a string configuration value loaded from a source.
func (c *Config) String(
	path string,
	def ...string,
) (string, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a string value from the local partial
	return c.partial.String(path, def...)
}

// List will retrieve a list configuration value loaded from a source.
func (c *Config) List(
	path string,
	def ...[]interface{},
) ([]interface{}, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a list value from the local partial
	return c.partial.List(path, def...)
}

// Partial will retrieve partial values loaded from a source.
func (c *Config) Partial(
	path string,
	def ...Partial,
) (*Partial, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to retrieve a partial value from the local partial
	return c.partial.Partial(path, def...)
}

// Populate will retrieve a config value loaded from a source.
func (c *Config) Populate(
	path string,
	data interface{},
	icase ...bool,
) (interface{}, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to populate a value from the local partial
	return c.partial.Populate(path, data, icase...)
}

// HasSource check if a source with a specific id has been registered.
func (c *Config) HasSource(
	id string,
) bool {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find a source with the requested id
	for _, ref := range c.sources {
		if ref.id == id {
			return true
		}
	}
	return false
}

// AddSource register a new source with a specific id with a given priority.
func (c *Config) AddSource(
	id string,
	priority int,
	src Source,
) error {
	// check the source argument reference
	if src == nil {
		return errNilPointer("source")
	}
	// check if there is already a registered source with the given id
	if c.HasSource(id) {
		return errDuplicateSource(id)
	}
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// add the source to the config and sort them so that the
	// data can be correctly merged
	c.sources = append(c.sources, sourceRef{id, priority, src})
	sort.Sort(sourceRefSorter(c.sources))
	// rebuild the local partial with the source's partial information
	c.rebuild()
	return nil
}

// RemoveSource remove a source from the registration list
// of the configuration. This will also update the configuration content and
// re-validate the observed paths.
func (c *Config) RemoveSource(
	id string,
) error {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find the requested source to be removed
	for i, ref := range c.sources {
		if ref.id != id {
			continue
		}
		// check if the source implements the closer interface
		if src, ok := ref.source.(io.Closer); ok {
			// close the removing source
			if e := src.Close(); e != nil {
				return e
			}
		}
		// remove the source from the config sources
		c.sources = append(c.sources[:i], c.sources[i+1:]...)
		// rebuild the local partial
		c.rebuild()
		return nil
	}
	return nil
}

// RemoveAllSources remove all the registered sources from the registration
// list of the configuration. This will also update the configuration content and
// re-validate the observed paths.
func (c *Config) RemoveAllSources() error {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// iterate through all the stored sources
	for _, ref := range c.sources {
		// check if the iterated source implements the close interface
		if src, ok := ref.source.(io.Closer); ok {
			// close the source
			if e := src.Close(); e != nil {
				return e
			}
		}
	}
	// recreate the sources array and rebuild the local partial
	c.sources = []sourceRef{}
	c.rebuild()
	return nil
}

// Source retrieve a previously registered source with a requested id.
func (c *Config) Source(
	id string,
) (Source, error) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find the requested source
	for _, ref := range c.sources {
		if ref.id == id {
			return ref.source, nil
		}
	}
	return nil, errSourceNotFound(id)
}

// SourcePriority set a priority value of a previously registered
// source with the specified id. This may change the defined values if there
// was an override process of the configuration paths of the changing source.
func (c *Config) SourcePriority(
	id string,
	priority int,
) error {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find the requested source to be updated
	for i, ref := range c.sources {
		if ref.id != id {
			continue
		}
		// redefine the stored source priority
		c.sources[i] = sourceRef{
			id:       ref.id,
			priority: priority,
			source:   ref.source,
		}
		// sort the sources and rebuild the local partial
		sort.Sort(sourceRefSorter(c.sources))
		c.rebuild()
		return nil
	}
	return errSourceNotFound(id)
}

// HasObserver check if there is an observer to a configuration value path.
func (c *Config) HasObserver(
	path string,
) bool {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// check if the requested observer is registered
	for _, oreg := range c.observers {
		if oreg.path == path {
			return true
		}
	}
	return false
}

// AddObserver register a new observer to a configuration path.
func (c *Config) AddObserver(
	path string,
	callback Observer,
) error {
	// validate the callback argument reference
	if callback == nil {
		return errNilPointer("callback")
	}
	// check if the requested path is present
	val, e := c.Get(path)
	if e != nil {
		return e
	}
	// if the founded value is a partial, clone it, so
	// it can be used for update checks
	if v, ok := val.(Partial); ok {
		val = v.Clone()
	}
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// register the requested observer with the current path value
	c.observers = append(c.observers, observerRef{path, val, callback})
	return nil
}

// RemoveObserver remove an observer to a configuration path.
func (c *Config) RemoveObserver(
	path string,
) {
	// lock the config for handling
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// try to find the observer to be removed
	for i, oreg := range c.observers {
		if oreg.path == path {
			// remove the found observer
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			return
		}
	}
}

func (c *Config) reload() error {
	// iterate through all stores sources
	reloaded := false
	for _, ref := range c.sources {
		// check if the iterated source is an observable source
		if s, ok := ref.source.(ObsSource); ok {
			// reload the source and update the reloaded flag if the request
			// resulted in a source info update
			updated, _ := s.Reload()
			reloaded = reloaded || updated
		}
	}
	// check if the iteration resulted in an update of any info
	if reloaded {
		// lock the config for handling
		c.mutex.Lock()
		defer c.mutex.Unlock()
		// rebuild the local partial with the new source info
		c.rebuild()
	}
	return nil
}

func (c *Config) rebuild() {
	// iterate through all the stored sources
	updated := Partial{}
	for _, ref := range c.sources {
		// retrieve the source stored partial information
		// and Merge it with all parsed sources
		cfg, _ := ref.source.Get("")
		updated.Merge(cfg.(Partial))
	}
	// store locally the resulting partial
	c.partial = &updated
	// iterate through all observers
	for id, observer := range c.observers {
		// retrieve the observer path value
		// and check if the current value differs from the previous one
		val, e := c.partial.Get(observer.path)
		if e == nil && !reflect.DeepEqual(observer.current, val) {
			// store the new value in the observer registry
			// and call the observer callback
			old := observer.current
			c.observers[id].current = val
			observer.callback(old, val)
		}
	}
}
