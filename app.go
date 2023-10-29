package slate

import (
	"fmt"
	"io"
	"reflect"

	"github.com/happyhippyhippo/slate/dig"
)

// ----------------------------------------------------------------------------
// defs
// ----------------------------------------------------------------------------

const (
	// ContainerID defines the slate package di service id base string.
	ContainerID = "slate"

	// EnvID defines the slate package environment variable base name.
	EnvID = "SLATE"
)

// ----------------------------------------------------------------------------
// errors
// ----------------------------------------------------------------------------

var (
	// ErrServiceContainer defines a di error.
	ErrServiceContainer = NewError("service di error")

	// ErrNonFunctionServiceFactory defines a service di registration error
	// that signals that the registration request was made with a
	// non-function service factory.
	ErrNonFunctionServiceFactory = NewError("non-function service factory")

	// ErrServiceFactoryWithoutResult defines a service di registration error
	// that signals that the registration request was made with a
	// function service factory that don't return a service.
	ErrServiceFactoryWithoutResult = NewError("service factory without result")

	// ErrServiceNotFound defines a service not found on the di.
	ErrServiceNotFound = NewError("service not found")
)

func errServiceContainer(
	e error,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrServiceContainer, fmt.Errorf("%w", e).Error(), ctx...)
}

func errNonFunctionServiceFactory(
	arg string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrNonFunctionServiceFactory, arg, ctx...)
}

func errServiceFactoryWithoutResult(
	arg string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrServiceFactoryWithoutResult, arg, ctx...)
}

func errServiceNotFound(
	arg string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrServiceNotFound, arg, ctx...)
}

// ----------------------------------------------------------------------------
// service Provider
// ----------------------------------------------------------------------------

type serviceContainerEntry struct {
	factory     interface{}
	getter      func() (any, error)
	remover     func() error
	reflectType reflect.Type
	tags        []string
	instance    interface{}
}

