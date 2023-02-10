package watchdog

import (
	"fmt"
)

// DefaultLogFormatter defines an instance to a watchdog logging message formatter.
type DefaultLogFormatter struct{}

var _ ILogFormatter = &DefaultLogFormatter{}

// Start format a watchdog starting signal message.
func (DefaultLogFormatter) Start(
	service string,
) string {
	return fmt.Sprintf(LogStartMessage, service)
}

// Error format a watchdog error signal message.
func (DefaultLogFormatter) Error(
	service string,
	e error,
) string {
	return fmt.Sprintf(LogErrorMessage, service, e)
}

// Done format a watchdog termination signal message.
func (DefaultLogFormatter) Done(
	service string,
) string {
	return fmt.Sprintf(LogDoneMessage, service)
}
