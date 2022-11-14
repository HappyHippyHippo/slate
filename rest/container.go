package rest

import (
	"github.com/happyhippyhippo/slate"
)

// GetEngine will try to retrieve the registered gin engine
// instance from the application service container.
func GetEngine(c slate.ServiceContainer) (IEngine, error) {
	instance, err := c.Get(ContainerEngineID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(IEngine)
	if !ok {
		return nil, errConversion(instance, "IEngine")
	}
	return i, nil
}
