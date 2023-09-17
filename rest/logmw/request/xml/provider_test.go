package xml

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate/rest/logmw"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
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
			t.Errorf("no decoator : %v", sut)
		}
	})

	t.Run("retrieving decorator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		container := slate.NewContainer()
		_ = (&Provider{}).Register(container)

		sut, e := container.Get(ID)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Error("didn't returned a reference to the REST engine")
		default:
			switch sut.(type) {
			case logmw.RequestReader:
			default:
				t.Error("didn't returned the decorator")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		container := slate.NewContainer()
		sut := &Provider{}
		_ = sut.Register(container)

		if e := sut.Boot(nil); e == nil {
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
