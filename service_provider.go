package slate

// ServiceProvider is an interface used to define the methods of an object that
// can be registered into a slate application and register elements in the
// application container and do some necessary boot actions on initialization.
type ServiceProvider interface {
	Register(ServiceContainer) error
	Boot(ServiceContainer) error
}
