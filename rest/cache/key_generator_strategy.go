package cache

import (
	sconfig "github.com/happyhippyhippo/slate/config"
)

// KeyGeneratorStrategy interface defines the methods of the key generator
// factory strategy that can validate creation requests and instantiation
// of particular type of key generator.
type KeyGeneratorStrategy interface {
	Accept(generatorType string) bool
	AcceptFromConfig(cfg sconfig.IConfig) bool
	Create(args ...interface{}) (KeyGenerator, error)
	CreateFromConfig(cfg sconfig.IConfig) (KeyGenerator, error)
}
