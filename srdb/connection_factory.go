package srdb

import (
	"fmt"
	"github.com/happyhippyhippo/slate/sconfig"
	"gorm.io/gorm"
)

// ConnectionFactory is a database connection generator.
type ConnectionFactory struct {
	config         sconfig.Manager
	dialectFactory *DialectFactory
	instances      map[string]*gorm.DB
}

// NewConnectionFactory instantiates a new connection factory instance.
func NewConnectionFactory(config sconfig.Manager, dialectFactory *DialectFactory) (*ConnectionFactory, error) {
	if config == nil {
		return nil, errNilPointer("config")
	}
	if dialectFactory == nil {
		return nil, errNilPointer("factory")
	}

	factory := &ConnectionFactory{
		config:         config,
		dialectFactory: dialectFactory,
		instances:      map[string]*gorm.DB{},
	}

	if ObserveConfig {
		_ = config.AddObserver(ConnectionsConfigPath, func(_ interface{}, _ interface{}) {
			for _, conn := range factory.instances {
				if db, _ := conn.DB(); db != nil {
					_ = db.Close()
				}
			}
			factory.instances = map[string]*gorm.DB{}
		})
	}

	return factory, nil
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
	} else if dialect, err := f.dialectFactory.Get(&cfg); err != nil {
		return nil, err
	} else if conn, err := gorm.Open(dialect, gcfg); err != nil {
		return nil, err
	} else {
		f.instances[name] = conn
		return conn, nil
	}
}
