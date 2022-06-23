package srlog

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/serror"
	"github.com/happyhippyhippo/slate/slog"
	"testing"
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
		case !container.Has(ContainerOkID):
			t.Errorf("didn't registered the ok status middleware : %v", p)
		case !container.Has(ContainerCreatedID):
			t.Errorf("didn't registered the created status middleware : %v", p)
		case !container.Has(ContainerNoContentID):
			t.Errorf("didn't registered the no content status middleware : %v", p)
		}
	})

	t.Run("error retrieving logger generating ok status middleware", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(slog.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerOkID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("successfully generate ok status middleware without decorators", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = false
		DecorateXML = false
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerOkID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate ok status middleware with json decorator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = true
		DecorateXML = false
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerOkID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate ok status middleware with xml decorator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = false
		DecorateXML = true
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerOkID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate ok status middleware with both json and xml decorators", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = true
		DecorateXML = true
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerOkID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("error retrieving logger generating created status middleware", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(slog.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerCreatedID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("successfully generate created status middleware without decorators", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = false
		DecorateXML = false
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerCreatedID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate created status middleware with json decorator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = true
		DecorateXML = false
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerCreatedID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate created status middleware with xml decorator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = false
		DecorateXML = true
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerCreatedID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate created status middleware with both json and xml decorators", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = true
		DecorateXML = true
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerCreatedID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("error retrieving logger generating no-content status middleware", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = container.Service(slog.ContainerID, func() (interface{}, error) {
			return nil, expected
		})

		if _, err := container.Get(ContainerNoContentID); err == nil {
			t.Error("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("successfully generate no-content status middleware without decorators", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = false
		DecorateXML = false
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerNoContentID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate no-content status middleware with json decorator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = true
		DecorateXML = false
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerNoContentID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate no-content status middleware with xml decorator", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = false
		DecorateXML = true
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerNoContentID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
			}
		}
	})

	t.Run("successfully generate no-content status middleware with both json and xml decorators", func(t *testing.T) {
		container := slate.ServiceContainer{}
		_ = (&Provider{}).Register(container)
		_ = (&slog.Provider{}).Register(container)

		DecorateJSON = true
		DecorateXML = true
		defer func() { DecorateJSON = true; DecorateXML = false }()

		middleware, err := container.Get(ContainerNoContentID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
		case middleware == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch middleware.(type) {
			case func(gin.HandlerFunc) gin.HandlerFunc:
			default:
				t.Error("didn't returned a valid middleware function")
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
