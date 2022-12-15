package slate

// IApplication defines the interface of a slate application instance.
type IApplication interface {
	IContainer
	IProvider

	Provide(provider IProvider) error
}

// Application defines the structure that controls the application
// container and container initialization.
type Application struct {
	Container

	providers []IProvider
	isBoot    bool
}

var _ IApplication = &Application{}

// NewApplication used to instantiate a new application.
func NewApplication() *Application {
	return &Application{
		Container: *NewContainer(),
		providers: []IProvider{},
		isBoot:    false,
	}
}

// Provide will register a new provider into the application used
// on the application boot.
func (a *Application) Provide(
	provider IProvider,
) error {
	// check provider argument
	if provider == nil {
		return errNilPointer("provider")
	}
	// call the provider registration method over the
	// application service container
	if e := provider.Register(a); e != nil {
		return e
	}
	// add the provider to the application provider slice
	a.providers = append(a.providers, provider)
	return nil
}

// Register is a no-op method, is present just to provide the application
// instance the interface of a provider, making it easy to define an
// application as a main provider instance.
func (a *Application) Register(_ ...IContainer) error {
	return nil
}

// Boot initializes the application if not initialized yet.
// The initialization of an application is the calling of the register method
// on all providers, after the registration of all objects in the container,
// the boot method of all providers will be executed.
func (a *Application) Boot(_ ...IContainer) (e error) {
	// boot panic fallback
	defer func() {
		if r := recover(); r != nil {
			if tr, ok := r.(error); !ok {
				panic(r)
			} else {
				e = tr
			}
		}
	}()
	// check if the application has already been booted
	if !a.isBoot {
		// call boot on all the registered providers
		for _, provider := range a.providers {
			if e := provider.Boot(a); e != nil {
				return e
			}
		}
		a.isBoot = true
	}
	return nil
}
