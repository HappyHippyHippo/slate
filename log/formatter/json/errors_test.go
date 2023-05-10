package json

import (
	"errors"
	"reflect"
	"testing"

	"github.com/happyhippyhippo/slate"
)

func Test_errNilPointer(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : invalid nil pointer"

	t.Run("creation without context", func(t *testing.T) {
		if e := errNilPointer(arg); !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("error not a instance of slate.ErrNilPointer")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errNilPointer(arg, context); !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("error not a instance of slate.ErrNilPointer")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}
