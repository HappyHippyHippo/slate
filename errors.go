package slate

import (
	"fmt"
)

var (
	// ErrNilPointer defines a nil pointer argument error.
	ErrNilPointer = NewError("invalid nil pointer")

	// ErrConversion defines a type conversion error.
	ErrConversion = NewError("invalid type conversion")

	// ErrContainer defines a container error.
	ErrContainer = NewError("service container error")

	// ErrNonFunctionFactory defines a service container registration error
	// that signals that the registration request was made with a
	// non-function service factory.
	ErrNonFunctionFactory = NewError("non-function service factory")

	// ErrFactoryWithoutResult defines a service container registration error
	// that signals that the registration request was made with a
	// function service factory that don't return a service.
	ErrFactoryWithoutResult = NewError("service factory without result")

	// ErrServiceNotFound defines a service not found on the container.
	ErrServiceNotFound = NewError("service not found")
)

func errNilPointer(
	arg string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrNilPointer, arg, ctx...)
}

func errContainer(
	e error,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrContainer, fmt.Errorf("%w", e).Error(), ctx...)
}

func errNonFunctionFactory(
	arg string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrNonFunctionFactory, arg, ctx...)
}

func errFactoryWithoutResult(
	arg string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrFactoryWithoutResult, arg, ctx...)
}

func errServiceNotFound(
	arg string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrServiceNotFound, arg, ctx...)
}
