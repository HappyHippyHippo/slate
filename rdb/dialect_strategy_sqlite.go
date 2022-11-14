package rdb

import (
	"fmt"
	"strings"

	sconfig "github.com/happyhippyhippo/slate/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type dialectStrategySqlite struct{}

var _ IDialectStrategy = &dialectStrategySqlite{}

// Accept check if the provided configuration should the handled as a mysql
// connection definition,
func (dialectStrategySqlite) Accept(name string) bool {
	return strings.EqualFold(strings.ToLower(name), DialectSqlite)
}

// Get instantiates the requested mysql connection dialect.
func (dialectStrategySqlite) Get(cfg sconfig.IConfig) (gorm.Dialector, error) {
	if dsn, e := cfg.String("host"); e != nil {
		return nil, e
	} else if params, e := cfg.Partial("params", nil); e != nil {
		return nil, e
	} else {
		if len(params) > 0 {
			dsn += "?"
			for key, value := range params {
				dsn += fmt.Sprintf("&%s=%v", key, value)
			}
		}

		return sqlite.Open(dsn), nil
	}
}
