package glog

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gfs"
)

// Provider defines the slate.log module service provider to be used on
// the application initialization to register the logging service.
type Provider struct{}

var _ slate.ServiceProvider = &Provider{}

// Register will register the logger package instances in the
// application container.
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Service(ContainerFormatterStrategyJSONID, func() (interface{}, error) {
		return &FormatterStrategyJSON{}, nil
	}, ContainerFormatterStrategyTag)

	_ = c.Service(ContainerFormatterFactoryID, func() (interface{}, error) {
		return &FormatterFactory{}, nil
	})

	_ = c.Service(ContainerStreamStrategyConsoleID, func() (interface{}, error) {
		factory, err := GetFormatterFactory(c)
		if err != nil {
			return nil, err
		}
		return NewStreamStrategyConsole(factory)
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamStrategyFileID, func() (interface{}, error) {
		if filesystem, err := gfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if formatterFactory, err := GetFormatterFactory(c); err != nil {
			return nil, err
		} else {
			return NewStreamStrategyFile(filesystem, formatterFactory)
		}
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamStrategRotatingFileID, func() (interface{}, error) {
		if filesystem, err := gfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if formatterFactory, err := GetFormatterFactory(c); err != nil {
			return nil, err
		} else {
			return NewStreamStrategyRotatingFile(filesystem, formatterFactory)
		}
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamFactoryID, func() (interface{}, error) {
		return &StreamFactory{}, nil
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		return NewLogger(), nil
	})

	_ = c.Service(ContainerLoaderID, func() (interface{}, error) {
		if config, err := gconfig.GetConfig(c); err != nil {
			return nil, err
		} else if logger, err := GetLogger(c); err != nil {
			return nil, err
		} else if streamFactory, err := GetStreamFactory(c); err != nil {
			return nil, err
		} else {
			return NewLoader(config, logger, streamFactory)
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
