package fs

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate"
	serror "github.com/happyhippyhippo/slate/error"
	"github.com/spf13/afero"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("register the file system", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Add(Provider{})

		system, e := app.Container.Get(ContainerID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error (%v)", e)
		case system == nil:
			t.Error("didn't returned the expected instance")
		default:
			switch system.(type) {
			case *afero.OsFs:
			default:
				t.Error("didn't returned the file system form the container")
			}
		}
	})
}

func Test_Provider_Boot(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (Provider{}).Boot(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("successful boot", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Add(Provider{})

		if e := app.Boot(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
