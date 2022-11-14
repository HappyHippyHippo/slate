package trigger

import (
	"fmt"

	serror "github.com/happyhippyhippo/slate/error"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serror.ErrNilPointer, arg)
}
