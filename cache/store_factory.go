package cache

import (
	"github.com/happyhippyhippo/slate/config"
)

// StoreFactory is a persistence Store generator based on a
// registered list of Store generation strategies.
type StoreFactory []StoreStrategy

// NewStoreFactory @todo doc
func NewStoreFactory() *StoreFactory {
	return &StoreFactory{}
}

// Register will register a new Store factory strategy to be used
// on creation requests.
func (f *StoreFactory) Register(
	strategy StoreStrategy,
) error {
	// check the strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// add the strategy to the Store factory strategy pool
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate and return a new Store loaded
// by a configuration instance.
func (f *StoreFactory) Create(
	cfg config.Partial,
) (Store, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// search in the factory strategy pool for one that would accept
	// to generate the requested Store with the requested type defined
	// in the given config
	for _, s := range *f {
		if s.Accept(cfg) {
			// return the creation of the requested Store
			return s.Create(cfg)
		}
	}
	return nil, errInvalidStore(cfg)
}
