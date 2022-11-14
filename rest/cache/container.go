package cache

import (
	"github.com/happyhippyhippo/slate"
	srest "github.com/happyhippyhippo/slate/rest"
)

// GetKeyGeneratorFactory will try to retrieve the registered key
// generator factory instance from the application service container.
func GetKeyGeneratorFactory(c slate.ServiceContainer) (*KeyGeneratorFactory, error) {
	instance, err := c.Get(ContainerKeyGeneratorFactoryID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*KeyGeneratorFactory)
	if !ok {
		return nil, errConversion(instance, "*cache.KeyGeneratorFactory")
	}
	return i, nil
}

// GetKeyGeneratorStrategies will try to retrieve the registered the list of
// key generator strategies instances from the application service container.
func GetKeyGeneratorStrategies(c slate.ServiceContainer) ([]KeyGeneratorStrategy, error) {
	tags, err := c.Tagged(ContainerKeyGeneratorStrategyTag)
	if err != nil {
		return nil, err
	}

	var list []KeyGeneratorStrategy
	for _, service := range tags {
		s, ok := service.(KeyGeneratorStrategy)
		if !ok {
			return nil, errConversion(service, "cache.KeyGeneratorStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetStoreFactory will try to retrieve the registered
// store factory instance from the application service container.
func GetStoreFactory(c slate.ServiceContainer) (IStoreFactory, error) {
	instance, err := c.Get(ContainerStoreFactoryID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(IStoreFactory)
	if !ok {
		return nil, errConversion(instance, "cache.IStoreFactory")
	}
	return i, nil
}

// GetStoreStrategies will try to retrieve the registered the list of
// store strategies instances from the application service container.
func GetStoreStrategies(c slate.ServiceContainer) ([]IStoreStrategy, error) {
	tags, err := c.Tagged(ContainerStoreStrategyTag)
	if err != nil {
		return nil, err
	}

	var list []IStoreStrategy
	for _, service := range tags {
		s, ok := service.(IStoreStrategy)
		if !ok {
			return nil, errConversion(service, "cache.IStoreStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetMiddlewareGenerator will try to retrieve the registered logging
// middleware for ok responses instance from the application service
// container.
func GetMiddlewareGenerator(c slate.ServiceContainer) (func(id string) (srest.Middleware, error), error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(func(id string) (srest.Middleware, error))
	if !ok {
		return nil, errConversion(instance, "func(id string) (srest.Middleware, error)")
	}
	return i, nil
}
