package srdb

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Provider defines the slate.rdb module service provider to be used on
// the application initialization to register the relational
// database services.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will register the rdb package instances in the
// application container
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Factory(ContainerConfigID, func() (interface{}, error) {
		return &gorm.Config{Logger: gormLogger.Discard}, nil
	})

	_ = c.Service(ContainerDialectStrategyMySQLID, func() (interface{}, error) {
		return &dialectStrategyMySQL{}, nil
	}, ContainerDialectStrategyTag)

	_ = c.Service(ContainerDialectStrategySqliteID, func() (interface{}, error) {
		return &dialectStrategySqlite{}, nil
	}, ContainerDialectStrategyTag)

	_ = c.Service(ContainerDialectFactoryID, func() (interface{}, error) {
		return &DialectFactory{}, nil
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		if cfg, err := sconfig.GetConfig(c); err != nil {
			return nil, err
		} else if dialectFactory, err := GetDialectFactory(c); err != nil {
			return nil, err
		} else {
			return newConnectionFactory(cfg, dialectFactory)
		}
	})

	_ = c.Factory(ContainerPrimaryID, func() (interface{}, error) {
		if connectionFactory, err := GetConnectionFactory(c); err != nil {
			return nil, err
		} else if cfg, err := GetConfig(c); err != nil {
			return nil, err
		} else {
			return connectionFactory.Get(Primary, cfg)
		}
	})

	return nil
}

// Boot will start the rdb package
func (p Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	if dialectFactory, err := GetDialectFactory(c); err != nil {
		return err
	} else if dialectStrategies, err := GetDialectStrategies(c); err != nil {
		return err
	} else {
		for _, strategy := range dialectStrategies {
			_ = dialectFactory.Register(strategy)
		}
	}

	return nil
}

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
func GetDialectFactory(c slate.ServiceContainer) (IDialectFactory, error) {
	instance, err := c.Get(ContainerDialectFactoryID)
	if err != nil {
		return nil, err
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
	tags, err := c.Tagged(ContainerDialectStrategyTag)
	if err != nil {
		return nil, err
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
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
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
