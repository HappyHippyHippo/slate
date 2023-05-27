package watchdog

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

func Test_errInvalidWatchdog(t *testing.T) {
	arg := &config.Partial{"field": "value"}
	context := map[string]interface{}{"field": "value"}
	message := "&map[field:value] : invalid watchdog config"

	t.Run("creation without context", func(t *testing.T) {
		if e := errInvalidWatchdog(arg); !errors.Is(e, ErrInvalidWatchdog) {
			t.Errorf("error not a instance of ErrInvalidWatchdog")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errInvalidWatchdog(arg, context); !errors.Is(e, ErrInvalidWatchdog) {
			t.Errorf("error not a instance of ErrInvalidWatchdog")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errDuplicateService(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : duplicate watchdog service"

	t.Run("creation without context", func(t *testing.T) {
		if e := errDuplicateService(arg); !errors.Is(e, ErrDuplicateService) {
			t.Errorf("error not a instance of ErrDuplicateService")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errDuplicateService(arg, context); !errors.Is(e, ErrDuplicateService) {
			t.Errorf("error not a instance of ErrDuplicateService")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(*slate.Error); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}
