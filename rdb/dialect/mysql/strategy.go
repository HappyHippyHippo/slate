//go:build mysql

package mysql

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/rdb"
)

const (
	// Type defines the value to be used to identify a
	// MySQL dialect.
	Type = "mysql"
)

// DialectStrategy define a MySQL dialect generation strategy instance.
type DialectStrategy struct{}

var _ rdb.DialectStrategy = &DialectStrategy{}

// NewDialectStrategy will instantiate a new mysql dialect creation strategy.
func NewDialectStrategy() *DialectStrategy {
	return &DialectStrategy{}
}

// Accept check if the provided configuration should the handled as a mysql
// connection definition,
func (DialectStrategy) Accept(
	cfg config.Partial,
) bool {
	// check config argument reference
	if cfg == nil {
		return false
	}
	// retrieve the connection dialect from the configuration
	dc := struct{ Dialect string }{}
	if _, e := cfg.Populate("", &dc); e != nil {
		return false
	}
	// only accepts a mysql dialect request
	return strings.EqualFold(strings.ToLower(dc.Dialect), Type)
}

// Create instantiates the requested mysql connection dialect.
func (DialectStrategy) Create(
	cfg config.Partial,
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
		Params   config.Partial
	}{
		Protocol: "tcp",
		Port:     3306,
	}
	if _, e := cfg.Populate("", &dc); e != nil {
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
