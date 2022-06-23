package srest

import (
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate"
)

// GetEngine will try to retrieve the registered gin engine
// instance from the application service container.
func GetEngine(c slate.ServiceContainer) (*gin.Engine, error) {
	instance, err := c.Get(ContainerEngineID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*gin.Engine)
	if !ok {
		return nil, errConversion(instance, "*gin.Engine")
	}
	return i, nil
}
