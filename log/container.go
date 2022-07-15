package log

import "github.com/happyhippyhippo/slate"

// GetFormatterFactory will try to retrieve the registered formatter
// factory instance from the application service container.
func GetFormatterFactory(c slate.ServiceContainer) (IFormatterFactory, error) {
	instance, e := c.Get(ContainerFormatterFactoryID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(IFormatterFactory)
	if !ok {
		return nil, errConversion(instance, "IFormatterFactory")
	}
	return i, nil
}

// GetFormatterStrategies will try to retrieve the registered the list of
// formatter strategies instances from the application service container.
func GetFormatterStrategies(c slate.ServiceContainer) ([]IFormatterStrategy, error) {
	tags, e := c.Tagged(ContainerFormatterStrategyTag)
	if e != nil {
		return nil, e
	}

	var list []IFormatterStrategy
	for _, service := range tags {
		s, ok := service.(IFormatterStrategy)
		if !ok {
			return nil, errConversion(service, "IFormatterStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetStreamFactory will try to retrieve the registered stream
// factory instance from the application service container.
func GetStreamFactory(c slate.ServiceContainer) (IStreamFactory, error) {
	instance, e := c.Get(ContainerStreamFactoryID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(IStreamFactory)
	if !ok {
		return nil, errConversion(instance, "IStreamFactory")
	}
	return i, nil
}

// GetStreamStrategies will try to retrieve the registered the list of
// stream strategies instances from the application service container.
func GetStreamStrategies(c slate.ServiceContainer) ([]IStreamStrategy, error) {
	tags, e := c.Tagged(ContainerStreamStrategyTag)
	if e != nil {
		return nil, e
	}

	var list []IStreamStrategy
	for _, service := range tags {
		s, ok := service.(IStreamStrategy)
		if !ok {
			return nil, errConversion(service, "IStreamStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// Get will try to retrieve the registered logger manager
// instance from the application service container.
func Get(c slate.ServiceContainer) (ILogger, error) {
	instance, e := c.Get(ContainerID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(ILogger)
	if !ok {
		return nil, errConversion(instance, "ILogger")
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
