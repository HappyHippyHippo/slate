package slog

import (
	"fmt"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serror"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serror.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", serror.ErrConversion, val, t)
}

func errInvalidFormat(format string) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidLogFormat, format)
}

func errInvalidLevel(level string) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidLogLevel, level)
}

func errDuplicateStream(id string) error {
	return fmt.Errorf("%w : %v", serror.ErrDuplicateLogStream, id)
}

func errInvalidStreamType(streamType string) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidLogStreamType, streamType)
}

func errInvalidStreamConfig(cfg sconfig.IConfig) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidLogStreamConfig, cfg)
}
