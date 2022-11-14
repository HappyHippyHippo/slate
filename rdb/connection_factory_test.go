package rdb

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	sconfig "github.com/happyhippyhippo/slate/config"
	serror "github.com/happyhippyhippo/slate/error"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func Test_NewConnectionFactory(t *testing.T) {
	t.Run("missing configuration", func(t *testing.T) {
		sut, e := NewConnectionFactory(nil, &DialectFactory{})
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("missing dialect factory", func(t *testing.T) {
		sut, e := NewConnectionFactory(sconfig.NewManager(0), nil)
		switch {
		case sut != nil:
			t.Error("return an unexpected valid connection factory instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("valid creation", func(t *testing.T) {
		if sut, e := NewConnectionFactory(sconfig.NewManager(0), &DialectFactory{}); sut == nil {
			t.Error("didn't returned the expected valid connection factory instance")
		} else if e != nil {
			t.Errorf("return the unexpected error : %v", e)
		}
	})

	t.Run("config change purge all stored connections", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		cfg1 := sconfig.Partial{"dialect": "sqlite", "host": ":memory:"}
		cfg2 := sconfig.Partial{"dialect": "sqlite", "host": ":memory:"}
		partial1 := sconfig.Partial{"slate": sconfig.Partial{"rdb": sconfig.Partial{"connections": sconfig.Partial{name: cfg1}}}}
		partial2 := sconfig.Partial{"slate": sconfig.Partial{"rdb": sconfig.Partial{"connections": sconfig.Partial{name + "salt": cfg2}}}}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(partial2, nil).MinTimes(1)
		dialectFactory := &DialectFactory{}
		_ = dialectFactory.Register(&dialectStrategySqlite{})
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id1", 0, source1)

		sut, _ := NewConnectionFactory(cfg, dialectFactory)

		_, _ = sut.Get(name, &gorm.Config{Logger: gormLogger.Discard})
		if len(sut.(*connectionFactory).instances) != 1 {
			t.Error("didn't store the requested connection instance")
		}

		_ = cfg.AddSource("id2", 10, source2)
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
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("slate.rdb.connections.primary").Return(false).Times(1)

		sut, _ := NewConnectionFactory(cfg, dialectFactory)

		conn, e := sut.Get("primary", &gorm.Config{})
		switch {
		case conn != nil:
			t.Error("return an unexpected valid connection instance")
		case e == nil:
			t.Error("didn't return the expected error")
		case !errors.Is(e, serror.ErrDatabaseConfigNotFound):
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrDatabaseConfigNotFound)
		}
	})

	t.Run("invalid requested connection configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		name := "primary"
		dialectFactory := NewMockDialectFactory(ctrl)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("slate.rdb.connections.primary").Return(nil, expected).Times(1)

		sut, _ := NewConnectionFactory(cfg, dialectFactory)

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
		partial := sconfig.Partial{"dialect": "invalid", "host": ":memory:"}
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&partial).Return(nil, expected).Times(1)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("slate.rdb.connections.primary").Return(partial, nil).Times(1)

		sut, _ := NewConnectionFactory(cfg, dialectFactory)

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
		partial := sconfig.Partial{"dialect": "invalid", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(expected).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&partial).Return(dialect, nil).Times(1)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("slate.rdb.connections.primary").Return(partial, nil).Times(1)

		sut, _ := NewConnectionFactory(cfg, dialectFactory)

		conn, e := sut.Get(name, &gorm.Config{Logger: gormLogger.Discard})
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
		partial := sconfig.Partial{"dialect": "invalid", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&partial).Return(dialect, nil).Times(1)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("slate.rdb.connections.primary").Return(partial, nil).Times(1)

		sut, _ := NewConnectionFactory(cfg, dialectFactory)

		if check, e := sut.Get(name, &gorm.Config{Logger: gormLogger.Discard}); check == nil {
			t.Error("didn't return the expected connection instance")
		} else if e != nil {
			t.Errorf("return the unexpected error : (%v)", e)
		}
	})

	t.Run("multiple requests only instantiate a single connection", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "primary"
		partial := sconfig.Partial{"dialect": "invalid", "host": ":memory:"}
		dialect := NewMockDialect(ctrl)
		dialect.EXPECT().Initialize(gomock.Any()).Return(nil).Times(1)
		dialectFactory := NewMockDialectFactory(ctrl)
		dialectFactory.EXPECT().Get(&partial).Return(dialect, nil).Times(1)
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("slate.rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("slate.rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("slate.rdb.connections.primary").Return(partial, nil).Times(1)

		sut, _ := NewConnectionFactory(cfg, dialectFactory)

		conn, _ := sut.Get(name, &gorm.Config{Logger: gormLogger.Discard})
		check, e := sut.Get(name, &gorm.Config{Logger: gormLogger.Discard})
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
