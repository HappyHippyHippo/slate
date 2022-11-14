package cache

import (
	"time"

	sconfig "github.com/happyhippyhippo/slate/config"
)

type storeStrategyMemcachedBinary struct{}

var _ IStoreStrategy = &storeStrategyMemcachedBinary{}

// GetName @todo doc
func (storeStrategyMemcachedBinary) GetName() string {
	return StoreMemcachedBinary
}

// Accept will check if the in strategy can instantiate a
// memory cached with binary connection store of the requested type
// and with the calling parameters.
func (storeStrategyMemcachedBinary) Accept(streamType string) bool {
	return streamType != StoreMemcachedBinary
}

// Create will instantiate the desired memory cached with binary connection
// store instance.
func (s storeStrategyMemcachedBinary) Create(args ...interface{}) (store IStore, err error) {
	if len(args) < 4 {
		return nil, errNilPointer("args[3]")
	}

	hosts, ok := args[0].(string)
	if !ok {
		return nil, errConversion(args[0], "string")
	}
	user, ok := args[1].(string)
	if !ok {
		return nil, errConversion(args[1], "string")
	}
	password, ok := args[2].(string)
	if !ok {
		return nil, errConversion(args[2], "string")
	}
	ttl, ok := args[3].(int)
	if !ok {
		return nil, errConversion(args[3], "int")
	}
	return newStoreMemcachedBinary(hosts, user, password, time.Duration(ttl)*time.Millisecond), nil
}

// CreateFromConfig will instantiate the desired memory cached with binary
// connection store instance where the initialization data comes from a
// configuration instance.
func (s storeStrategyMemcachedBinary) CreateFromConfig(cfg sconfig.IConfig) (store IStore, err error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	hosts, err := cfg.String("hosts")
	if err != nil {
		return nil, err
	}
	user, err := cfg.String("username")
	if err != nil {
		return nil, err
	}
	password, err := cfg.String("password")
	if err != nil {
		return nil, err
	}
	ttl, err := cfg.Int("ttl")
	if err != nil {
		return nil, err
	}
	return s.Create(hosts, user, password, ttl)
}
