package watchdog

import (
	"fmt"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
)

func errNilPointer(
	arg string,
) error {
	return fmt.Errorf("%w : %v", err.NilPointer, arg)
}

func errConversion(
	val interface{},
	t string,
) error {
	return fmt.Errorf("%w : %v to %v", err.Conversion, val, t)
}

func errInvalidConfig(
	cfg config.IConfig,
) error {
	return fmt.Errorf("%w : %v", err.InvalidWatchdogConfig, cfg)
}

func errDuplicateService(
	service string,
) error {
	return fmt.Errorf("%w : %v", err.DuplicateWatchdogService, service)
}
