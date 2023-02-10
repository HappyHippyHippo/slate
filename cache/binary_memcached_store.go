package cache

import (
	"time"

	"github.com/memcachier/mc/v3"
)

// BinaryMemcachedStore represents the cache with memcached persistence using
// the binary protocol
type BinaryMemcachedStore struct {
	*mc.Client
	defaultExpiration time.Duration
}

// NewBinaryMemcachedStore returns a BinaryMemcachedStore
func NewBinaryMemcachedStore(
	hostList,
	username,
	password string,
	defaultExpiration time.Duration,
) *BinaryMemcachedStore {
	// return the initialized memcached store struct
	return &BinaryMemcachedStore{
		Client:            mc.NewMC(hostList, username, password),
		defaultExpiration: defaultExpiration,
	}
}

// NewBinaryMemcachedStoreWithConfig returns a BinaryMemcachedStore
// using the provided configuration
func NewBinaryMemcachedStoreWithConfig(
	hostList,
	username,
	password string,
	defaultExpiration time.Duration,
	config *mc.Config,
) *BinaryMemcachedStore {
	// return the initialized memcached store struct
	return &BinaryMemcachedStore{
		Client:            mc.NewMCwithConfig(hostList, username, password, config),
		defaultExpiration: defaultExpiration,
	}
}

// Set (see Store interface)
func (s *BinaryMemcachedStore) Set(
	key string,
	value interface{},
	expires time.Duration,
) error {
	exp := s.getExpiration(expires)
	b, err := serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Set(key, string(b), 0, exp, 0)
	return convertBinaryMemcachedError(key, err)
}

// Add (see Store interface)
func (s *BinaryMemcachedStore) Add(
	key string,
	value interface{},
	expires time.Duration,
) error {
	exp := s.getExpiration(expires)
	b, err := serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Add(key, string(b), 0, exp)
	return convertBinaryMemcachedError(key, err)
}

// Replace (see Store interface)
func (s *BinaryMemcachedStore) Replace(
	key string,
	value interface{},
	expires time.Duration,
) error {
	exp := s.getExpiration(expires)
	b, err := serialize(value)
	if err != nil {
		return err
	}
	_, err = s.Client.Replace(key, string(b), 0, exp, 0)
	return convertBinaryMemcachedError(key, err)
}

// Get (see Store interface)
func (s *BinaryMemcachedStore) Get(
	key string,
	value interface{},
) error {
	val, _, _, err := s.Client.Get(key)
	if err != nil {
		return convertBinaryMemcachedError(key, err)
	}
	return deserialize([]byte(val), value)
}

// Delete (see Store interface)
func (s *BinaryMemcachedStore) Delete(
	key string,
) error {
	return convertBinaryMemcachedError(key, s.Client.Del(key))
}

// Increment (see Store interface)
func (s *BinaryMemcachedStore) Increment(
	key string,
	delta uint64,
) (uint64, error) {
	n, _, err := s.Client.Incr(key, delta, 0, 0xffffffff, 0)
	return n, convertBinaryMemcachedError(key, err)
}

// Decrement (see Store interface)
func (s *BinaryMemcachedStore) Decrement(
	key string,
	delta uint64,
) (uint64, error) {
	n, _, err := s.Client.Decr(key, delta, 0, 0xffffffff, 0)
	return n, convertBinaryMemcachedError(key, err)
}

// Flush (see Store interface)
func (s *BinaryMemcachedStore) Flush() error {
	return convertBinaryMemcachedError("", s.Client.Flush(0))
}

// getExpiration converts a cache expiration in the form of a
// time.Duration to a valid memcached expiration either in seconds (<30 days)
// or a Unix timestamp (>30 days)
func (s *BinaryMemcachedStore) getExpiration(
	expires time.Duration,
) uint32 {
	// check for statically defined expiration values
	switch expires {
	case DEFAULT:
		expires = s.defaultExpiration
	case FOREVER:
		expires = time.Duration(0)
	}
	// convert to a memcached uint expiration value
	exp := uint32(expires.Seconds())
	if exp > 60*60*24*30 { // > 30 days
		exp += uint32(time.Now().Unix())
	}
	return exp
}

func convertBinaryMemcachedError(
	key string,
	err error,
) error {
	switch err {
	case nil:
		return nil
	case mc.ErrNotFound:
		return errMiss(key)
	case mc.ErrValueNotStored:
		return errNotStored(key)
	case mc.ErrKeyExists:
		return errNotStored(key)
	}
	return err
}
