package cache

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

var (
	// ErrConfigNotFound defines an error that signal that the
	// configuration to the requested Store was not found.
	ErrConfigNotFound = fmt.Errorf("cache store config not found")

	// ErrInvalidStore defines an error that signal that the
	// given cache configuration was unable to be parsed correctly.
	ErrInvalidStore = fmt.Errorf("invalid cache store config")

	// ErrMiss @todo doc
	ErrMiss = fmt.Errorf("cache key not found")

	// ErrNotStored @todo doc
	ErrNotStored = fmt.Errorf("cache element not stored")
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

func errInvalidStore(
	cfg config.Partial,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrInvalidStore, fmt.Sprintf("%v", cfg), ctx...)
}
