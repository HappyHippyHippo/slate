package config

import (
	"time"

	"github.com/happyhippyhippo/slate"
	sfs "github.com/happyhippyhippo/slate/fs"
)

// Provider defines the config module service provider to be used on
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
		return &decoderFactory{}, nil
	})

	_ = c.Service(ContainerSourceStrategyFileID, func() (interface{}, error) {
		if filesystem, e := sfs.Get(c); e != nil {
			return nil, e
		} else if decoderFactory, e := GetDecoderFactory(c); e != nil {
			return nil, e
		} else {
			return newSourceStrategyFile(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyFileObservableID, func() (interface{}, error) {
		if filesystem, e := sfs.Get(c); e != nil {
			return nil, e
		} else if decoderFactory, e := GetDecoderFactory(c); e != nil {
			return nil, e
		} else {
			return newSourceStrategyObservableFile(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyDirID, func() (interface{}, error) {
		if filesystem, e := sfs.Get(c); e != nil {
			return nil, e
		} else if decoderFactory, e := GetDecoderFactory(c); e != nil {
			return nil, e
		} else {
			return newSourceStrategyDir(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyRestID, func() (interface{}, error) {
		decoderFactory, e := GetDecoderFactory(c)
		if e != nil {
			return nil, e
		}
		return newSourceStrategyRest(decoderFactory)
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyRestObservableID, func() (interface{}, error) {
		decoderFactory, e := GetDecoderFactory(c)
		if e != nil {
			return nil, e
		}
		return newSourceStrategyObservableRest(decoderFactory)
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyEnvID, func() (interface{}, error) {
		return &sourceStrategyEnv{}, nil
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyAggregateID, func() (interface{}, error) {
		partials, e := GetSourceContainerPartials(c)
		if e != nil {
			return nil, e
		}

		return &sourceStrategyAggregate{partials: partials}, nil
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceFactoryID, func() (interface{}, error) {
		return &sourceFactory{}, nil
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		return NewManager(time.Duration(ObserveFrequency) * time.Second), nil
	})

	_ = c.Service(ContainerLoaderID, func() (interface{}, error) {
		if config, e := Get(c); e != nil {
			return nil, e
		} else if sourceFactory, e := GetSourceFactory(c); e != nil {
			return nil, e
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

	if decoderFactory, e := GetDecoderFactory(c); e != nil {
		return e
	} else if strategies, e := GetDecoderStrategies(c); e != nil {
		return e
	} else {
		for _, strategy := range strategies {
			_ = decoderFactory.Register(strategy)
		}
	}

	if sourceFactory, e := GetSourceFactory(c); e != nil {
		return e
	} else if strategies, e := GetSourceStrategies(c); e != nil {
		return e
	} else {
		for _, strategy := range strategies {
			_ = sourceFactory.Register(strategy)
		}
	}

	if !LoaderActive {
		return nil
	}

	l, e := GetLoader(c)
	if e != nil {
		return e
	}
	return l.Load()
}
