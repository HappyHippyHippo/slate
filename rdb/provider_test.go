package rdb

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/fs"
	"github.com/pkg/errors"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't return the expected err")
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
		case !container.Has(ContainerConfigID):
			t.Errorf("didn't register the connection configuration : %v", sut)
		case !container.Has(ContainerDialectStrategyMySQLID):
			t.Errorf("didn't register the mysql dialect strategy : %v", sut)
		case !container.Has(ContainerDialectStrategySqliteID):
			t.Errorf("didn't register the slite dialect strategy : %v", sut)
		case !container.Has(ContainerDialectFactoryID):
			t.Errorf("didn't register the dialect factory : %v", sut)
		case !container.Has(ContainerID):
			t.Errorf("didn't register the connection factory : %v", sut)
		case !container.Has(ContainerPrimaryID):
			t.Errorf("didn't register the primary connection handler : %v", sut)
		}
	})

	t.Run("retrieving connection configuration", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		cfg, e := container.Get(ContainerConfigID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case cfg == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("always return a new rdb connection cfg every call", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		cfg1, _ := container.Get(ContainerConfigID)
		cfg2, _ := container.Get(ContainerConfigID)

		if cfg1 == cfg2 {
			t.Error("multiple calls returned the same connection cfg instance")
		}
	})

	t.Run("retrieving mysql dialect strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		service, e := container.Get(ContainerDialectStrategyMySQLID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving sqlite dialect strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		service, e := container.Get(ContainerDialectStrategySqliteID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving dialect factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		service, e := container.Get(ContainerDialectFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving configuration when retrieving connection factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerID); e == nil {
			t.Error("didn't returned the expected err")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid configuration instance on retrieving the connection factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerID); e == nil {
			t.Error("didn't returned the expected err")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("error retrieving dialect factory when retrieving connection factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerID); e == nil {
			t.Error("didn't returned the expected err")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid dialect factory instance on retrieving the connection factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerID); e == nil {
			t.Error("didn't returned the expected err")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("retrieving connection factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&config.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		factory, e := container.Get(ContainerID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
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

		if _, e := container.Get(ContainerPrimaryID); e == nil {
			t.Error("didn't returned the expected err")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid connection factory when retrieving primary connection", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(ContainerID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerPrimaryID); e == nil {
			t.Error("didn't returned the expected err")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("error retrieving connection configuration when retrieving primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&config.Provider{}).Register(container)
		_ = container.Service(ContainerConfigID, func() (interface{}, error) {
			return nil, expected
		})

		if _, e := container.Get(ContainerPrimaryID); e == nil {
			t.Error("didn't returned the expected err")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid connection configuration when retrieving primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&config.Provider{}).Register(container)
		_ = container.Service(ContainerConfigID, func() (interface{}, error) {
			return "string", nil
		})

		if _, e := container.Get(ContainerPrimaryID); e == nil {
			t.Error("didn't returned the expected err")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("valid primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dFactory := NewMockDialectFactory(ctrl)
		dFactory.EXPECT().Get(&partial).Return(dialect, nil).Times(1)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return dFactory, nil
		})
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.primary").Return(partial, nil).Times(1)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		if check, e := container.Get(ContainerPrimaryID); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if check == nil {
			t.Error("didn't return a valid reference")
		}
	})

	t.Run("valid primary connection with overridden primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)

		primary := "other_primary"
		Primary = primary
		defer func() { Primary = "primary" }()

		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dFactory := NewMockDialectFactory(ctrl)
		dFactory.EXPECT().Get(&partial).Return(dialect, nil).Times(1)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return dFactory, nil
		})
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.other_primary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.other_primary").Return(partial, nil).Times(1)
		_ = container.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		if check, e := container.Get(ContainerPrimaryID); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if check == nil {
			t.Error("didn't return a valid reference")
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected err")
		} else if !errors.Is(e, err.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrNilPointer)
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

		if e := provider.Boot(container); e == nil {
			t.Error("didn't returned the expected err")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("invalid dialect factory", func(t *testing.T) {
		container := slate.ServiceContainer{}
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(ContainerDialectFactoryID, func() (interface{}, error) {
			return "string", nil
		})

		if e := provider.Boot(container); e == nil {
			t.Error("didn't returned the expected err")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("error retrieving dialect factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, ContainerDialectStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected err")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("invalid dialect factory strategy", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return "string", nil
		}, ContainerDialectStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected err")
		} else if !errors.Is(e, err.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrConversion)
		}
	})

	t.Run("run boot", func(t *testing.T) {
		container := slate.ServiceContainer{}
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
