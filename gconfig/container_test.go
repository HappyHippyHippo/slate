package gconfig

import (
	"errors"
	"fmt"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/gerror"
	"github.com/happyhippyhippo/slate/gfs"
	"testing"
)

func Test_GetDecoderFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetDecoderFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non decoder factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerDecoderFactoryID, func() (any, error) {
			return "string", nil
		})

		s, err := GetDecoderFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid decoder factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetDecoderFactory(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetDecoderStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		e := fmt.Errorf("dummy message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, e
		}, ContainerDecoderStrategyTag)

		s, err := GetDecoderStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, e):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non decoder strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerDecoderStrategyTag)

		s, err := GetDecoderStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid decoder strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetDecoderStrategies(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetSourceFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetSourceFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non source factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerSourceFactoryID, func() (any, error) {
			return "string", nil
		})

		s, err := GetSourceFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid decoder factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetSourceFactory(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetSourceStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		e := fmt.Errorf("dummy message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, e
		}, ContainerSourceStrategyTag)

		s, err := GetSourceStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, e):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non source strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerSourceStrategyTag)

		s, err := GetSourceStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid source strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&gfs.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		s, err := GetSourceStrategies(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetConfig(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetConfig(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non config instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		s, err := GetConfig(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid decoder factory instance returned", func(t *testing.T) {
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

func Test_GetLoader(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetLoader(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non config instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerLoaderID, func() (any, error) {
			return "string", nil
		})

		s, err := GetLoader(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, gerror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid decoder factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, err := GetLoader(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
