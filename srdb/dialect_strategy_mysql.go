package srdb

import (
	"fmt"
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

// DialectStrategyMySQL is a MySQL dialect interface generator.
type DialectStrategyMySQL struct{}

var _ DialectStrategy = &DialectStrategyMySQL{}

// Accept check if the provided configuration should the handled as a mysql
// connection definition,
func (DialectStrategyMySQL) Accept(dialect string) bool {
	return strings.EqualFold(strings.ToLower(dialect), DialectMySQL)
}

// Get instantiates the requested mysql connection dialect.
func (DialectStrategyMySQL) Get(config sconfig.Config) (gorm.Dialector, error) {
	if username, err := config.String("username"); err != nil {
		return nil, err
	} else if password, err := config.String("password"); err != nil {
		return nil, err
	} else if protocol, err := config.String("protocol", "tcp"); err != nil {
		return nil, err
	} else if host, err := config.String("host"); err != nil {
		return nil, err
	} else if port, err := config.Int("port", 3306); err != nil {
		return nil, err
	} else if schema, err := config.String("schema"); err != nil {
		return nil, err
	} else {
		dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", username, password, protocol, host, port, schema)

		params, err := config.Partial("params", sconfig.Partial{})
		if err != nil {
			return nil, err
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
