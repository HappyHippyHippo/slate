//go:build inmemory

package inmemory

import (
	"errors"
	"reflect"
	"testing"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/cache"
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

func Test_errInvalidStore(t *testing.T) {
	arg := config.Partial{"field": "value"}
	context := map[string]interface{}{"field": "value"}
	message := "map[field:value] : invalid cache store config"

	t.Run("creation without context", func(t *testing.T) {
		if e := errInvalidStore(arg); !errors.Is(e, cache.ErrInvalidStore) {
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
		if e := errInvalidStore(arg, context); !errors.Is(e, cache.ErrInvalidStore) {
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

func Test_errMiss(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : cache key not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errMiss(arg); !errors.Is(e, cache.ErrMiss) {
			t.Errorf("error not a instance of ErrMiss")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errMiss(arg, context); !errors.Is(e, cache.ErrMiss) {
			t.Errorf("error not a instance of ErrMiss")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errNotStored(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : cache element not stored"

	t.Run("creation without context", func(t *testing.T) {
		if e := errNotStored(arg); !errors.Is(e, cache.ErrNotStored) {
			t.Errorf("error not a instance of ErrNotStored")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errNotStored(arg, context); !errors.Is(e, cache.ErrNotStored) {
			t.Errorf("error not a instance of ErrNotStored")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}
