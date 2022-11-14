package log

import (
	"github.com/happyhippyhippo/slate"
	sconfig "github.com/happyhippyhippo/slate/config"
	sfs "github.com/happyhippyhippo/slate/fs"
)

// Provider defines the log module service provider to be used on
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
		formatterFactory, e := GetFormatterFactory(c)
		if e != nil {
			return nil, e
		}
		return newStreamStrategyConsole(formatterFactory)
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamStrategyFileID, func() (interface{}, error) {
		if filesystem, e := sfs.Get(c); e != nil {
			return nil, e
		} else if formatterFactory, e := GetFormatterFactory(c); e != nil {
			return nil, e
		} else {
			return newStreamStrategyFile(filesystem, formatterFactory)
		}
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamStrategyRotatingFileID, func() (interface{}, error) {
		if filesystem, e := sfs.Get(c); e != nil {
			return nil, e
		} else if formatterFactory, e := GetFormatterFactory(c); e != nil {
			return nil, e
		} else {
			return newStreamStrategyRotatingFile(filesystem, formatterFactory)
		}
	}, ContainerStreamStrategyTag)

	_ = c.Service(ContainerStreamFactoryID, func() (interface{}, error) {
		return &StreamFactory{}, nil
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		return NewLogger(), nil
	})

	_ = c.Service(ContainerLoaderID, func() (interface{}, error) {
		if cfg, e := sconfig.Get(c); e != nil {
			return nil, e
		} else if logger, e := Get(c); e != nil {
			return nil, e
		} else if sFactory, e := GetStreamFactory(c); e != nil {
			return nil, e
		} else {
			return newLoader(cfg, logger, sFactory)
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

	if formatterFactory, e := GetFormatterFactory(c); e != nil {
		return e
	} else if strategies, e := GetFormatterStrategies(c); e != nil {
		return e
	} else {
		for _, strategy := range strategies {
			_ = formatterFactory.Register(strategy)
		}
	}

	if sFactory, e := GetStreamFactory(c); e != nil {
		return e
	} else if strategies, e := GetStreamStrategies(c); e != nil {
		return e
	} else {
		for _, strategy := range strategies {
			_ = sFactory.Register(strategy)
		}
	}

	if !LoaderActive {
		return nil
	}

	loader, e := GetLoader(c)
	if e != nil {
		return e
	}
	return loader.Load()
}
