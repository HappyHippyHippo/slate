//go:build mysql

package slate

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// RdbMySqlDialectCreatorContainerID defines the id to be used as the
	// Provider registration id of a MySql dialect creator.
	RdbMySqlDialectCreatorContainerID = RdbDialectCreatorTag + ".mysql"

	// RdbTypeMySql defines the value to be used to identify a
	// MySql dialect.
	RdbTypeMySql = "mysql"
)

// ----------------------------------------------------------------------------
// rdb mysql dialect creator
// ----------------------------------------------------------------------------

// RdbMySqlDialectCreator define a mysql dialect creator service.
type RdbMySqlDialectCreator struct{}

var _ RdbDialectCreator = &RdbMySqlDialectCreator{}

// NewRdbMySqlDialectCreator will instantiate a new mysql dialect creator.
func NewRdbMySqlDialectCreator() *RdbMySqlDialectCreator {
	return &RdbMySqlDialectCreator{}
}

// Accept check if the provided configuration should the handled as a MySql
// connection definition,
func (RdbMySqlDialectCreator) Accept(
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
	return strings.EqualFold(strings.ToLower(dc.Dialect), RdbTypeMySql)
}

// Create instantiates the requested sqlite connection dialect.
func (RdbMySqlDialectCreator) Create(
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
		Protocol string
		Host     string
		Port     int
		Schema   string
		Params   ConfigPartial
	}{
		Protocol: "tcp",
		Port:     3306,
	}
	if _, e := config.Populate("", &sConfig); e != nil {
		return nil, e
	}
	// compose the connection DSN string
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s",
		sConfig.Username,
		sConfig.Password,
		sConfig.Protocol,
		sConfig.Host,
		sConfig.Port,
		sConfig.Schema,
	)
	// add the extra parameters to the generated DSN string
	if len(sConfig.Params) > 0 {
		dsn += "?"
		for key, value := range sConfig.Params {
			dsn += fmt.Sprintf("&%s=%v", key, value)
		}
	}
	// create the connection dialect instance
	return mysql.Open(dsn), nil
}

// ----------------------------------------------------------------------------
// rdb sqlite service register
// ----------------------------------------------------------------------------

// RdbMySqlServiceRegister defines the service provider to be used on
// the application initialization to register the sqlite relational
// database services.
type RdbMySqlServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &RdbMySqlServiceRegister{}

// NewRdbMySqlServiceRegister will generate a new sqlite services
// registry instance
func NewRdbMySqlServiceRegister(
	app ...*App,
) *RdbMySqlServiceRegister {
	return &RdbMySqlServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will register the sqlite relational database module services in the
// application Provider.
func (sr RdbMySqlServiceRegister) Provide(
	container *ServiceContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("Provider")
	}
	// register the services
	_ = container.Add(RdbMySqlDialectCreatorContainerID, NewRdbMySqlDialectCreator, RdbDialectCreatorTag)
	return nil
}
