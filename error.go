package slate

import (
	"errors"
	"fmt"
)

// IError @todo doc {
type IError interface {
	error
	Unwrap() error
	Context() map[string]interface{}
}

// Error @todo doc
type Error struct {
	err     error
	context map[string]interface{}
}

var _ error = &Error{}
var _ IError = &Error{}

// NewError @todo doc
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

// NewErrorFrom @todo doc
func NewErrorFrom(
	err error,
	msg string,
	ctx ...map[string]interface{},
) error {
	e := NewError(msg, ctx...)
	e.(*Error).err = fmt.Errorf("%s : %w", e.Error(), err)
	return e
}

// Error @todo doc
func (e *Error) Error() string {
	return e.err.Error()
}

// Unwrap @todo doc
func (e *Error) Unwrap() error {
	return errors.Unwrap(e.err)
}

// Context @todo doc
func (e *Error) Context() map[string]interface{} {
	return e.context
}
