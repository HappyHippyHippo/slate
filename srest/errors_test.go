package srest

import (
	"errors"
	"github.com/happyhippyhippo/slate/serror"
	"testing"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if err := errNilPointer(arg); !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("error not a instance of ErrNilPointer")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrConversion(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		typ := "dummy type"
		expected := "invalid type conversion : dummy value to dummy type"

		if err := errConversion(arg, typ); !errors.Is(err, serror.ErrConversion) {
			t.Errorf("error not a instance of ErrConversion")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}
