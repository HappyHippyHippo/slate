package strigger

import (
	"fmt"
	"github.com/happyhippyhippo/slate/serror"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serror.ErrNilPointer, arg)
}
