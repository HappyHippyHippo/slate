package smigration

import "github.com/happyhippyhippo/slate"

// GetDao will try to retrieve a new migration DAO instances
// from the application service container.
func GetDao(c slate.ServiceContainer) (*Dao, error) {
	instance, err := c.Get(ContainerDaoID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*Dao)
	if !ok {
		return nil, errConversion(instance, "*Dao")
	}
	return i, nil
}

// GetMigrator will try to retrieve a new migrator instance
// from the application service container.
func GetMigrator(c slate.ServiceContainer) (*Migrator, error) {
	instance, err := c.Get(ContainerID)
	if err != nil {
		return nil, err
	}

	i, ok := instance.(*Migrator)
	if !ok {
		return nil, errConversion(instance, "*Migrator")
	}
	return i, nil
}

// GetMigrations will try to retrieve the registered the list of
// migration instances from the application service container.
func GetMigrations(c slate.ServiceContainer) ([]Migration, error) {
	tags, err := c.Tagged(ContainerMigrationTag)
	if err != nil {
		return nil, err
	}

	var list []Migration
	for _, service := range tags {
		s, ok := service.(Migration)
		if !ok {
			return nil, errConversion(service, "Migration")
		}
		list = append(list, s)
	}
	return list, nil
}
