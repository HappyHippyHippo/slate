package grdb

import (
	"github.com/happyhippyhippo/slate"
	"gorm.io/gorm"
)

// GetConfig will try to retrieve a new default gorm config instance
// from the application service container.
func GetConfig(c slate.ServiceContainer) (*gorm.Config, error) {
	instance, err := c.Get(ContainerConfigID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*gorm.Config)
	if !ok {
		return nil, errConversion(instance, "*gorm.Config")
	}
	return i, nil
}

// GetDialectFactory will try to retrieve the registered dialect
// factory instance from the application service container.
func GetDialectFactory(c slate.ServiceContainer) (*DialectFactory, error) {
	instance, err := c.Get(ContainerDialectFactoryID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*DialectFactory)
	if !ok {
		return nil, errConversion(instance, "*DialectFactory")
	}
	return i, nil
}

// GetDialectStrategies will try to retrieve the registered the list of
// dialect strategies instances from the application service container.
func GetDialectStrategies(c slate.ServiceContainer) ([]DialectStrategy, error) {
	tags, err := c.Tagged(ContainerDialectStrategyTag)
	if err != nil {
		return nil, err
	}

	var list []DialectStrategy
	for _, service := range tags {
		s, ok := service.(DialectStrategy)
		if !ok {
			return nil, errConversion(service, "DialectStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}

// GetConnectionFactory will try to retrieve the registered connection
// factory instance from the application service container.
func GetConnectionFactory(c slate.ServiceContainer) (*ConnectionFactory, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*ConnectionFactory)
	if !ok {
		return nil, errConversion(instance, "*ConnectionFactory")
	}
	return i, nil
}

// GetPrimaryConnection will try to retrieve the registered connection
// factory instance from the application service container.
func GetPrimaryConnection(c slate.ServiceContainer) (*gorm.DB, error) {
	instance, err := c.Get(ContainerPrimaryID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*gorm.DB)
	if !ok {
		return nil, errConversion(instance, "*gorm.DB")
	}
	return i, nil
}
