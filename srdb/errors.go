package srdb

import (
	"fmt"
	"github.com/happyhippyhippo/slate/serror"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serror.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", serror.ErrConversion, val, t)
}

func errDatabaseConfigNotFound(name string) error {
	return fmt.Errorf("%w : %v", serror.ErrDatabaseConfigNotFound, name)
}

func errUnknownDatabaseDialect(dialect string) error {
	return fmt.Errorf("%w : %v", serror.ErrUnknownDatabaseDialect, dialect)
}
