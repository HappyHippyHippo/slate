package smigration

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/srdb"
)

// Provider defines the slate.migration module service provider to be used on
// the application initialization to register the migrations service.
type Provider struct{}

var _ slate.ServiceProvider = &Provider{}

// Register will register the migration package instances in the
// application container
func (p Provider) Register(c slate.ServiceContainer) error {
	if c == nil {
		return errNilPointer("container")
	}

	_ = c.Service(ContainerDaoID, func() (interface{}, error) {
		connFactory, err := srdb.GetConnectionFactory(c)
		if err != nil {
			return nil, err
		}

		connConfig, err := srdb.GetConfig(c)
		if err != nil {
			return nil, err
		}

		conn, err := connFactory.Get(Database, connConfig)
		if err != nil {
			return nil, err
		}

		return NewDao(conn)
	})

	_ = c.Service(ContainerID, func() (interface{}, error) {
		dao, err := GetDao(c)
		if err != nil {
			return nil, err
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

	migrator, err := GetMigrator(c)
	if err != nil {
		return err
	}

	migrations, err := GetMigrations(c)
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		_ = migrator.AddMigration(migration)
	}

	return migrator.Migrate()
}
