package rdb

import (
	"github.com/happyhippyhippo/slate"
	senv "github.com/happyhippyhippo/slate/env"
)

const (
	// DialectMySQL defines the value to be used to identify a
	// MySQL dialect.
	DialectMySQL = "mysql"

	// DialectSqlite defines the value to be used to identify a
	// Sqlite dialect.
	DialectSqlite = "sqlite"
)

const (
	// ContainerID defines the id to be used as the container
	// registration id of a relational database connection factory instance,
	// and as a base id of all other relational database package instances
	// registered in the application container.
	ContainerID = slate.ContainerID + ".rdb"

	// ContainerConfigID defines the id to be used as the container
	// registration id of the relational database connection configuration
	// instance.
	ContainerConfigID = ContainerID + ".config"

	// ContainerDialectStrategyTag defines the tag to be assigned to all
	// container relational database dialect strategies.
	ContainerDialectStrategyTag = ContainerID + ".dialect.strategy"

	// ContainerDialectStrategyMySQLID defines the id to be used
	// as the container registration id of the relational database connection
	// MySQL dialect instance.
	ContainerDialectStrategyMySQLID = ContainerID + ".dialect.strategy.mysql"

	// ContainerDialectStrategySqliteID defines the id to be used
	// as the container registration id of the relational database connection
	// sqlite dialect instance.
	ContainerDialectStrategySqliteID = ContainerID + ".dialect.strategy.sqlite"

	// ContainerDialectFactoryID defines the id to be used as the
	// container registration id of the relational database connection dialect
	// factory instance.
	ContainerDialectFactoryID = ContainerID + ".dialect.factory"

	// ContainerPrimaryID defines the id to be used as the container
	// registration id of primary relational database instance.
	ContainerPrimaryID = ContainerID + ".primary"
)

const (
	// EnvID defines the rdb package base environment variable name.
	EnvID = slate.EnvID + "_RDB"
)

var (
	// Primary contains the name given to the primary connection.
	Primary = senv.String(EnvID+"_PRIMARY", "primary")

	// ConnectionsConfigPath contains the configuration path that holds the
	// relational database connection configurations.
	ConnectionsConfigPath = senv.String(EnvID+"_CONNECTIONS_CONFIG_PATH", "slate.rdb.connections")

	// ObserveConfig defines the connection factory cfg observing flag
	// used to register in the cfg object an observer of the connection
	// cfg entries list, so it can reset the connections pool.
	ObserveConfig = senv.Bool(EnvID+"_OBSERVE_CONFIG", true)
)
