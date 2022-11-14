package cache

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	sconfig "github.com/happyhippyhippo/slate/config"
	srest "github.com/happyhippyhippo/slate/rest"
	srerror "github.com/happyhippyhippo/slate/rest/error"
)

// NewMiddlewareGenerator @todo doc
func NewMiddlewareGenerator(cfg sconfig.IManager, keygenFactory *KeyGeneratorFactory, storeFactory IStoreFactory) (func(id string) (srest.Middleware, error), error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	if keygenFactory == nil {
		return nil, errNilPointer("keygenFactory")
	}
	if storeFactory == nil {
		return nil, errNilPointer("storeFactory")
	}

	return func(id string) (srest.Middleware, error) {
		cfg, e := cfg.Partial(fmt.Sprintf(ConfigPathEndpointCacheFormat, id))
		if e != nil {
			return nil, e
		}

		active, err := cfg.Bool("active", false)
		if err != nil {
			return nil, err
		}

		if !active {
			return func(next gin.HandlerFunc) gin.HandlerFunc {
				return func(c *gin.Context) {
					next(c)
				}
			}, nil
		}

		keygenCfg, err := cfg.Partial("keygen", sconfig.Partial{"type": DefaultKeyGenerator})
		if err != nil {
			return nil, err
		}

		keygen, err := keygenFactory.CreateFromConfig(&keygenCfg)
		if err != nil {
			return nil, err
		}

		storeType, _ := cfg.String("store", DefaultStore)
		store, err := storeFactory.Create(storeType)
		if err != nil {
			return nil, err
		}

		ttl, err := cfg.Int("ttl", DefaultTTL)
		if err != nil {
			return nil, err
		}
		expire := time.Duration(ttl) * time.Second

		return func(next gin.HandlerFunc) gin.HandlerFunc {
			return func(c *gin.Context) {
				var cache responseCache

				key := keygen(c.Request)

				if err := store.Get(key, &cache); err != nil {
					if !errors.Is(err, srerror.ErrCacheMiss) {
						log.Println(err.Error())
					}
					// replace writer
					writer := newCachedWriter(store, expire, c.Writer, key)
					c.Writer = writer

					next(c)

					// Drop caches of aborted contexts
					if c.IsAborted() {
						_ = store.Delete(key)
					}
				} else {
					c.Writer.WriteHeader(cache.Status)
					for k, vals := range cache.Header {
						for _, v := range vals {
							c.Writer.Header().Set(k, v)
						}
					}
					_, _ = c.Writer.Write(cache.Data)
				}
			}
		}, nil
	}, nil
}
