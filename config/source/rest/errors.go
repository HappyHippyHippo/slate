package rest

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

var (
	// ErrConfigNotFound defines a rest response
	// config not found error.
	ErrConfigNotFound = fmt.Errorf("rest config source config not found")

	// ErrTimestampNotFound defines a rest response
	// timestamp not found error.
	ErrTimestampNotFound = fmt.Errorf("rest config source timestamp not found")
)

func errNilPointer(
	arg string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(slate.ErrNilPointer, arg, ctx...)
}

func errInvalidSource(
	partial *config.Partial,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(config.ErrInvalidSource, fmt.Sprintf("%v", partial), ctx...)
}

func errConfigNotFound(
	arg string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrConfigNotFound, arg, ctx...)
}

func errTimestampNotFound(
	arg string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrTimestampNotFound, arg, ctx...)
}
