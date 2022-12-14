package rdb

import (
	"fmt"

	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// IConnectionFactory defines the interface of a connection factory instance.
type IConnectionFactory interface {
	Get(name string, gormCfg *gorm.Config) (*gorm.DB, error)
}

// connectionFactory is a database connection generator.
type connectionFactory struct {
	cfg            config.IManager
	dialectFactory IDialectFactory
	instances      map[string]*gorm.DB
}

var _ IConnectionFactory = &connectionFactory{}

// NewConnectionFactory will instantiate a new relational
// database connection factory instance.
func NewConnectionFactory(
	cfg config.IManager,
	dialectFactory IDialectFactory,
) (IConnectionFactory, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}
	// check dialect factory argument reference
	if dialectFactory == nil {
		return nil, errNilPointer("dialectFactory")
	}
	// instantiate the connection factory
	factory := &connectionFactory{
		cfg:            cfg,
		dialectFactory: dialectFactory,
		instances:      map[string]*gorm.DB{},
	}
	// check if is to observe connection configuration changes
	if ObserveConfig {
		// add an observer to the connections config
		_ = cfg.AddObserver(ConnectionsConfigPath, func(_ interface{}, _ interface{}) {
			// close all the current opened connections
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
// and apply the extra connection cfg also given as arguments.
func (f *connectionFactory) Get(
	name string,
	gormCfg *gorm.Config,
) (*gorm.DB, error) {
	// generate the configuration path of the requested connection
	name = fmt.Sprintf("%s.%s", ConnectionsConfigPath, name)
	// check if the connection as already been created and return it
	if conn, ok := f.instances[name]; ok {
		return conn, nil
	}
	// check if there is a configuration for the requested connection
	if !f.cfg.Has(name) {
		return nil, errConfigNotFound(name)
	}
	// obtain the connection configuration
	cfg, e := f.cfg.Config(name)
	if e != nil {
		return nil, e
	}
	// get a dialect instance for the connection
	dialect, e := f.dialectFactory.Get(cfg)
	if e != nil {
		return nil, e
	}
	// open the connection
	conn, e := gorm.Open(dialect, gormCfg)
	if e != nil {
		return nil, e
	}
	// store and return the generated connection
	f.instances[name] = conn
	return conn, nil
}
