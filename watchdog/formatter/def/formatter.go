package def

import (
	"fmt"

	"github.com/happyhippyhippo/slate/watchdog"
)

// Formatter defines an instance to a watchdog logging message formatter.
type Formatter struct{}

var _ watchdog.ILogFormatter = &Formatter{}

// Start format a watchdog starting signal message.
func (Formatter) Start(
	service string,
) string {
	return fmt.Sprintf(watchdog.LogStartMessage, service)
}

// Error format a watchdog error signal message.
func (Formatter) Error(
	service string,
	e error,
) string {
	return fmt.Sprintf(watchdog.LogErrorMessage, service, e)
}

// Done format a watchdog termination signal message.
func (Formatter) Done(
	service string,
) string {
	return fmt.Sprintf(watchdog.LogDoneMessage, service)
}
