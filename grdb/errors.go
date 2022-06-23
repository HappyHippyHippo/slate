package grdb

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

func errDatabaseConfigNotFound(name string) error {
	return fmt.Errorf("%w : %v", gerror.ErrDatabaseConfigNotFound, name)
}

func errUnknownDatabaseDialect(dialect string) error {
	return fmt.Errorf("%w : %v", gerror.ErrUnknownDatabaseDialect, dialect)
}
