package rdb

import (
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

const (
	// UnknownDialect defines the value to be used to
	// identify an unknown dialect.
	UnknownDialect = "unknown"
)

// DialectStrategy defines the interface to a gorm rdb
// dialect instantiation strategy, based on a configuration.
type DialectStrategy interface {
	Accept(*config.Partial) bool
	Create(*config.Partial) (gorm.Dialector, error)
}
