//go:build sqlite

package sqlite

import (
	"fmt"
	"strings"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/rdb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	// Type defines the value to be used to identify a
	// Sqlite dialect.
	Type = "sqlite"
)

type dialectConfig struct {
	Host   string
	Params config.Partial
}

// DialectStrategy define a Sqlite dialect generation strategy instance.
type DialectStrategy struct{}

var _ rdb.DialectStrategy = &DialectStrategy{}

// NewDialectStrategy will instantiate a new sqlite dialect creation strategy.
func NewDialectStrategy() *DialectStrategy {
	return &DialectStrategy{}
}

// Accept check if the provided configuration should the handled as a mysql
// connection definition,
func (DialectStrategy) Accept(
	cfg *config.Partial,
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
	return strings.EqualFold(strings.ToLower(dc.Dialect), Type)
}

// Create instantiates the requested mysql connection dialect.
func (DialectStrategy) Create(
	cfg *config.Partial,
) (gorm.Dialector, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// retrieve the data from the configuration
	dc := dialectConfig{}
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
