package cache

import (
	"fmt"

	sconfig "github.com/happyhippyhippo/slate/config"
	serror "github.com/happyhippyhippo/slate/error"
	srerror "github.com/happyhippyhippo/slate/rest/error"
)

func errNilPointer(arg string) error {
	return fmt.Errorf("%w : %v", serror.ErrNilPointer, arg)
}

func errConversion(val interface{}, t string) error {
	return fmt.Errorf("%w : %v to %v", serror.ErrConversion, val, t)
}

func errInvalidKeyGeneratorType(arg string) error {
	return fmt.Errorf("%w : %v", srerror.ErrInvalidKeyGeneratorType, arg)
}

func errInvalidKeyGeneratorPartial(cfg sconfig.IConfig) error {
	return fmt.Errorf("%w : %v", srerror.ErrInvalidKeyGeneratorPartial, cfg)
}

func errInvalidStoreType(arg string) error {
	return fmt.Errorf("%w : %v", srerror.ErrInvalidStoreType, arg)
}

func errInvalidStorePartial(cfg sconfig.IConfig) error {
	return fmt.Errorf("%w : %v", srerror.ErrInvalidStorePartial, cfg)
}

func errCacheMiss(arg string) error {
	return fmt.Errorf("%w : %v", srerror.ErrCacheMiss, arg)
}

func errCacheNotStored(arg string) error {
	return fmt.Errorf("%w : %v", srerror.ErrCacheNotStored, arg)
}

func errCacheOpNotSupport(arg string) error {
	return fmt.Errorf("%w : %v", srerror.ErrCacheOpNotSupport, arg)
}
