package srdb

import (
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/gorm"
)

// IDialectStrategy defines the interface to a gorm srdb
// dialect instantiation strategy, based on a configuration.
type IDialectStrategy interface {
	Accept(dialect string) bool
	Get(config sconfig.IConfig) (gorm.Dialector, error)
}
