package gfs

import (
	"fmt"
	"github.com/happyhippyhippo/slate/gerror"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", gerror.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", gerror.ErrConversion, val, t)
}
