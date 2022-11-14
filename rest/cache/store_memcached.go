package cache

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type storeMemcached struct {
	*memcache.Client
	defaultExpiration time.Duration
}

func newStoreMemcached(hostList []string, defaultExpiration time.Duration) IStore {
	return &storeMemcached{memcache.New(hostList...), defaultExpiration}
}

// Set (see IStore interface)
func (c *storeMemcached) Set(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Set, key, value, expires)
}

// Add (see IStore interface)
func (c *storeMemcached) Add(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Add, key, value, expires)
}

// Replace (see IStore interface)
func (c *storeMemcached) Replace(key string, value interface{}, expires time.Duration) error {
	return c.invoke((*memcache.Client).Replace, key, value, expires)
}

// Get (see IStore interface)
func (c *storeMemcached) Get(key string, value interface{}) error {
	item, err := c.Client.Get(key)
	if err != nil {
		return convertMemcacheError(key, err)
	}
	return deserialize(item.Value, value)
}

// Delete (see IStore interface)
func (c *storeMemcached) Delete(key string) error {
	return convertMemcacheError(key, c.Client.Delete(key))
}

// Increment (see IStore interface)
func (c *storeMemcached) Increment(key string, delta uint64) (uint64, error) {
	newValue, err := c.Client.Increment(key, delta)
	return newValue, convertMemcacheError(key, err)
}

// Decrement (see IStore interface)
func (c *storeMemcached) Decrement(key string, delta uint64) (uint64, error) {
	newValue, err := c.Client.Decrement(key, delta)
	return newValue, convertMemcacheError(key, err)
}

// Flush (see IStore interface)
func (c *storeMemcached) Flush() error {
	return errCacheOpNotSupport("Flush")
}

func (c *storeMemcached) invoke(storeFn func(*memcache.Client, *memcache.Item) error,
	key string, value interface{}, expire time.Duration) error {

	switch expire {
	case time.Duration(0):
		expire = c.defaultExpiration
	case time.Duration(-1):
		expire = time.Duration(0)
	}

	b, err := serialize(value)
	if err != nil {
		return err
	}
	return convertMemcacheError(key, storeFn(c.Client, &memcache.Item{
		Key:        key,
		Value:      b,
		Expiration: int32(expire / time.Second),
	}))
}

func convertMemcacheError(key string, err error) error {
	switch err {
	case nil:
		return nil
	case memcache.ErrCacheMiss:
		return errCacheMiss(key)
	case memcache.ErrNotStored:
		return errCacheNotStored(key)
	}

	return err
}
