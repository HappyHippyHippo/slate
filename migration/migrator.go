package migration

import (
	"sort"
)

type dao interface {
	Last() (Record, error)
	Up(version uint64) (Record, error)
	Down(last Record) error
}

// Migrator defines a new migration manager instance.
type Migrator struct {
	dao        dao
	migrations []Migration
}

// NewMigrator will instantiate a new Migrator instance.
func NewMigrator(
	dao *Dao,
) (*Migrator, error) {
	// check DAO argument reference
	if dao == nil {
		return nil, errNilPointer("dao")
	}
	// instantiate the migration manager
	return &Migrator{
		dao:        dao,
		migrations: []Migration{},
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
		if migrationVersion := migration.Version(); current.Version < migrationVersion {
			// execute the migration
			if e := migration.Up(); e != nil {
				return e
			}
			// register the executed migration
			if current, e = m.dao.Up(migrationVersion); e != nil {
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
		if migrationVersion := migration.Version(); current.Version < migrationVersion {
			// execute the migration
			if e := migration.Up(); e != nil {
				return e
			}
			// register the executed migration
			_, e = m.dao.Up(migrationVersion)
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
		if migrationVersion := migration.Version(); current.Version == migrationVersion {
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
