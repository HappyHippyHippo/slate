package slate

const (
	// ID defines the slate package container service id base string.
	ID = "slate"

	// EnvID defines the slate package environment variable base name.
	EnvID = "SLATE"
)

// Provider is an interface used to define the methods of an object
// that can be registered into a slate application and register elements in
// the application container and do some necessary boot actions on
// initialization.
type Provider interface {
	Register(container *Container) error
	Boot(container *Container) error
}
