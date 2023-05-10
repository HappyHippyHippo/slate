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
	container slate.IContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// add the connection configuration service
	_ = container.Service(ConfigID, func() *gorm.Config {
		return &gorm.Config{Logger: logger.Discard}
	})
	// add dialect strategies and factory
	_ = container.Service(DialectFactoryID, NewDialectFactory)
	// add the connection factory and pool
	_ = container.Service(ConnectionFactoryID, NewConnectionFactory)
	_ = container.Service(ID, NewConnectionPool)
	// add the primary connection auxiliary service
	_ = container.Service(ConnectionPrimaryID, func(
		connectionFactory IConnectionPool,
		cfg *gorm.Config,
	) (*gorm.DB, error) {
		return connectionFactory.Get(Primary, cfg)
	})
	return nil
}

// Boot will start the rdb package
func (p Provider) Boot(
	container slate.IContainer,
) (e error) {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}

	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	// populate the container dialect factory with all
	// registered dialect strategies
	dialectFactory := p.getDialectFactory(container)
	for _, s := range p.getDialectStrategies(container) {
		_ = dialectFactory.Register(s)
	}
	return nil
}

func (Provider) getDialectFactory(
	container slate.IContainer,
) IDialectFactory {
	// retrieve the factory entry
	instance, e := container.Get(DialectFactoryID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	i, ok := instance.(IDialectFactory)
	if !ok {
		panic(errConversion(instance, "rdb.IDialectFactory"))
	}
	return i
}

func (Provider) getDialectStrategies(
	container slate.IContainer,
) []IDialectStrategy {
	// retrieve the strategies entries
	tags, e := container.Tag(DialectStrategyTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved strategies
	var list []IDialectStrategy
	for _, service := range tags {
		s, ok := service.(IDialectStrategy)
		if !ok {
			panic(errConversion(service, "rdb.IDialectStrategy"))
		}
		list = append(list, s)
	}
	return list
}
