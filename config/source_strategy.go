package config

const (
	// SourceStrategyUnknown defines the value to be used to declare an
	// unknown config source type.
	SourceStrategyUnknown = "unknown"
)

// ISourceStrategy interface defines the methods of the source
// factory strategy that will be used instantiate a particular source type.
type ISourceStrategy interface {
	Accept(config IConfig) bool
	Create(config IConfig) (ISource, error)
}
