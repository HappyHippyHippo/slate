package gmigration

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/genv"
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

var (
	// Database defines the default migration manager rdb name.
	Database = genv.String(
		"SLATE_GMIGRATOR_DATABASE",
		"primary")

	// AutoMigrate defines the flag that will be used to enable
	// the auto migration on application boot.
	AutoMigrate = genv.Bool(
		"SLATE_GMIGRATOR_AUTO_MIGRATE",
		true)
)
