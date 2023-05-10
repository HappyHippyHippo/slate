package dir

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/fs"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		sut := &Provider{}
		_ = sut.Register(nil)

		if e := sut.Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
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
			t.Errorf("didn't registered the strategy : %v", sut)
		}
	})

	t.Run("retrieving strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&fs.Provider{}).Register(container)
		_ = (&config.Provider{}).Register(container)
		_ = (&Provider{}).Register(container)

		strategy, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case strategy == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch strategy.(type) {
			case *SourceStrategy:
			default:
				t.Error("didn't return a strategy reference")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("run boot", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
