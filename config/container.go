package config

import "github.com/happyhippyhippo/slate"

// GetDecoderFactory will try to retrieve the registered decoder
// factory instance from the application service container.
func GetDecoderFactory(c slate.ServiceContainer) (IDecoderFactory, error) {
	instance, e := c.Get(ContainerDecoderFactoryID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(IDecoderFactory)
	if !ok {
		return nil, errConversion(instance, "IDecoderFactory")
	}
	return i, nil
}

// GetDecoderStrategies will try to retrieve the registered the list of decoder
// strategies instances from the application service container.
func GetDecoderStrategies(c slate.ServiceContainer) ([]IDecoderStrategy, error) {
	tags, e := c.Tagged(ContainerDecoderStrategyTag)
	if e != nil {
		return nil, e
	}

	var list []IDecoderStrategy
	for _, service := range tags {
		s, ok := service.(IDecoderStrategy)
		if !ok {
			return nil, errConversion(service, "IDecoderStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetSourceFactory will try to retrieve the registered source
// factory instance from the application service container.
func GetSourceFactory(c slate.ServiceContainer) (ISourceFactory, error) {
	instance, e := c.Get(ContainerSourceFactoryID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(ISourceFactory)
	if !ok {
		return nil, errConversion(instance, "ISourceFactory")
	}
	return i, nil
}

// GetSourceStrategies will try to retrieve the registered the list of source
// strategies instances from the application service container.
func GetSourceStrategies(c slate.ServiceContainer) ([]ISourceStrategy, error) {
	tags, e := c.Tagged(ContainerSourceStrategyTag)
	if e != nil {
		return nil, e
	}

	var list []ISourceStrategy
	for _, service := range tags {
		s, ok := service.(ISourceStrategy)
		if !ok {
			return nil, errConversion(service, "ISourceStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetSourceContainerPartials will try to retrieve the registered the list
// of source partials instances from the application service container.
func GetSourceContainerPartials(c slate.ServiceContainer) ([]IConfig, error) {
	tags, e := c.Tagged(ContainerSourceContainerPartialTag)
	if e != nil {
		return nil, e
	}

	var list []IConfig
	for _, service := range tags {
		s, ok := service.(IConfig)
		if !ok {
			return nil, errConversion(service, "IConfig")
		}
		list = append(list, s)
	}
	return list, nil
}

// Get will try to retrieve the registered config
// instance from the application service container.
func Get(c slate.ServiceContainer) (IManager, error) {
	instance, e := c.Get(ContainerID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(IManager)
	if !ok {
		return nil, errConversion(instance, "IManager")
	}
	return i, nil
}

// GetLoader will try to retrieve the registered loader
// instance from the application service container.
func GetLoader(c slate.ServiceContainer) (ILoader, error) {
	instance, e := c.Get(ContainerLoaderID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(ILoader)
	if !ok {
		return nil, errConversion(instance, "ILoader")
	}
	return i, nil
}
