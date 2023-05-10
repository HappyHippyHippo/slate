package rest

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

func Test_errInvalidSource(t *testing.T) {
	arg := &config.Config{"field": "value"}
	context := map[string]interface{}{"field": "value"}
	message := "&map[field:value] : invalid config source"

	t.Run("creation without context", func(t *testing.T) {
		if e := errInvalidSource(arg); !errors.Is(e, config.ErrInvalidSource) {
			t.Errorf("error not a instance of config.ErrInvalidSource")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errInvalidSource(arg, context); !errors.Is(e, config.ErrInvalidSource) {
			t.Errorf("error not a instance of config.ErrInvalidSource")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errConfigNotFound(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : rest config source config not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errConfigNotFound(arg); !errors.Is(e, ErrConfigNotFound) {
			t.Errorf("error not a instance of ErrConfigNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
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
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errTimestampNotFound(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : rest config source timestamp not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errTimestampNotFound(arg); !errors.Is(e, ErrTimestampNotFound) {
			t.Errorf("error not a instance of ErrTimestampNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errTimestampNotFound(arg, context); !errors.Is(e, ErrTimestampNotFound) {
			t.Errorf("error not a instance of ErrTimestampNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}
