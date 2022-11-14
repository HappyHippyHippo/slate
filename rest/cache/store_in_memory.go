package cache

import (
	"reflect"
	"time"

	"github.com/robfig/go-cache"
)

type storeInMemory struct {
	cache.Cache
}

func newStoreInMemory(defaultExpiration time.Duration) IStore {
	return &storeInMemory{*cache.New(defaultExpiration, time.Minute)}
}

// Get (see IStore interface)
func (c *storeInMemory) Get(key string, value interface{}) error {
	val, found := c.Cache.Get(key)
	if !found {
		return errCacheMiss(key)
	}

	v := reflect.ValueOf(value)
	if v.Type().Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		return nil
	}

	return errCacheNotStored(key)
}

// Set (see IStore interface)
func (c *storeInMemory) Set(key string, value interface{}, expires time.Duration) error {
	// NOTE: go-cache understands the values of DEFAULT and FOREVER
	c.Cache.Set(key, value, expires)
	return nil
}

// Add (see IStore interface)
func (c *storeInMemory) Add(key string, value interface{}, expires time.Duration) error {
	err := c.Cache.Add(key, value, expires)
	if err == cache.ErrKeyExists {
		return errCacheNotStored(key)
	}
	return err
}

// Replace (see IStore interface)
func (c *storeInMemory) Replace(key string, value interface{}, expires time.Duration) error {
	if err := c.Cache.Replace(key, value, expires); err != nil {
		return errCacheNotStored(key)
	}
	return nil
}

// Delete (see IStore interface)
func (c *storeInMemory) Delete(key string) error {
	if found := c.Cache.Delete(key); !found {
		return errCacheMiss(key)
	}
	return nil
}

// Increment (see IStore interface)
func (c *storeInMemory) Increment(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Increment(key, n)
	if err == cache.ErrCacheMiss {
		return 0, errCacheMiss(key)
	}
	return newValue, err
}

// Decrement (see IStore interface)
func (c *storeInMemory) Decrement(key string, n uint64) (uint64, error) {
	newValue, err := c.Cache.Decrement(key, n)
	if err == cache.ErrCacheMiss {
		return 0, errCacheMiss(key)
	}
	return newValue, err
}

// Flush (see IStore interface)
func (c *storeInMemory) Flush() error {
	c.Cache.Flush()
	return nil
}
