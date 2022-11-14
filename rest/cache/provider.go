package cache

import (
	"github.com/happyhippyhippo/slate"
	sconfig "github.com/happyhippyhippo/slate/config"
)

// Provider defines the slate.rest.cache module service provider to be used on
// the application initialization to register the cache services.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will register the configuration section instances in the
// application container.
func (Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Service(ContainerKeyGeneratorStrategyURIID, func() (interface{}, error) {
		return &keyGeneratorStrategyURI{}, nil
	}, ContainerKeyGeneratorStrategyTag)

	_ = c.Service(ContainerKeyGeneratorFactoryID, func() (interface{}, error) {
		return &KeyGeneratorFactory{}, nil
	})

	_ = c.Service(ContainerStoreStrategyInMemoryID, func() (interface{}, error) {
		return &storeStrategyInMemory{}, nil
	}, ContainerStoreStrategyTag)

	_ = c.Service(ContainerStoreStrategyMemcachedID, func() (interface{}, error) {
		return &storeStrategyMemcached{}, nil
	}, ContainerStoreStrategyTag)

	_ = c.Service(ContainerStoreStrategyRedisID, func() (interface{}, error) {
		return &storeStrategyRedis{}, nil
	}, ContainerStoreStrategyTag)

	_ = c.Service(ContainerStoreFactoryID, func() (interface{}, error) {
		cfg, e := sconfig.Get(c)
		if e != nil {
			return nil, e
		}

		cfgStore, e := cfg.Partial(ConfigPathStores)
		if e != nil {
			return nil, e
		}

		return NewStoreFactory(&cfgStore)
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		cfg, e := sconfig.Get(c)
		if e != nil {
			return nil, e
		}

		keygenFact, err := GetKeyGeneratorFactory(c)
		if err != nil {
			return nil, err
		}

		storeFact, err := GetStoreFactory(c)
		if err != nil {
			return nil, err
		}

		return NewMiddlewareGenerator(cfg, keygenFact, storeFact)
	})

	return nil
}

// Boot will start the configuration config instance by calling the
// configuration loader with the defined provider base entry information.
func (Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	if keyGeneratorFactory, err := GetKeyGeneratorFactory(c); err != nil {
		return err
	} else if keyGeneratorStrategies, err := GetKeyGeneratorStrategies(c); err != nil {
		return err
	} else {
		for _, strategy := range keyGeneratorStrategies {
			_ = keyGeneratorFactory.Register(strategy)
		}
	}

	if storeFactory, err := GetStoreFactory(c); err != nil {
		return err
	} else if storeStrategies, err := GetStoreStrategies(c); err != nil {
		return err
	} else {
		for _, strategy := range storeStrategies {
			_ = storeFactory.Register(strategy)
		}
	}

	return nil
}
