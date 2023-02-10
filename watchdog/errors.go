package watchdog

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

var (
	// ErrInvalidWatchdog defines an error that signal that the
	// given watchdog configuration was unable to be parsed correctly.
	ErrInvalidWatchdog = fmt.Errorf("invalid watchdog config")

	// ErrDuplicateService defines an error that signal that the
	// given watchdog service is already registered.
	ErrDuplicateService = fmt.Errorf("duplicate watchdog service")
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

func errInvalidWatchdog(
	cfg config.IConfig,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrInvalidWatchdog, fmt.Sprintf("%v", cfg), ctx...)
}

func errDuplicateService(
	service string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrDuplicateService, service, ctx...)
}
