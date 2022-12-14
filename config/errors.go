package config

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

func errPathNotFound(
	path string,
) error {
	return fmt.Errorf("%w : %v", err.ConfigPathNotFound, path)
}

func errInvalidFormat(
	format string,
) error {
	return fmt.Errorf("%w : %v", err.InvalidConfigFormat, format)
}

func errInvalidSource(
	sourceType string,
) error {
	return fmt.Errorf("%w : %v", err.InvalidConfigSource, sourceType)
}

func errRestPathNotFound(
	path string,
) error {
	return fmt.Errorf("%w : %v", err.ConfigRestPathNotFound, path)
}

func errInvalidSourceData(
	cfg IConfig,
) error {
	return fmt.Errorf("%w : %v", err.InvalidConfigSourceData, cfg)
}

func errSourceNotFound(
	id string,
) error {
	return fmt.Errorf("%w : %v", err.ConfigSourceNotFound, id)
}

func errDuplicateSource(
	id string,
) error {
	return fmt.Errorf("%w : %v", err.DuplicateConfigSource, id)
}
