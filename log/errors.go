package log

import (
	"fmt"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", err.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", err.ErrConversion, val, t)
}

func errInvalidFormat(format string) error {
	return fmt.Errorf("%w : %v", err.ErrInvalidLogFormat, format)
}

func errInvalidLevel(level string) error {
	return fmt.Errorf("%w : %v", err.ErrInvalidLogLevel, level)
}

func errDuplicateStream(id string) error {
	return fmt.Errorf("%w : %v", err.ErrDuplicateLogStream, id)
}

func errInvalidStreamType(streamType string) error {
	return fmt.Errorf("%w : %v", err.ErrInvalidLogStreamType, streamType)
}

func errInvalidStreamConfig(cfg config.IConfig) error {
	return fmt.Errorf("%w : %v", err.ErrInvalidLogStreamConfig, cfg)
}
