package xml

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/rest/logmw"
	"github.com/happyhippyhippo/slate/rest/logmw/response"
)

const (
	// ID defines the id to be used as the container
	// registration id of a logging middleware instance factory function.
	ID = response.ID + ".xml"
)

// Provider defines the slate.rest.log module service provider to be used on
// the application initialization to register the logging middleware service.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the log middleware package instances in the
// application container
func (Provider) Register(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(ID, func() (logmw.ResponseReader, error) {
		return NewDecorator(response.NewReader(), nil)
	})
	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (p Provider) Boot(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
