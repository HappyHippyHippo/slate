package env

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
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

func errInvalidSource(
	partial *config.Partial,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(config.ErrInvalidSource, fmt.Sprintf("%v", partial), ctx...)
}
