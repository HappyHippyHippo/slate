package gconfig

import (
	"github.com/happyhippyhippo/slate/gtrigger"
	"io"
	"reflect"
	"sort"
	"sync"
	"time"
)

// Manager defined an interface to a instance that manages configuration
type Manager interface {
	Config
	Sourced
	Observable
}

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

// config defines the instance of a configuration managing structure.
type manager struct {
	mutex     sync.Locker
	sources   []sourceRef
	observers []observerRef
	partial   *Partial
	loader    gtrigger.Trigger
}

var _ Config = &manager{}
var _ Sourced = &manager{}
var _ Observable = &manager{}
var _ Manager = &manager{}

// NewConfig instantiate a new configuration object.
// This object will manage a series of sources, alongside of the ability of
// registration of configuration path/values observer callbacks that will be
// called whenever the value has changed.
func NewConfig(period time.Duration) Manager {
	c := &manager{
		mutex:     &sync.Mutex{},
		sources:   []sourceRef{},
		observers: []observerRef{},
		partial:   &Partial{},
		loader:    nil,
	}

	if period != 0 {
		c.loader, _ = gtrigger.NewRecurring(period, func() error {
			return c.reload()
		})
	}

	return c
}

// Close terminates the config instance.
// This will stop the observer trigger and call close on all registered sources.
func (c *manager) Close() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, ref := range c.sources {
		if source, ok := ref.source.(io.Closer); ok {
			if err := source.Close(); err != nil {
				return err
			}
		}
	}

	if c.loader != nil {
		if err := c.loader.Close(); err != nil {
			return err
		}
		c.loader = nil
	}

	return nil
}

// Has will check if a path has been loaded.
// This means that if the values has been loaded by any registered source.
func (c *manager) Has(path string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.partial.Has(path)
}

// Get will retrieve a configuration value loaded from a source.
func (c *manager) Get(path string, def ...interface{}) (interface{}, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(def) > 0 {
		return c.partial.Get(path, def[0])
	}
	return c.partial.Get(path)
}

// Bool will retrieve a bool configuration value loaded from a source.
func (c *manager) Bool(path string, def ...bool) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.partial.Bool(path, def...)
}

// Int will retrieve an integer configuration value loaded from a source.
func (c *manager) Int(path string, def ...int) (int, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.partial.Int(path, def...)
}

// Float will retrieve a floating point configuration value loaded from a source.
func (c *manager) Float(path string, def ...float64) (float64, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.partial.Float(path, def...)
}

// String will retrieve a string configuration value loaded from a source.
func (c *manager) String(path string, def ...string) (string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.partial.String(path, def...)
}

// List will retrieve a list configuration value loaded from a source.
func (c *manager) List(path string, def ...[]interface{}) ([]interface{}, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.partial.List(path, def...)
}

// Partial will retrieve a config configuration value loaded from a source.
func (c *manager) Partial(path string, def ...Partial) (Partial, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.partial.Partial(path, def...)
}

// HasSource check if a source with a specific id has been registered.
func (c *manager) HasSource(id string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, ref := range c.sources {
		if ref.id == id {
			return true
		}
	}
	return false
}

// AddSource register a new source with a specific id with a given priority.
func (c *manager) AddSource(id string, priority int, src Source) error {
	if src == nil {
		return errNilPointer("src")
	}

	if c.HasSource(id) {
		return errDuplicateConfigSource(id)
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.sources = append(c.sources, sourceRef{id, priority, src})
	sort.Sort(sourceRefSorter(c.sources))
	c.rebuild()

	return nil
}

// RemoveSource remove a source from the registration list
// of the configuration. This will also update the configuration content and
// re-validate the observed paths.
func (c *manager) RemoveSource(id string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i, ref := range c.sources {
		if ref.id == id {
			if source, ok := ref.source.(io.Closer); ok {
				if err := source.Close(); err != nil {
					return err
				}
			}
			c.sources = append(c.sources[:i], c.sources[i+1:]...)
			c.rebuild()
			return nil
		}
	}
	return nil
}

// RemoveAllSources remove all the registered sources from the registration
// list of the configuration. This will also update the configuration content and
// re-validate the observed paths.
func (c *manager) RemoveAllSources() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, ref := range c.sources {
		if source, ok := ref.source.(io.Closer); ok {
			if err := source.Close(); err != nil {
				return err
			}
		}
	}
	c.sources = []sourceRef{}
	c.rebuild()

	return nil
}

// Source retrieve a previously registered source with a requested id.
func (c *manager) Source(id string) (Source, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, ref := range c.sources {
		if ref.id == id {
			return ref.source, nil
		}
	}

	return nil, errConfigSourceNotFound(id)
}

// SourcePriority set a priority value of a previously registered
// source with the specified id. This may change the defined values if there
// was an override process of the configuration paths of the changing source.
func (c *manager) SourcePriority(id string, priority int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i, ref := range c.sources {
		if ref.id == id {
			c.sources[i] = sourceRef{
				id:       ref.id,
				priority: priority,
				source:   ref.source,
			}
			sort.Sort(sourceRefSorter(c.sources))
			c.rebuild()

			return nil
		}
	}

	return errConfigSourceNotFound(id)
}

// HasObserver check if there is an observer to a configuration value path.
func (c *manager) HasObserver(path string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, oreg := range c.observers {
		if oreg.path == path {
			return true
		}
	}
	return false
}

// AddObserver register a new observer to a configuration path.
func (c *manager) AddObserver(path string, callback Observer) error {
	if callback == nil {
		return errNilPointer("callback")
	}

	val, err := c.Get(path)
	if err != nil {
		return err
	}

	if v, ok := val.(Partial); ok {
		val = v.Clone()
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.observers = append(c.observers, observerRef{path, val, callback})

	return nil
}

// RemoveObserver remove an observer to a configuration path.
func (c *manager) RemoveObserver(path string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i, oreg := range c.observers {
		if oreg.path == path {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			return
		}
	}
}

func (c *manager) reload() error {
	reloaded := false
	for _, ref := range c.sources {
		if s, ok := ref.source.(SourceObservable); ok {
			updated, _ := s.Reload()
			reloaded = reloaded || updated
		}
	}

	if reloaded {
		c.mutex.Lock()
		defer c.mutex.Unlock()

		c.rebuild()
	}

	return nil
}

func (c *manager) rebuild() {
	updated := Partial{}
	for _, ref := range c.sources {
		cfg, _ := ref.source.Get("")
		updated.merge(cfg.(Partial))
	}

	c.partial = &updated

	for id, observer := range c.observers {
		val, err := c.partial.Get(observer.path)
		if err == nil && !reflect.DeepEqual(observer.current, val) {
			old := observer.current
			c.observers[id].current = val
			observer.callback(old, val)
		}
	}
}
