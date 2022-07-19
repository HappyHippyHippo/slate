package smigration

import "sort"

// IMigrator defines an object that handles the
// persistence layer migrations.
type IMigrator interface {
	AddMigration(migration IMigration) error
	Current() (uint64, error)
	Migrate() error
	Up() error
	Down() error
}

type migrator struct {
	dao        IDao
	migrations []IMigration
}

var _ IMigrator = &migrator{}

func newMigrator(dao IDao) (IMigrator, error) {
	if dao == nil {
		return nil, errNilPointer("dao")
	}

	return &migrator{
		dao:        dao,
		migrations: []IMigration{},
	}, nil
}

// AddMigration registers a migration into the migration manager.
func (m *migrator) AddMigration(migration IMigration) error {
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
func (m migrator) Current() (uint64, error) {
	current, e := m.dao.Last()
	if e != nil {
		return 0, e
	}

	return current.Version, nil
}

// Migrate execute all migrations that are yet to be executed.
func (m migrator) Migrate() error {
	if len(m.migrations) == 0 {
		return nil
	}

	current, e := m.dao.Last()
	if e != nil {
		return e
	}

	for _, migration := range m.migrations {
		if migrationVersion := migration.Version(); current.Version < migrationVersion {
			if e := migration.Up(); e != nil {
				return e
			}

			current, e = m.dao.Up(migrationVersion)
			if e != nil {
				return e
			}
		}
	}

	return nil
}

// Up will try to execute the next migration in queue to be executed.
func (m migrator) Up() error {
	if len(m.migrations) == 0 {
		return nil
	}

	current, e := m.dao.Last()
	if e != nil {
		return e
	}

	for _, migration := range m.migrations {
		if migrationVersion := migration.Version(); current.Version < migrationVersion {
			if e := migration.Up(); e != nil {
				return e
			}

			_, e = m.dao.Up(migrationVersion)
			return e
		}
	}

	return nil
}

// Down will try to revert the last migration executed.
func (m migrator) Down() error {
	if len(m.migrations) == 0 {
		return nil
	}

	current, e := m.dao.Last()
	if e != nil {
		return e
	}

	for _, migration := range m.migrations {
		if migrationVersion := migration.Version(); current.Version == migrationVersion {
			if e := migration.Down(); e != nil {
				return e
			}

			return m.dao.Down(current)
		}
	}

	return nil
}
