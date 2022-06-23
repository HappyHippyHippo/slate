package grdb

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"testing"
)

func Test_NewConnectionFactory(t *testing.T) {
	t.Run("missing configuration", func(t *testing.T) {
		factory, err := NewConnectionFactory(nil, &DialectFactory{})
		switch {
		case factory != nil:
			t.Error("return an unexpected valid connection factory instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("missing dialect factory", func(t *testing.T) {
		factory, err := NewConnectionFactory(gconfig.NewConfig(0), nil)
		switch {
		case factory != nil:
			t.Error("return an unexpected valid connection factory instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("valid creation", func(t *testing.T) {
		if factory, err := NewConnectionFactory(gconfig.NewConfig(0), &DialectFactory{}); factory == nil {
			t.Error("didn't returned the expected valid connection factory instance")
		} else if err != nil {
			t.Errorf("return the unexpected error : %v", err)
		}
	})

	t.Run("config change purge all stored connections", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		partial1 := gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					name: gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id1", 0, source1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		connFactory, _ := NewConnectionFactory(cfg, dialectFactory)

		_, _ = connFactory.Get(name, &gorm.Config{Logger: gormLogger.Discard})

		if len(connFactory.instances) != 1 {
			t.Error("didn't store the requested connection instance")
		}

		partial2 := gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					name + "salt": gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(partial2, nil).MinTimes(1)
		_ = cfg.AddSource("id2", 10, source2)

		if len(connFactory.instances) != 0 {
			t.Error("didn't removed the stored connection instances")
		}
	})
}

func Test_ConnectionFactory_Get(t *testing.T) {
	t.Run("missing requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		partial1 := gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					name: gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id1", 0, source1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		connFactory, _ := NewConnectionFactory(cfg, dialectFactory)

		conn, err := connFactory.Get(name+"salt", &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrDatabaseConfigNotFound):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrDatabaseConfigNotFound)
		}
	})

	t.Run("invalid requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		partial1 := gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					name: "string",
				},
			},
		}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id1", 0, source1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		connFactory, _ := NewConnectionFactory(cfg, dialectFactory)

		conn, err := connFactory.Get(name, &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrConversion)
		}
	})

	t.Run("error instantiating dialect", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		partial1 := gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					name: gconfig.Partial{"dialect": "invalid", "host": ":memory:"},
				},
			},
		}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id1", 0, source1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		connFactory, _ := NewConnectionFactory(cfg, dialectFactory)

		conn, err := connFactory.Get(name, &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case !errors.Is(err, gerror.ErrUnknownDatabaseDialect):
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrUnknownDatabaseDialect)
		}
	})

	t.Run("error instantiating connector", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("unable to open database file: no such file or directory")
		name := "primary"
		partial1 := gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					name: gconfig.Partial{"dialect": "sqlite", "host": "//////invalid"},
				},
			},
		}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id1", 0, source1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		connFactory, _ := NewConnectionFactory(cfg, dialectFactory)

		conn, err := connFactory.Get(name, &gorm.Config{Logger: gormLogger.Discard})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case err == nil:
			t.Error("didn't return an expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expected (%v)", err, expected)
		}
	})

	t.Run("valid connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		partial1 := gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					name: gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id1", 0, source1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		connFactory, _ := NewConnectionFactory(cfg, dialectFactory)

		if check, err := connFactory.Get(name, &gorm.Config{Logger: gormLogger.Discard}); check == nil {
			t.Error("didn't return the expected connection instance")
		} else if err != nil {
			t.Errorf("return the unexpected error : (%v)", err)
		}
	})

	t.Run("multiple requests only instantiate a single connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		partial1 := gconfig.Partial{
			"rdb": gconfig.Partial{
				"connections": gconfig.Partial{
					name: gconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		cfg := gconfig.NewConfig(0)
		_ = cfg.AddSource("id1", 0, source1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&DialectStrategySqlite{})
		connFactory, _ := NewConnectionFactory(cfg, dialectFactory)

		conn, _ := connFactory.Get(name, &gorm.Config{Logger: gormLogger.Discard})
		check, err := connFactory.Get(name, &gorm.Config{Logger: gormLogger.Discard})
		switch {
		case check == nil:
			t.Error("didn't return the expected connection instance")
		case err != nil:
			t.Errorf("return the unexpected error : (%v)", err)
		case check != conn:
			t.Error("didn't returned the same instance")
		}
	})
}
