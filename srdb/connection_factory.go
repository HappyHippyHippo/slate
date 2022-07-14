package srdb

import (
	"fmt"
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/gorm"
)

// IConnectionFactory defines the interface of a connection factory instance.
type IConnectionFactory interface {
	Get(name string, gcfg *gorm.Config) (*gorm.DB, error)
}

// ConnectionFactory is a database connection generator.
type ConnectionFactory struct {
	config    sconfig.IManager
	dFactory  IDialectFactory
	instances map[string]*gorm.DB
}

var _ IConnectionFactory = &ConnectionFactory{}

func newConnectionFactory(config sconfig.IManager, dFactory IDialectFactory) (IConnectionFactory, error) {
	if config == nil {
		return nil, errNilPointer("config")
	}
	if dFactory == nil {
		return nil, errNilPointer("cFactory")
	}

	cFactory := &ConnectionFactory{
		config:    config,
		dFactory:  dFactory,
		instances: map[string]*gorm.DB{},
	}

	if ObserveConfig {
		_ = config.AddObserver(ConnectionsConfigPath, func(_ interface{}, _ interface{}) {
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
// and apply the extra connection config also given as arguments.
func (f *ConnectionFactory) Get(name string, gcfg *gorm.Config) (*gorm.DB, error) {
	name = fmt.Sprintf("%s.%s", ConnectionsConfigPath, name)

	if conn, ok := f.instances[name]; ok {
		return conn, nil
	}
	if !f.config.Has(name) {
		return nil, errDatabaseConfigNotFound(name)
	}

	if cfg, err := f.config.Partial(name); err != nil {
		return nil, err
	} else if dialect, err := f.dFactory.Get(&cfg); err != nil {
		return nil, err
	} else if conn, err := gorm.Open(dialect, gcfg); err != nil {
		return nil, err
	} else {
		f.instances[name] = conn
		return conn, nil
	}
}
