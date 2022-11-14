package migration

import (
	"github.com/happyhippyhippo/slate"
	senv "github.com/happyhippyhippo/slate/env"
)

const (
	// ContainerID defines the id to be used as the container
	// registration id of a migrator instance, and as a base id of all
	// other migration package instances registered in the application
	// container.
	ContainerID = slate.ContainerID + ".migrator"

	// ContainerDaoID defines the id to be used as the container
	// registration id of the migrator DAO.
	ContainerDaoID = ContainerID + ".dao"

	// ContainerMigrationTag defines the default tag to be used
	// to identify a migration entry in the container.
	ContainerMigrationTag = ContainerID + ".migration"
)

const (
	// EnvID defines the migration package base environment variable name.
	EnvID = slate.EnvID + "_MIGRATION"
)

var (
	// Database defines the default migration manager rdb name.
	Database = senv.String(EnvID+"_DATABASE", "primary")

	// AutoMigrate defines the flag that will be used to enable
	// the auto migration on application boot.
	AutoMigrate = senv.Bool(EnvID+"_AUTO_MIGRATE", true)
)
