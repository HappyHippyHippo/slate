package slate

// IServiceProvider is an interface used to define the methods of an object
// that can be registered into a slate application and register elements in
// the application container and do some necessary boot actions on
// initialization.
type IServiceProvider interface {
	Register(ServiceContainer) error
	Boot(ServiceContainer) error
}
