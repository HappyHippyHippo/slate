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

var _ slate.ServiceProvider = &Provider{}

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
		return &DialectStrategyMySQL{}, nil
	}, ContainerDialectStrategyTag)

	_ = c.Service(ContainerDialectStrategySqliteID, func() (interface{}, error) {
		return &DialectStrategySqlite{}, nil
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
			return NewConnectionFactory(cfg, dialectFactory)
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
