package slog

import (
	"errors"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serror"
	"testing"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if err := errNilPointer(arg); !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("error not a instance of ErrNilPointer")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrConversion(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		typ := "dummy type"
		expected := "invalid type conversion : dummy value to dummy type"

		if err := errConversion(arg, typ); !errors.Is(err, serror.ErrConversion) {
			t.Errorf("error not a instance of ErrConversion")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidFormat(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid output format : dummy argument"

		if err := errInvalidFormat(arg); !errors.Is(err, serror.ErrInvalidLogFormat) {
			t.Errorf("error not a instance of ErrInvalidFormat")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidLevel(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid logger level : dummy argument"

		if err := errInvalidLevel(arg); !errors.Is(err, serror.ErrInvalidLogLevel) {
			t.Errorf("error not a instance of ErrInvalidLevel")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrDuplicateStream(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "stream already registered : dummy argument"

		if err := errDuplicateStream(arg); !errors.Is(err, serror.ErrDuplicateLogStream) {
			t.Errorf("error not a instance of ErrDuplicateStream")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidStreamType(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid stream type : dummy argument"

		if err := errInvalidStreamType(arg); !errors.Is(err, serror.ErrInvalidLogStreamType) {
			t.Errorf("error not a instance of ErrInvalidStreamType")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidStreamConfig(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := &sconfig.Partial{}

		if err := errInvalidStreamConfig(arg); !errors.Is(err, serror.ErrInvalidLogStreamConfig) {
			t.Errorf("error not a instance of ErrInvalidStreamConfig")
		}
	})
}
