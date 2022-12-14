package slate

import (
	"fmt"
	"github.com/happyhippyhippo/slate/err"
)

func errNilPointer(
	arg string,
) error {
	return fmt.Errorf("%w : %v", err.NilPointer, arg)
}

func errConversion(
	val interface{},
	t string,
) error {
	return fmt.Errorf("%w : %v to %v", err.Conversion, val, t)
}

func errFactoryWithoutResult(
	arg string,
) error {
	return fmt.Errorf("%w : %v", err.FactoryWithoutResult, arg)
}

func errNonFunctionFactory(
	arg string,
) error {
	return fmt.Errorf("%w : %v", err.NonFunctionFactory, arg)
}

func errServiceNotFound(
	arg string,
) error {
	return fmt.Errorf("%w : %v", err.ServiceNotFound, arg)
}

func errContainer(
	e ...error,
) error {
	return fmt.Errorf("%w : %v", err.Container, e)
}
