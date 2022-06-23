package grdb

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
	"github.com/happyhippyhippo/slate/gfs"
	"github.com/pkg/errors"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if err := (&Provider{}).Register(nil); err == nil {
			t.Error("didn't return the expected error")
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
		case !container.Has(ContainerConfigID):
			t.Errorf("didn't register the connection configuration : %v", p)
		case !container.Has(ContainerDialectStrategyMySQLID):
			t.Errorf("didn't register the mysql dialect strategy : %v", p)
		case !container.Has(ContainerDialectStrategySqliteID):
			t.Errorf("didn't register the slite dialect strategy : %v", p)
		case !container.Has(ContainerDialectFactoryID):
			t.Errorf("didn't register the dialect factory : %v", p)
		case !container.Has(ContainerID):
			t.Errorf("didn't register the connection factory : %v", p)
		case !container.Has(ContainerPrimaryID):
			t.Errorf("didn't register the primary connection handler : %v", p)
		}
	})

	t.Run("retrieving connection configuration", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		cfg, err := container.Get(ContainerConfigID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case cfg == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("always return a new rdb connection config every call", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		cfg1, _ := container.Get(ContainerConfigID)
		cfg2, _ := container.Get(ContainerConfigID)

		if cfg1 == cfg2 {
			t.Error("multiple calls returned the same connection config instance")
		}
	})

	t.Run("retrieving mysql dialect strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		service, err := container.Get(ContainerDialectStrategyMySQLID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving sqlite dialect strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		service, err := container.Get(ContainerDialectStrategySqliteID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving dialect factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		service, err := container.Get(ContainerDialectFactoryID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving configuration when retrieving connection factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(gconfig.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid configuration instance on retrieving the connection factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(gconfig.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error retrieving dialect factory when retrieving connection factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid dialect factory instance on retrieving the connection factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("retrieving connection factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&gconfig.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		factory, err := container.Get(ContainerID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case factory == nil:
			t.Error("didn't return a valid reference")
		}
	})

	t.Run("error retrieving connection factory when retrieving primary connection", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerPrimaryID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid connection factory when retrieving primary connection", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerPrimaryID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error retrieving connection configuration when retrieving primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&gconfig.Provider{}).Register(container)
		_ = container.Service(ContainerConfigID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerPrimaryID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("invalid connection configuration when retrieving primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&gconfig.Provider{}).Register(container)
		_ = container.Service(ContainerConfigID, func() (interface{}, error) {
			return "string", nil
		})

		if _, err := container.Get(ContainerPrimaryID); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("valid primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return dialectFactory, nil
		})
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

		if check, err := container.Get(ContainerPrimaryID); err != nil {
			t.Errorf("returned the unexpected error (%v)", err)
		} else if check == nil {
			t.Error("didn't return a valid reference")
		}
	})

	t.Run("valid primary connection with overridden primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		primary := "other_primary"
		Primary = primary
		defer func() { Primary = "primary" }()

		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return dialectFactory, nil
		})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(
			gconfig.Partial{
				"rdb": gconfig.Partial{
					"connections": gconfig.Partial{
						"other_primary": gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
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

		if check, err := container.Get(ContainerPrimaryID); err != nil {
			t.Errorf("returned the unexpected error (%v)", err)
		} else if check == nil {
			t.Error("didn't return a valid reference")
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if err := (&Provider{}).Boot(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("error retrieving dialect factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("invalid dialect factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error retrieving dialect factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&gfs.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerDialectStrategyTag)

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("invalid dialect factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&gfs.Provider{}).Register(container)
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerDialectStrategyTag)

		if err := provider.Boot(container); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("run boot", func(t *testing.T) {
		container := slate.ServiceContainer{}
		provider := &Provider{}
		_ = provider.Register(container)

		if err := provider.Boot(container); err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})
}
