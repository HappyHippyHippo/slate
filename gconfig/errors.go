package gconfig

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

func errConfigSourceNotFound(id string) error {
	return fmt.Errorf("%w : %v", gerror.ErrConfigSourceNotFound, id)
}

func errDuplicateConfigSource(id string) error {
	return fmt.Errorf("%w : %v", gerror.ErrDuplicateConfigSource, id)
}

func errConfigPathNotFound(path string) error {
	return fmt.Errorf("%w : %v", gerror.ErrConfigPathNotFound, path)
}

func errConfigRemotePathNotFound(path string) error {
	return fmt.Errorf("%w : %v", gerror.ErrConfigRemotePathNotFound, path)
}

func errInvalidConfigDecoderFormat(format string) error {
	return fmt.Errorf("%w : %v", gerror.ErrInvalidConfigDecoderFormat, format)
}

func errInvalidConfigSourceType(stype string) error {
	return fmt.Errorf("%w : %v", gerror.ErrInvalidConfigSourceType, stype)
}

func errInvalidConfigSourcePartial(cfg Config) error {
	return fmt.Errorf("%w : %v", gerror.ErrInvalidConfigSourcePartial, cfg)
}
