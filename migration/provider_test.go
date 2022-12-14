package migration

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/rdb"
	"gorm.io/gorm"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(ID):
			t.Errorf("didn't registered the migrator : %v", sut)
		case !container.Has(DaoID):
			t.Errorf("didn't registered the migrator DAO : %v", sut)
		}
	})

	t.Run("error retrieving db connection factory when retrieving migrator DAO", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(rdb.ID, func() (rdb.IConnectionFactory, error) { return nil, expected })

		if _, e := container.Get(DaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving db connection config when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(rdb.ConfigID, func() (*gorm.Config, error) { return nil, expected })

		if _, e := container.Get(DaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving connection when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		if _, e := container.Get(DaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		partial := config.Config{"dialect": "sqlite", "host": ":memory:"}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&partial, nil).Times(1)
		_ = container.Service(config.ID, func() config.IManager { return cfgManager })

		sut, e := container.Get(DaoID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the migrator DAO")
		default:
			switch sut.(type) {
			case *Dao:
			default:
				t.Error("didn't returned a migrator DAO reference")
			}
		}
	})

	t.Run("retrieving migrator DAO with db name from env", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		Database = "secondary"
		defer func() { Database = "primary" }()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		partial := config.Config{"dialect": "sqlite", "host": ":memory:"}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.secondary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.secondary").Return(&partial, nil).Times(1)
		_ = container.Service(config.ID, func() config.IManager { return cfgManager })

		sut, e := container.Get(DaoID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the migrator DAO")
		default:
			switch sut.(type) {
			case *Dao:
			default:
				t.Error("didn't returned a migrator DAO reference")
			}
		}
	})

	t.Run("invalid migrator DAO when retrieving the migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DaoID, func() (IDao, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Container)
		}
	})

	t.Run("error retrieving the migrator DAO when retrieving the migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DaoID, func() (IDao, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Container)
		}
	})

	t.Run("retrieving migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		partial := config.Config{"dialect": "sqlite", "host": ":memory:"}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&partial, nil).Times(1)
		_ = container.Service(config.ID, func() (config.IManager, error) { return cfgManager, nil })

		sut, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the migrator")
		default:
			switch sut.(type) {
			case *Migrator:
			default:
				t.Error("didn't returned a migrator reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("disable auto migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		AutoMigrate = false
		defer func() { AutoMigrate = true }()

		container := slate.NewContainer()
		sut := &Provider{}

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected serr, (%v)", e)
		}
	})

	t.Run("disable migrator auto migration by environment variable", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		AutoMigrate = false
		defer func() { AutoMigrate = true }()

		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected serr, (%v)", e)
		}
	})

	t.Run("error on retrieving migrator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ID, func() (IMigrator, error) { return nil, expected })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Container)
		}
	})

	t.Run("invalid retrieved migrator", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ID, func() (interface{}, error) { return "string", nil })

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("error on retrieving migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)

		expected := fmt.Errorf("error message")
		partial := config.Config{"dialect": "sqlite", "host": ":memory:"}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&partial, nil).Times(1)
		_ = container.Service(config.ID, func() (config.IManager, error) { return cfgManager, nil })
		_ = container.Service("id", func() (interface{}, error) { return nil, expected }, MigrationTag)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Container) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Container)
		}
	})

	t.Run("invalid migration on retrieving migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)

		partial := config.Config{"dialect": "sqlite", "host": ":memory:"}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&partial, nil).Times(1)
		_ = container.Service(config.ID, func() (config.IManager, error) { return cfgManager, nil })
		_ = container.Service("id", func() (interface{}, error) { return "string", nil }, MigrationTag)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("running migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)

		partial := config.Config{"dialect": "sqlite", "host": ":memory:"}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&partial, nil).Times(1)
		_ = container.Service(config.ID, func() (config.IManager, error) { return cfgManager, nil })

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Times(1)
		_ = container.Service("id", func() (interface{}, error) {
			return migration, nil
		}, MigrationTag)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})
}
