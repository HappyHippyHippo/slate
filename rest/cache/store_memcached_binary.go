package cache

import (
	"time"

	"github.com/memcachier/mc/v3"
)

type storeMemcachedBinary struct {
	*mc.Client
	defaultExpiration time.Duration
}

func newStoreMemcachedBinary(hostList, username, password string, defaultExpiration time.Duration) IStore {
	return &storeMemcachedBinary{mc.NewMC(hostList, username, password), defaultExpiration}
}

func newStoreMemcachedBinaryWithConfig(hostList, username, password string, defaultExpiration time.Duration, config *mc.Config) IStore {
	return &storeMemcachedBinary{mc.NewMCwithConfig(hostList, username, password, config), defaultExpiration}
}

// Set (see IStore interface)
func (s *storeMemcachedBinary) Set(key string, value interface{}, expires time.Duration) error {
	exp := s.getExpiration(expires)
	b, err := serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Set(key, string(b), 0, exp, 0)
	return convertMcError(key, err)
}

// Add (see IStore interface)
func (s *storeMemcachedBinary) Add(key string, value interface{}, expires time.Duration) error {
	exp := s.getExpiration(expires)
	b, err := serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Add(key, string(b), 0, exp)
	return convertMcError(key, err)
}

// Replace (see IStore interface)
func (s *storeMemcachedBinary) Replace(key string, value interface{}, expires time.Duration) error {
	exp := s.getExpiration(expires)
	b, err := serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Replace(key, string(b), 0, exp, 0)
	return convertMcError(key, err)
}

// Get (see IStore interface)
func (s *storeMemcachedBinary) Get(key string, value interface{}) error {
	val, _, _, err := s.Client.Get(key)
	if err != nil {
		return convertMcError(key, err)
	}
	return deserialize([]byte(val), value)
}

// Delete (see IStore interface)
func (s *storeMemcachedBinary) Delete(key string) error {
	return convertMcError(key, s.Client.Del(key))
}

// Increment (see IStore interface)
func (s *storeMemcachedBinary) Increment(key string, delta uint64) (uint64, error) {
	n, _, err := s.Client.Incr(key, delta, 0, 0xffffffff, 0)
	return n, convertMcError(key, err)
}

// Decrement (see IStore interface)
func (s *storeMemcachedBinary) Decrement(key string, delta uint64) (uint64, error) {
	n, _, err := s.Client.Decr(key, delta, 0, 0xffffffff, 0)
	return n, convertMcError(key, err)
}

// Flush (see IStore interface)
func (s *storeMemcachedBinary) Flush() error {
	return convertMcError("", s.Client.Flush(0))
}

// getExpiration converts a gin-contrib/cache expiration in the form of a
// time.Duration to a valid memcached expiration either in seconds (<30 days)
// or a Unix timestamp (>30 days)
func (s *storeMemcachedBinary) getExpiration(expires time.Duration) uint32 {
	switch expires {
	case time.Duration(0):
		expires = s.defaultExpiration
	case time.Duration(-1):
		expires = time.Duration(0)
	}
	exp := uint32(expires.Seconds())
	if exp > 60*60*24*30 { // > 30 days
		exp += uint32(time.Now().Unix())
	}
	return exp
}

func convertMcError(key string, err error) error {
	switch err {
	case nil:
		return nil
	case mc.ErrNotFound:
		return errCacheMiss(key)
	case mc.ErrValueNotStored:
		return errCacheNotStored(key)
	case mc.ErrKeyExists:
		return errCacheNotStored(key)
	}
	return err
}
