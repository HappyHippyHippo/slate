package trigger

import (
	"errors"
	"testing"

	serror "github.com/happyhippyhippo/slate/error"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, serror.ErrNilPointer) {
			t.Errorf("error not a instance of NilPointer")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}