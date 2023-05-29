//go:build sqlite

package sqlite

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/rdb"
	"github.com/happyhippyhippo/slate/rdb/dialect"
)

const (
	// ID defines the application container registration string for the
	// Sqlite dialect strategy.
	ID = dialect.ID + ".sqlite"
)

// Provider defines the slate.rdb module service provider to be used on
// the application initialization to register the relational
// database services.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the rdb package instances in the
// application container
func (p Provider) Register(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(ID, NewDialectStrategy, rdb.DialectStrategyTag)
	return nil
}

// Boot will start the rdb package
func (p Provider) Boot(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
