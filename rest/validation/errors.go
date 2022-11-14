package validation

import (
	"fmt"
	serror "github.com/happyhippyhippo/slate/error"
	srerror "github.com/happyhippyhippo/slate/rest/error"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serror.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", serror.ErrConversion, val, t)
}

func errTranslatorNotFound(translator string) error {
	return fmt.Errorf("%w : %v", srerror.ErrTranslatorNotFound, translator)
}
