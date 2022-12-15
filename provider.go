package slate

const (
	// ID defines the slate package container id base string.
	ID = "slate"

	// EnvID defines the slate package base environment variable name.
	EnvID = "SLATE"
)

// IProvider is an interface used to define the methods of an object
// that can be registered into a slate application and register elements in
// the application container and do some necessary boot actions on
// initialization.
type IProvider interface {
	Register(...IContainer) error
	Boot(...IContainer) error
}
