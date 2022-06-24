package sconfig

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sfs"
	"time"
)

// Provider defines the slate.config module service provider to be used on
// the application initialization to register the config service.
type Provider struct{}

var _ slate.ServiceProvider = &Provider{}

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
			return NewSourceStrategyFile(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyFileObservableID, func() (interface{}, error) {
		if filesystem, err := sfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if decoderFactory, err := GetDecoderFactory(c); err != nil {
			return nil, err
		} else {
			return NewSourceStrategyObservableFile(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyDirID, func() (interface{}, error) {
		if filesystem, err := sfs.GetFileSystem(c); err != nil {
			return nil, err
		} else if decoderFactory, err := GetDecoderFactory(c); err != nil {
			return nil, err
		} else {
			return NewSourceStrategyDir(filesystem, decoderFactory)
		}
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyRestID, func() (interface{}, error) {
		decoderFactory, err := GetDecoderFactory(c)
		if err != nil {
			return nil, err
		}
		return NewSourceStrategyRest(decoderFactory)
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyRestObservableID, func() (interface{}, error) {
		decoderFactory, err := GetDecoderFactory(c)
		if err != nil {
			return nil, err
		}
		return NewSourceStrategyObservableRest(decoderFactory)
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceStrategyEnvID, func() (interface{}, error) {
		return &sourceStrategyEnv{}, nil
	}, ContainerSourceStrategyTag)

	_ = c.Service(ContainerSourceFactoryID, func() (interface{}, error) {
		return &SourceFactory{}, nil
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		return NewConfig(time.Duration(ObserveFrequency) * time.Second), nil
	})

	_ = c.Service(ContainerLoaderID, func() (interface{}, error) {
		if config, err := GetConfig(c); err != nil {
			return nil, err
		} else if sourceFactory, err := GetSourceFactory(c); err != nil {
			return nil, err
		} else {
			return NewLoader(config, sourceFactory)
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
