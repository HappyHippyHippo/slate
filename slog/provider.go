package slog

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/sfs"
)

// Provider defines the slate.log module service provider to be used on
// the application initialization to register the logging service.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will register the logger package instances in the
// application container.
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Service(ContainerFormatterStrategyJSONID, func() (interface{}, error) {
		return &formatterStrategyJSON{}, nil
	}, ContainerFormatterStrategyTag)

	_ = c.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
		return &FormatterFactory{}, nil
	})

	_ = c.Service(ContainerStreamStrategyConsoleID, func() (interface{}, error) {
		factory, err := GetFormatterFactory(c)
		if err != nil {
			return nil, err
		}
		return newStreamStrategyConsole(factory)
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamStrategyFileID, func() (interface{}, error) {
		if filesystem, err := sfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if formatterFactory, err := GetFormatterFactory(c); err != nil {
			return nil, err
		} else {
			return newStreamStrategyFile(filesystem, formatterFactory)
		}
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamStrategRotatingFileID, func() (interface{}, error) {
		if filesystem, err := sfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if formatterFactory, err := GetFormatterFactory(c); err != nil {
			return nil, err
		} else {
			return newStreamStrategyRotatingFile(filesystem, formatterFactory)
		}
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamFactoryID, func() (interface{}, error) {
		return &StreamFactory{}, nil
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		return newLogger(), nil
	})

	_ = c.Service(ContainerLoaderID, func() (interface{}, error) {
		if config, err := sconfig.GetConfig(c); err != nil {
			return nil, err
		} else if logger, err := GetLogger(c); err != nil {
			return nil, err
		} else if streamFactory, err := GetStreamFactory(c); err != nil {
			return nil, err
		} else {
			return newLoader(config, logger, streamFactory)
		}
	})

	return nil
}

// Boot will start the logger package config instance by calling the
// logger loader with the defined provider base entry information.
func (p Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	if formatterFactory, err := GetFormatterFactory(c); err != nil {
		return err
	} else if formatterStrategies, err := GetFormatterStrategies(c); err != nil {
		return err
	} else {
		for _, strategy := range formatterStrategies {
			_ = formatterFactory.Register(strategy)
		}
	}

	if streamFactory, err := GetStreamFactory(c); err != nil {
		return err
	} else if streamStrategies, err := GetStreamStrategies(c); err != nil {
		return err
	} else {
		for _, strategy := range streamStrategies {
			_ = streamFactory.Register(strategy)
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

// GetFormatterFactory will try to retrieve the registered formatter
// factory instance from the application service container.
func GetFormatterFactory(c slate.ServiceContainer) (IFormatterFactory, error) {
	instance, err := c.Get(ContainerFormatterFactoryID)
	if err != nil {
		return nil, err
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
	tags, err := c.Tagged(ContainerFormatterStrategyTag)
	if err != nil {
		return nil, err
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
	instance, err := c.Get(ContainerStreamFactoryID)
	if err != nil {
		return nil, err
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
	tags, err := c.Tagged(ContainerStreamStrategyTag)
	if err != nil {
		return nil, err
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

// GetLogger will try to retrieve the registered logger manager
// instance from the application service container.
func GetLogger(c slate.ServiceContainer) (ILogger, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
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