func newServiceContainerEntry(
	container *ServiceContainer,
	id string,
	factory interface{},
	reflectType reflect.Type,
	tags ...string,
) serviceContainerEntry {
	var getter func() (any, error)
	var remover func() error
	var entry serviceContainerEntry
	ctor := reflectType.Out(0)
	// define a service creation method
	getter = func() (any, error) {
		// check if the service has been already been instantiated
		if container.entries[id].instance != nil {
			return container.entries[id].instance, nil
		}
		// instantiate the service
		results, e := container.di.Get(ctor)
		if e != nil {
			return nil, errServiceContainer(e)
		}
		instance := results[0]
		// store the instance in the registration entry if needed
		container.entries[id] = serviceContainerEntry{
			factory:     factory,
			getter:      getter,
			remover:     remover,
			reflectType: ctor,
			tags:        tags,
			instance:    instance,
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
		// remove the factory from the instantiation di
		_ = container.di.Remove(factory)
		// remove the registration entry
		delete(container.entries, id)
		return nil
	}
	// store the factory in the instantiation di
	_ = container.di.Provide(factory)
	// return a populated service di entry
	entry = serviceContainerEntry{
		factory:     factory,
		getter:      getter,
		remover:     remover,
		reflectType: reflectType.Out(0),
		tags:        tags,
		instance:    nil,
	}
	return entry
}

func (e *serviceContainerEntry) hasTag(
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

// ServiceContainer defines the structure that hold the application service
// factories and initialized services.
type ServiceContainer struct {
	entries map[string]serviceContainerEntry
	di      *dig.Container
}

// NewServiceContainer used to instantiate a new application service di.
func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{
		entries: map[string]serviceContainerEntry{},
		di:      dig.New(),
	}
}

// Close clean up the di from all the stored objects.
// If the object has been already instantiated and implements the Closable
// interface, then the Close method will be called upon the removing instance.
func (c *ServiceContainer) Close() error {
	// remove all the elements from the di
	if e := c.Clear(); e != nil {
		return e
	}
	// release the di allocated memory
	c.entries = nil
	c.di = nil
	return nil
}

// Has will check if a service is registered with the requested id.
// This does not mean that is instantiated. The instantiation is just executed
// when the instance is requested for the first time.
func (c *ServiceContainer) Has(
	id string,
) bool {
	// check if the entry exists
	_, ok := c.entries[id]
	return ok
}

// Add will register a service factory used to create a service instance
// generated by the given factory. This factory method will only be called
// once, meaning that everytime the service is requested, is always returned
// the same instance.
// If any service was registered previously with the requested id, then the
// service will be removed by calling the Remove method before the storing
// of the new service factory.
func (c *ServiceContainer) Add(
	id string,
	factory interface{},
	tags ...string,
) error {
	// check if the factory argument is a valid pointer
	if factory == nil {
		return errNilPointer("factory")
	}
	// check if the passed factory is a valid function
	reflectType := reflect.TypeOf(factory)
	if reflectType.Kind() != reflect.Func {
		return errNonFunctionServiceFactory(reflectType.Name())
	}
	// check if the factory does return (at least) a value
	if reflectType.NumOut() == 0 {
		return errServiceFactoryWithoutResult(reflectType.Name())
	}
	// check if there is an entry with the requested id
	if _, ok := c.entries[id]; ok {
		// remove the previously registered service
		if e := c.Remove(id); e != nil {
			return e
		}
	}
	// store the entry registry
	c.entries[id] = newServiceContainerEntry(c, id, factory, reflectType, tags...)
	return nil
}

// Get will retrieve the requested service from the di.
// If the object has not yet been instantiated, then the factory method
// will be executed to instantiate it.
func (c *ServiceContainer) Get(
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

// Tag will retrieve the list of entries connections that where registered
// with the request teg.
func (c *ServiceContainer) Tag(
	tag string,
) ([]any, error) {
	var result []any

	// search all the registered entries for the requested tag
	for id, entry := range c.entries {
		if entry.hasTag(tag) {
			instance := entry.instance
			// check if the service as been instantiated already
			if instance == nil {
				// instantiate the service from the instantiation di
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

// Remove will eliminate the service from the di.
// If the service has been already instantiated and implements the Closable
// interface, then the Close method will be called on the removing instance.
func (c *ServiceContainer) Remove(
	id string,
) error {
	// check if the service is registered
	entry, ok := c.entries[id]
	if !ok {
		return nil
	}
	// remove the service from the di
	return entry.remover()
}

// Clear will eliminate all the registered services from the di.
func (c *ServiceContainer) Clear() error {
	// remove all the registration entries
	for id := range c.entries {
		if e := c.Remove(id); e != nil {
			return e
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// service provider
// ----------------------------------------------------------------------------

// ServiceProvider defines the base application provider registry
// interface of an object used to register services in the application di.
type ServiceProvider interface {
	Provide(container *ServiceContainer) error
	Boot(container *ServiceContainer) error
}

// ----------------------------------------------------------------------------
// service register
// ----------------------------------------------------------------------------

// ServiceRegister defines the base application provider registry object
// used to register services in the application di Provider
type ServiceRegister struct {
	App *App
}

var _ ServiceProvider = &ServiceRegister{}

// NewServiceRegister will instantiate a new provider instance
func NewServiceRegister(
	app ...*App,
) *ServiceRegister {
	register := &ServiceRegister{}
	if len(app) > 0 {
		register.App = app[0]
	}
	return register
}

// Provide method will register the services in the application di
func (ServiceRegister) Provide(
	container *ServiceContainer,
) error {
	if container == nil {
		return errNilPointer("Provider")
	}
	return nil
}

// Boot will execute the registry instance boot process
func (ServiceRegister) Boot(
	container *ServiceContainer,
) error {
	if container == nil {
		return errNilPointer("Provider")
	}
	return nil
}

// ----------------------------------------------------------------------------
// application
// ----------------------------------------------------------------------------

// App defines the structure that controls the application
// di and di initialization.
type App struct {
	ServiceContainer

	providers []ServiceProvider
	isBoot    bool
}

// NewApp used to instantiate a new application.
func NewApp() *App {
	return &App{
		ServiceContainer: *NewServiceContainer(),
		providers:        []ServiceProvider{},
		isBoot:           false,
	}
}

// Provide will register a new provider into the application and by doing so,
// allow to the provider register his services in the application service
// Provider.
// At boot will also call the provider boot method so that the provider
// services and also be initialized.
func (a *App) Provide(
	provider ServiceProvider,
) error {
	// check provider argument
	if provider == nil {
		return errNilPointer("provider")
	}
	// call the provider registration method over the
	// application service di
	if e := provider.Provide(&a.ServiceContainer); e != nil {
		return e
	}
	// add the provider to the application provider slice
	a.providers = append(a.providers, provider)
	return nil
}

// Boot initializes the application if not initialized yet.
// The initialization of an application is made by calling of the Provide
// method on all providers, after the registration of all services in the di,
// the boot method of all providers will be executed.
func (a *App) Boot() error {
	// check if the application has already been booted
	if !a.isBoot {
		// call boot on all the registered providers
		for _, provider := range a.providers {
			if e := provider.Boot(&a.ServiceContainer); e != nil {
				return e
			}
		}
		a.isBoot = true
	}
	return nil
}
