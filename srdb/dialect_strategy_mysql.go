package srdb

import (
	"fmt"
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type dialectStrategyMySQL struct{}

var _ IDialectStrategy = &dialectStrategyMySQL{}

// Accept check if the provided configuration should the handled as a mysql
// connection definition,
func (dialectStrategyMySQL) Accept(dialect string) bool {
	return strings.EqualFold(strings.ToLower(dialect), DialectMySQL)
}

// Get instantiates the requested mysql connection dialect.
func (dialectStrategyMySQL) Get(cfg sconfig.IConfig) (gorm.Dialector, error) {
	if username, e := cfg.String("username"); e != nil {
		return nil, e
	} else if password, e := cfg.String("password"); e != nil {
		return nil, e
	} else if protocol, e := cfg.String("protocol", "tcp"); e != nil {
		return nil, e
	} else if host, e := cfg.String("host"); e != nil {
		return nil, e
	} else if port, e := cfg.Int("port", 3306); e != nil {
		return nil, e
	} else if schema, e := cfg.String("schema"); e != nil {
		return nil, e
	} else {
		dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, protocol, host, port, schema)

		params, e := cfg.Partial("params", sconfig.Partial{})
		if e != nil {
			return nil, e
		}

		if len(params) > 0 {
			dsn += "?"
			for key, value := range params {
				dsn += fmt.Sprintf("&%s=%v", key, value)
			}
		}

		return mysql.Open(dsn), nil
	}
}
