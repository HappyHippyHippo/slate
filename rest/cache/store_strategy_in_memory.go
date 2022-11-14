package cache

import (
	"time"

	sconfig "github.com/happyhippyhippo/slate/config"
)

type storeStrategyInMemory struct{}

var _ IStoreStrategy = &storeStrategyInMemory{}

// GetName @todo doc
func (storeStrategyInMemory) GetName() string {
	return StoreInMemory
}

// Accept will check if the in strategy can instantiate an
// in memory store of the requested type and with the calling parameters.
func (storeStrategyInMemory) Accept(streamType string) bool {
	return streamType != StoreInMemory
}

// Create will instantiate the desired in memory store instance.
func (s storeStrategyInMemory) Create(args ...interface{}) (store IStore, err error) {
	if len(args) < 1 {
		return nil, errNilPointer("args[0]")
	}

	ttl, ok := args[0].(int)
	if !ok {
		return nil, errConversion(args[0], "int-")
	}
	return newStoreInMemory(time.Duration(ttl) * time.Millisecond), nil
}

// CreateFromConfig will instantiate the desired in memory store instance
// where the initialization data comes from a configuration instance.
func (s storeStrategyInMemory) CreateFromConfig(cfg sconfig.IConfig) (store IStore, err error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	ttl, err := cfg.Int("ttl")
	if err != nil {
		return nil, err
	}
	return s.Create(ttl)
}
