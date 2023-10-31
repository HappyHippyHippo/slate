package slate

import (
	"errors"
	"sort"
	"time"

	"gorm.io/gorm"
)

// ----------------------------------------------------------------------------
// defs
// ----------------------------------------------------------------------------

const (
	// MigratorContainerID defines the id to be used as the Provider
	// registration id of a migrator instance, and as a base id of all
	// other migration module services registered in the application
	// Provider.
	MigratorContainerID = ContainerID + ".migrator"

	// MigratorDAOContainerID defines the id to be used as the Provider
	// registration id of the migrator DAO.
	MigratorDAOContainerID = MigratorContainerID + ".dao"

	// MigratorMigrationTag defines the tag to be used
	// to identify a migration entry in the Provider.
	MigratorMigrationTag = MigratorContainerID + ".migration"

	// MigratorAllMigrationsContainerID defines the id to be used to
	// retrieve the list of all migrations from the application Provider.
	MigratorAllMigrationsContainerID = MigratorMigrationTag + ".all"

	// MigratorEnvID defines the migrator module base environment variable name.
	MigratorEnvID = EnvID + "_MIGRATION"
)

var (
	// MigratorDatabase defines the migration manager rdb connection name.
	MigratorDatabase = EnvString(MigratorEnvID+"_DATABASE", "primary")

	// MigratorAutoMigrate defines the flag that will be used to enable
	// the auto migration on application boot.
	MigratorAutoMigrate = EnvBool(MigratorEnvID+"_AUTO_MIGRATE", true)
)

// ----------------------------------------------------------------------------
// migrator record
// ----------------------------------------------------------------------------

// MigratorRecord defines the rdb record that stores a migration.
type MigratorRecord struct {
	ID uint `json:"id" xml:"id" gorm:"primaryKey"`

	Version uint64 `json:"model" xml:"model"`

	CreatedAt time.Time  `json:"createdAt" xml:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" xml:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" xml:"deletedAt" sql:"index"`
}

// TableName defines the table name to be used to manage the migrations.
func (MigratorRecord) TableName() string {
	return "__migrations"
}

// ----------------------------------------------------------------------------
// migrator DAO
// ----------------------------------------------------------------------------

// MigratorDao defines an migration record DAO instance responsible
// to manager the installed migrations.
type MigratorDao struct {
	conn *gorm.DB
}

// NewMigratorDao will instantiate a new migration DAO instance.
func NewMigratorDao(
	conn *gorm.DB,
) (*MigratorDao, error) {
	// check db argument reference
	if conn == nil {
		return nil, errNilPointer("conn")
	}
	// execute the dao auto migration to guarantee the table existence
	if e := conn.AutoMigrate(&MigratorRecord{}); e != nil {
		return nil, e
	}
	// instantiate the DAO
	return &MigratorDao{conn: conn}, nil
}

// Last will retrieve the last registered migration
func (d MigratorDao) Last() (MigratorRecord, error) {
	// retrieve the last record from the version control table
	model := MigratorRecord{}
	result := d.conn.
		Order("created_at desc").
		FirstOrInit(&model, MigratorRecord{Version: 0})
	// check for retrieval error
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return MigratorRecord{}, result.Error
		}
	}
	return model, nil
}

// Up will register a new executed migration
func (d MigratorDao) Up(
	version uint64,
) (MigratorRecord, error) {
	// add the new version info into the database
	model := MigratorRecord{Version: version}
	if result := d.conn.Create(&model); result.Error != nil {
		return MigratorRecord{}, result.Error
	}
	return model, nil
}

// Down will remove the last migration register
func (d MigratorDao) Down(
	last MigratorRecord,
) error {
	// check if there is info of the last entry
	if last.Version != 0 {
		// remove the last version entry from the database
		if result := d.conn.Unscoped().Delete(&last); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// migration
// ----------------------------------------------------------------------------

// Migration defines an interface that all migrations must obey.
type Migration interface {
	Version() uint64
	Up() error
	Down() error
}

// ----------------------------------------------------------------------------
// migrator
// ----------------------------------------------------------------------------

// Migrator defines a new migration manager instance.
type Migrator struct {
	dao        *MigratorDao
	migrations []Migration
}

// NewMigrator will instantiate a new Migrator instance.
func NewMigrator(
	dao *MigratorDao,
	migrations []Migration,
) (*Migrator, error) {
	// check migration version dao argument reference
	if dao == nil {
		return nil, errNilPointer("dao")
	}
	// sort the migrations
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version() < migrations[j].Version()
	})
	// instantiate the migration manager
	return &Migrator{
		dao:        dao,
		migrations: migrations,
	}, nil
}

// AddMigration registers a migration into the migration manager.
func (m *Migrator) AddMigration(
	migration Migration,
) error {
	// check the migration argument reference
	if migration == nil {
		return errNilPointer("migration")
	}
	// add the migration to the manager migrations pool
	m.migrations = append(m.migrations, migration)
	// sort the migrations pool by version
	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Version() < m.migrations[j].Version()
	})
	return nil
}

