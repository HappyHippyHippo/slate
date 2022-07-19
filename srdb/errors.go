package srdb

import (
	"fmt"
	"github.com/happyhippyhippo/slate/serr"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serr.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", serr.ErrConversion, val, t)
}

func errDatabaseConfigNotFound(name string) error {
	return fmt.Errorf("%w : %v", serr.ErrDatabaseConfigNotFound, name)
}

func errUnknownDatabaseDialect(dialect string) error {
	return fmt.Errorf("%w : %v", serr.ErrUnknownDatabaseDialect, dialect)
}
