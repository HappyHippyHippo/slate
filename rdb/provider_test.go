package rdb

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/fs"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't return the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !container.Has(ConfigID):
			t.Errorf("no connection configuration : %v", sut)
		case !container.Has(DialectFactoryID):
			t.Errorf("no dialect factory : %v", sut)
		case !container.Has(ID):
			t.Errorf("no connection factory : %v", sut)
		case !container.Has(ConnectionPrimaryID):
			t.Errorf("no primary connection handler : %v", sut)
		}
	})

	t.Run("retrieving connection configuration", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		cfg, e := container.Get(ConfigID)
		switch {
		case e != nil:
			t.Errorf("unexpected error (%v)", e)
		case cfg == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving dialect factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		service, e := container.Get(DialectFactoryID)
		switch {
		case e != nil:
			t.Errorf("unexpected error (%v)", e)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving configuration when retrieving connection factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		_ = container.Service(config.ID, func() (*config.Config, error) {
			return nil, expected
		})

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving dialect factory when retrieving connection factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DialectFactoryID, func() (*DialectFactory, error) {
			return nil, expected
		})

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving connection factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		factory, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("unexpected error (%v)", e)
		case factory == nil:
			t.Error("didn't return a valid reference")
		}
	})

	t.Run("error retrieving connection factory when retrieving primary connection", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		_ = container.Service(ID, func() (*ConnectionPool, error) {
			return nil, expected
		})

		if _, e := container.Get(ConnectionPrimaryID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving connection configuration when retrieving primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		_ = (&config.Provider{}).Register(container)
		_ = container.Service(ConfigID, func() (*gorm.Config, error) {
			return nil, expected
		})

		if _, e := container.Get(ConnectionPrimaryID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("valid primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		rdbCfg := config.Partial{"dialect": "invalid", "host": ":memory:"}
		partial := config.Partial{}
		_, _ = partial.Set("slate.rdb.connections.primary", rdbCfg)
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectStrategy := NewMockDialectStrategy(ctrl)
		dialectStrategy.EXPECT().Accept(rdbCfg).Return(true).Times(1)
		dialectStrategy.EXPECT().Create(rdbCfg).Return(dialect, nil).Times(1)
		dialectFactory := NewDialectFactory()
		_ = dialectFactory.Register(dialectStrategy)
		_ = container.Service(DialectFactoryID, func() *DialectFactory {
			return dialectFactory
		})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 1, source)
		_ = container.Service(config.ID, func() *config.Config {
			return cfg
		})

		if check, e := container.Get(ConnectionPrimaryID); e != nil {
			t.Errorf("unexpected error (%v)", e)
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

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		rdbCfg := config.Partial{"dialect": "invalid", "host": ":memory:"}
		partial := config.Partial{}
		_, _ = partial.Set("slate.rdb.connections.other_primary", rdbCfg)
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectStrategy := NewMockDialectStrategy(ctrl)
		dialectStrategy.EXPECT().Accept(rdbCfg).Return(true).Times(1)
		dialectStrategy.EXPECT().Create(rdbCfg).Return(dialect, nil).Times(1)
		dialectFactory := NewDialectFactory()
		_ = dialectFactory.Register(dialectStrategy)
		_ = container.Service(DialectFactoryID, func() *DialectFactory {
			return dialectFactory
		})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 1, source)
		_ = container.Service(config.ID, func() *config.Config {
			return cfg
		})

		if check, e := container.Get(ConnectionPrimaryID); e != nil {
			t.Errorf("unexpected error (%v)", e)
		} else if check == nil {
			t.Error("didn't return a valid reference")
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error retrieving dialect factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(DialectFactoryID, func() (*DialectFactory, error) {
			return nil, expected
		})

		if e := provider.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid dialect factory", func(t *testing.T) {
		container := slate.NewContainer()
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(DialectFactoryID, func() interface{} {
			return "string"
		})

		if e := provider.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving dialect factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) {
			return nil, expected
		}, DialectStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid dialect factory strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() interface{} {
			return "string"
		}, DialectStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("run boot", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()

		dialectStrategy := NewMockDialectStrategy(ctrl)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() interface{} {
			return dialectStrategy
		}, DialectStrategyTag)

		if e := sut.Boot(container); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})
}
