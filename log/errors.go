package log

import (
	"fmt"

	"github.com/happyhippyhippo/slate/config"
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

func errInvalidFormat(
	format string,
) error {
	return fmt.Errorf("%w : %v", err.InvalidLogFormat, format)
}

func errInvalidLevel(
	level string,
) error {
	return fmt.Errorf("%w : %v", err.InvalidLogLevel, level)
}

func errInvalidType(
	streamType string,
) error {
	return fmt.Errorf("%w : %v", err.InvalidLogStream, streamType)
}

func errInvalidConfig(
	cfg config.IConfig,
) error {
	return fmt.Errorf("%w : %v", err.InvalidLogConfig, cfg)
}

func errDuplicateStream(
	id string,
) error {
	return fmt.Errorf("%w : %v", err.DuplicateLogStream, id)
}
