package config

import (
	"fmt"
	"github.com/happyhippyhippo/slate/err"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", err.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", err.ErrConversion, val, t)
}

func errConfigSourceNotFound(id string) error {
	return fmt.Errorf("%w : %v", err.ErrConfigSourceNotFound, id)
}

func errDuplicateConfigSource(id string) error {
	return fmt.Errorf("%w : %v", err.ErrDuplicateConfigSource, id)
}

func errConfigPathNotFound(path string) error {
	return fmt.Errorf("%w : %v", err.ErrConfigPathNotFound, path)
}

func errConfigRestPathNotFound(path string) error {
	return fmt.Errorf("%w : %v", err.ErrConfigRestPathNotFound, path)
}

func errInvalidConfigDecoderFormat(format string) error {
	return fmt.Errorf("%w : %v", err.ErrInvalidConfigDecoderFormat, format)
}

func errInvalidConfigSourceType(sourceType string) error {
	return fmt.Errorf("%w : %v", err.ErrInvalidConfigSourceType, sourceType)
}

func errInvalidConfigSourcePartial(cfg IConfig) error {
	return fmt.Errorf("%w : %v", err.ErrInvalidConfigSourcePartial, cfg)
}
