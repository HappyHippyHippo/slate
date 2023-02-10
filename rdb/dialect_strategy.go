package rdb

import (
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

const (
	// UnknownDialectType defines the value to be used to
	// identify an unknown dialect.
	UnknownDialectType = "unknown"
)

// IDialectStrategy defines the interface to a gorm rdb
// dialect instantiation strategy, based on a configuration.
type IDialectStrategy interface {
	Accept(config.IConfig) bool
	Get(config.IConfig) (gorm.Dialector, error)
}
