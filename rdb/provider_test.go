package rdb

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/fs"
	"gorm.io/gorm"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Register(); e == nil {
			t.Error("didn't return the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't return the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(ConfigID):
			t.Errorf("didn't register the connection configuration : %v", sut)
		case !container.Has(MySQLDialectStrategyID):
			t.Errorf("didn't register the mysql dialect strategy : %v", sut)
		case !container.Has(SqliteDialectStrategyID):
			t.Errorf("didn't register the slite dialect strategy : %v", sut)
		case !container.Has(DialectFactoryID):
			t.Errorf("didn't register the dialect factory : %v", sut)
		case !container.Has(ID):
			t.Errorf("didn't register the connection factory : %v", sut)
		case !container.Has(ConnectionPrimaryID):
			t.Errorf("didn't register the primary connection handler : %v", sut)
		}
	})

	t.Run("retrieving connection configuration", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		cfg, e := container.Get(ConfigID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case cfg == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving mysql dialect strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		service, e := container.Get(MySQLDialectStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving sqlite dialect strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		service, e := container.Get(SqliteDialectStrategyID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("retrieving dialect factory", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		service, e := container.Get(DialectFactoryID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case service == nil:
			t.Error("didn't returned a valid reference")
		}
	})

	t.Run("error retrieving configuration when retrieving connection factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		_ = container.Service(config.ID, func() (config.IManager, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving dialect factory when retrieving connection factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&config.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)
		_ = container.Service(DialectFactoryID, func() (IDialectFactory, error) { return nil, expected })

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
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
			t.Errorf("returned the unexpected error (%v)", e)
		case factory == nil:
			t.Error("didn't return a valid reference")
		}
	})

	t.Run("error retrieving connection factory when retrieving primary connection", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		_ = container.Service(ID, func() (IConnectionPool, error) { return nil, expected })

		if _, e := container.Get(ConnectionPrimaryID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving connection configuration when retrieving primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)
		_ = (&config.Provider{}).Register(container)
		_ = container.Service(ConfigID, func() (*gorm.Config, error) { return nil, expected })

		if _, e := container.Get(ConnectionPrimaryID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("valid primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		cfg := config.Config{"dialect": "sqlite", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&cfg).Return(dialect, nil).Times(1)
		_ = container.Service(DialectFactoryID, func() IDialectFactory { return dialectFactory })
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)
		_ = container.Service(config.ID, func() config.IManager { return cfgManager })

		if check, e := container.Get(ConnectionPrimaryID); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if check == nil {
			t.Error("didn't return a valid reference")
		}
	})

	t.Run("valid primary connection with overridden primary connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		primary := "other_primary"
		Primary = primary
		defer func() { Primary = "primary" }()

		cfg := config.Config{"dialect": "sqlite", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&cfg).Return(dialect, nil).Times(1)
		_ = container.Service(DialectFactoryID, func() IDialectFactory { return dialectFactory })
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.other_primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.other_primary").Return(&cfg, nil).Times(1)
		_ = container.Service(config.ID, func() config.IManager { return cfgManager })

		if check, e := container.Get(ConnectionPrimaryID); e != nil {
			t.Errorf("returned the unexpected error (%v)", e)
		} else if check == nil {
			t.Error("didn't return a valid reference")
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("no argument", func(t *testing.T) {
		if e := (&Provider{}).Boot(); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error retrieving dialect factory", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(DialectFactoryID, func() (IDialectFactory, error) { return nil, expected })

		if e := provider.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid dialect factory", func(t *testing.T) {
		container := slate.NewContainer()
		provider := &Provider{}
		_ = provider.Register(container)
		_ = container.Service(DialectFactoryID, func() interface{} { return "string" })

		if e := provider.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error retrieving dialect factory strategy", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) { return nil, expected }, DialectStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("invalid dialect factory strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service("id", func() (interface{}, error) { return "string", nil }, DialectStrategyTag)

		if e := sut.Boot(container); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("run boot", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
