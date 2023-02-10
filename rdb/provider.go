package rdb

import (
	"github.com/happyhippyhippo/slate"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	// ID defines the id to be used as the container
	// registration id of a relational database connection pool instance,
	// and as a base id of all other relational database package instances
	// registered in the application container.
	ID = slate.ID + ".rdb"

	// ConfigID defines the id to be used as the container
	// registration id of the relational database connection configuration
	// instance.
	ConfigID = ID + ".config"

	// DialectStrategyTag defines the tag to be assigned to all
	// container relational database dialect strategies.
	DialectStrategyTag = ID + ".dialect.strategy"

	// MySQLDialectStrategyID defines the id to be used
	// as the container registration id of the relational database connection
	// MySQL dialect instance.
	MySQLDialectStrategyID = ID + ".dialect.strategy.mysql"

	// SqliteDialectStrategyID defines the id to be used
	// as the container registration id of the relational database connection
	// sqlite dialect instance.
	SqliteDialectStrategyID = ID + ".dialect.strategy.sqlite"

	// DialectFactoryID defines the id to be used as the
	// container registration id of the relational database connection dialect
	// factory instance.
	DialectFactoryID = ID + ".dialect.factory"

	// ConnectionFactoryID defines the id to be used as the container
	// registration id of the connection factory instance.
	ConnectionFactoryID = ID + ".connection.factory"

	// ConnectionPrimaryID defines the id to be used as the container
	// registration id of primary relational database instance.
	ConnectionPrimaryID = ID + ".connection.primary"
)

// Provider defines the slate.rdb module service provider to be used on
// the application initialization to register the relational
// database services.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the rdb package instances in the
// application container
func (p Provider) Register(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// add the connection configuration service
	_ = container[0].Service(ConfigID, func() *gorm.Config {
		return &gorm.Config{Logger: logger.Discard}
	})
	// add dialect strategies and factory
	_ = container[0].Service(MySQLDialectStrategyID, NewMySQLDialectStrategy, DialectStrategyTag)
	_ = container[0].Service(SqliteDialectStrategyID, NewSqliteDialectStrategy, DialectStrategyTag)
	_ = container[0].Service(DialectFactoryID, NewDialectFactory)
	// add the connection factory and pool
	_ = container[0].Service(ConnectionFactoryID, NewConnectionFactory)
	_ = container[0].Service(ID, NewConnectionPool)
	// add the primary connection auxiliary service
	_ = container[0].Service(ConnectionPrimaryID, func(
		connectionFactory IConnectionPool,
		cfg *gorm.Config,
	) (*gorm.DB, error) {
		return connectionFactory.Get(Primary, cfg)
	})
	return nil
}

// Boot will start the rdb package
func (p Provider) Boot(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// populate the container dialect factory with all
	// registered dialect strategies
	dialectFactory, e := p.getDialectFactory(container[0])
	if e != nil {
		return e
	}
	strategies, e := p.getDialectStrategies(container[0])
	if e != nil {
		return e
	}
	for _, strategy := range strategies {
		_ = dialectFactory.Register(strategy)
	}
	return nil
}

func (Provider) getDialectFactory(
	container slate.IContainer,
) (IDialectFactory, error) {
	// retrieve the factory entry
	instance, e := container.Get(DialectFactoryID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	i, ok := instance.(IDialectFactory)
	if !ok {
		return nil, errConversion(instance, "rdb.IDialectFactory")
	}
	return i, nil
}

func (Provider) getDialectStrategies(
	container slate.IContainer,
) ([]IDialectStrategy, error) {
	// retrieve the strategies entries
	tags, e := container.Tag(DialectStrategyTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved strategies
	var list []IDialectStrategy
	for _, service := range tags {
		s, ok := service.(IDialectStrategy)
		if !ok {
			return nil, errConversion(service, "rdb.IDialectStrategy")
		}
		list = append(list, s)
	}
	return list, nil
}
