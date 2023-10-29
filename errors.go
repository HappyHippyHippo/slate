package slate

import (
	"errors"
	"fmt"
)

var (
	// ErrNilPointer defines a nil pointer argument error.
	ErrNilPointer = NewError("invalid nil pointer")

	// ErrConversion defines a type conversion error.
	ErrConversion = NewError("invalid type conversion")
)

// Error defines a contextualized error
type Error struct {
	err     error
	context map[string]interface{}
}

var _ error = &Error{}

// NewError will instantiate a new error instance
func NewError(
	msg string,
	ctx ...map[string]interface{},
) error {
	e := &Error{err: errors.New(msg)}
	if len(ctx) != 0 {
		e.context = ctx[0]
	}
	return e
}

// NewErrorFrom create a new contextualized error instance
// from an existing error
func NewErrorFrom(
	err error,
	msg string,
	ctx ...map[string]interface{},
) error {
	e := NewError(msg, ctx...)
	e.(*Error).err = fmt.Errorf("%s : %w", e.Error(), err)
	return e
}

// Error retrieve the error information from the error instance
func (e *Error) Error() string {
	return e.err.Error()
}

// Unwrap will try to unwrap the error information
func (e *Error) Unwrap() error {
	return errors.Unwrap(e.err)
}

// Context will retrieve the error context information
func (e *Error) Context() map[string]interface{} {
	return e.context
}

func errNilPointer(
	arg string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrNilPointer, arg, ctx...)
}

func errConversion(
	val interface{},
	t string,
	ctx ...map[string]interface{},
) error {
	return NewErrorFrom(ErrConversion, fmt.Sprintf("%v to %s", val, t), ctx...)
}
