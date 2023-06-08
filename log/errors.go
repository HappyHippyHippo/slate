package log

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

var (
	// ErrInvalidFormat defines an error that signal an invalid
	// logger format.
	ErrInvalidFormat = fmt.Errorf("invalid output log format")

	// ErrInvalidLevel defines an error that signal an invalid
	// logger Level.
	ErrInvalidLevel = fmt.Errorf("invalid log Level")

	// ErrInvalidStream defines an error that signal that the
	// given logger stream configuration was unable to be parsed correctly
	// enabling the logger stream generation.
	ErrInvalidStream = fmt.Errorf("invalid log stream")

	// ErrStreamNotFound defines an error that signal that the
	// given stream was not found in the logger manager.
	ErrStreamNotFound = fmt.Errorf("log stream not found")

	// ErrDuplicateStream defines an error that signal that the
	// requested logger stream to be registered have an id of an already
	// registered logger stream.
	ErrDuplicateStream = fmt.Errorf("log stream already registered")
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

func errInvalidFormat(
	format string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrInvalidFormat, format, ctx...)
}

func errInvalidLevel(
	level string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrInvalidLevel, level, ctx...)
}

func errInvalidStream(
	cfg config.Partial,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrInvalidStream, fmt.Sprintf("%v", cfg), ctx...)
}

func errStreamNotFound(
	id string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrStreamNotFound, id, ctx...)
}

func errDuplicateStream(
	id string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrDuplicateStream, id, ctx...)
}
