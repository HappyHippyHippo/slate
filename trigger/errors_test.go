package trigger

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate/err"
)

func Test_errNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, err.NilPointer) {
			t.Errorf("error not a instance of err.NilPointer")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}
