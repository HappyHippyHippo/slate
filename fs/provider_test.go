package fs

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/err"
	"github.com/spf13/afero"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if e := (Provider{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) err when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("register the file system", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Provider(Provider{})

		system, e := app.Get(ID)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected err (%v)", e)
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
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) err when expected (%v)", e, err.NilPointer)
		}
	})

	t.Run("successful boot", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Provider(Provider{})

		if e := app.Boot(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}
