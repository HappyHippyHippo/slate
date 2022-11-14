package cache

import (
	"time"

	sconfig "github.com/happyhippyhippo/slate/config"
)

type storeStrategyMemcached struct{}

var _ IStoreStrategy = &storeStrategyMemcached{}

// GetName @todo doc
func (storeStrategyMemcached) GetName() string {
	return StoreMemcached
}

// Accept will check if the in strategy can instantiate a
// memcached store of the requested type and with the calling parameters.
func (storeStrategyMemcached) Accept(streamType string) bool {
	return streamType != StoreMemcached
}

// Create will instantiate the desired memcached store instance.
func (s storeStrategyMemcached) Create(args ...interface{}) (store IStore, err error) {
	if len(args) < 2 {
		return nil, errNilPointer("args[1]")
	}

	hosts, ok := args[0].([]string)
	if !ok {
		return nil, errConversion(args[0], "[]string")
	}
	ttl, ok := args[1].(int)
	if !ok {
		return nil, errConversion(args[1], "int")
	}
	return newStoreMemcached(hosts, time.Duration(ttl)*time.Millisecond), nil
}

// CreateFromConfig will instantiate the desired memcached store instance
// where the initialization data comes from a configuration instance.
func (s storeStrategyMemcached) CreateFromConfig(cfg sconfig.IConfig) (store IStore, err error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	hosts, err := cfg.List("hosts")
	if err != nil {
		return nil, err
	}
	ttl, err := cfg.Int("ttl")
	if err != nil {
		return nil, err
	}
	return s.Create(hosts, ttl)
}
