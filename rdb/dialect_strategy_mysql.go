package rdb

import (
	"fmt"
	"strings"

	"github.com/happyhippyhippo/slate/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDialectStrategy define a MySQL dialect generation strategy instance.
type MySQLDialectStrategy struct{}

var _ IDialectStrategy = &MySQLDialectStrategy{}

// Accept check if the provided configuration should the handled as a mysql
// connection definition,
func (MySQLDialectStrategy) Accept(
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
	return strings.EqualFold(strings.ToLower(dc.Dialect), DialectMySQL)
}

// Get instantiates the requested mysql connection dialect.
func (MySQLDialectStrategy) Get(
	cfg config.IConfig,
) (gorm.Dialector, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	dc := struct {
		Username string
		Password string
		Protocol string
		Host     string
		Port     int
		Schema   string
		Params   config.Config
	}{Protocol: "tcp", Port: 3306}
	_, e := cfg.Populate("", &dc)
	if e != nil {
		return nil, e
	}
	// compose the connection DSN string
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s",
		dc.Username,
		dc.Password,
		dc.Protocol,
		dc.Host,
		dc.Port,
		dc.Schema,
	)
	// add the extra parameters to the generated DSN string
	if len(dc.Params) > 0 {
		dsn += "?"
		for key, value := range dc.Params {
			dsn += fmt.Sprintf("&%s=%v", key, value)
		}
	}
	// create the connection dialect instance
	return mysql.Open(dsn), nil
}
