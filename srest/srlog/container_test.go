package srlog

import (
	"errors"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/happyhippyhippo/slate/slog"
	"testing"
)

func Test_GetMiddlewareOk(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetMiddlewareOk(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non middleware instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerOkID, func() (any, error) {
			return "string", nil
		})

		s, err := GetMiddlewareOk(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid middleware instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&slog.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		s, err := GetMiddlewareOk(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetMiddlewareCreated(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetMiddlewareCreated(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non middleware instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerCreatedID, func() (any, error) {
			return "string", nil
		})

		s, err := GetMiddlewareCreated(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid middleware instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&slog.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		s, err := GetMiddlewareCreated(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}

func Test_GetMiddlewareNoContent(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetMiddlewareNoContent(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrServiceNotFound):
			t.Error("returned the error is not of the expected a service not found error")
		}
	})

	t.Run("non middleware instance", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = c.Service(ContainerNoContentID, func() (any, error) {
			return "string", nil
		})

		s, err := GetMiddlewareNoContent(c)
		switch {
		case s != nil:
			t.Error("returned an unexpectedly valid instance of a service")
		case err == nil:
			t.Error("didn't returned an expected error")
		case !errors.Is(err, serror.ErrConversion):
			t.Error("returned the error is not of the expected a conversion error")
		}
	})

	t.Run("valid middleware instance returned", func(t *testing.T) {
		c := slate.ServiceContainer{}
		_ = (&slog.Provider{}).Register(c)
		_ = (&Provider{}).Register(c)

		s, err := GetMiddlewareNoContent(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
