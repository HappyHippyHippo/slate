package slate

// IApplication defines the interface of a slate application instance.
type IApplication interface {
	Add(provider IServiceProvider) error
	Boot() error
}

// Application defines the structure that controls the application
// container and container initialization.
type Application struct {
	Container ServiceContainer
	providers []IServiceProvider
	isBoot    bool
}

var _ IApplication = &Application{}

// NewApplication used to instantiate a new application.
func NewApplication() *Application {
	return &Application{
		Container: ServiceContainer{},
		providers: []IServiceProvider{},
		isBoot:    false,
	}
}

// Add will register a new provider into the application used
// on the application boot.
func (a *Application) Add(provider IServiceProvider) error {
	if provider == nil {
		return errNilPointer("provider")
	}

	if e := provider.Register(a.Container); e != nil {
		return e
	}
	a.providers = append(a.providers, provider)

	return nil
}

// Boot initializes the application if not initialized yet.
// The initialization of an application is the calling of the register method
// on all providers, after the registration of all objects in the container,
// the boot method of all providers will be executed.
func (a *Application) Boot() (e error) {
	defer func() {
		if r := recover(); r != nil {
			if tr, ok := r.(error); !ok {
				panic(r)
			} else {
				e = tr
			}
		}
	}()

	if !a.isBoot {
		for _, provider := range a.providers {
			if e := provider.Boot(a.Container); e != nil {
				return e
			}
		}
		a.isBoot = true
	}
	return nil
}
