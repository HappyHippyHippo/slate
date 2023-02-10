package validation

import (
	"fmt"

	"github.com/happyhippyhippo/slate"
)

var (
	// ErrTranslatorNotFound @todo doc
	ErrTranslatorNotFound = fmt.Errorf("translator not found")
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

func errTranslatorNotFound(
	translator string,
	ctx ...map[string]interface{},
) error {
	return slate.NewErrorFrom(ErrTranslatorNotFound, translator, ctx...)
}
