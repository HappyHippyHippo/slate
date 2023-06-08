package rdb

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

var (
	// ErrConfigNotFound defines an error that signal that the
	// configuration to the requested database connection was not found.
	ErrConfigNotFound = fmt.Errorf("database config not found")

	// ErrUnknownDialect defines an error that signal that the
	// requested database connection configured dialect is unknown.
	ErrUnknownDialect = fmt.Errorf("unknown database dialect")
)

func errNilPointer(
	arg string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(slate.ErrNilPointer, arg, ctx...)
}

func errConversion(
	val interface{},
	t string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(slate.ErrConversion, fmt.Sprintf("%v to %s", val, t), ctx...)
}

func errConfigNotFound(
	name string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrConfigNotFound, name, ctx...)
}

func errUnknownDialect(
	cfg config.Partial,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrUnknownDialect, fmt.Sprintf("%v", cfg), ctx...)
}
