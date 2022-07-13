package sconfig

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sfs"
	"time"
)

// Provider defines the slate.config module service provider to be used on
// the application initialization to register the config service.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will register the configuration section instances in the
// application container.
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Service(ContainerDecoderStrategyYAMLID, func() (interface{}, error) {
		return &decoderStrategyYAML{}, nil
	}, ContainerDecoderStrategyTag)

	_ = c.Service(ContainerDecoderStrategyJSONID, func() (interface{}, error) {
		return &decoderStrategyJSON{}, nil
	}, ContainerDecoderStrategyTag)

	_ = c.Service(ContainerDecoderFactoryID, func() (interface{}, error) {
		return &DecoderFactory{}, nil
	})

	_ = c.Service(ContainerSourceStrategyFileID, func() (interface{}, error) {
		if filesystem, err := sfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if decoderFactory, err := GetDecoderFactory(c); err != nil {
			return nil, err
		} else {
			return newSourceStrategyFile(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyFileObservableID, func() (interface{}, error) {
		if filesystem, err := sfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if decoderFactory, err := GetDecoderFactory(c); err != nil {
			return nil, err
		} else {
			return newSourceStrategyObservableFile(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyDirID, func() (interface{}, error) {
		if filesystem, err := sfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if decoderFactory, err := GetDecoderFactory(c); err != nil {
			return nil, err
		} else {
			return newSourceStrategyDir(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyRestID, func() (interface{}, error) {
		decoderFactory, err := GetDecoderFactory(c)
		if err != nil {
			return nil, err
		}
		return newSourceStrategyRest(decoderFactory)
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyRestObservableID, func() (interface{}, error) {
		decoderFactory, err := GetDecoderFactory(c)
		if err != nil {
			return nil, err
		}
		return newSourceStrategyObservableRest(decoderFactory)
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyEnvID, func() (interface{}, error) {
		return &sourceStrategyEnv{}, nil
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyContainerID, func() (interface{}, error) {
		partials, e := GetSourceContainerPartials(c)
		if e != nil {
			return nil, e
		}

		return &sourceStrategyContainer{partials: partials}, nil
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceFactoryID, func() (interface{}, error) {
		return &SourceFactory{}, nil
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		return NewManager(time.Duration(ObserveFrequency) * time.Second), nil
	})

	_ = c.Service(ContainerLoaderID, func() (interface{}, error) {
		if config, err := GetConfig(c); err != nil {
			return nil, err
		} else if sourceFactory, err := GetSourceFactory(c); err != nil {
			return nil, err
		} else {
			return newLoader(config, sourceFactory)
		}
	})

	return nil
}

// Boot will start the configuration config instance by calling the
// configuration loader with the defined provider base entry information.
func (p Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	if decoderFactory, err := GetDecoderFactory(c); err != nil {
		return err
	} else if decoderStrategies, err := GetDecoderStrategies(c); err != nil {
		return err
	} else {
		for _, strategy := range decoderStrategies {
			_ = decoderFactory.Register(strategy)
		}
	}

	if sourceFactory, err := GetSourceFactory(c); err != nil {
		return err
	} else if sourceStrategies, err := GetSourceStrategies(c); err != nil {
		return err
	} else {
		for _, strategy := range sourceStrategies {
			_ = sourceFactory.Register(strategy)
		}
	}

	if !LoaderActive {
		return nil
	}

	loader, err := GetLoader(c)
	if err != nil {
		return err
	}
	return loader.Load()
}

// GetDecoderFactory will try to retrieve the registered decoder
// factory instance from the application service container.
func GetDecoderFactory(c slate.ServiceContainer) (IDecoderFactory, error) {
	instance, err := c.Get(ContainerDecoderFactoryID)
	if err != nil {
		return nil, err
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
	tags, err := c.Tagged(ContainerDecoderStrategyTag)
	if err != nil {
		return nil, err
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
	instance, err := c.Get(ContainerSourceFactoryID)
	if err != nil {
		return nil, err
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
	tags, err := c.Tagged(ContainerSourceStrategyTag)
	if err != nil {
		return nil, err
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
	tags, err := c.Tagged(ContainerSourceContainerPartialTag)
	if err != nil {
		return nil, err
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

// GetConfig will try to retrieve the registered config
// instance from the application service container.
func GetConfig(c slate.ServiceContainer) (IManager, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
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
	instance, err := c.Get(ContainerLoaderID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(ILoader)
	if !ok {
		return nil, errConversion(instance, "ILoader")
	}
	return i, nil
}
