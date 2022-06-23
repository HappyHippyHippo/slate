package gconfig

// SourceStrategy interface defines the methods of the source
// factory strategy that will be used instantiate a particular source type.
type SourceStrategy interface {
	Accept(stype string) bool
	AcceptFromConfig(cfg Config) bool
	Create(args ...interface{}) (Source, error)
	CreateFromConfig(cfg Config) (Source, error)
}
