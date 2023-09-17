//go:build inmemory

package inmemory

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/cache"
	"github.com/happyhippyhippo/slate/cache/store"
)

const (
	// ID defines the id to be used as the container
	// registration id of a cache pool instance, as a base id of all other
	// cache package instances registered in the application container.
	ID = store.ID + ".in_memory"
)

// Provider defines the slate.cache module service provider to be used on
// the application initialization to register the caching services.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the cache package instances in the
// application container.
func (p Provider) Register(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(ID, NewStoreStrategy, cache.StoreStrategyTag)
	return nil
}

// Boot will start the cache package.
func (p Provider) Boot(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
