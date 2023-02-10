package rdb

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_NewConnectionPool(t *testing.T) {
	t.Run("missing configuration", func(t *testing.T) {
		sut, e := NewConnectionPool(nil, &ConnectionFactory{})
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("missing connection factory", func(t *testing.T) {
		sut, e := NewConnectionPool(config.NewManager(0), nil)
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("valid creation", func(t *testing.T) {
		if sut, e := NewConnectionPool(config.NewManager(0), &ConnectionFactory{}); sut == nil {
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
		config1 := config.Config{"slate": config.Config{"rdb": config.Config{"connections": config.Config{name: cfg1}}}}
		config2 := config.Config{"slate": config.Config{"rdb": config.Config{"connections": config.Config{name + "salt": cfg2}}}}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(config1, nil).MinTimes(1)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(config2, nil).MinTimes(1)
		dialectFactory := NewDialectFactory()
		_ = dialectFactory.Register(NewSqliteDialectStrategy())
		connectionFactory, _ := NewConnectionFactory(dialectFactory)
		cfgManager := config.NewManager(0)
		_ = cfgManager.AddSource("id1", 0, source1)

		sut, _ := NewConnectionPool(cfgManager, connectionFactory)

		_, _ = sut.Get(name, &gorm.Config{Logger: logger.Discard})
		if len(sut.(*connectionPool).instances) != 1 {
			t.Error("didn't store the requested connection instance")
		}

		_ = cfgManager.AddSource("id2", 10, source2)
		if len(sut.(*connectionPool).instances) != 0 {
			t.Error("didn't removed the stored connection instances")
		}
	})
}

func Test_ConnectionPool_Get(t *testing.T) {
	t.Run("missing requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		dialectFactory := NewMockDialectFactory(ctrl)
		connectionFactory := &ConnectionFactory{dialectFactory: dialectFactory}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(false).Times(1)

		sut, _ := NewConnectionPool(cfgManager, connectionFactory)

		conn, e := sut.Get("primary", &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, ErrConfigNotFound):
			t.Errorf("returned the (%v) error when expected (%v)", e, ErrConfigNotFound)
		}
	})

	t.Run("invalid requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		name := "primary"
		dialectFactory := NewMockDialectFactory(ctrl)
		connectionFactory := &ConnectionFactory{dialectFactory: dialectFactory}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(nil, expected).Times(1)

		sut, _ := NewConnectionPool(cfgManager, connectionFactory)

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
		connectionFactory := &ConnectionFactory{dialectFactory: dialectFactory}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)

		sut, _ := NewConnectionPool(cfgManager, connectionFactory)

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
		connectionFactory := &ConnectionFactory{dialectFactory: dialectFactory}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)

		sut, _ := NewConnectionPool(cfgManager, connectionFactory)

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
		connectionFactory := &ConnectionFactory{dialectFactory: dialectFactory}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)

		sut, _ := NewConnectionPool(cfgManager, connectionFactory)

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
		connectionFactory := &ConnectionFactory{dialectFactory: dialectFactory}
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfgManager.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfgManager.EXPECT().Config("slate.rdb.connections.primary").Return(&cfg, nil).Times(1)

		sut, _ := NewConnectionPool(cfgManager, connectionFactory)

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
