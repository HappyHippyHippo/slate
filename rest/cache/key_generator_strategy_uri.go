package cache

import (
	sconfig "github.com/happyhippyhippo/slate/config"
)

type keyGeneratorStrategyURI struct{}

var _ KeyGeneratorStrategy = &keyGeneratorStrategyURI{}

// Accept will check if the console stream factory strategy can instantiate a
// stream of the requested type and with the calling parameters.
func (keyGeneratorStrategyURI) Accept(generatorType string) bool {
	return generatorType == KeyGeneratorURI
}

// AcceptFromConfig will check if the stream factory strategy can instantiate
// a stream where the data to check comes from a configuration partial
// instance.
func (s keyGeneratorStrategyURI) AcceptFromConfig(cfg sconfig.IConfig) bool {
	if cfg == nil {
		return false
	}

	if generatorType, err := cfg.String("type"); err == nil {
		return s.Accept(generatorType)
	}

	return false
}

// Create will instantiate the desired stream instance.
func (keyGeneratorStrategyURI) Create(_ ...interface{}) (KeyGenerator, error) {
	return keyGeneratorURI, nil
}

// CreateFromConfig will instantiate the desired stream instance where
// the initialization data comes from a configuration instance.
func (s keyGeneratorStrategyURI) CreateFromConfig(cfg sconfig.IConfig) (stream KeyGenerator, err error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	return s.Create()
}
