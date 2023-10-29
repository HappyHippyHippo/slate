package slate

import (
	"errors"
	"testing"

	"github.com/spf13/afero"
)

func Test_FileSystem(t *testing.T) {
	t.Run("NewFileSystemServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if NewFileSystemServiceRegister(nil) == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewFileSystemServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil Provider", func(t *testing.T) {
			if e := NewFileSystemServiceRegister(nil).Provide(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("register the file system", func(t *testing.T) {
			app := NewApp()
			_ = app.Provide(NewFileSystemServiceRegister(app))

			system, e := app.Get(FileSystemContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case system == nil:
				t.Error("didn't returned the expected instance")
			default:
				switch system.(type) {
				case *afero.OsFs:
				default:
					t.Error("didn't returned the file system")
				}
			}
		})
	})

	t.Run("Boot", func(t *testing.T) {
		t.Run("nil Provider", func(t *testing.T) {
			if e := NewFileSystemServiceRegister(nil).Boot(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("successful boot", func(t *testing.T) {
			app := NewApp()
			_ = app.Provide(NewFileSystemServiceRegister(app))

			if e := app.Boot(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})
}
