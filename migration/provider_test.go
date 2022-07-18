package migration

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/rdb"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(ContainerID):
			t.Errorf("didn't registered the migrator : %v", sut)
		case !container.Has(ContainerDaoID):
			t.Errorf("didn't registered the migrator DAO : %v", sut)
		}
	})

	t.Run("error retrieving db connection factory when retrieving migrator DAO", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(rdb.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerDaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid db connection factory when retrieving migrator DAO", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(rdb.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerDaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("error retrieving db connection config when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(rdb.ContainerConfigID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerDaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid db connection config when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(rdb.ContainerConfigID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerDaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("error retrieving connection when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		if _, e := container.Get(ContainerDaoID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrDatabaseConfigNotFound) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrDatabaseConfigNotFound)
		}
	})

	t.Run("retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.primary").Return(partial, nil).Times(1)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		sut, e := container.Get(ContainerDaoID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpectederror (%v)", e)
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

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.secondary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.secondary").Return(partial, nil).Times(1)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		sut, e := container.Get(ContainerDaoID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpectederror (%v)", e)
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

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDaoID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("error retrieving the migrator DAO when retrieving the migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDaoID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerID); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("retrieving migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.primary").Return(partial, nil).Times(1)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		sut, e := container.Get(ContainerID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpectederror (%v)", e)
		case sut == nil:
			t.Error("didn't returned a reference to the migrator")
		default:
			switch sut.(type) {
			case *migrator:
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
		} else if !errors.Is(e, err.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("disable auto migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		AutoMigrate = false
		defer func() { AutoMigrate = true }()

		container := slate.ServiceContainer{}
		sut := &Provider{}

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected err, (%v)", e)
		}
	})

	t.Run("disable migrator auto migration by environment variable", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		AutoMigrate = false
		defer func() { AutoMigrate = true }()

		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpected err, (%v)", e)
		}
	})

	t.Run("error on retrieving migrator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid retrieved migrator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("error on retrieving migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)

		expected := fmt.Errorf("error message")
		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.primary").Return(partial, nil).Times(1)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerMigrationTag)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid migration on retrieving migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)

		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.primary").Return(partial, nil).Times(1)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerMigrationTag)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("running migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Register(container)
		_ = (&rdb.Provider{}).Boot(container)

		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.primary").Return(partial, nil).Times(1)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Times(1)
		_ = container.Service("id", func() (interface{}, error) {
			return migration, nil
		}, ContainerMigrationTag)

		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the unexpectederror : %v", e)
		}
	})
}
