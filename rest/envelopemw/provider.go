package envelopemw

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/rest"
)

const (
	// ID defines the default id used to register
	// the application envelope middleware and related services.
	ID = rest.ID + ".envelope"
)

// Provider defines the default envelope provider to be used on
// the application initialization to register the file system adapter service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will add to the container a new file system adapter instance.
func (p Provider) Register(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// register the envelope middleware generator
	_ = container[0].Service(ID, NewMiddlewareGenerator)
	return nil
}

// Boot (no-op).
func (Provider) Boot(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	return nil
}
