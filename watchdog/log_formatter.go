package watchdog

import (
	"fmt"
)

// ILogFormatter defines an interface to a watchdog logging message formatter.
type ILogFormatter interface {
	Start(service string) string
	Error(service string, e error) string
	Done(service string) string
}

// LogFormatter defines an instance to a watchdog logging message formatter.
type LogFormatter struct{}

var _ ILogFormatter = &LogFormatter{}

// Start format a watchdog starting signal message.
func (LogFormatter) Start(
	service string,
) string {
	return fmt.Sprintf(LogStartMessage, service)
}

// Error format a watchdog error signal message.
func (LogFormatter) Error(
	service string,
	e error,
) string {
	return fmt.Sprintf(LogErrorMessage, service, e)
}

// Done format a watchdog termination signal message.
func (LogFormatter) Done(
	service string,
) string {
	return fmt.Sprintf(LogDoneMessage, service)
}
