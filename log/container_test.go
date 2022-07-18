package log

import (
	"errors"
	"fmt"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/fs"
	"testing"
)

func Test_GetFormatterFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		sut, e := GetFormatterFactory(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non formatter factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerFormatterFactoryID, func() (any, error) {
			return "string", nil
		})

		sut, e := GetFormatterFactory(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid formatter factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		sut, e := GetFormatterFactory(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetFormatterStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		expected := fmt.Errorf("dummy message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, expected
		}, ContainerFormatterStrategyTag)

		sut, e := GetFormatterStrategies(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, expected):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non formatter strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerFormatterStrategyTag)

		sut, e := GetFormatterStrategies(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid formatter strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		sut, e := GetFormatterStrategies(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetStreamFactory(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		sut, e := GetStreamFactory(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non stream factory instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerStreamFactoryID, func() (any, error) {
			return "string", nil
		})

		sut, e := GetStreamFactory(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid stream factory instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		sut, e := GetStreamFactory(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetStreamStrategies(t *testing.T) {
	t.Run("tagged retrieval error", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return nil, expected
		}, ContainerStreamStrategyTag)

		sut, e := GetStreamStrategies(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, expected):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("non stream strategy tagged service", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service("dummy", func() (any, error) {
			return "string", nil
		}, ContainerStreamStrategyTag)

		sut, e := GetStreamStrategies(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned the error is not of the expected error")
		}
	})

	t.Run("valid stream strategy list returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&fs.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		sut, e := GetStreamStrategies(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetLogger(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		sut, e := Get(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non logger instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerID, func() (any, error) {
			return "string", nil
		})

		sut, e := Get(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid logger instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&Provider{}).Register(c)

		sut, e := Get(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}

func Test_GetLoader(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		sut, e := GetLoader(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non loader instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerLoaderID, func() (any, error) {
			return "string", nil
		})

		sut, e := GetLoader(c)
		switch {
		case sut != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case e == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(e, err.ErrConversion):
			t.Error("returned error is not of the expected conversion error")
		}
	})

	t.Run("valid loader instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&fs.Provider{}).Register(c)
		_ = (&config.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		sut, e := GetLoader(c)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}
