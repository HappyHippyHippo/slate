package rdb

import (
	"fmt"

	"github.com/happyhippyhippo/slate/err"
)

func errNilPointer(
	arg string,
) error {
	return fmt.Errorf("%w : %v", err.NilPointer, arg)
}

func errConversion(
	val interface{},
	t string,
) error {
	return fmt.Errorf("%w : %v to %v", err.Conversion, val, t)
}

func errConfigNotFound(
	name string,
) error {
	return fmt.Errorf("%w : %v", err.DatabaseConfigNotFound, name)
}

func errUnknownDialect(
	dialect string,
) error {
	return fmt.Errorf("%w : %v", err.UnknownDatabaseDialect, dialect)
}
