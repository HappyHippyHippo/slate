package srdb

import (
	"github.com/happyhippyhippo/slate"
	"gorm.io/gorm"
)

// GetConfig will try to retrieve a new default gorm cfg instance
// from the application service container.
func GetConfig(c slate.ServiceContainer) (*gorm.Config, error) {
	instance, e := c.Get(ContainerConfigID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(*gorm.Config)
	if !ok {
		return nil, errConversion(instance, "*gorm.Config")
	}
	return i, nil
}

// GetDialectFactory will try to retrieve the registered dialect
// factory instance from the application service container.
func GetDialectFactory(c slate.ServiceContainer) (IDialectFactory, error) {
	instance, e := c.Get(ContainerDialectFactoryID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(IDialectFactory)
	if !ok {
		return nil, errConversion(instance, "IDialectFactory")
	}
	return i, nil
}

// GetDialectStrategies will try to retrieve the registered the list of
// dialect strategies instances from the application service container.
func GetDialectStrategies(c slate.ServiceContainer) ([]IDialectStrategy, error) {
	tags, e := c.Tagged(ContainerDialectStrategyTag)
	if e != nil {
		return nil, e
	}

	var list []IDialectStrategy
	for _, service := range tags {
		s, ok := service.(IDialectStrategy)
		if !ok {
			return nil, errConversion(service, "IDialectStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetConnectionFactory will try to retrieve the registered connection
// factory instance from the application service container.
func GetConnectionFactory(c slate.ServiceContainer) (IConnectionFactory, error) {
	instance, e := c.Get(ContainerID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(IConnectionFactory)
	if !ok {
		return nil, errConversion(instance, "IConnectionFactory")
	}
	return i, nil
}

// GetPrimaryConnection will try to retrieve the registered connection
// factory instance from the application service container.
func GetPrimaryConnection(c slate.ServiceContainer) (*gorm.DB, error) {
	instance, e := c.Get(ContainerPrimaryID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(*gorm.DB)
	if !ok {
		return nil, errConversion(instance, "*gorm.DB")
	}
	return i, nil
}
