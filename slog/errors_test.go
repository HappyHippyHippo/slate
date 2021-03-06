package slog

import (
	"errors"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serr"
	"testing"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, serr.ErrNilPointer) {
			t.Errorf("error not a instance of NilPointer")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrConversion(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		typ := "dummy type"
		expected := "invalid type conversion : dummy value to dummy type"

		if e := errConversion(arg, typ); !errors.Is(e, serr.ErrConversion) {
			t.Errorf("error not a instance of Conversion")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidFormat(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid output format : dummy argument"

		if e := errInvalidFormat(arg); !errors.Is(e, serr.ErrInvalidLogFormat) {
			t.Errorf("error not a instance of ErrInvalidFormat")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidLevel(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid logger level : dummy argument"

		if e := errInvalidLevel(arg); !errors.Is(e, serr.ErrInvalidLogLevel) {
			t.Errorf("error not a instance of ErrInvalidLevel")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrDuplicateStream(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "stream already registered : dummy argument"

		if e := errDuplicateStream(arg); !errors.Is(e, serr.ErrDuplicateLogStream) {
			t.Errorf("error not a instance of ErrDuplicateStream")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidStreamType(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid stream type : dummy argument"

		if e := errInvalidStreamType(arg); !errors.Is(e, serr.ErrInvalidLogStreamType) {
			t.Errorf("error not a instance of ErrInvalidStreamType")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidStreamConfig(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := &sconfig.Partial{}

		if e := errInvalidStreamConfig(arg); !errors.Is(e, serr.ErrInvalidLogStreamConfig) {
			t.Errorf("error not a instance of ErrInvalidStreamConfig")
		}
	})
}
