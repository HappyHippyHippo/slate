package slog

import (
	"fmt"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serr"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serr.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", serr.ErrConversion, val, t)
}

func errInvalidFormat(format string) error {
	return fmt.Errorf("%w : %v", serr.ErrInvalidLogFormat, format)
}

func errInvalidLevel(level string) error {
	return fmt.Errorf("%w : %v", serr.ErrInvalidLogLevel, level)
}

func errDuplicateStream(id string) error {
	return fmt.Errorf("%w : %v", serr.ErrDuplicateLogStream, id)
}

func errInvalidStreamType(streamType string) error {
	return fmt.Errorf("%w : %v", serr.ErrInvalidLogStreamType, streamType)
}

func errInvalidStreamConfig(cfg sconfig.IConfig) error {
	return fmt.Errorf("%w : %v", serr.ErrInvalidLogStreamConfig, cfg)
}
