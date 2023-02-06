package validation

import (
	"fmt"

	"github.com/happyhippyhippo/slate/err"
)

func errNilPointer(
	arg string,
) error {
	return fmt.Errorf("%w : %v", err.NilPointer, arg)
}

func errConversion(
	val interface{},
	t string,
) error {
	return fmt.Errorf("%w : %v to %v", err.Conversion, val, t)
}

func errTranslatorNotFound(
	translator string,
) error {
	return fmt.Errorf("%w : %v", err.TranslatorNotFound, translator)
}
