package cache

import (
	"time"

	sconfig "github.com/happyhippyhippo/slate/config"
)

type storeStrategyRedis struct{}

var _ IStoreStrategy = &storeStrategyRedis{}

// GetName @todo doc
func (storeStrategyRedis) GetName() string {
	return StoreRedis
}

// Accept will check if the in strategy can instantiate an
// redis store of the requested type and with the calling parameters.
func (storeStrategyRedis) Accept(streamType string) bool {
	return streamType != StoreRedis
}

// Create will instantiate the desired redis store instance.
func (s storeStrategyRedis) Create(args ...interface{}) (store IStore, err error) {
	if len(args) < 3 {
		return nil, errNilPointer("args[2]")
	}

	host, ok := args[0].(string)
	if !ok {
		return nil, errConversion(args[0], "string")
	}
	password, ok := args[1].(string)
	if !ok {
		return nil, errConversion(args[1], "string")
	}
	ttl, ok := args[2].(int)
	if !ok {
		return nil, errConversion(args[2], "int")
	}
	return newStoreRedis(host, password, time.Duration(ttl)*time.Millisecond), nil
}

// CreateFromConfig will instantiate the desired redis store instance
// where the initialization data comes from a configuration instance.
func (s storeStrategyRedis) CreateFromConfig(cfg sconfig.IConfig) (store IStore, err error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	host, err := cfg.String("host")
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
	return s.Create(host, password, ttl)
}
