package srdb

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Provider defines the slate.srdb module service provider to be used on
// the application initialization to register the relational
// database services.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will register the srdb package instances in the
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
		if cfg, e := sconfig.Get(c); e != nil {
			return nil, e
		} else if dFactory, e := GetDialectFactory(c); e != nil {
			return nil, e
		} else {
			return newConnectionFactory(cfg, dFactory)
		}
	})

	_ = c.Factory(ContainerPrimaryID, func() (interface{}, error) {
		if cFactory, e := GetConnectionFactory(c); e != nil {
			return nil, e
		} else if cfg, e := GetConfig(c); e != nil {
			return nil, e
		} else {
			return cFactory.Get(Primary, cfg)
		}
	})

	return nil
}

// Boot will start the srdb package
func (p Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	if dFactory, e := GetDialectFactory(c); e != nil {
		return e
	} else if strategies, e := GetDialectStrategies(c); e != nil {
		return e
	} else {
		for _, strategy := range strategies {
			_ = dFactory.Register(strategy)
		}
	}

	return nil
}
