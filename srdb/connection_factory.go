package srdb

import (
	"fmt"
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/gorm"
)

// IConnectionFactory defines the interface of a connection factory instance.
type IConnectionFactory interface {
	Get(name string, gormCfg *gorm.Config) (*gorm.DB, error)
}

// connectionFactory is a database connection generator.
type connectionFactory struct {
	cfg       sconfig.IManager
	dFactory  IDialectFactory
	instances map[string]*gorm.DB
}

var _ IConnectionFactory = &connectionFactory{}

func newConnectionFactory(cfg sconfig.IManager, dFactory IDialectFactory) (IConnectionFactory, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	if dFactory == nil {
		return nil, errNilPointer("cFactory")
	}

	cFactory := &connectionFactory{
		cfg:       cfg,
		dFactory:  dFactory,
		instances: map[string]*gorm.DB{},
	}

	if ObserveConfig {
		_ = cfg.AddObserver(ConnectionsConfigPath, func(_ interface{}, _ interface{}) {
			for _, conn := range cFactory.instances {
				if db, _ := conn.DB(); db != nil {
					_ = db.Close()
				}
			}
			cFactory.instances = map[string]*gorm.DB{}
		})
	}

	return cFactory, nil
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
	} else if dialect, e := f.dFactory.Get(&cfg); e != nil {
		return nil, e
	} else if conn, e := gorm.Open(dialect, gormCfg); e != nil {
		return nil, e
	} else {
		f.instances[name] = conn
		return conn, nil
	}
}
