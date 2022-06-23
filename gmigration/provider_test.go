package gmigration

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
	"github.com/happyhippyhippo/slate/grdb"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if err := (&Provider{}).Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}

		err := p.Register(container)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case !container.Has(ContainerID):
			t.Errorf("didn't registered the migrator : %v", p)
		case !container.Has(ContainerDaoID):
			t.Errorf("didn't registered the migrator DAO : %v", p)
		}
	})

	t.Run("error retrieving db connection factory when retrieving migrator DAO", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(grdb.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerDaoID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid db connection factory when retrieving migrator DAO", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(grdb.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerDaoID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error retrieving db connection config when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(grdb.ContainerConfigID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerDaoID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid db connection config when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(grdb.ContainerConfigID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerDaoID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error retrieving connection when retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		if _, err := container.Get(ContainerDaoID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrDatabaseConfigNotFound) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrDatabaseConfigNotFound)
		}
	})

	t.Run("retrieving migrator DAO", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(
			gconfig.Partial{
				"rdb": gconfig.Partial{
					"connections": gconfig.Partial{
						"primary": gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
					},
				},
			},
			nil,
		).Times(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(gconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		sut, err := container.Get(ContainerDaoID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
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
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					"secondary": gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}, nil,
		).Times(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(gconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		sut, err := container.Get(ContainerDaoID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
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
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDaoID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error retrieving the migrator DAO when retrieving the migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDaoID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("retrieving migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Boot(container)
		_ = (&Provider{}).Register(container)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					"primary": gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}, nil,
		).Times(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(gconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		sut, err := container.Get(ContainerID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
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
		if err := (&Provider{}).Boot(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("disable auto migration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		AutoMigrate = false
		defer func() { AutoMigrate = true }()
		container := slate.ServiceContainer{}
		p := &Provider{}

		if err := p.Boot(container); err != nil {
			t.Errorf("returned the unexpected error, (%v)", err)
		}
	})

	t.Run("disable migrator auto migration by environment variable", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		AutoMigrate = false
		defer func() { AutoMigrate = true }()
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)

		if err := p.Boot(container); err != nil {
			t.Errorf("returned the unexpected error, (%v)", err)
		}
	})

	t.Run("error on retrieving migrator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid retrieved migrator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}
		_ = p.Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error on retrieving migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")

		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Boot(container)
		p := &Provider{}
		_ = p.Register(container)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					"primary": gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}, nil,
		).Times(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(gconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerMigrationTag)

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid migration on retrieving migrations", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Boot(container)
		p := &Provider{}
		_ = p.Register(container)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					"primary": gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}, nil,
		).Times(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(gconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerMigrationTag)

		if err := p.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("running migrator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Register(container)
		_ = (&grdb.Provider{}).Boot(container)
		p := &Provider{}
		_ = p.Register(container)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					"primary": gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}, nil,
		).Times(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(gconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		migration := NewMockMigration(ctrl)
		migration.EXPECT().Version().Return(uint64(1)).Times(1)
		migration.EXPECT().Up().Times(1)
		_ = container.Service("id", func() (interface{}, error) {
			return migration, nil
		}, ContainerMigrationTag)

		if err := p.Boot(container); err != nil {
			t.Errorf("returned the unexpected error : %v", err)
		}
	})
}
