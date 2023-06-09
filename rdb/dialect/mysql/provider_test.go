//go:build mysql

package mysql

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (&Provider{}).Register(nil); e == nil {
			t.Error("didn't return the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expected (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register components", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}

		e := sut.Register(container)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !container.Has(ID):
			t.Errorf("no strategy : %v", sut)
		}
	})

	t.Run("retrieving strategy", func(t *testing.T) {
		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		service, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("unexpected error (%v)", e)
		case service == nil:
			t.Error("didn't returned a valid reference")
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

	t.Run("run boot", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(container); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})
}
