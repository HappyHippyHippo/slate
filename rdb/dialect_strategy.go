package rdb

import (
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

const (
	// DialectMySQL defines the value to be used to identify a
	// MySQL dialect.
	DialectMySQL = "mysql"

	// DialectSqlite defines the value to be used to identify a
	// Sqlite dialect.
	DialectSqlite = "sqlite"
)

// IDialectStrategy defines the interface to a gorm rdb
// dialect instantiation strategy, based on a configuration.
type IDialectStrategy interface {
	Accept(config.IConfig) bool
	Get(config.IConfig) (gorm.Dialector, error)
}
