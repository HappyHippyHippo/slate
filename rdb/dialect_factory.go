package rdb

import (
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// IDialectFactory defines the interface of a connection dialect instance.
type IDialectFactory interface {
	Register(strategy IDialectStrategy) error
	Get(config.IConfig) (gorm.Dialector, error)
}

// DialectFactory defines an object that will generate a database
// dialect interface based on a registered list of dialect
// generation strategies.
type DialectFactory []IDialectStrategy

var _ IDialectFactory = &DialectFactory{}

// NewDialectFactory will instantiate a new relational database
// dialect factory instance.
func NewDialectFactory() IDialectFactory {
	return &DialectFactory{}
}

// Register will register a new dialect factory strategy to be used
// on requesting to create a dialect.
func (f *DialectFactory) Register(
	strategy IDialectStrategy,
) error {
	// check strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// store the strategy in the factory strategy pool
	*f = append(*f, strategy)
	return nil
}

// Get generates a new connection dialect interface defined by the
// configuration parameters stored in the configuration partial marked by
// the given name.
func (f DialectFactory) Get(
	cfg config.IConfig,
) (gorm.Dialector, error) {
	// check the config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// search for a strategy that can create a dialect instance for the
	// dialect name retrieved from the configuration
	for _, strategy := range f {
		if strategy.Accept(cfg) {
			// generate the dialect instance
			return strategy.Get(cfg)
		}
	}
	return nil, errUnknownDialect(cfg)
}
