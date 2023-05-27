package config

const (
	// UnknownSource defines the value to be used to declare an
	// unknown partial source type.
	UnknownSource = "unknown"
)

// SourceStrategy interface defines the methods of the source
// factory strategy that will be used instantiate a particular source type.
type SourceStrategy interface {
	Accept(partial *Partial) bool
	Create(partial *Partial) (Source, error)
}
