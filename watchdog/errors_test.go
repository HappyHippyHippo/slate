package watchdog

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

func Test_errInvalidConfig(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := &config.Config{}

		if e := errInvalidConfig(arg); !errors.Is(e, err.InvalidWatchdogConfig) {
			t.Errorf("error not a instance of err.InvalidWatchdogConfig")
		}
	})
}

func Test_errDuplicateService(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "service"

		if e := errDuplicateService(arg); !errors.Is(e, err.DuplicateWatchdogService) {
			t.Errorf("error not a instance of err.DuplicateWatchdogService")
		}
	})
}
