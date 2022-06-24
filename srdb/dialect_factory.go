package srdb

import (
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/gorm"
	"strings"
)

// DialectFactory defines an object that will generate a database
// dialect interface based on a registered list of dialect
// generation strategies.
type DialectFactory []DialectStrategy

// Register will register a new dialect factory strategy to be used
// on requesting to create a dialect.
func (f *DialectFactory) Register(strategy DialectStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	*f = append(*f, strategy)

	return nil
}

// Get generates a new connection dialect interface defined by the
// configuration parameters stored in the configuration partial marked by
// the given name.
func (f DialectFactory) Get(config sconfig.Config) (gorm.Dialector, error) {
	if config == nil {
		return nil, errNilPointer("config")
	}

	dialect, err := config.String("dialect", "")
	if err != nil {
		return nil, err
	}

	name := strings.ToLower(dialect)
	for _, strategy := range f {
		if strategy.Accept(name) {
			return strategy.Get(config)
		}
	}

	return nil, errUnknownDatabaseDialect(name)
}
