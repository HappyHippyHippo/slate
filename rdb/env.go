package rdb

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/env"
)

const (
	// EnvID defines the rdb package base environment variable name.
	EnvID = slate.EnvID + "_RDB"
)

var (
	// Primary contains the name given to the primary connection.
	Primary = env.String(EnvID+"_PRIMARY", "primary")

	// ConnectionsConfigPath contains the configuration path that holds the
	// relational database connection configurations.
	ConnectionsConfigPath = env.String(EnvID+"_CONNECTIONS_CONFIG_PATH", "slate.rdb.connections")

	// ObserveConfig defines the connection factory config observing flag
	// used to register in the config object an observer of the connection
	// config entries list, so it can reset the connections pool.
	ObserveConfig = env.Bool(EnvID+"_OBSERVE_CONFIG", true)
)
