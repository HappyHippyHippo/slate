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

var _ slate.Provider = &Provider{}

// Register will add to the container a new file system adapter instance.
func (p Provider) Register(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	_ = container.Service(ID, NewMiddlewareGenerator)
	return nil
}

// Boot (no-op).
func (Provider) Boot(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	return nil
}
