package rest

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate"
	serror "github.com/happyhippyhippo/slate/error"
	slog "github.com/happyhippyhippo/slate/log"
)

func Test_GetEngine(t *testing.T) {
	t.Run("not registered service", func(t *testing.T) {
		c := slate.ServiceContainer{}

		s, err := GetEngine(c)
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
		_ = c.Service(ContainerEngineID, func() (any, error) {
			return "string", nil
		})

		s, err := GetEngine(c)
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
		_ = c.Service(ContainerEngineID, func() (any, error) {
			return gin.New(), nil
		})

		s, err := GetEngine(c)
		switch {
		case s == nil:
			t.Error("didn't returned the expected valid instance of a service")
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
