package cache

import (
	"github.com/happyhippyhippo/slate"
)

const (
	// ID defines the id to be used as the container
	// registration id of a cache pool instance, as a base id of all other
	// cache package instances registered in the application container.
	ID = slate.ID + ".cache"

	// StoreStrategyTag defines the tag to be assigned to all
	// container Store strategies.
	StoreStrategyTag = ID + ".Store.strategy"

	// StoreFactoryID defines the id to be used as
	//	// the container registration id of a Store factory instance.
	StoreFactoryID = ID + ".Store.factory"
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
	// add store strategies and factory
	_ = container.Service(StoreFactoryID, NewStoreFactory)
	// add store pool instance
	_ = container.Service(ID, NewStorePool)
	return nil
}

// Boot will start the cache package.
func (p Provider) Boot(
	container *slate.Container,
) (e error) {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}

	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	// populate the container Store factory with
	// all registered Store strategies
	storeFactory := p.getStoreFactory(container)
	for _, strategy := range p.getStoreStrategies(container) {
		_ = storeFactory.Register(strategy)
	}
	return nil
}

func (Provider) getStoreFactory(
	container *slate.Container,
) *StoreFactory {
	// retrieve the factory entry
	entry, e := container.Get(StoreFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	instance, ok := entry.(*StoreFactory)
	if ok {
		return instance
	}
	panic(errConversion(entry, "*cache.StoreFactory"))
}

func (Provider) getStoreStrategies(
	container *slate.Container,
) []StoreStrategy {
	// retrieve the strategies entries
	entries, e := container.Tag(StoreStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var strategies []StoreStrategy
	for _, entry := range entries {
		if instance, ok := entry.(StoreStrategy); ok {
			strategies = append(strategies, instance)
		}
	}
	return strategies
}
