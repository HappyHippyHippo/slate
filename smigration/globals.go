package smigration

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/senv"
)

const (
	// ContainerID defines the id to be used as the container
	// registration id of a migrator instance, and as a base id of all
	// other smigration package instances registered in the application
	// container.
	ContainerID = slate.ContainerID + ".migrator"

	// ContainerDaoID defines the id to be used as the container
	// registration id of the migrator DAO.
	ContainerDaoID = ContainerID + ".dao"

	// ContainerMigrationTag defines the default tag to be used
	// to identify a smigration entry in the container.
	ContainerMigrationTag = ContainerID + ".smigration"
)

const (
	// EnvID defines the slate.smigration package base environment variable name.
	EnvID = slate.EnvID + "_SMIGRATION"
)

var (
	// Database defines the default smigration manager srdb name.
	Database = senv.String(EnvID+"_DATABASE", "primary")

	// AutoMigrate defines the flag that will be used to enable
	// the auto smigration on application boot.
	AutoMigrate = senv.Bool(EnvID+"_AUTO_MIGRATE", true)
)
