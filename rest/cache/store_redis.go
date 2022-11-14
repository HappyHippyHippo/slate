package cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type storeRedis struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
}

func newStoreRedis(host, password string, defaultExpiration time.Duration) IStore {
	var pool = &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			// the redis protocol should probably be made sett-able
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			} else {
				// check with PING
				if _, err := c.Do("PING"); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// custom connection test method
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
	return &storeRedis{pool, defaultExpiration}
}

func newStoreRedisWithPool(pool *redis.Pool, defaultExpiration time.Duration) *storeRedis {
	return &storeRedis{pool, defaultExpiration}
}

// Set (see IStore interface)
func (c *storeRedis) Set(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	return c.invoke(conn.Do, key, value, expires)
}

// Add (see IStore interface)
func (c *storeRedis) Add(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	if exists(conn, key) {
		return errCacheNotStored(key)
	}
	return c.invoke(conn.Do, key, value, expires)
}

// Replace (see IStore interface)
func (c *storeRedis) Replace(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		return errCacheNotStored(key)
	}
	err := c.invoke(conn.Do, key, value, expires)
	if value == nil {
		return errCacheNotStored(key)
	}

	return err

}

// Get (see IStore interface)
func (c *storeRedis) Get(key string, ptrValue interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()
	raw, err := conn.Do("GET", key)
	if raw == nil {
		return errCacheMiss(key)
	}
	item, err := redis.Bytes(raw, err)
	if err != nil {
		return err
	}
	return deserialize(item, ptrValue)
}

func exists(conn redis.Conn, key string) bool {
	retval, _ := redis.Bool(conn.Do("EXISTS", key))
	return retval
}

// Delete (see IStore interface)
func (c *storeRedis) Delete(key string) error {
	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		return errCacheMiss(key)
	}
	_, err := conn.Do("DEL", key)
	return err
}

// Increment (see IStore interface)
func (c *storeRedis) Increment(key string, delta uint64) (uint64, error) {
	conn := c.pool.Get()
	defer conn.Close()
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
	val, err := conn.Do("GET", key)
	if val == nil {
		return 0, errCacheMiss(key)
	}
	if err == nil {
		currentVal, err := redis.Int64(val, nil)
		if err != nil {
			return 0, err
		}
		sum := currentVal + int64(delta)
		_, err = conn.Do("SET", key, sum)
		if err != nil {
			return 0, err
		}
		return uint64(sum), nil
	}

	return 0, err
}

// Decrement (see IStore interface)
func (c *storeRedis) Decrement(key string, delta uint64) (newValue uint64, err error) {
	conn := c.pool.Get()
	defer conn.Close()
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that, hence the exists call
	if !exists(conn, key) {
		return 0, errCacheMiss(key)
	}
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
	currentVal, err := redis.Int64(conn.Do("GET", key))
	if err == nil && delta > uint64(currentVal) {
		tempint, err := redis.Int64(conn.Do("DECRBY", key, currentVal))
		return uint64(tempint), err
	}
	tempint, err := redis.Int64(conn.Do("DECRBY", key, delta))
	return uint64(tempint), err
}

// Flush (see IStore interface)
func (c *storeRedis) Flush() error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("FLUSHALL")
	return err
}

func (c *storeRedis) invoke(f func(string, ...interface{}) (interface{}, error),
	key string, value interface{}, expires time.Duration) error {

	switch expires {
	case time.Duration(0):
		expires = c.defaultExpiration
	case time.Duration(-1):
		expires = time.Duration(0)
	}

	b, err := serialize(value)
	if err != nil {
		return err
	}

	if expires > 0 {
		_, err := f("SETEX", key, int32(expires/time.Second), b)
		return err
	}

	_, err = f("SET", key, b)
	return err

}
