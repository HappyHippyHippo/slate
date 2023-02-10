package migration

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/rdb"
	"gorm.io/gorm"
)

const (
	// ID defines the id to be used as the container
	// registration id of a migrator instance, and as a base id of all
	// other migration package instances registered in the application
	// container.
	ID = slate.ID + ".migrator"

	// DaoID defines the id to be used as the container
	// registration id of the migrator DAO.
	DaoID = ID + ".dao"

	// MigrationTag defines the default tag to be used
	// to identify a migration entry in the container.
	MigrationTag = ID + ".migration"
)

// Provider defines the slate.migration module service provider to be used on
// the application initialization to register the migrations service.
type Provider struct{}

var _ slate.IProvider = &Provider{}

// Register will register the migration package instances in the
// application container
func (p Provider) Register(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// add the migration DAO
	_ = container[0].Service(DaoID, func(connPool rdb.IConnectionPool, cfg *gorm.Config) (IDao, error) {
		// retrieve the connection instance to be given to the
		// version control DAO instance
		conn, e := connPool.Get(Database, cfg)
		if e != nil {
			return nil, e
		}
		// instantiate the required DAO
		return NewDao(conn)
	})
	// add the migration manager
	_ = container[0].Service(ID, NewMigrator)
	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (p Provider) Boot(
	container ...slate.IContainer,
) error {
	// check container argument reference
	if len(container) == 0 || container[0] == nil {
		return errNilPointer("container")
	}
	// check the application auto migration flag
	if !AutoMigrate {
		return nil
	}
	// retrieve the migration manager
	migrator, e := p.getMigrator(container[0])
	if e != nil {
		return e
	}
	// retrieve the list of migrations of the application
	migrations, e := p.getMigrations(container[0])
	if e != nil {
		return e
	}
	// add all the found migrations into the migration manager
	for _, migration := range migrations {
		_ = migrator.AddMigration(migration)
	}
	// execute the migrations
	return migrator.Migrate()
}

func (p Provider) getMigrator(
	container slate.IContainer,
) (IMigrator, error) {
	// retrieve the manager entry
	instance, e := container.Get(ID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	i, ok := instance.(IMigrator)
	if !ok {
		return nil, errConversion(instance, "migration.IMigrator")
	}
	return i, nil
}

func (p Provider) getMigrations(
	container slate.IContainer,
) ([]IMigration, error) {
	// retrieve the migrations entries
	tags, e := container.Tag(MigrationTag)
	if e != nil {
		return nil, e
	}
	// type check the retrieved migrations
	var list []IMigration
	for _, service := range tags {
		s, ok := service.(IMigration)
		if !ok {
			return nil, errConversion(service, "migration.IMigration")
		}
		list = append(list, s)
	}
	return list, nil
}
