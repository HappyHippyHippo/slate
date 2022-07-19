package sconfig

import (
	"fmt"
	"github.com/happyhippyhippo/slate/serr"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serr.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", serr.ErrConversion, val, t)
}

func errConfigSourceNotFound(id string) error {
	return fmt.Errorf("%w : %v", serr.ErrConfigSourceNotFound, id)
}

func errDuplicateConfigSource(id string) error {
	return fmt.Errorf("%w : %v", serr.ErrDuplicateConfigSource, id)
}

func errConfigPathNotFound(path string) error {
	return fmt.Errorf("%w : %v", serr.ErrConfigPathNotFound, path)
}

func errConfigRestPathNotFound(path string) error {
	return fmt.Errorf("%w : %v", serr.ErrConfigRestPathNotFound, path)
}

func errInvalidConfigDecoderFormat(format string) error {
	return fmt.Errorf("%w : %v", serr.ErrInvalidConfigDecoderFormat, format)
}

func errInvalidConfigSourceType(sourceType string) error {
	return fmt.Errorf("%w : %v", serr.ErrInvalidConfigSourceType, sourceType)
}

func errInvalidConfigSourcePartial(cfg IConfig) error {
	return fmt.Errorf("%w : %v", serr.ErrInvalidConfigSourcePartial, cfg)
}
