package log

import (
	"errors"
	"testing"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
)

func Test_errNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, err.NilPointer) {
			t.Errorf("error not a instance of err.NilPointer")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errConversion(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		typ := "dummy type"
		expected := "invalid type conversion : dummy value to dummy type"

		if e := errConversion(arg, typ); !errors.Is(e, err.Conversion) {
			t.Errorf("error not a instance of err.Conversion")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errInvalidFormat(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid output log format : dummy argument"

		if e := errInvalidFormat(arg); !errors.Is(e, err.InvalidLogFormat) {
			t.Errorf("err not a instance of err.InvalidFormat")
		} else if e.Error() != expected {
			t.Errorf("err message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errInvalidLevel(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid log level : dummy argument"

		if e := errInvalidLevel(arg); !errors.Is(e, err.InvalidLogLevel) {
			t.Errorf("error not a instance of err.InvalidLogLevel")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errInvalidType(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid log stream type : dummy argument"

		if e := errInvalidType(arg); !errors.Is(e, err.InvalidLogStream) {
			t.Errorf("error not a instance of err.InvalidLogStream")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errInvalidConfig(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := &config.Config{}

		if e := errInvalidConfig(arg); !errors.Is(e, err.InvalidLogConfig) {
			t.Errorf("error not a instance of err.InvalidLogConfig")
		}
	})
}
