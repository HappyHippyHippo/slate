package rdb

import (
	"fmt"

	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// IConnectionPool defines the interface of a connection pool instance.
type IConnectionPool interface {
	Get(name string, gormCfg *gorm.Config) (*gorm.DB, error)
}

// connectionPool is a database connection pool and generator.
type connectionPool struct {
	cfg               config.IManager
	connectionFactory IConnectionFactory
	instances         map[string]*gorm.DB
}

var _ IConnectionPool = &connectionPool{}

// NewConnectionPool will instantiate a new relational
// database connection pool instance.
func NewConnectionPool(
	cfg config.IManager,
	factory IConnectionFactory,
) (IConnectionPool, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// check ConnectionFactory argument reference
	if factory == nil {
		return nil, errNilPointer("factory")
	}
	// instantiate the connection pool instance
	pool := &connectionPool{
		cfg:               cfg,
		connectionFactory: factory,
		instances:         map[string]*gorm.DB{},
	}
	// check if is to observe connection configuration changes
	if ObserveConfig {
		// add an observer to the connections config
		_ = cfg.AddObserver(ConnectionsConfigPath, func(_ interface{}, _ interface{}) {
			// close all the currently opened connections
			for _, conn := range pool.instances {
				if db, _ := conn.DB(); db != nil {
					_ = db.Close()
				}
			}
			// clear the storing pool
			pool.instances = map[string]*gorm.DB{}
		})
	}
	return pool, nil
}

// Get execute the process of the connection creation based on the
// base configuration defined by the given name of the connection,
// and apply the extra connection cfg also given as arguments.
func (f *connectionPool) Get(
	name string,
	gormCfg *gorm.Config,
) (*gorm.DB, error) {
	// check if the connection as already been created and return it
	if conn, ok := f.instances[name]; ok {
		return conn, nil
	}
	// generate the configuration path of the requested connection
	path := fmt.Sprintf("%s.%s", ConnectionsConfigPath, name)
	// check if there is a configuration for the requested connection
	if !f.cfg.Has(path) {
		return nil, errConfigNotFound(path)
	}
	// obtain the connection configuration
	cfg, e := f.cfg.Config(path)
	if e != nil {
		return nil, e
	}
	// create the connection
	conn, e := f.connectionFactory.Create(cfg, gormCfg)
	if e != nil {
		return nil, e
	}
	// store the connection instance
	f.instances[name] = conn
	return conn, nil
}
