package fs

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate"
	serror "github.com/happyhippyhippo/slate/error"
)

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
			t.Error("returned error is not of the expected a service not found error")
		}
	})

	t.Run("non afero file system adapter instance", func(t *testing.T) {
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

	t.Run("valid afero file system instance returned", func(t *testing.T) {
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
