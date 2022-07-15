package log

import (
	"errors"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"testing"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, err.ErrNilPointer) {
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

		if e := errConversion(arg, typ); !errors.Is(e, err.ErrConversion) {
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

		if e := errInvalidFormat(arg); !errors.Is(e, err.ErrInvalidLogFormat) {
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

		if e := errInvalidLevel(arg); !errors.Is(e, err.ErrInvalidLogLevel) {
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

		if e := errDuplicateStream(arg); !errors.Is(e, err.ErrDuplicateLogStream) {
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

		if e := errInvalidStreamType(arg); !errors.Is(e, err.ErrInvalidLogStreamType) {
			t.Errorf("error not a instance of ErrInvalidStreamType")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrInvalidStreamConfig(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := &config.Partial{}

		if e := errInvalidStreamConfig(arg); !errors.Is(e, err.ErrInvalidLogStreamConfig) {
			t.Errorf("error not a instance of ErrInvalidStreamConfig")
		}
	})
}
