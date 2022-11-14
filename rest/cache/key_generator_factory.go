package cache

import (
	sconfig "github.com/happyhippyhippo/slate/config"
)

// KeyGeneratorFactory is a key generator factory based on a
// registered list of generators generation strategies.
type KeyGeneratorFactory []KeyGeneratorStrategy

// Register will register a new key generator strategy to be used
// on creation requests.
func (f *KeyGeneratorFactory) Register(strategy KeyGeneratorStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	*f = append(*f, strategy)

	return nil
}

// Create will instantiate and return a new key generator.
func (f KeyGeneratorFactory) Create(streamType string, args ...interface{}) (KeyGenerator, error) {
	for _, s := range f {
		if s.Accept(streamType) {
			return s.Create(args...)
		}
	}
	return nil, errInvalidKeyGeneratorType("streamType")
}

// CreateFromConfig will instantiate and return a new key generator loaded
// by a configuration instance.
func (f KeyGeneratorFactory) CreateFromConfig(cfg sconfig.IConfig) (KeyGenerator, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	for _, s := range f {
		if s.AcceptFromConfig(cfg) {
			return s.CreateFromConfig(cfg)
		}
	}
	return nil, errInvalidKeyGeneratorPartial(cfg)
}
