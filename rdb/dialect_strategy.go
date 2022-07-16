package rdb

import (
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// IDialectStrategy defines the interface to a gorm rdb
// dialect instantiation strategy, based on a configuration.
type IDialectStrategy interface {
	Accept(dialect string) bool
	Get(config config.IConfig) (gorm.Dialector, error)
}
