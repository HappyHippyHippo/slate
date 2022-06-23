package gfs

import (
	"errors"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/gerror"
	"github.com/spf13/afero"
	"testing"
)

func Test_Provider_Register(t *testing.T) {
	t.Run("nil container", func(t *testing.T) {
		if err := (Provider{}).Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("register the file system", func(t *testing.T) {
		app := slate.NewApplication()
		_ = app.Add(Provider{})

		system, err := app.Container.Get(ContainerID)
		switch {
		case err != nil:
			t.Errorf("returned the unexpected error (%v)", err)
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
		if err := (Provider{}).Boot(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expected (%v)", err, gerror.ErrNilPointer)
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
