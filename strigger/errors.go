package strigger

import (
	"fmt"
	"github.com/happyhippyhippo/slate/serr"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serr.ErrNilPointer, arg)
}
