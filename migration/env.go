package migration

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/env"
)

const (
	// EnvID defines the slate.migration package base environment variable name.
	EnvID = slate.EnvID + "_MIGRATION"
)

var (
	// Database defines the simple migration manager rdb name.
	Database = env.String(EnvID+"_DATABASE", "primary")

	// AutoMigrate defines the flag that will be used to enable
	// the auto migration on application boot.
	AutoMigrate = env.Bool(EnvID+"_AUTO_MIGRATE", true)
)
