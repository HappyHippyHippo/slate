package srdb

import (
	"errors"
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

func Test_ErrDatabaseConfigNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "database config not found : dummy argument"

		if err := errDatabaseConfigNotFound(arg); !errors.Is(err, serror.ErrDatabaseConfigNotFound) {
			t.Errorf("error not a instance of ErrConfigNotFound")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}

func Test_ErrUnknownDatabaseDialect(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "unknown database dialect : dummy argument"

		if err := errUnknownDatabaseDialect(arg); !errors.Is(err, serror.ErrUnknownDatabaseDialect) {
			t.Errorf("error not a instance of ErrUnknownDialect")
		} else if err.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", err, expected)
		}
	})
}
