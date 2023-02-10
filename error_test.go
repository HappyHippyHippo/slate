package slate

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_NewError(t *testing.T) {
	t.Run("creation without context", func(t *testing.T) {
		msg := "error message"

		if e := NewError(msg); e.Error() != msg {
			t.Errorf("error message (%v) not same as expected (%v)", e, msg)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		msg := "error message"
		context := map[string]interface{}{"field": "value"}

		if e := NewError(msg, context); e.Error() != msg {
			t.Errorf("error message (%v) not same as expected (%v)", e, msg)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_NewErrorFrom(t *testing.T) {
	t.Run("creation without context", func(t *testing.T) {
		from := fmt.Errorf("source error")
		msg := "error message"
		expected := "error message : source error"

		if e := NewErrorFrom(from, msg); e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		} else if !errors.Is(e, from) {
			t.Errorf("didn't returned a wrapped error from source (%v)", from)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		from := fmt.Errorf("source error")
		msg := "error message"
		context := map[string]interface{}{"field": "value"}
		expected := "error message : source error"

		if e := NewErrorFrom(from, msg, context); e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		} else if !errors.Is(e, from) {
			t.Errorf("didn't returned a wrapped error from source (%v)", from)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_Error_Error(t *testing.T) {
	t.Run("retrieve error message from non-wrapped error", func(t *testing.T) {
		msg := "error message"
		if e := NewError(msg); e.Error() != msg {
			t.Errorf("error message (%v) not same as expected (%v)", e, msg)
		}
	})

	t.Run("retrieve error message from wrapped error", func(t *testing.T) {
		from := fmt.Errorf("source error")
		msg := "error message"
		expected := "error message : source error"
		if e := NewErrorFrom(from, msg); e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_Error_Unwrap(t *testing.T) {
	t.Run("retrieve nil from non-wrapped error", func(t *testing.T) {
		e := NewError("error message")
		if check := errors.Unwrap(e); check != nil {
			t.Errorf("returned the unexpeceted wrapped error : %v", check)
		}
	})

	t.Run("retrieve error message from wrapped error", func(t *testing.T) {
		from := fmt.Errorf("source error")
		e := NewErrorFrom(from, "error message")
		if check := errors.Unwrap(e); check == nil {
			t.Errorf("didn't returned the expeceted wrapped error")
		} else if check.Error() != from.Error() {
			t.Errorf("error message from wrapped error (%v) not same as expected (%v)", check, from)
		}
	})
}

func Test_Error_Context(t *testing.T) {
	t.Run("retrieve nil from no context error", func(t *testing.T) {
		e := NewError("error message")
		if check := e.(IError).Context(); check != nil {
			t.Errorf("returned the unexpeceted context : %v", check)
		}
	})

	t.Run("retrieve assigned error context", func(t *testing.T) {
		context := map[string]interface{}{"field": "value"}
		e := NewError("error message", context)
		if check := e.(IError).Context(); check == nil {
			t.Errorf("didn't returned the expeceted context reference")
		} else if !reflect.DeepEqual(check, context) {
			t.Errorf("returned (%v) context when expecting (%v)", check, context)
		}
	})
}
