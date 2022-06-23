package srlog

import (
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate"
)

// GetMiddlewareOk will try to retrieve the registered logging middleware
// for ok responses instance from the application service container.
func GetMiddlewareOk(c slate.ServiceContainer) (func(next gin.HandlerFunc) gin.HandlerFunc, error) {
	instance, err := c.Get(ContainerOkID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(func(next gin.HandlerFunc) gin.HandlerFunc)
	if !ok {
		return nil, errConversion(instance, "func(next gin.HandlerFunc) gin.HandlerFunc")
	}
	return i, nil
}

// GetMiddlewareCreated will try to retrieve the registered logging middleware
// for created responses instance from the application service container.
func GetMiddlewareCreated(c slate.ServiceContainer) (func(next gin.HandlerFunc) gin.HandlerFunc, error) {
	instance, err := c.Get(ContainerCreatedID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(func(next gin.HandlerFunc) gin.HandlerFunc)
	if !ok {
		return nil, errConversion(instance, "func(next gin.HandlerFunc) gin.HandlerFunc")
	}
	return i, nil
}

// GetMiddlewareNoContent will try to retrieve the registered logging middleware
// for no-content responses instance from the application service container.
func GetMiddlewareNoContent(c slate.ServiceContainer) (func(next gin.HandlerFunc) gin.HandlerFunc, error) {
	instance, err := c.Get(ContainerNoContentID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(func(next gin.HandlerFunc) gin.HandlerFunc)
	if !ok {
		return nil, errConversion(instance, "func(next gin.HandlerFunc) gin.HandlerFunc")
	}
	return i, nil
}
