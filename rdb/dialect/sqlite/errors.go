//go:build sqlite

package sqlite

import (
	"github.com/happyhippyhippo/slate"
)

func errNilPointer(
	arg string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(slate.ErrNilPointer, arg, ctx...)
}