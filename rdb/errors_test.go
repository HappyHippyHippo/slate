package rdb

import (
	"errors"
	"testing"

	serror "github.com/happyhippyhippo/slate/error"
)

func Test_ErrNilPointer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "invalid nil pointer : dummy argument"

		if e := errNilPointer(arg); !errors.Is(e, serror.ErrNilPointer) {
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

		if e := errConversion(arg, typ); !errors.Is(e, serror.ErrConversion) {
			t.Errorf("error not a instance of Conversion")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrDatabaseConfigNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "database config not found : dummy argument"

		if e := errDatabaseConfigNotFound(arg); !errors.Is(e, serror.ErrDatabaseConfigNotFound) {
			t.Errorf("error not a instance of ErrConfigNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_ErrUnknownDatabaseDialect(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "unknown database dialect : dummy argument"

		if e := errUnknownDatabaseDialect(arg); !errors.Is(e, serror.ErrUnknownDatabaseDialect) {
			t.Errorf("error not a instance of ErrUnknownDialect")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}
