package migration

import (
	"gorm.io/gorm"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/rdb"
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

	// MigrationTag defines the simple tag to be used
	// to identify a migration entry in the container.
	MigrationTag = ID + ".migration"
)

// Provider defines the slate.migration module service provider to be used on
// the application initialization to register the migrations service.
type Provider struct{}

var _ slate.Provider = &Provider{}

// Register will register the migration package instances in the
// application container
func (p Provider) Register(
	container *slate.Container,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}
	// add the migration DAO
	_ = container.Service(DaoID, func(connPool *rdb.ConnectionPool, cfg *gorm.Config) (*Dao, error) {
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
	_ = container.Service(ID, NewMigrator)
	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (p Provider) Boot(
	container *slate.Container,
) (e error) {
	// check container argument reference
	if container == nil {
		return errNilPointer("container")
	}

	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	// check the application auto migration flag
	if !AutoMigrate {
		return nil
	}
	// add all the found migrations into the migration manager
	migrator := p.getMigrator(container)
	for _, migration := range p.getMigrations(container) {
		_ = migrator.AddMigration(migration)
	}
	// execute the migrations
	return migrator.Migrate()
}

func (p Provider) getMigrator(
	container *slate.Container,
) *Migrator {
	// retrieve the manager entry
	instance, e := container.Get(ID)
	if e != nil {
		panic(e)
	}
	// validate the retrieved entry type
	if i, ok := instance.(*Migrator); ok {
		return i
	}
	panic(errConversion(instance, "*migration.Migrator"))
}

func (p Provider) getMigrations(
	container *slate.Container,
) []Migration {
	// retrieve the migrations entries
	tags, e := container.Tag(MigrationTag)
	if e != nil {
		panic(e)
	}
	// type check the retrieved migrations
	var list []Migration
	for _, service := range tags {
		s, ok := service.(Migration)
		if !ok {
			panic(errConversion(service, "migration.Migration"))
		}
		list = append(list, s)
	}
	return list
}
