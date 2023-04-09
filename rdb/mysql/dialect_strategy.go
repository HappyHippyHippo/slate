//go:build mysql

package mysql

import (
	"fmt"
	"strings"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/rdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// DialectType defines the value to be used to identify a
	// MySQL dialect.
	DialectType = "mysql"
)

type dialectConfig struct {
	Username string
	Password string
	Protocol string
	Host     string
	Port     int
	Schema   string
	Params   config.Config
}

// DialectStrategy define a MySQL dialect generation strategy instance.
type DialectStrategy struct{}

var _ rdb.IDialectStrategy = &DialectStrategy{}

// NewDialectStrategy @todo doc
func NewDialectStrategy() *DialectStrategy {
	return &DialectStrategy{}
}

// Accept check if the provided configuration should the handled as a mysql
// connection definition,
func (DialectStrategy) Accept(
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
	return strings.EqualFold(strings.ToLower(dc.Dialect), DialectType)
}

// Get instantiates the requested mysql connection dialect.
func (DialectStrategy) Get(
	cfg config.IConfig,
) (gorm.Dialector, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	dc := dialectConfig{Protocol: "tcp", Port: 3306}
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
