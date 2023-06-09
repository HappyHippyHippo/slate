package rdb

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_NewConnectionPool(t *testing.T) {
	t.Run("missing configuration", func(t *testing.T) {
		sut, e := NewConnectionPool(nil, &ConnectionFactory{})
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("missing connection factory", func(t *testing.T) {
		sut, e := NewConnectionPool(config.NewConfig(), nil)
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("valid creation", func(t *testing.T) {
		connectionsFactory, _ := NewConnectionFactory(NewDialectFactory())

		if sut, e := NewConnectionPool(config.NewConfig(), connectionsFactory); sut == nil {
			t.Error("didn't returned the expected valid connection factory ")
		} else if e != nil {
			t.Errorf("return the unexpected error : %v", e)
		}
	})

	t.Run("config change purge all stored connections", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		cfg1 := config.Partial{"dialect": "sqlite", "host": ":memory1:"}
		cfg2 := config.Partial{"dialect": "sqlite", "host": ":memory2:"}
		gormCfg := &gorm.Config{Logger: logger.Discard, DisableAutomaticPing: true}
		config1 := config.Partial{}
		_, _ = config1.Set("slate.rdb.connections", config.Partial{name: cfg1})
		config2 := config.Partial{}
		_, _ = config2.Set("slate.rdb.connections", config.Partial{name + "salt": cfg2})
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(config1, nil).MinTimes(1)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(config2, nil).MinTimes(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source1)

		conn, _ := gorm.Open(sqlite.Open(":memory:"), nil)
		connectionCreator := NewMockConnectionCreator(ctrl)
		connectionCreator.EXPECT().Create(cfg1, gormCfg).Return(conn, nil).Times(1)

		cf, _ := NewConnectionFactory(NewDialectFactory())
		sut, _ := NewConnectionPool(cfg, cf)
		sut.connectionCreator = connectionCreator

		_, _ = sut.Get(name, gormCfg)
		if len(sut.instances) != 1 {
			t.Error("didn't store the requested connection instance")
		}

		_ = cfg.AddSource("id2", 10, source2)
		if len(sut.instances) != 0 {
			t.Error("didn't removed the stored connection instances")
		}
	})
}

func Test_ConnectionPool_Get(t *testing.T) {
	t.Run("missing requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Has("slate.rdb.connections.primary").
			Return(false).
			Times(1)
		connectionCreator := NewMockConnectionCreator(ctrl)

		cf, _ := NewConnectionFactory(NewDialectFactory())
		sut, _ := NewConnectionPool(config.NewConfig(), cf)
		sut.connectionCreator = connectionCreator
		sut.configurer = configurer

		conn, e := sut.Get("primary", &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, ErrConfigNotFound):
			t.Errorf("(%v) when expected (%v)", e, ErrConfigNotFound)
		}
	})

	t.Run("invalid requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Has("slate.rdb.connections.name").
			Return(true).
			Times(1)
		configurer.
			EXPECT().
			Partial("slate.rdb.connections.name").
			Return(nil, expected).
			Times(1)
		connectionCreator := NewMockConnectionCreator(ctrl)

		cf, _ := NewConnectionFactory(NewDialectFactory())
		sut, _ := NewConnectionPool(config.NewConfig(), cf)
		sut.connectionCreator = connectionCreator
		sut.configurer = configurer

		conn, e := sut.Get("name", &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expected (%v)", e, expected)
		}
	})

	t.Run("error instantiating connector", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		partial := config.Partial{"data": "string"}
		gormCfg := &gorm.Config{Logger: logger.Discard}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Has("slate.rdb.connections.name").
			Return(true).
			Times(1)
		configurer.
			EXPECT().
			Partial("slate.rdb.connections.name").
			Return(partial, nil).
			Times(1)
		connectionCreator := NewMockConnectionCreator(ctrl)
		connectionCreator.
			EXPECT().
			Create(partial, gormCfg).
			Return(nil, expected).
			Times(1)

		cf, _ := NewConnectionFactory(NewDialectFactory())
		sut, _ := NewConnectionPool(config.NewConfig(), cf)
		sut.connectionCreator = connectionCreator
		sut.configurer = configurer

		conn, e := sut.Get("name", gormCfg)
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expected (%v)", e, expected)
		}
	})

	t.Run("valid connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"data": "string"}
		gormCfg := &gorm.Config{Logger: logger.Discard}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Has("slate.rdb.connections.name").
			Return(true).
			Times(1)
		configurer.
			EXPECT().
			Partial("slate.rdb.connections.name").
			Return(partial, nil).
			Times(1)
		connectionCreator := NewMockConnectionCreator(ctrl)
		connectionCreator.
			EXPECT().
			Create(partial, gormCfg).
			Return(&gorm.DB{}, nil).
			Times(1)

		cf, _ := NewConnectionFactory(NewDialectFactory())
		sut, _ := NewConnectionPool(config.NewConfig(), cf)
		sut.connectionCreator = connectionCreator
		sut.configurer = configurer

		if check, e := sut.Get("name", gormCfg); check == nil {
			t.Error("didn't return the expected connection instance")
		} else if e != nil {
			t.Errorf("return the unexpected error : (%v)", e)
		}
	})

	t.Run("multiple requests only instantiate a single connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"data": "string"}
		gormCfg := &gorm.Config{Logger: logger.Discard}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Has("slate.rdb.connections.name").
			Return(true).
			Times(1)
		configurer.
			EXPECT().
			Partial("slate.rdb.connections.name").
			Return(partial, nil).
			Times(1)
		connectionCreator := NewMockConnectionCreator(ctrl)
		connectionCreator.
			EXPECT().
			Create(partial, gormCfg).
			Return(&gorm.DB{}, nil).
			Times(1)

		cf, _ := NewConnectionFactory(NewDialectFactory())
		sut, _ := NewConnectionPool(config.NewConfig(), cf)
		sut.connectionCreator = connectionCreator
		sut.configurer = configurer

		conn, _ := sut.Get("name", gormCfg)
		check, e := sut.Get("name", gormCfg)
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
