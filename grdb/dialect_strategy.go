package grdb

import (
	"github.com/happyhippyhippo/slate/gconfig"
	"gorm.io/gorm"
)

// DialectStrategy defines the interface to a gorm rdb
// dialect instantiation strategy, based on a configuration.
type DialectStrategy interface {
	Accept(dialect string) bool
	Get(config gconfig.Config) (gorm.Dialector, error)
}
