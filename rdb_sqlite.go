//go:build sqlite

package slate

import (
	"fmt"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	// RdbSqliteDialectCreatorContainerID defines the id to be used as the
	// Provider registration id of a sqlite dialect creator.
	RdbSqliteDialectCreatorContainerID = RdbDialectCreatorTag + ".sqlite"

	// RdbTypeSqlite defines the value to be used to identify a
	// Sqlite dialect.
	RdbTypeSqlite = "sqlite"
)

// ----------------------------------------------------------------------------
// rdb sqlite dialect creator
// ----------------------------------------------------------------------------

// RdbSqliteDialectCreator define a sqlite dialect creator service.
type RdbSqliteDialectCreator struct{}

var _ RdbDialectCreator = &RdbSqliteDialectCreator{}

// NewRdbSqliteDialectCreator will instantiate a new sqlite dialect creator.
func NewRdbSqliteDialectCreator() *RdbSqliteDialectCreator {
	return &RdbSqliteDialectCreator{}
}

// Accept check if the provided configuration should the handled as a sqlite
// connection definition,
func (RdbSqliteDialectCreator) Accept(
	config *ConfigPartial,
) bool {
	// check config argument reference
	if config == nil {
		return false
	}
	// retrieve the connection dialect from the configuration
	dc := struct{ Dialect string }{}
	if _, e := config.Populate("", &dc); e != nil {
		return false
	}
	// only accepts a sqlite dialect request
	return strings.EqualFold(strings.ToLower(dc.Dialect), RdbTypeSqlite)
}

// Create instantiates the requested sqlite connection dialect.
func (RdbSqliteDialectCreator) Create(
	config *ConfigPartial,
) (gorm.Dialector, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		Host   string
		Params ConfigPartial
	}{}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// add the extra parameters to the generated DSN string
	if len(sConfig.Params) > 0 {
		sConfig.Host += "?"
		for key, value := range sConfig.Params {
			sConfig.Host += fmt.Sprintf("&%s=%v", key, value)
		}
	}
	// create the connection dialect instance
	return sqlite.Open(sConfig.Host), nil
}

// ----------------------------------------------------------------------------
// rdb sqlite service register
// ----------------------------------------------------------------------------

// RdbSqliteServiceRegister defines the service provider to be used on
// the application initialization to register the sqlite relational
// database services.
type RdbSqliteServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &RdbSqliteServiceRegister{}

// NewRdbSqliteServiceRegister will generate a new sqlite services
// registry instance
func NewRdbSqliteServiceRegister(
	app ...*App,
) *RdbSqliteServiceRegister {
	return &RdbSqliteServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will register the sqlite relational database module services in the
// application Provider.
func (sr RdbSqliteServiceRegister) Provide(
	container *ServiceContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// register the services
	_ = container.Add(RdbSqliteDialectCreatorContainerID, NewRdbSqliteDialectCreator, RdbDialectCreatorTag)
	return nil
}
