package config

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
		} else if te, ok := e.(slate.IError); !ok {
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
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errPathNotFound(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : config path not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errPathNotFound(arg); !errors.Is(e, ErrPathNotFound) {
			t.Errorf("error not a instance of ErrPathNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errPathNotFound(arg, context); !errors.Is(e, ErrPathNotFound) {
			t.Errorf("error not a instance of ErrPathNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errInvalidFormat(t *testing.T) {
	arg := UnknownDecoderFormat
	context := map[string]interface{}{"field": "value"}
	message := "unknown : invalid config format"

	t.Run("creation without context", func(t *testing.T) {
		if e := errInvalidFormat(arg); !errors.Is(e, ErrInvalidFormat) {
			t.Errorf("error not a instance of ErrInvalidFormat")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errInvalidFormat(arg, context); !errors.Is(e, ErrInvalidFormat) {
			t.Errorf("error not a instance of ErrInvalidFormat")
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
	arg := &Config{"field": "value"}
	context := map[string]interface{}{"field": "value"}
	message := "&map[field:value] : invalid config source"

	t.Run("creation without context", func(t *testing.T) {
		if e := errInvalidSource(arg); !errors.Is(e, ErrInvalidSource) {
			t.Errorf("error not a instance of ErrInvalidSource")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errInvalidSource(arg, context); !errors.Is(e, ErrInvalidSource) {
			t.Errorf("error not a instance of ErrInvalidSource")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errRestConfigNotFound(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : rest config not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errRestConfigNotFound(arg); !errors.Is(e, ErrRestConfigNotFound) {
			t.Errorf("error not a instance of ErrRestConfigNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errRestConfigNotFound(arg, context); !errors.Is(e, ErrRestConfigNotFound) {
			t.Errorf("error not a instance of ErrRestConfigNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errRestTimestampNotFound(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : rest timestamp not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errRestTimestampNotFound(arg); !errors.Is(e, ErrRestTimestampNotFound) {
			t.Errorf("error not a instance of ErrRestTimestampNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errRestTimestampNotFound(arg, context); !errors.Is(e, ErrRestTimestampNotFound) {
			t.Errorf("error not a instance of ErrRestTimestampNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errSourceNotFound(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : config source not found"

	t.Run("creation without context", func(t *testing.T) {
		if e := errSourceNotFound(arg); !errors.Is(e, ErrSourceNotFound) {
			t.Errorf("error not a instance of ErrSourceNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errSourceNotFound(arg, context); !errors.Is(e, ErrSourceNotFound) {
			t.Errorf("error not a instance of ErrSourceNotFound")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}

func Test_errDuplicateSource(t *testing.T) {
	arg := "dummy argument"
	context := map[string]interface{}{"field": "value"}
	message := "dummy argument : config source already registered"

	t.Run("creation without context", func(t *testing.T) {
		if e := errDuplicateSource(arg); !errors.Is(e, ErrDuplicateSource) {
			t.Errorf("error not a instance of ErrDuplicateSource")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if te.Context() != nil {
			t.Errorf("didn't stored a nil value context")
		}
	})

	t.Run("creation with context", func(t *testing.T) {
		if e := errDuplicateSource(arg, context); !errors.Is(e, ErrDuplicateSource) {
			t.Errorf("error not a instance of ErrDuplicateSource")
		} else if e.Error() != message {
			t.Errorf("error message (%v) not same as expected (%v)", e, message)
		} else if te, ok := e.(slate.IError); !ok {
			t.Errorf("didn't returned a slate error instance")
		} else if check := te.Context(); !reflect.DeepEqual(check, context) {
			t.Errorf("context (%v) not same as expected (%v)", check, context)
		}
	})
}
