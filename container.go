package slate

import (
	"io"
	"reflect"

	"github.com/happyhippyhippo/slate/dig"
)

type containerEntry struct {
	factory  interface{}
	getter   func() (any, error)
	remover  func() error
	ctype    reflect.Type
	tags     []string
	instance interface{}
}

func newContainerServiceEntry(
	container *Container,
	id string,
	factory interface{},
	ctype reflect.Type,
	tags ...string,
) containerEntry {
	var getter func() (any, error)
	var remover func() error
	var entry containerEntry
	ctor := ctype.Out(0)
	// define a service creation method
	getter = func() (any, error) {
		// check if the service has been already been instantiated
		if container.entries[id].instance != nil {
			return container.entries[id].instance, nil
		}
		// instantiate the service
		results, e := container.container.Get(ctor)
		if e != nil {
			return nil, errContainer(e)
		}
		instance := results[0]
		// store the instance in the registration entry if needed
		container.entries[id] = containerEntry{
			factory:  factory,
			getter:   getter,
			remover:  remover,
			ctype:    ctor,
			tags:     tags,
			instance: instance,
		}
		return instance, nil
	}
	// define a service removal method
	remover = func() error {
		// check if the entry has been created
		if container.entries[id].instance != nil {
			// check if the instance implements the closer interface
			if closer, ok := container.entries[id].instance.(io.Closer); ok {
				if e := closer.Close(); e != nil {
					return e
				}
			}
		}
		// remove the factory from the instantiation container
		_ = container.container.Remove(factory)
		// remove the registration entry
		delete(container.entries, id)
		return nil
	}
	// store the factory in the instantiation container
	_ = container.container.Provide(factory)
	// return a populated service container entry
	entry = containerEntry{
		factory:  factory,
		getter:   getter,
		remover:  remover,
		ctype:    ctype.Out(0),
		tags:     tags,
		instance: nil,
	}
	return entry
}

func (e *containerEntry) hasTag(
	tag string,
) bool {
	// search for the requested tag in the entry tag list
	found := false
	for _, t := range e.tags {
		if t == tag {
			found = true
		}
	}
	return found
}

// IContainer defines the interface of a slate
// application service container instance.
type IContainer interface {
	io.Closer

	Has(id string) bool
	Service(id string, factory interface{}, tags ...string) error
	Get(id string) (any, error)
	Tag(tag string) ([]any, error)
	Remove(id string) error
	Clear() error
}

// Container defines the structure that hold the application service
// factories and initialized services.
type Container struct {
	entries   map[string]containerEntry
	container *dig.Container
}

var _ IContainer = &Container{}

// NewContainer used to instantiate a new application service container.
func NewContainer() *Container {
	return &Container{
		entries:   map[string]containerEntry{},
		container: dig.New(),
	}
}

// Close clean up the container from all the stored objects.
// If the object has been already instantiated and implements the Closable
// interface, then the Close method will be called upon the removing instance.
func (c *Container) Close() error {
	// remove all the elements from the container
	if e := c.Clear(); e != nil {
		return e
	}
	// release the container allocated memory
	c.entries = nil
	c.container = nil
	return nil
}

// Has will check if an object is registered with the requested id.
// This does not mean that is instantiated. The instantiation is just executed
// when the instance is requested for the first time.
func (c *Container) Has(
	id string,
) bool {
	// check if the entry exists
	_, ok := c.entries[id]
	return ok
}

// Service will register a service factory used to return an instance
// generated by the given factory. This factory method will only be called
// once, meaning that everytime the service is requested, is always returned
// the same instance.
// If any object was registered previously with the requested id, then the
// object will be removed by calling the Remove method previously the storing
// of the new object factory.
func (c *Container) Service(
	id string,
	factory interface{},
	tags ...string,
) error {
	// check if the factory argument is a valid pointer
	if factory == nil {
		return errNilPointer("factory")
	}
	// check if the passed factory is a valid function
	t := reflect.TypeOf(factory)
	if t.Kind() != reflect.Func {
		return errNonFunctionFactory(t.Name())
	}
	// check if the factory does return (at least) a value
	if t.NumOut() == 0 {
		return errFactoryWithoutResult(t.Name())
	}
	// check if there is an entry with the requested id
	if _, ok := c.entries[id]; ok {
		// remove the previously registered service
		if e := c.Remove(id); e != nil {
			return e
		}
	}
	// store the entry registry
	c.entries[id] = newContainerServiceEntry(c, id, factory, t, tags...)
	return nil
}

// Get will retrieve the requested object from the container.
// If the object has not yet been instantiated, then the factory method
// will be executed to instantiate it.
func (c *Container) Get(
	id string,
) (any, error) {
	// check if there is a registry with the requested id
	entry, ok := c.entries[id]
	if !ok {
		return nil, errServiceNotFound(id)
	}
	// retrieve the service
	return entry.getter()
}

// Tag will retrieve the list of entries instances that where registered
// with a tag list containing the request teg.
func (c *Container) Tag(
	tag string,
) ([]any, error) {
	var result []any
	// search all the registered entries for the requested tag
	for id, entry := range c.entries {
		if entry.hasTag(tag) {
			instance := entry.instance
			// check if the service as been instantiated already
			if instance == nil {
				// instantiate the service from the instantiation container
				i, e := c.Get(id)
				if e != nil {
					return nil, e
				}
				instance = i
			}
			// store the tagged service instance
			result = append(result, instance)
		}
	}
	return result, nil
}

// Remove will eliminate the object from the container.
// If the object has been already instantiated and implements the Closable
// interface, then the Close method will be called on the removing instance.
func (c *Container) Remove(
	id string,
) error {
	// check if the service is registered
	entry, ok := c.entries[id]
	if !ok {
		return nil
	}
	// remove the service from the container
	return entry.remover()
}

// Clear will eliminate all the registered object from the container.
func (c *Container) Clear() error {
	// remove all the registration entries
	for id := range c.entries {
		if e := c.Remove(id); e != nil {
			return e
		}
	}
	return nil
}
