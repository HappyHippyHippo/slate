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
	cfg config.IConfig,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(config.ErrInvalidSource, fmt.Sprintf("%v", cfg), ctx...)
}
