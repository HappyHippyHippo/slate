package migration

import (
	"fmt"
	"github.com/happyhippyhippo/slate/err"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", err.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", err.ErrConversion, val, t)
}
