package rdb

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_NewConnectionFactory(t *testing.T) {
	t.Run("missing configuration", func(t *testing.T) {
		sut, e := NewConnectionFactory(nil, &DialectFactory{})
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("missing dialect factory", func(t *testing.T) {
		sut, e := NewConnectionFactory(config.NewManager(0), nil)
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("valid creation", func(t *testing.T) {
		if sut, e := NewConnectionFactory(config.NewManager(0), &DialectFactory{}); sut == nil {
			t.Error("didn't returned the expected valid connection factory instance")
		} else if e != nil {
			t.Errorf("return the unexpected error : %v", e)
		}
	})

	t.Run("config change purge all stored connections", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		cfg1 := config.Config{"dialect": "sqlite", "host": ":memory:"}
		cfg2 := config.Config{"dialect": "sqlite", "host": ":memory:"}
		Config1 := config.Config{"slate": config.Config{"rdb": config.Config{"connections": config.Config{name: cfg1}}}}
		Config2 := config.Config{"slate": config.Config{"rdb": config.Config{"connections": config.Config{name + "salt": cfg2}}}}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(Config1, nil).MinTimes(1)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(Config2, nil).MinTimes(1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&SqliteDialectStrategy{})
		cfgManager := config.NewManager(0)
		_ = cfgManager.AddSource("id1", 0, source1)

		sut, _ := NewConnectionFactory(cfgManager, dialectFactory)

		_, _ = sut.Get(name, &gorm.Config{Logger: logger.Discard})
		if len(sut.(*connectionFactory).instances) != 1 {
			t.Error("didn't store the requested connection instance")
		}

		_ = cfgManager.AddSource("id2", 10, source2)
		if len(sut.(*connectionFactory).instances) != 0 {
			t.Error("didn't removed the stored connection instances")
		}
	})
}

func Test_ConnectionFactory_Get(t *testing.T) {
	t.Run("missing requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dialectFactory := NewMockDialectFactory(ctrl)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(false).Times(1)

		sut, _ := NewConnectionFactory(cfgManager, dialectFactory)

		conn, e := sut.Get("primary", &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, err.DatabaseConfigNotFound):
			t.Errorf("returned the (%v) error when expected (%v)", e, err.DatabaseConfigNotFound)
		}
	})

	t.Run("invalid requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		name := "primary"
		dialectFactory := NewMockDialectFactory(ctrl)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(nil, expected).Times(1)

		sut, _ := NewConnectionFactory(cfgManager, dialectFactory)

		conn, e := sut.Get(name, &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("error instantiating dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		name := "primary"
		cfg := config.Config{"dialect": "invalid", "host": ":memory:"}
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&cfg).Return(nil, expected).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)

		sut, _ := NewConnectionFactory(cfgManager, dialectFactory)

		conn, e := sut.Get(name, &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("error instantiating connector", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		name := "primary"
		cfg := config.Config{"dialect": "invalid", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(expected).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&cfg).Return(dialect, nil).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)

		sut, _ := NewConnectionFactory(cfgManager, dialectFactory)

		conn, e := sut.Get(name, &gorm.Config{Logger: logger.Discard})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expected (%v)", e, expected)
		}
	})

	t.Run("valid connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		cfg := config.Config{"dialect": "invalid", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&cfg).Return(dialect, nil).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)

		sut, _ := NewConnectionFactory(cfgManager, dialectFactory)

		if check, e := sut.Get(name, &gorm.Config{Logger: logger.Discard}); check == nil {
			t.Error("didn't return the expected connection instance")
		} else if e != nil {
			t.Errorf("return the unexpected error : (%v)", e)
		}
	})

	t.Run("multiple requests only instantiate a single connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		cfg := config.Config{"dialect": "invalid", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&cfg).Return(dialect, nil).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)

		sut, _ := NewConnectionFactory(cfgManager, dialectFactory)

		conn, _ := sut.Get(name, &gorm.Config{Logger: logger.Discard})
		check, e := sut.Get(name, &gorm.Config{Logger: logger.Discard})
		switch {
		case check == nil:
			t.Error("didn't return the expected connection instance")
		case e != nil:
			t.Errorf("return the unexpected error : (%v)", e)
		case check != conn:
			t.Error("didn't returned the same instance")
		}
	})
}
