package gtrigger

import (
	"fmt"
	"github.com/happyhippyhippo/slate/gerror"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", gerror.ErrNilPointer, arg)
}
