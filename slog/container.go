package slog

import (
	"github.com/happyhippyhippo/slate"
)

// GetFormatterFactory will try to retrieve the registered formatter
// factory instance from the application service container.
func GetFormatterFactory(c slate.ServiceContainer) (*FormatterFactory, error) {
	instance, err := c.Get(ContainerFormatterFactoryID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*FormatterFactory)
	if !ok {
		return nil, errConversion(instance, "*FormatterFactory")
	}
	return i, nil
}

// GetFormatterStrategies will try to retrieve the registered the list of
// formatter strategies instances from the application service container.
func GetFormatterStrategies(c slate.ServiceContainer) ([]FormatterStrategy, error) {
	tags, err := c.Tagged(ContainerFormatterStrategyTag)
	if err != nil {
		return nil, err
	}

	var list []FormatterStrategy
	for _, service := range tags {
		s, ok := service.(FormatterStrategy)
		if !ok {
			return nil, errConversion(service, "FormatterStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetStreamFactory will try to retrieve the registered stream
// factory instance from the application service container.
func GetStreamFactory(c slate.ServiceContainer) (*StreamFactory, error) {
	instance, err := c.Get(ContainerStreamFactoryID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*StreamFactory)
	if !ok {
		return nil, errConversion(instance, "*StreamFactory")
	}
	return i, nil
}

// GetStreamStrategies will try to retrieve the registered the list of
// stream strategies instances from the application service container.
func GetStreamStrategies(c slate.ServiceContainer) ([]StreamStrategy, error) {
	tags, err := c.Tagged(ContainerStreamStrategyTag)
	if err != nil {
		return nil, err
	}

	var list []StreamStrategy
	for _, service := range tags {
		s, ok := service.(StreamStrategy)
		if !ok {
			return nil, errConversion(service, "StreamStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetLogger will try to retrieve the registered logger manager
// instance from the application service container.
func GetLogger(c slate.ServiceContainer) (*Logger, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*Logger)
	if !ok {
		return nil, errConversion(instance, "*Logger")
	}
	return i, nil
}

// GetLoader will try to retrieve the registered loader
// instance from the application service container.
func GetLoader(c slate.ServiceContainer) (*Loader, error) {
	instance, err := c.Get(ContainerLoaderID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*Loader)
	if !ok {
		return nil, errConversion(instance, "*Loader")
	}
	return i, nil
}