// Current returns the version of the last executed migration.
func (m *Migrator) Current() (uint64, error) {
	// get the current/last applied migration
	current, e := m.dao.Last()
	if e != nil {
		return 0, e
	}
	return current.Version, nil
}

// Migrate execute all migrations that are yet to be executed.
func (m *Migrator) Migrate() error {
	// check if there is migrations registered
	if len(m.migrations) == 0 {
		return nil
	}
	// get the current/last applied migration
	current, e := m.dao.Last()
	if e != nil {
		return e
	}
	// iterate through all registered migrations
	for _, migration := range m.migrations {
		// check if the iterated migration has a higher version from
		// the last applied one
		if v := migration.Version(); current.Version < v {
			// execute the migration
			if e := migration.Up(); e != nil {
				return e
			}
			// register the executed migration
			if current, e = m.dao.Up(v); e != nil {
				return e
			}
		}
	}
	return nil
}

// Up will try to execute the next migration in queue to be executed.
func (m *Migrator) Up() error {
	// check if there is migrations registered
	if len(m.migrations) == 0 {
		return nil
	}
	// get the current/last applied migration
	current, e := m.dao.Last()
	if e != nil {
		return e
	}
	// iterate through all registered migrations
	for _, migration := range m.migrations {
		// check if the iterated migration has a higher version from
		// the last applied one
		if v := migration.Version(); current.Version < v {
			// execute the migration
			if e := migration.Up(); e != nil {
				return e
			}
			// register the executed migration
			_, e = m.dao.Up(v)
			return e
		}
	}
	return nil
}

// Down will try to revert the last migration executed.
func (m *Migrator) Down() error {
	// check if there is migrations registered
	if len(m.migrations) == 0 {
		return nil
	}
	// get the current/last applied migration
	current, e := m.dao.Last()
	if e != nil {
		return e
	}
	// iterate through all registered migrations
	for _, migration := range m.migrations {
		// check if the iterated migration is the currently applied migration
		if v := migration.Version(); current.Version == v {
			// execute the undo action of the migration
			if e := migration.Down(); e != nil {
				return e
			}
			// remove the record of the execution of the migration
			return m.dao.Down(current)
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// migrator service register
// ----------------------------------------------------------------------------

// MigratorServiceRegister defines the slate.migration module service provider to be used on
// the application initialization to register the migrations service.
type MigratorServiceRegister struct {
	ServiceRegister
}

var _ ServiceProvider = &MigratorServiceRegister{}

// NewMigratorServiceRegister will generate a new registry instance
func NewMigratorServiceRegister(
	app ...*App,
) *MigratorServiceRegister {
	return &MigratorServiceRegister{
		ServiceRegister: *NewServiceRegister(app...),
	}
}

// Provide will register the migration package connections in the
// application Provider
func (sr MigratorServiceRegister) Provide(
	container *ServiceContainer,
) error {
	// check container argument reference
	if container == nil {
		return errNilPointer("Provider")
	}
	// register the services
	_ = container.Add(MigratorDAOContainerID, sr.getDAO())
	_ = container.Add(MigratorAllMigrationsContainerID, sr.getMigrations(container))
	_ = container.Add(MigratorContainerID, NewMigrator)
	return nil
}

// Boot will start the migration package
// If the auto migration is defined as true, ether by global variable or
// by environment variable, the migrator will automatically try to migrate
// to the last registered migration
func (sr MigratorServiceRegister) Boot(
	container *ServiceContainer,
) (e error) {
	// check container argument reference
	if container == nil {
		return errNilPointer("Provider")
	}
	// check the application auto migration flag
	if !MigratorAutoMigrate {
		return nil
	}
	// execute the migrations
	migrator, e := sr.getMigrator(container)
	if e != nil {
		return e
	}
	return migrator.Migrate()
}

func (MigratorServiceRegister) getMigrator(
	container *ServiceContainer,
) (*Migrator, error) {
	// retrieve the manager entry
	entry, e := container.Get(MigratorContainerID)
	if e != nil {
		return nil, e
	}
	// validate the retrieved entry type
	if instance, ok := entry.(*Migrator); ok {
		return instance, nil
	}
	return nil, errConversion(entry, "*Migrator")
}

func (MigratorServiceRegister) getDAO() func(pool *RdbConnectionPool, config *gorm.Config) (*MigratorDao, error) {
	return func(pool *RdbConnectionPool, config *gorm.Config) (*MigratorDao, error) {
		conn, e := pool.Get(MigratorDatabase, config)
		if e != nil {
			return nil, e
		}
		return NewMigratorDao(conn)
	}
}

func (MigratorServiceRegister) getMigrations(
	container *ServiceContainer,
) func() []Migration {
	return func() []Migration {
		// retrieve all the migrations from the Provider
		var migrations []Migration
		entries, _ := container.Tag(MigratorMigrationTag)
		for _, entry := range entries {
			// type check the retrieved service
			s, ok := entry.(Migration)
			if ok {
				migrations = append(migrations, s)
			}
		}
		return migrations
	}
}
