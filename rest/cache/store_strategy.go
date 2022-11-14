package cache

import sconfig "github.com/happyhippyhippo/slate/config"

// IStoreStrategy interface defines the methods of the store
// factory strategy that can validate creation requests and instantiation
// of particular type of store.
type IStoreStrategy interface {
	GetName() string
	Accept(sourceType string) bool
	Create(arg ...interface{}) (IStore, error)
	CreateFromConfig(cfg sconfig.IConfig) (IStore, error)
}
