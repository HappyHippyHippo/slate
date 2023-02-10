package rdb

import (
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

// IConnectionFactory defines the interface of a connection factory instance.
type IConnectionFactory interface {
	Create(cfg config.IConfig, gormCfg *gorm.Config) (*gorm.DB, error)
}

// ConnectionFactory is a database connection generator.
type ConnectionFactory struct {
	dialectFactory IDialectFactory
}

var _ IConnectionFactory = &ConnectionFactory{}

// NewConnectionFactory will instantiate a new relational
// database connection factory instance.
func NewConnectionFactory(
	dialectFactory IDialectFactory,
) (IConnectionFactory, error) {
	// check dialect factory argument reference
	if dialectFactory == nil {
		return nil, errNilPointer("dialectFactory")
	}
	// instantiate the connection factory
	return &ConnectionFactory{
		dialectFactory: dialectFactory,
	}, nil
}

// Create execute the process of the connection creation based on the
// base configuration defined by the given name of the connection,
// and apply the extra connection cfg also given as arguments.
func (f *ConnectionFactory) Create(
	cfg config.IConfig,
	gormCfg *gorm.Config,
) (*gorm.DB, error) {
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
	return conn, nil
}
