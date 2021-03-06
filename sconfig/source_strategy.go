package sconfig

// ISourceStrategy interface defines the methods of the source
// dFactory strategy that will be used instantiate a particular source type.
type ISourceStrategy interface {
	Accept(sourceType string) bool
	AcceptFromConfig(cfg IConfig) bool
	Create(args ...interface{}) (ISource, error)
	CreateFromConfig(cfg IConfig) (ISource, error)
}
