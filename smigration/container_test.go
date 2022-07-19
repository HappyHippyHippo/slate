package smigration

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serr"
	"github.com/happyhippyhippo/slate/srdb"
	"testing"
)

func Test_GetDao(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, e := GetDao(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serr.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non dao instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerDaoID, func() (any, error) {
			return "string", nil
		})

		s, e := GetDao(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid dao instance returned", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&sconfig.Provider{}).Register(container)
		_ = (&srdb.Provider{}).Register(container)
		_ = (&srdb.Provider{}).Boot(container)
		p := &Provider{}
		_ = p.Register(container)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(sconfig.Partial{
			"rdb": sconfig.Partial{
				"connections": sconfig.Partial{
					"primary": sconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}, nil,
		).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		s, e := GetDao(container)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetMigrator(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, e := GetMigrator(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serr.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non migrator instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		s, e := GetMigrator(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serr.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid dao instance returned", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.ServiceContainer{}
		_ = (&sconfig.Provider{}).Register(container)
		_ = (&srdb.Provider{}).Register(container)
		_ = (&srdb.Provider{}).Boot(container)
		p := &Provider{}
		_ = p.Register(container)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(sconfig.Partial{
			"rdb": sconfig.Partial{
				"connections": sconfig.Partial{
					"primary": sconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
				},
			},
		}, nil,
		).Times(1)
		cfg := sconfig.NewManager(0)
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		s, e := GetDao(container)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
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

		s, e := GetMigrations(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, e):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non migration tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerMigrationTag)

		s, e := GetMigrations(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serr.ErrConversion):
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

		s, e := GetMigrations(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}
