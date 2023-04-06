//go:build mysql

package mysql

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/rdb"
)

const (
	// ID defines the id to be used as the container
	// registration id of a relational database connection pool instance,
	// and as a base id of all other relational database package instances
	// registered in the application container.
	ID = rdb.ID + ".mysql"

	// DialectStrategyID defines the id to be used
	// as the container registration id of the relational database connection
	// sqlite dialect instance.
	DialectStrategyID = ID + ".dialect.strategy"
)

// Provider defines the slate.rdb module service provider to be used on
// the application initialization to register the relational
// database services.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the rdb package instances in the
// application container
func (p Provider) Register(
	container slate.IContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(DialectStrategyID, NewMySQLDialectStrategy, rdb.DialectStrategyTag)
	return nil
}

// Boot will start the rdb package
func (p Provider) Boot(
	container slate.IContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
