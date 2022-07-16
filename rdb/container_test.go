package rdb

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/fs"
	"testing"
)

func Test_GetConfig(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		sut, e := GetConfig(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found err")
		}
	})

	t.Run("non gorm cfg instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerConfigID, func() (any, error) {
			return "string", nil
		})

		sut, e := GetConfig(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned the error is not of the expected a conversion err")
		}
	})

	t.Run("valid gorm cfg instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		sut, e := GetConfig(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetDialectFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		sut, e := GetDialectFactory(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found err")
		}
	})

	t.Run("non dialect factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerDialectFactoryID, func() (any, error) {
			return "string", nil
		})

		sut, e := GetDialectFactory(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned the error is not of the expected a conversion err")
		}
	})

	t.Run("valid dialect factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		sut, e := GetDialectFactory(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetDialectStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		e := fmt.Errorf("dummy message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, e
		}, ContainerDialectStrategyTag)

		sut, e := GetDialectStrategies(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, e):
			t.Error("returned the error is not of the expected err")
		}
	})

	t.Run("non dialect strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerDialectStrategyTag)

		sut, e := GetDialectStrategies(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned the error is not of the expected err")
		}
	})

	t.Run("valid dialect strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		sut, e := GetDialectStrategies(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetConnectionFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		sut, e := GetConnectionFactory(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found err")
		}
	})

	t.Run("non connection factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		sut, e := GetConnectionFactory(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned the error is not of the expected a conversion err")
		}
	})

	t.Run("valid connection factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&fs.Provider{}).Register(c)
		_ = (&config.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		sut, e := GetConnectionFactory(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetPrimaryConnection(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		sut, e := GetPrimaryConnection(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found err")
		}
	})

	t.Run("non connection instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&fs.Provider{}).Register(c)
		_ = (&config.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)
		_ = c.Service(ContainerPrimaryID, func() (any, error) {
			return "string", nil
		})

		sut, e := GetPrimaryConnection(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected err")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned the error is not of the expected a conversion err")
		}
	})

	t.Run("valid connection instance returned", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		c := slate.ServiceContainer{}
		_ = (&fs.Provider{}).Register(c)
		_ = (&config.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)
		_ = (&Provider{}).Boot(c)

		partial := config.Partial{"dialect": "sqlite", "host": ":memory:"}
		cfg := NewMockConfigManager(ctrl)
		cfg.EXPECT().AddObserver("rdb.connections", gomock.Any()).Return(nil).Times(1)
		cfg.EXPECT().Has("rdb.connections.primary").Return(true).Times(1)
		cfg.EXPECT().Partial("rdb.connections.primary").Return(partial, nil).Times(1)
		_ = c.Service(config.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		sut, e := GetPrimaryConnection(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}
