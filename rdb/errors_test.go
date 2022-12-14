package rdb

import (
	"errors"
	"testing"

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

func Test_errConfigNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "database config not found : dummy argument"

		if e := errConfigNotFound(arg); !errors.Is(e, err.DatabaseConfigNotFound) {
			t.Errorf("error not a instance of err.DatabaseConfigNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errUnknownDialect(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "unknown database dialect : dummy argument"

		if e := errUnknownDialect(arg); !errors.Is(e, err.UnknownDatabaseDialect) {
			t.Errorf("error not a instance of err.UnknownDatabaseDialect")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}
