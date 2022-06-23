package sconfig

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

func errConfigSourceNotFound(id string) error {
	return fmt.Errorf("%w : %v", serror.ErrConfigSourceNotFound, id)
}

func errDuplicateConfigSource(id string) error {
	return fmt.Errorf("%w : %v", serror.ErrDuplicateConfigSource, id)
}

func errConfigPathNotFound(path string) error {
	return fmt.Errorf("%w : %v", serror.ErrConfigPathNotFound, path)
}

func errConfigRemotePathNotFound(path string) error {
	return fmt.Errorf("%w : %v", serror.ErrConfigRemotePathNotFound, path)
}

func errInvalidConfigDecoderFormat(format string) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidConfigDecoderFormat, format)
}

func errInvalidConfigSourceType(stype string) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidConfigSourceType, stype)
}

func errInvalidConfigSourcePartial(cfg Config) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidConfigSourcePartial, cfg)
}
