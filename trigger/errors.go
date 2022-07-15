package trigger

import (
	"fmt"
	"github.com/happyhippyhippo/slate/err"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", err.ErrNilPointer, arg)
}
