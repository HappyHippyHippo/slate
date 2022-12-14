package slate

import (
	"errors"
	"fmt"
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

func Test_errFactoryWithoutResult(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		expected := "factory without result : dummy value"

		if e := errFactoryWithoutResult(arg); !errors.Is(e, err.FactoryWithoutResult) {
			t.Errorf("error not a instance of err.FactoryWithoutResult")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errNonFunctionFactory(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy value"
		expected := "non-function factory : dummy value"

		if e := errNonFunctionFactory(arg); !errors.Is(e, err.NonFunctionFactory) {
			t.Errorf("error not a instance of err.NonFunctionFactory")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errServiceNotFound(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := "dummy argument"
		expected := "service not found : dummy argument"

		if e := errServiceNotFound(arg); !errors.Is(e, err.ServiceNotFound) {
			t.Errorf("error not a instance of err.ServiceNotFound")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}

func Test_errContainer(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		arg := []error{fmt.Errorf("dummy argument 1"), fmt.Errorf("dummy argument 2")}
		expected := "service container error : [dummy argument 1 dummy argument 2]"

		if e := errContainer(arg...); !errors.Is(e, err.Container) {
			t.Errorf("error not a instance of err.Container")
		} else if e.Error() != expected {
			t.Errorf("error message (%v) not same as expected (%v)", e, expected)
		}
	})
}
