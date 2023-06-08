package rdb

import (
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
)

type dialectCreator interface {
	Create(cfg config.Partial) (gorm.Dialector, error)
}

// ConnectionFactory is a database connection generator.
type ConnectionFactory struct {
	dialectCreator dialectCreator
}

var _ connectionCreator = &ConnectionFactory{}

// NewConnectionFactory will instantiate a new relational
// database connection factory instance.
func NewConnectionFactory(
	dialectCreator *DialectFactory,
) (*ConnectionFactory, error) {
	// check dialect factory argument reference
	if dialectCreator == nil {
		return nil, errNilPointer("dialectCreator")
	}
	// instantiate the connection factory
	return &ConnectionFactory{
		dialectCreator: dialectCreator,
	}, nil
}

// Create execute the process of the connection creation based on the
// base configuration defined by the given name of the connection,
// and apply the extra connection config also given as arguments.
func (f *ConnectionFactory) Create(
	cfg config.Partial,
	gormCfg *gorm.Config,
) (*gorm.DB, error) {
	// get a dialect instance for the connection
	dialect, e := f.dialectCreator.Create(cfg)
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
