package srdb

import (
	"fmt"
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
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
	if dsn, err := cfg.String("host"); err != nil {
		return nil, err
	} else if params, err := cfg.Partial("params", nil); err != nil {
		return nil, err
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
