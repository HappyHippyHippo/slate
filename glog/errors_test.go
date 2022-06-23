package glog

import (
	"errors"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
	"testing"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if err := errNilPointer(arg); !errors.Is(err, gerror.ErrNilPointer) {
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

		if err := errConversion(arg, typ); !errors.Is(err, gerror.ErrConversion) {
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

		if err := errInvalidFormat(arg); !errors.Is(err, gerror.ErrInvalidLogFormat) {
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

		if err := errInvalidLevel(arg); !errors.Is(err, gerror.ErrInvalidLogLevel) {
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

		if err := errDuplicateStream(arg); !errors.Is(err, gerror.ErrDuplicateLogStream) {
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

		if err := errInvalidStreamType(arg); !errors.Is(err, gerror.ErrInvalidLogStreamType) {
			t.Errorf("error not a instance of ErrInvalidStreamType")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrInvalidStreamConfig(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := &gconfig.Partial{}

		if err := errInvalidStreamConfig(arg); !errors.Is(err, gerror.ErrInvalidLogStreamConfig) {
			t.Errorf("error not a instance of ErrInvalidStreamConfig")
		}
	})
}
