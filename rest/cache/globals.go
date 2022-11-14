package cache

import (
	senv "github.com/happyhippyhippo/slate/env"
	srest "github.com/happyhippyhippo/slate/rest"
)

const (
	// KeyGeneratorURI @todo doc
	KeyGeneratorURI = "uri"

	// StoreInMemory @todo doc
	StoreInMemory = "in_memory"

	// StoreMemcached @todo doc
	StoreMemcached = "memcached"

	// StoreMemcachedBinary @todo doc
	StoreMemcachedBinary = "memcached.binary"

	// StoreRedis @todo doc
	StoreRedis = "redis"
)

const (
	// ContainerID defines the id to be used as the container
	// registration id of a cache middleware generator instance, as a base id
	// of all other cache package instances registered in the application
	// container.
	ContainerID = srest.ContainerID + ".cache"

	// ContainerKeyGeneratorStrategyTag defines the tag to be assigned
	// to all container key generator strategies.
	ContainerKeyGeneratorStrategyTag = ContainerID + ".keygenerator.strategy"

	// ContainerKeyGeneratorStrategyURIID defines the tag to be assigned
	// to the URI key generator.
	ContainerKeyGeneratorStrategyURIID = ContainerID + ".keygenerator.strategy.uri"

	// ContainerKeyGeneratorFactoryID defines the tag to be assigned
	// to the key generator factory.
	ContainerKeyGeneratorFactoryID = ContainerID + ".keygenerator.factory"

	// ContainerStoreStrategyTag defines the tag to be assigned
	// to all container store strategies.
	ContainerStoreStrategyTag = ContainerID + ".store.strategy"

	// ContainerStoreStrategyInMemoryID defines the tag to be assigned
	// to the in-memory store.
	ContainerStoreStrategyInMemoryID = ContainerID + ".store.strategy.inmemory"

	// ContainerStoreStrategyMemcachedID defines the tag to be assigned
	// to the memcached store.
	ContainerStoreStrategyMemcachedID = ContainerID + ".store.strategy.memcached"

	// ContainerStoreStrategyMemcachedBinaryID defines the tag to be assigned
	// to the memcached binary connection store.
	ContainerStoreStrategyMemcachedBinaryID = ContainerID + ".store.strategy.memcached_binary"

	// ContainerStoreStrategyRedisID defines the tag to be assigned
	// to the redis store.
	ContainerStoreStrategyRedisID = ContainerID + ".store.strategy.redis"

	// ContainerStoreFactoryID defines the tag to be assigned
	// to the store factory.
	ContainerStoreFactoryID = ContainerID + ".store.factory"
)

const (
	// EnvID defines the slaterest.cache package base environment variable name.
	EnvID = srest.EnvID + "_CACHE"
)

var (
	// ConfigPathStores @todo doc
	ConfigPathStores = senv.String(EnvID+"_CONFIG_PATH_STORES", "slate.cache.stores")

	// ConfigPathEndpointCacheFormat defines the format of the configuration
	// path where the endpoint return status value can be retrieved.
	ConfigPathEndpointCacheFormat = senv.String(EnvID+"_CONFIG_PATH_ENDPOINT_CACHE_FORMAT", "slate.endpoints.%s.cache")

	// DefaultKeyGenerator @todo doc
	DefaultKeyGenerator = senv.String(EnvID+"_DEFAULT_KEY_GENERATOR", KeyGeneratorURI)

	// DefaultStore @todo doc
	DefaultStore = senv.String(EnvID+"_DEFAULT_STORE", StoreInMemory)

	// DefaultTTL @todo doc
	DefaultTTL = senv.Int(EnvID+"_DEFAULT_TTL", 300)
)
