package slate

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_errNilPointer(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : invalid nil pointer"

	t.Run("creation without context", func(t *testing.T) {
		if e := errNilPointer(arg); !errors.Is(e, ErrNilPointer) {
			t.Errorf("error not a instance of ErrNilPointer")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errNilPointer(arg, context); !errors.Is(e, ErrNilPointer) {
			t.Errorf("error not a instance of ErrNilPointer")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errConversion(t *testing.T) {
	arg := "dummy value"
	typ := "dummy type"
	context := map[string]interface{}{"field": "value"}
	message := "dummy value to dummy type : invalid type conversion"

	t.Run("creation without context", func(t *testing.T) {
		if e := errConversion(arg, typ); !errors.Is(e, ErrConversion) {
			t.Errorf("error not a instance of ErrConversion")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errConversion(arg, typ, context); !errors.Is(e, ErrConversion) {
			t.Errorf("error not a instance of ErrConversion")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errContainer(t *testing.T) {
	arg := fmt.Errorf("dummy argument 1")
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument 1 : service container error"

	t.Run("creation without context", func(t *testing.T) {
		if e := errContainer(arg); !errors.Is(e, ErrContainer) {
			t.Errorf("error not a instance of ErrContainer")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errContainer(arg, context); !errors.Is(e, ErrContainer) {
			t.Errorf("error not a instance of ErrContainer")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errNonFunctionFactory(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : non-function service factory"

	t.Run("creation without context", func(t *testing.T) {
		if e := errNonFunctionFactory(arg); !errors.Is(e, ErrNonFunctionFactory) {
			t.Errorf("error not a instance of ErrNonFunctionFactory")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errNonFunctionFactory(arg, context); !errors.Is(e, ErrNonFunctionFactory) {
			t.Errorf("error not a instance of ErrNonFunctionFactory")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errFactoryWithoutResult(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : service factory without result"

	t.Run("creation without context", func(t *testing.T) {
		if e := errFactoryWithoutResult(arg); !errors.Is(e, ErrFactoryWithoutResult) {
			t.Errorf("error not a instance of ErrFactoryWithoutResult")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errFactoryWithoutResult(arg, context); !errors.Is(e, ErrFactoryWithoutResult) {
			t.Errorf("error not a instance of ErrFactoryWithoutResult")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errServiceNotFound(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : service not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errServiceNotFound(arg); !errors.Is(e, ErrServiceNotFound) {
			t.Errorf("error not a instance of ErrServiceNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errServiceNotFound(arg, context); !errors.Is(e, ErrServiceNotFound) {
			t.Errorf("error not a instance of ErrServiceNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}
