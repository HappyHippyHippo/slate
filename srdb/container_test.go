package srdb

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/happyhippyhippo/slate/sfs"
	"testing"
)

func Test_GetConfig(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetConfig(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non gorm config instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerConfigID, func() (any, error) {
			return "string", nil
		})

		s, err := GetConfig(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid gorm config instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetConfig(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetDialectFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetDialectFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non dialect factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerDialectFactoryID, func() (any, error) {
			return "string", nil
		})

		s, err := GetDialectFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid dialect factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetDialectFactory(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
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

		s, err := GetDialectStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, e):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non dialect strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerDialectStrategyTag)

		s, err := GetDialectStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid dialect strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetDialectStrategies(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetConnectionFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetConnectionFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non connection factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		s, err := GetConnectionFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid connection factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&sfs.Provider{}).Register(c)
		_ = (&sconfig.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		s, err := GetConnectionFactory(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetPrimaryConnection(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetPrimaryConnection(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non connection instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&sfs.Provider{}).Register(c)
		_ = (&sconfig.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)
		_ = c.Service(ContainerPrimaryID, func() (any, error) {
			return "string", nil
		})

		s, err := GetPrimaryConnection(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid connection instance returned", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		c := slate.ServiceContainer{}
		_ = (&sfs.Provider{}).Register(c)
		_ = (&sconfig.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)
		_ = (&Provider{}).Boot(c)

		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(
			sconfig.Partial{
				"rdb": sconfig.Partial{
					"connections": sconfig.Partial{
						"primary": sconfig.Partial{"dialect": "sqlite", "host": ":memory:"},
					},
				},
			},
			nil,
		).Times(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		_ = c.Service(sconfig.ContainerID, func() (interface{}, error) {
			return cfg, nil
		})

		s, err := GetPrimaryConnection(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
