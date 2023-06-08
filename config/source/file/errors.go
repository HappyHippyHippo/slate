package file

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

func errInvalidSource(
	partial config.Partial,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(config.ErrInvalidSource, fmt.Sprintf("%v", partial), ctx...)
}
