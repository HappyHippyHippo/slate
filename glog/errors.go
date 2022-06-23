package glog

import (
	"fmt"
	"github.com/happyhippyhippo/slate/gconfig"
	"github.com/happyhippyhippo/slate/gerror"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", gerror.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", gerror.ErrConversion, val, t)
}

func errInvalidFormat(format string) error {
	return fmt.Errorf("%w : %v", gerror.ErrInvalidLogFormat, format)
}

func errInvalidLevel(level string) error {
	return fmt.Errorf("%w : %v", gerror.ErrInvalidLogLevel, level)
}

func errDuplicateStream(id string) error {
	return fmt.Errorf("%w : %v", gerror.ErrDuplicateLogStream, id)
}

func errInvalidStreamType(stype string) error {
	return fmt.Errorf("%w : %v", gerror.ErrInvalidLogStreamType, stype)
}

func errInvalidStreamConfig(cfg gconfig.Config) error {
	return fmt.Errorf("%w : %v", gerror.ErrInvalidLogStreamConfig, cfg)
}
