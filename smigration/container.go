package smigration

import "github.com/happyhippyhippo/slate"

// GetDao will try to retrieve a new migration DAO instances
// from the application service container.
func GetDao(c slate.ServiceContainer) (IDao, error) {
	instance, e := c.Get(ContainerDaoID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(IDao)
	if !ok {
		return nil, errConversion(instance, "smigration.IDao")
	}
	return i, nil
}

// GetMigrator will try to retrieve a new migrator instance
// from the application service container.
func GetMigrator(c slate.ServiceContainer) (IMigrator, error) {
	instance, e := c.Get(ContainerID)
	if e != nil {
		return nil, e
	}

	i, ok := instance.(IMigrator)
	if !ok {
		return nil, errConversion(instance, "smigration.IMigrator")
	}
	return i, nil
}

// GetMigrations will try to retrieve the registered the list of
// migration instances from the application service container.
func GetMigrations(c slate.ServiceContainer) ([]IMigration, error) {
	tags, e := c.Tagged(ContainerMigrationTag)
	if e != nil {
		return nil, e
	}

	var list []IMigration
	for _, service := range tags {
		s, ok := service.(IMigration)
		if !ok {
			return nil, errConversion(service, "smigration.IMigration")
		}
		list = append(list, s)
	}
	return list, nil
}
