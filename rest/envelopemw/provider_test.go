package envelopemw

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case !container.Has(ID):
			t.Errorf("no generator : %v", sut)
		}
	})

	t.Run("error retrieving config manager when retrieving the generator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(config.ID, func() (*config.Config, error) {
			return nil, expected
		})

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("error retrieving logger when retrieving the generator", func(t *testing.T) {
		expected := fmt.Errorf("error message")
		container := slate.NewContainer()

		sut := &Provider{}
		_ = sut.Register(container)
		_ = container.Service(log.ID, func() (*log.Log, error) {
			return nil, expected
		})

		if _, e := container.Get(ID); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrContainer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrContainer)
		}
	})

	t.Run("retrieving generator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&log.Provider{}).Register(container)
		_ = (&config.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.accept", []interface{}{"application/json"})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		_ = container.Service(config.ID, func() *config.Config {
			return cfg
		})

		sut, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a reference to the generator")
		default:
			switch sut.(type) {
			case MiddlewareGenerator:
			default:
				t.Error("didn't returned a generator reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("successful boot", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Provide(Provider{})

		if e := app.Boot(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
