package rdb

import (
	"fmt"
	"strings"

	"github.com/happyhippyhippo/slate/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SqliteDialectStrategy define a Sqlite dialect generation strategy instance.
type SqliteDialectStrategy struct{}

var _ IDialectStrategy = &SqliteDialectStrategy{}

type sqliteDialectConfig struct {
	Host   string
	Params config.Config
}

// Accept check if the provided configuration should the handled as a mysql
// connection definition,
func (SqliteDialectStrategy) Accept(
	cfg config.IConfig,
) bool {
	// check config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the connection dialect from the configuration
	dc := struct{ Dialect string }{}
	_, e := cfg.Populate("", &dc)
	if e != nil {
		return false
	}
	// only accepts a mysql dialect request
	return strings.EqualFold(strings.ToLower(dc.Dialect), DialectSqlite)
}

// Get instantiates the requested mysql connection dialect.
func (SqliteDialectStrategy) Get(
	cfg config.IConfig,
) (gorm.Dialector, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	dc := sqliteDialectConfig{}
	_, e := cfg.Populate("", &dc)
	if e != nil {
		return nil, e
	}
	// add the extra parameters to the generated DSN string
	if len(dc.Params) > 0 {
		dc.Host += "?"
		for key, value := range dc.Params {
			dc.Host += fmt.Sprintf("&%s=%v", key, value)
		}
	}
	// create the connection dialect instance
	return sqlite.Open(dc.Host), nil
}
