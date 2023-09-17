//go:build inmemory

package inmemory

import (
	"errors"
	"reflect"
	"time"

	gocache "github.com/robfig/go-cache"

	"github.com/happyhippyhippo/slate/cache"
	"github.com/happyhippyhippo/slate/cache/store"
)

type client interface {
	Set(key string, value interface{}, expire time.Duration)
	Add(key string, value interface{}, expire time.Duration) error
	Replace(key string, value interface{}, expire time.Duration) error
	Get(key string) (interface{}, bool)
	Increment(key string, delta uint64) (uint64, error)
	Decrement(key string, delta uint64) (uint64, error)
	Delete(key string) bool
	Flush()
}

var _ client = &gocache.Cache{}

// Store represents the cache with memory persistence
type Store struct {
	store.Store
	client client
}

var _ cache.Store = &Store{}

// NewStore returns a Store
func NewStore(
	Expiration time.Duration,
) *Store {
	// return the initialized in-memory store struct
	return &Store{
		Store:  store.Store{Expiration: Expiration},
		client: gocache.New(Expiration, time.Minute),
	}
}

// Get (see Store interface)
func (c *Store) Get(
	key string,
	value interface{},
) error {
	// retrieve the element from the store
	val, found := c.client.Get(key)
	if !found {
		return errMiss(key)
	}
	// try to store the value in the pointer argument
	v := reflect.ValueOf(value)
	if v.Type().Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		return nil
	}
	// signal error while storing the value
	return errNotStored(key)
}

// Set (see Store interface)
func (c *Store) Set(
	key string,
	value interface{},
	expire time.Duration,
) error {
	// store the value in the memory persistence layer
	c.client.Set(key, value, expire)
	return nil
}

// Add (see Store interface)
func (c *Store) Add(
	key string,
	value interface{},
	expire time.Duration,
) error {
	// add the value to the memory, and signal error storing if the
	// key already exists in the memory persistence layer
	err := c.client.Add(key, value, expire)
	if errors.Is(err, gocache.ErrKeyExists) {
		return errNotStored(key)
	}
	return err
}

// Replace (see Store interface)
func (c *Store) Replace(
	key string,
	value interface{},
	expire time.Duration,
) error {
	// try to replace an existing value in memory
	if err := c.client.Replace(key, value, expire); err != nil {
		return errNotStored(key)
	}
	return nil
}

// Delete (see Store interface)
func (c *Store) Delete(
	key string,
) error {
	// try to remove a value stored in memory marked with the requested key
	if found := c.client.Delete(key); !found {
		return errMiss(key)
	}
	return nil
}

// Increment (see Store interface)
func (c *Store) Increment(
	key string,
	n uint64,
) (uint64, error) {
	// try to increment a value stored in memory or signal a cache miss
	// if not present
	newValue, err := c.client.Increment(key, n)
	if errors.Is(err, gocache.ErrCacheMiss) {
		return 0, errMiss(key)
	}
	return newValue, err
}

// Decrement (see Store interface)
func (c *Store) Decrement(
	key string,
	n uint64,
) (uint64, error) {
	// try to decrement a value stored in memory or signal a cache miss
	// if not present
	newValue, err := c.client.Decrement(key, n)
	if errors.Is(err, gocache.ErrCacheMiss) {
		return 0, errMiss(key)
	}
	return newValue, err
}

// Flush (see Store interface)
func (c *Store) Flush() error {
	// flush the cache
	c.client.Flush()
	return nil
}
