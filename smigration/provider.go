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

		return newDao(conn)
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

// GetDao will try to retrieve a new migration DAO instances
// from the application service container.
func GetDao(c slate.ServiceContainer) (IDao, error) {
	instance, err := c.Get(ContainerDaoID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(IDao)
	if !ok {
		return nil, errConversion(instance, "IDao")
	}
	return i, nil
}

// GetMigrator will try to retrieve a new migrator instance
// from the application service container.
func GetMigrator(c slate.ServiceContainer) (IMigrator, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(IMigrator)
	if !ok {
		return nil, errConversion(instance, "IMigrator")
	}
	return i, nil
}

// GetMigrations will try to retrieve the registered the list of
// migration instances from the application service container.
func GetMigrations(c slate.ServiceContainer) ([]IMigration, error) {
	tags, err := c.Tagged(ContainerMigrationTag)
	if err != nil {
		return nil, err
	}

	var list []IMigration
	for _, service := range tags {
		s, ok := service.(IMigration)
		if !ok {
			return nil, errConversion(service, "IMigration")
		}
		list = append(list, s)
	}
	return list, nil
}
