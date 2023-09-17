package cache

import (
	"errors"
	"reflect"
	"testing"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
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
		} else if te, ok := e.(*slate.Error); !ok {
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
		} else if te, ok := e.(*slate.Error); !ok {
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
		if e := errConversion(arg, typ); !errors.Is(e, slate.ErrConversion) {
			t.Errorf("error not a instance of slate.ErrConversion")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errConversion(arg, typ, context); !errors.Is(e, slate.ErrConversion) {
			t.Errorf("error not a instance of slate.ErrConversion")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errConfigNotFound(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : cache store config not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errConfigNotFound(arg); !errors.Is(e, ErrConfigNotFound) {
			t.Errorf("error not a instance of ErrConfigNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errConfigNotFound(arg, context); !errors.Is(e, ErrConfigNotFound) {
			t.Errorf("error not a instance of ErrConfigNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errInvalidStore(t *testing.T) {
	arg := config.Partial{"field": "value"}
	context := map[string]interface{}{"field": "value"}
	message := "map[field:value] : invalid cache store config"

	t.Run("creation without context", func(t *testing.T) {
		if e := errInvalidStore(arg); !errors.Is(e, ErrInvalidStore) {
			t.Errorf("error not a instance of ErrInvalidStore")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errInvalidStore(arg, context); !errors.Is(e, ErrInvalidStore) {
			t.Errorf("error not a instance of ErrInvalidStore")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}
