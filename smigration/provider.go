package smigration

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/srdb"
)

// Provider defines the slate.migration module service provider to be used on
// the application initialization to register the migrations service.
type Provider struct{}

var _ slate.IServiceProvider = &Provider{}

// Register will register the migration package instances in the
// application container
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Service(ContainerDaoID, func() (interface{}, error) {
		connFactory, e := srdb.GetConnectionFactory(c)
		if e != nil {
			return nil, e
		}

		connConfig, e := srdb.GetConfig(c)
		if e != nil {
			return nil, e
		}

		conn, e := connFactory.Get(Database, connConfig)
		if e != nil {
			return nil, e
		}

		return newDao(conn)
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		dao, e := GetDao(c)
		if e != nil {
			return nil, e
		}

		return newMigrator(dao)
	})

	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (p Provider) Boot(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	if !AutoMigrate {
		return nil
	}

	migrator, e := GetMigrator(c)
	if e != nil {
		return e
	}

	migrations, e := GetMigrations(c)
	if e != nil {
		return e
	}

	for _, migration := range migrations {
		_ = migrator.AddMigration(migration)
	}

	return migrator.Migrate()
}
