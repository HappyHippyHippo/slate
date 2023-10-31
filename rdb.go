package slate

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ----------------------------------------------------------------------------
// defs
// ----------------------------------------------------------------------------

const (
	// RdbContainerID defines the id to be used as the Provider
	// registration id of a relational database connection pool instance,
	// and as a base id of all other relational database package connections
	// registered in the application Provider.
	RdbContainerID = ContainerID + ".rdb"

	// RdbConfigContainerID defines the id to be used as the Provider
	// registration id of the relational database connection configuration
	// instance.
	RdbConfigContainerID = RdbContainerID + ".config"

	// RdbDialectContainerID defines the base if to all rdb dialect
	// related services
	RdbDialectContainerID = RdbContainerID + ".dialect"

	// RdbDialectCreatorTag defines the tag to be assigned to all
	// Provider relational database dialect creators.
	RdbDialectCreatorTag = RdbDialectContainerID + ".creator"

	// RdbAllDialectCreatorsContainerID defines the id to be used as the
	// Provider registration id of an aggregate dialect creators list.
	RdbAllDialectCreatorsContainerID = RdbDialectCreatorTag + ".all"

	// RdbDialectFactoryContainerID defines the id to be used as the
	// Provider registration id of the relational database connection
	// dialect factory service.
	RdbDialectFactoryContainerID = RdbDialectContainerID + ".factory"

	// RdbConnectionContainerID defines the base if to all rdb connection
	// related services
	RdbConnectionContainerID = RdbContainerID + ".connection"

	// RdbConnectionFactoryContainerID defines the id to be used as the
	// Provider registration id of the connection factory service.
	RdbConnectionFactoryContainerID = RdbConnectionContainerID + ".factory"

	// RdbPrimaryConnectionContainerID defines the id to be used as the
	// Provider registration id of the primary relational database service.
	RdbPrimaryConnectionContainerID = RdbConnectionContainerID + ".primary"

	// RdbEnvID defines the base environment variable name for all
	// relational database related environment variables.
	RdbEnvID = EnvID + "_RDB"
)

var (
	// RdbPrimary contains the name given to the primary connection.
	RdbPrimary = EnvString(RdbEnvID+"_PRIMARY", "primary")

	// RdbConnectionsConfigPath contains the configuration path that holds the
	// relational database connection configurations.
	RdbConnectionsConfigPath = EnvString(RdbEnvID+"_CONNECTIONS_CONFIG_PATH", "slate.rdb.connections")

	// RdbObserveConfig defines the connection factory config observing flag
	// used to register in the config manager as an observer of the connection
	// config entries list, so it can reset the connections pool as reconfigure
	// the connections.
	RdbObserveConfig = EnvBool(RdbEnvID+"_OBSERVE_CONFIG", true)
)

// ----------------------------------------------------------------------------
// errors
// ----------------------------------------------------------------------------

var (
	// ErrUnknownRdbDialect defines an error that signal that the
	// requested database connection configured dialect is unknown.
	ErrUnknownRdbDialect = fmt.Errorf("unknown database dialect")
)

func errUnknownRdbDialect(
	config ConfigPartial,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrUnknownRdbDialect, fmt.Sprintf("%v", config), ctx...)
}

// ----------------------------------------------------------------------------
// rdb dialect creator
// ----------------------------------------------------------------------------

// RdbDialectCreator defines the interface to a gorm rdb
// dialect instantiation strategy, based on a configuration.
type RdbDialectCreator interface {
	Accept(config *ConfigPartial) bool
	Create(config *ConfigPartial) (gorm.Dialector, error)
}

// ----------------------------------------------------------------------------
// rdb dialect factory
// ----------------------------------------------------------------------------

// RdbDialectFactory defines an object that will generate a database
// dialect interface based on a registered list of dialect creators.
type RdbDialectFactory []RdbDialectCreator

// NewRdbDialectFactory will instantiate a new relational database
// dialect factory service.
func NewRdbDialectFactory(
	creators []RdbDialectCreator,
) *RdbDialectFactory {
	factory := &RdbDialectFactory{}
	for _, creator := range creators {
		*factory = append(*factory, creator)
	}
	return factory
}

// Create generates a new connection dialect interface defined by the
// configuration parameters stored in the configuration partial marked by
// the given name.
func (f *RdbDialectFactory) Create(
	config *ConfigPartial,
) (gorm.Dialector, error) {
	// check the config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// search for a creator that can create a dialect instance for the
	// dialect name retrieved from the configuration
	for _, creator := range *f {
		if creator.Accept(config) {
			// generate the dialect instance
			return creator.Create(config)
		}
	}
	return nil, errUnknownRdbDialect(*config)
}

// ----------------------------------------------------------------------------
// rdb connection factory
// ----------------------------------------------------------------------------

// RdbConnectionFactory is a database connection generator.
type RdbConnectionFactory struct {
	dialectFactory *RdbDialectFactory
}

// NewRdbConnectionFactory will instantiate a new relational
// database connection factory instance.
func NewRdbConnectionFactory(
	dialectFactory *RdbDialectFactory,
) (*RdbConnectionFactory, error) {
	// check dialect factory argument reference
	if dialectFactory == nil {
		return nil, errNilPointer("dialectFactory")
	}
	// instantiate the connection factory
	return &RdbConnectionFactory{
		dialectFactory: dialectFactory,
	}, nil
}

