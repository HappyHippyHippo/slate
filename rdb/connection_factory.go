package rdb

import (
	"fmt"

	sconfig "github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// IConnectionFactory defines the interface of a connection factory instance.
type IConnectionFactory interface {
	Get(name string, gormCfg *gorm.Config) (*gorm.DB, error)
}

// connectionFactory is a database connection generator.
type connectionFactory struct {
	cfg            sconfig.IManager
	dialectFactory IDialectFactory
	instances      map[string]*gorm.DB
}

var _ IConnectionFactory = &connectionFactory{}

// NewConnectionFactory will instantiate a new relational
// database connection factory instance.
func NewConnectionFactory(cfg sconfig.IManager, dialectFactory IDialectFactory) (IConnectionFactory, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	if dialectFactory == nil {
		return nil, errNilPointer("dialectFactory")
	}

	connFactory := &connectionFactory{
		cfg:            cfg,
		dialectFactory: dialectFactory,
		instances:      map[string]*gorm.DB{},
	}

	if ObserveConfig {
		_ = cfg.AddObserver(ConnectionsConfigPath, func(_ interface{}, _ interface{}) {
			for _, conn := range connFactory.instances {
				if db, _ := conn.DB(); db != nil {
					_ = db.Close()
				}
			}
			connFactory.instances = map[string]*gorm.DB{}
		})
	}

	return connFactory, nil
}

// Get execute the process of the connection creation based on the
// base configuration defined by the given name of the connection,
// and apply the extra connection cfg also given as arguments.
func (f *connectionFactory) Get(name string, gormCfg *gorm.Config) (*gorm.DB, error) {
	name = fmt.Sprintf("%s.%s", ConnectionsConfigPath, name)

	if conn, ok := f.instances[name]; ok {
		return conn, nil
	}
	if !f.cfg.Has(name) {
		return nil, errDatabaseConfigNotFound(name)
	}

	if cfg, e := f.cfg.Partial(name); e != nil {
		return nil, e
	} else if dialect, e := f.dialectFactory.Get(&cfg); e != nil {
		return nil, e
	} else if conn, e := gorm.Open(dialect, gormCfg); e != nil {
		return nil, e
	} else {
		f.instances[name] = conn
		return conn, nil
	}
}
