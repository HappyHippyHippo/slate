package config

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
)

var (
	// ErrPathNotFound defines a path in Config not found error.
	ErrPathNotFound = fmt.Errorf("config path not found")

	// ErrInvalidFormat defines an error that signal an
	// unexpected/unknown config source decoder format.
	ErrInvalidFormat = fmt.Errorf("invalid config format")

	// ErrInvalidSource defines an error that signal an
	// unexpected/unknown config source type.
	ErrInvalidSource = fmt.Errorf("invalid config source")

	// ErrRestConfigNotFound defines a rest response
	// config not found error.
	ErrRestConfigNotFound = fmt.Errorf("rest config source config not found")

	// ErrRestTimestampNotFound defines a rest response
	// timestamp not found error.
	ErrRestTimestampNotFound = fmt.Errorf("rest config source timestamp not found")

	// ErrSourceNotFound defines a source config source not found error.
	ErrSourceNotFound = fmt.Errorf("config source not found")

	// ErrDuplicateSource defines a duplicate config source
	// registration attempt.
	ErrDuplicateSource = fmt.Errorf("config source already registered")
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

func errPathNotFound(
	path string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrPathNotFound, path, ctx...)
}

func errInvalidFormat(
	format string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrInvalidFormat, format, ctx...)
}

func errInvalidSource(
	cfg IConfig,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrInvalidSource, fmt.Sprintf("%v", cfg), ctx...)
}

func errRestConfigNotFound(
	path string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrRestConfigNotFound, path, ctx...)
}

func errRestTimestampNotFound(
	path string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrRestTimestampNotFound, path, ctx...)
}

func errSourceNotFound(
	id string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrSourceNotFound, id, ctx...)
}

func errDuplicateSource(
	id string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrDuplicateSource, id, ctx...)
}
