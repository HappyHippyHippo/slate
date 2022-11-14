package config

import (
	"errors"
	"fmt"
	"testing"

	"github.com/happyhippyhippo/slate"
	serror "github.com/happyhippyhippo/slate/error"
	sfs "github.com/happyhippyhippo/slate/fs"
)

func Test_GetDecoderFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, e := GetDecoderFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non decoder factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerDecoderFactoryID, func() (any, error) {
			return "string", nil
		})

		s, e := GetDecoderFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid decoder factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, e := GetDecoderFactory(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetDecoderStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, expected
		}, ContainerDecoderStrategyTag)

		s, e := GetDecoderStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, expected):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non decoder strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerDecoderStrategyTag)

		s, e := GetDecoderStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid decoder strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, e := GetDecoderStrategies(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetSourceFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, e := GetSourceFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non source factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerSourceFactoryID, func() (any, error) {
			return "string", nil
		})

		s, e := GetSourceFactory(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid decoder factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, e := GetSourceFactory(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetSourceStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, expected
		}, ContainerSourceStrategyTag)

		s, e := GetSourceStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, expected):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non source strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerSourceStrategyTag)

		s, e := GetSourceStrategies(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid source strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&sfs.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		s, e := GetSourceStrategies(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetSourceContainerPartials(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, expected
		}, ContainerSourceContainerPartialTag)

		s, e := GetSourceContainerPartials(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, expected):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non partial tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerSourceContainerPartialTag)

		s, e := GetSourceContainerPartials(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid config list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)
		_ = c.Service("dummy", func() (any, error) {
			return &Partial{}, nil
		}, ContainerSourceContainerPartialTag)

		s, e := GetSourceContainerPartials(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_Get(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, e := Get(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non config instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		s, e := Get(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid decoder factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, e := Get(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetLoader(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, e := GetLoader(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non config instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerLoaderID, func() (any, error) {
			return "string", nil
		})

		s, e := GetLoader(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, serror.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid decoder factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		s, e := GetLoader(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}