// Create execute the process of the connection creation based on
// parameters given by the passed config partial.
func (f *RdbConnectionFactory) Create(
	config *ConfigPartial,
	gormConfig *gorm.Config,
) (*gorm.DB, error) {
	// get a dialect instance for the connection
	dialect, e := f.dialectFactory.Create(config)
	if e != nil {
		return nil, e
	}
	// open the connection
	conn, e := gorm.Open(dialect, gormConfig)
	if e != nil {
		return nil, e
	}
	return conn, nil
}

// ----------------------------------------------------------------------------
// rdb connection pool
// ----------------------------------------------------------------------------

// RdbConnectionPool is a database connection pool and generator.
type RdbConnectionPool struct {
	config            *Config
	connectionFactory *RdbConnectionFactory
	connections       map[string]*gorm.DB
}

// NewRdbConnectionPool will instantiate a new relational
// database connection pool instance.
func NewRdbConnectionPool(
	config *Config,
	connectionFactory *RdbConnectionFactory,
) (*RdbConnectionPool, error) {
	// check config argument reference
	if config == nil {
		return nil, errNilPointer("config")
	}
	// check ConnectionFactory argument reference
	if connectionFactory == nil {
		return nil, errNilPointer("connectionFactory")
	}
	// instantiate the connection pool instance
	pool := &RdbConnectionPool{
		config:            config,
		connectionFactory: connectionFactory,
		connections:       map[string]*gorm.DB{},
	}
	// check if is to observe connection configuration changes
	if RdbObserveConfig {
		// add an observer to the connections config
		_ = config.AddObserver(RdbConnectionsConfigPath, func(_ interface{}, _ interface{}) {
			// close all the currently opened connections
			for _, conn := range pool.connections {
				if db, e := conn.DB(); db != nil && e == nil {
					_ = db.Close()
				}
			}
			// clear the storing pool
			pool.connections = map[string]*gorm.DB{}
		})
	}
	return pool, nil
}

// Get execute the process of the connection creation based on the
// base configuration defined by the given name of the connection,
// and apply the extra connection config also given as arguments.
func (f *RdbConnectionPool) Get(
	name string,
	gormConfig *gorm.Config,
) (*gorm.DB, error) {
	// check if the connection as already been created and return it
	if conn, ok := f.connections[name]; ok {
		return conn, nil
	}
	// generate the configuration path of the requested connection
	path := fmt.Sprintf("%s.%s", RdbConnectionsConfigPath, name)
	// check if there is a configuration for the requested connection
	if !f.config.Has(path) {
		return nil, errConfigPathNotFound(path)
	}
	// obtain the connection configuration
	config, e := f.config.Partial(path)
	if e != nil {
		return nil, e
	}
	// create the connection
	conn, e := f.connectionFactory.Create(&config, gormConfig)
	if e != nil {
		return nil, e
	}
	// store the connection instance
	f.connections[name] = conn
	return conn, nil
}

// ----------------------------------------------------------------------------
// rdb service register
// ----------------------------------------------------------------------------

// RdbServiceRegister defines the service provider to be used on
// the application initialization to register the relational
// database services.
type RdbServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &RdbServiceRegister{}

// NewRdbServiceRegister will generate a new service registry instance
func NewRdbServiceRegister(
	app ...*App,
) *RdbServiceRegister {
	return &RdbServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will register the relational database module services in the
// application Provider.
func (sr RdbServiceRegister) Provide(
	container *ServiceContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// register the services
	_ = container.Add(RdbConfigContainerID, sr.getDefaultConfig())
	_ = container.Add(RdbAllDialectCreatorsContainerID, sr.getDialectCreators(container))
	_ = container.Add(RdbDialectFactoryContainerID, NewRdbDialectFactory)
	_ = container.Add(RdbConnectionFactoryContainerID, NewRdbConnectionFactory)
	_ = container.Add(RdbContainerID, NewRdbConnectionPool)
	_ = container.Add(RdbPrimaryConnectionContainerID, sr.getPrimaryConnection())
	return nil
}

func (RdbServiceRegister) getDefaultConfig() func() *gorm.Config {
	return func() *gorm.Config {
		return &gorm.Config{Logger: logger.Discard}
	}
}

func (RdbServiceRegister) getDialectCreators(
	container *ServiceContainer,
) func() []RdbDialectCreator {
	return func() []RdbDialectCreator {
		// retrieve all the dialect creators from the Provider
		var creators []RdbDialectCreator
		entries, _ := container.Tag(RdbDialectCreatorTag)
		for _, entry := range entries {
			// type check the retrieved service
			s, ok := entry.(RdbDialectCreator)
			if ok {
				creators = append(creators, s)
			}
		}
		return creators
	}
}

func (RdbServiceRegister) getPrimaryConnection() func(pool *RdbConnectionPool, config *gorm.Config) (*gorm.DB, error) {
	return func(pool *RdbConnectionPool, config *gorm.Config) (*gorm.DB, error) {
		return pool.Get(RdbPrimary, config)
	}
}
