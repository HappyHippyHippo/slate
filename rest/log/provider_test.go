package log

import (
	"errors"
	"fmt"
	sconfig "github.com/happyhippyhippo/slate/config"
	"testing"

	"github.com/happyhippyhippo/slate"
	serror "github.com/happyhippyhippo/slate/error"
	slog "github.com/happyhippyhippo/slate/log"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		p := &Provider{}

		if err := p.Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.ServiceContainer{}
		p := &Provider{}

		err := p.Register(container)
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case !container.Has(ContainerID):
		}
	})

	t.Run("error retrieving config requesting middleware generator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(sconfig.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("error retrieving logger requesting middleware generator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&sconfig.Provider{}).Register(container)
		_ = container.Service(slog.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("successfully generate middleware generator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&sconfig.Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = false
		DecorateXML = false
		defer func() { DecorateJSON = true; DecorateXML = false }()

		generator, err := container.Get(ContainerID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case generator == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch generator.(type) {
			case MiddlewareGenerator:
			default:
				t.Error("didn't returned a valid middleware generator function")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if err := (Provider{}).Boot(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("successful boot", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Add(Provider{})

		if err := app.Boot(); err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})
}
