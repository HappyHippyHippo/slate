package gmigration

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
	"github.com/happyhippyhippo/slate/grdb"
	"github.com/pkg/errors"
	"testing"
)

func Test_GetDao(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetDao(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non dao instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerDaoID, func() (any, error) {
			return "string", nil
		})

		s, err := GetDao(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid dao instance returned", func(t *testing.T) {
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

		s, err := GetDao(container)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetMigrator(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetMigrator(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non migrator instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		s, err := GetMigrator(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid dao instance returned", func(t *testing.T) {
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

		s, err := GetDao(container)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetMigrations(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		e := fmt.Errorf("dummy message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, e
		}, ContainerMigrationTag)

		s, err := GetMigrations(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, e):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non migration tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerMigrationTag)

		s, err := GetMigrations(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid migration list returned", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)
		_ = c.Service("dummy", func() (any, error) {
			return NewMockMigration(ctrl), nil
		}, ContainerMigrationTag)

		s, err := GetMigrations(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
