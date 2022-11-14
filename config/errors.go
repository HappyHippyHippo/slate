package config

import (
	"fmt"

	serror "github.com/happyhippyhippo/slate/error"
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

func errConfigRestPathNotFound(path string) error {
	return fmt.Errorf("%w : %v", serror.ErrConfigRestPathNotFound, path)
}

func errInvalidConfigDecoderFormat(format string) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidConfigDecoderFormat, format)
}

func errInvalidConfigSourceType(sourceType string) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidConfigSourceType, sourceType)
}

func errInvalidConfigSourcePartial(cfg IConfig) error {
	return fmt.Errorf("%w : %v", serror.ErrInvalidConfigSourcePartial, cfg)
}
