package console

import (
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func errNilPointer(
	arg string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(slate.ErrNilPointer, arg, ctx...)
}

func errInvalidLevel(
	level string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(log.ErrInvalidLevel, level, ctx...)
}
