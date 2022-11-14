package cache

import (
	sconfig "github.com/happyhippyhippo/slate/config"
	"io"
)

// IStoreFactory @todo doc
type IStoreFactory interface {
	Register(strategy IStoreStrategy) error
	Create(storeType string) (IStore, error)
}

type storeFactoryReg struct {
	storeType string
	strategy  IStoreStrategy
	instance  IStore
}

type storeFactory struct {
	cfg    sconfig.IConfig
	stores map[string]storeFactoryReg
}

// NewStoreFactory @todo doc
func NewStoreFactory(cfg sconfig.IConfig) (IStoreFactory, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	return &storeFactory{
		cfg:    cfg,
		stores: map[string]storeFactoryReg{},
	}, nil
}

// Register will register a new store strategy to be used
// on creation requests.
func (f *storeFactory) Register(strategy IStoreStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	storeType := strategy.GetName()

	if reg, ok := f.stores[storeType]; ok && reg.instance != nil {
		if closer, ok := reg.instance.(io.Closer); ok {
			if e := closer.Close(); e != nil {
				return e
			}
		}
	}

	f.stores[storeType] = storeFactoryReg{
		storeType: storeType,
		strategy:  strategy,
		instance:  nil,
	}

	return nil
}

// Create will instantiate and return a new store.
func (f storeFactory) Create(storeType string) (IStore, error) {
	reg, ok := f.stores[storeType]
	if !ok {
		return nil, errInvalidStoreType("storeType")
	}

	if reg.instance == nil {
		cfg, e := f.cfg.Partial(storeType)
		if e != nil {
			return nil, e
		}

		store, e := reg.strategy.CreateFromConfig(&cfg)
		if e != nil {
			return nil, e
		}

		f.stores[storeType] = storeFactoryReg{
			storeType: storeType,
			strategy:  reg.strategy,
			instance:  store,
		}
	}

	return f.stores[storeType].instance, nil
}
