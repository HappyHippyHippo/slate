//go:build postgres

package slate

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	// RdbPostgresDialectCreatorContainerID defines the id to be used as the
	// Provider registration id of a Postgres dialect creator.
	RdbPostgresDialectCreatorContainerID = RdbDialectCreatorTag + ".postgres"

	// RdbTypePostgres defines the value to be used to identify a
	// Postgres dialect.
	RdbTypePostgres = "postgres"
)

// ----------------------------------------------------------------------------
// rdb postgres dialect creator
// ----------------------------------------------------------------------------

// RdbPostgresDialectCreator define a postgres dialect creator service.
type RdbPostgresDialectCreator struct{}

var _ RdbDialectCreator = &RdbPostgresDialectCreator{}

// NewRdbPostgresDialectCreator will instantiate a new postgres dialect creator.
func NewRdbPostgresDialectCreator() *RdbPostgresDialectCreator {
	return &RdbPostgresDialectCreator{}
}

// Accept check if the provided configuration should the handled as a Postgres
// connection definition,
func (RdbPostgresDialectCreator) Accept(
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
	return strings.EqualFold(strings.ToLower(dc.Dialect), RdbTypePostgres)
}

// Create instantiates the requested sqlite connection dialect.
func (RdbPostgresDialectCreator) Create(
	config *ConfigPartial,
) (gorm.Dialector, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// retrieve the data from the configuration
	sConfig := struct {
		Username string
		Password string
		Host     string
		Port     int
		Schema   string
		Params   ConfigPartial
	}{
		Port: 5432,
	}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// compose the connection DSN string
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		sConfig.Username,
		sConfig.Password,
		sConfig.Host,
		sConfig.Port,
		sConfig.Schema,
	)
	// add the extra parameters to the generated DSN string
	if len(sConfig.Params) > 0 {
		dsn += " "
		for key, value := range sConfig.Params {
			dsn += fmt.Sprintf(" %s=%v", key, value)
		}
	}
	// create the connection dialect instance
	return postgres.Open(dsn), nil
}

// ----------------------------------------------------------------------------
// rdb sqlite service register
// ----------------------------------------------------------------------------

// RdbPostgresServiceRegister defines the service provider to be used on
// the application initialization to register the sqlite relational
// database services.
type RdbPostgresServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &RdbPostgresServiceRegister{}

// NewRdbPostgresServiceRegister will generate a new sqlite services
// registry instance
func NewRdbPostgresServiceRegister(
	app ...*App,
) *RdbPostgresServiceRegister {
	return &RdbPostgresServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will register the sqlite relational database module services in the
// application Provider.
func (sr RdbPostgresServiceRegister) Provide(
	container *ServiceContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// register the services
	_ = container.Add(RdbPostgresDialectCreatorContainerID, NewRdbPostgresDialectCreator, RdbDialectCreatorTag)
	return nil
}
