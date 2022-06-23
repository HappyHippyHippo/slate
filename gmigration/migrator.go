package gmigration

import "sort"

// Migrator defines the structure of an application persistence layer
// migration manager.
type Migrator struct {
	dao        *Dao
	migrations []Migration
}

// NewMigrator instantiates a new migration manager instance.
func NewMigrator(dao *Dao) (*Migrator, error) {
	if dao == nil {
		return nil, errNilPointer("dao")
	}

	return &Migrator{
		dao:        dao,
		migrations: []Migration{},
	}, nil
}

// AddMigration registers a migration into the migration manager.
func (m *Migrator) AddMigration(migration Migration) error {
	if migration == nil {
		return errNilPointer("migration")
	}

	m.migrations = append(m.migrations, migration)

	sort.Slice(m.migrations, func(i, j int) bool {
		return m.migrations[i].Version() < m.migrations[j].Version()
	})

	return nil
}

// Current returns the version of the last executed migration.
func (m Migrator) Current() (uint64, error) {
	current, err := m.dao.Last()
	if err != nil {
		return 0, err
	}

	return current.Version, nil
}

// Migrate execute all migrations that are yet to be executed.
func (m Migrator) Migrate() error {
	if len(m.migrations) == 0 {
		return nil
	}

	current, err := m.dao.Last()
	if err != nil {
		return err
	}

	for _, migration := range m.migrations {
		if migrationVersion := migration.Version(); current.Version < migrationVersion {
			if err := migration.Up(); err != nil {
				return err
			}

			current, err = m.dao.Up(migrationVersion)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Up will try to execute the next migration in queue to be executed.
func (m Migrator) Up() error {
	if len(m.migrations) == 0 {
		return nil
	}

	current, err := m.dao.Last()
	if err != nil {
		return err
	}

	for _, migration := range m.migrations {
		if migrationVersion := migration.Version(); current.Version < migrationVersion {
			if err := migration.Up(); err != nil {
				return err
			}

			_, err = m.dao.Up(migrationVersion)
			return err
		}
	}

	return nil
}

// Down will try to revert the last migration executed.
func (m Migrator) Down() error {
	if len(m.migrations) == 0 {
		return nil
	}

	current, err := m.dao.Last()
	if err != nil {
		return err
	}

	for _, migration := range m.migrations {
		if migrationVersion := migration.Version(); current.Version == migrationVersion {
			if err := migration.Down(); err != nil {
				return err
			}

			return m.dao.Down(current)
		}
	}

	return nil
}
