package slate

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_Error(t *testing.T) {
	t.Run("newError", func(t *testing.T) {
		t.Run("creation without context", func(t *testing.T) {
			msg := "error message"

			if e := NewError(msg); e.Error() != msg {
				t.Errorf("error message (%v) not same as expected (%v)", e, msg)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			msg := "error message"
			context := map[string]interface{}{"field": "value"}

			if e := NewError(msg, context); e.Error() != msg {
				t.Errorf("error message (%v) not same as expected (%v)", e, msg)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})

	t.Run("newErrorFrom", func(t *testing.T) {
		t.Run("creation without context", func(t *testing.T) {
			from := fmt.Errorf("supplier error")
			msg := "error message"
			expected := "error message : supplier error"

			if e := NewErrorFrom(from, msg); e.Error() != expected {
				t.Errorf("error message (%v) not same as expected (%v)", e, expected)
			} else if !errors.Is(e, from) {
				t.Errorf("didn't returned a wrapped error from supplier (%v)", from)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			from := fmt.Errorf("supplier error")
			msg := "error message"
			context := map[string]interface{}{"field": "value"}
			expected := "error message : supplier error"

			if e := NewErrorFrom(from, msg, context); e.Error() != expected {
				t.Errorf("error message (%v) not same as expected (%v)", e, expected)
			} else if !errors.Is(e, from) {
				t.Errorf("didn't returned a wrapped error from supplier (%v)", from)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})

	t.Run("Error", func(t *testing.T) {
		t.Run("retrieve error message from non-wrapped error", func(t *testing.T) {
			msg := "error message"
			if e := NewError(msg); e.Error() != msg {
				t.Errorf("error message (%v) not same as expected (%v)", e, msg)
			}
		})

		t.Run("retrieve error message from wrapped error", func(t *testing.T) {
			from := fmt.Errorf("supplier error")
			msg := "error message"
			expected := "error message : supplier error"
			if e := NewErrorFrom(from, msg); e.Error() != expected {
				t.Errorf("error message (%v) not same as expected (%v)", e, expected)
			}
		})
	})

	t.Run("Unwarp", func(t *testing.T) {
		t.Run("retrieve nil from non-wrapped error", func(t *testing.T) {
			e := NewError("error message")
			if check := errors.Unwrap(e); check != nil {
				t.Errorf("returned the unexpeceted wrapped error : %v", check)
			}
		})

		t.Run("retrieve error message from wrapped error", func(t *testing.T) {
			from := fmt.Errorf("supplier error")
			e := NewErrorFrom(from, "error message")
			if check := errors.Unwrap(e); check == nil {
				t.Errorf("didn't returned the expeceted wrapped error")
			} else if check.Error() != from.Error() {
				t.Errorf("error message from wrapped error (%v) not same as expected (%v)", check, from)
			}
		})
	})

	t.Run("Context", func(t *testing.T) {
		t.Run("retrieve nil from no context error", func(t *testing.T) {
			e := NewError("error message")
			if check := e.(*Error).Context(); check != nil {
				t.Errorf("returned the unexpeceted context : %v", check)
			}
		})

		t.Run("retrieve assigned error context", func(t *testing.T) {
			context := map[string]interface{}{"field": "value"}
			e := NewError("error message", context)
			if check := e.(*Error).Context(); check == nil {
				t.Errorf("didn't returned the expeceted context reference")
			} else if !reflect.DeepEqual(check, context) {
				t.Errorf("returned (%v) context when expecting (%v)", check, context)
			}
		})
	})
}

func Test_err(t *testing.T) {
	t.Run("errNilPointer", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : invalid nil pointer"

		t.Run("creation without context", func(t *testing.T) {
			if e := errNilPointer(arg); !errors.Is(e, ErrNilPointer) {
				t.Errorf("error not a instance of ErrNilPointer")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errNilPointer(arg, context); !errors.Is(e, ErrNilPointer) {
				t.Errorf("error not a instance of ErrNilPointer")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})

	t.Run("errConversion", func(t *testing.T) {
		arg := "dummy value"
		typ := "dummy type"
		context := map[string]interface{}{"field": "value"}
		message := "dummy value to dummy type : invalid type conversion"

		t.Run("creation without context", func(t *testing.T) {
			if e := errConversion(arg, typ); !errors.Is(e, ErrConversion) {
				t.Errorf("error not a instance of ErrConversion")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errConversion(arg, typ, context); !errors.Is(e, ErrConversion) {
				t.Errorf("error not a instance of ErrConversion")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})
}
