package log

import (
	"github.com/happyhippyhippo/slate"
)

// GetMiddlewareGenerator will try to retrieve the registered logging
// middleware for ok responses instance from the application service container.
func GetMiddlewareGenerator(c slate.ServiceContainer) (MiddlewareGenerator, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(MiddlewareGenerator)
	if !ok {
		return nil, errConversion(instance, "log.MiddlewareGenerator")
	}
	return i, nil
}
