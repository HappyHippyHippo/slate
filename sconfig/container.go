package sconfig

import (
	"github.com/happyhippyhippo/slate"
)

// GetDecoderFactory will try to retrieve the registered decoder
// factory instance from the application service container.
func GetDecoderFactory(c slate.ServiceContainer) (*DecoderFactory, error) {
	instance, err := c.Get(ContainerDecoderFactoryID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*DecoderFactory)
	if !ok {
		return nil, errConversion(instance, "*DecoderFactory")
	}
	return i, nil
}

// GetDecoderStrategies will try to retrieve the registered the list of decoder
// strategies instances from the application service container.
func GetDecoderStrategies(c slate.ServiceContainer) ([]DecoderStrategy, error) {
	tags, err := c.Tagged(ContainerDecoderStrategyTag)
	if err != nil {
		return nil, err
	}

	var list []DecoderStrategy
	for _, service := range tags {
		s, ok := service.(DecoderStrategy)
		if !ok {
			return nil, errConversion(service, "DecoderStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetSourceFactory will try to retrieve the registered source
// factory instance from the application service container.
func GetSourceFactory(c slate.ServiceContainer) (*SourceFactory, error) {
	instance, err := c.Get(ContainerSourceFactoryID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*SourceFactory)
	if !ok {
		return nil, errConversion(instance, "*SourceFactory")
	}
	return i, nil
}

// GetSourceStrategies will try to retrieve the registered the list of source
// strategies instances from the application service container.
func GetSourceStrategies(c slate.ServiceContainer) ([]SourceStrategy, error) {
	tags, err := c.Tagged(ContainerSourceStrategyTag)
	if err != nil {
		return nil, err
	}

	var list []SourceStrategy
	for _, service := range tags {
		s, ok := service.(SourceStrategy)
		if !ok {
			return nil, errConversion(service, "SourceStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetConfig will try to retrieve the registered config
// instance from the application service container.
func GetConfig(c slate.ServiceContainer) (Manager, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(Manager)
	if !ok {
		return nil, errConversion(instance, "Manager")
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
