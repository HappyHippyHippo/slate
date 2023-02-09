package watchdog

import (
	"github.com/happyhippyhippo/slate/log"
)

// ILogAdapter defines an interface to a watchdog logging adapter
// used to mediate the watchdogs logging messages, a formatter and
// the application logger.
type ILogAdapter interface {
	Start() error
	Error(e error) error
	Done() error
}

// LogAdapter define an instance a watchdog logging adapter.
type LogAdapter struct {
	name       string
	channel    string
	startLevel log.Level
	errorLevel log.Level
	doneLevel  log.Level
	logger     log.ILog
	formatter  ILogFormatter
}

var _ ILogAdapter = &LogAdapter{}

// NewLogAdapter will create a new watchdog logging adapter.
func NewLogAdapter(
	name,
	channel string,
	startLevel,
	errorLevel,
	doneLevel log.Level,
	logger log.ILog,
	formatter ILogFormatter,
) (*LogAdapter, error) {
	// check log argument instance
	if logger == nil {
		return nil, errNilPointer("logger")
	}
	// check log formatter argument instance
	if formatter == nil {
		return nil, errNilPointer("formatter")
	}
	// return the created log adapter instance
	return &LogAdapter{
		name:       name,
		channel:    channel,
		startLevel: startLevel,
		errorLevel: errorLevel,
		doneLevel:  doneLevel,
		logger:     logger,
		formatter:  formatter,
	}, nil
}

// Start will format and redirect the start logging message to
// the application logger.
func (a *LogAdapter) Start() error {
	// propagate the logging signal to the adapter stored log instance
	return a.logger.Signal(a.channel, a.startLevel, a.formatter.Start(a.name), nil)
}

// Error will format and redirect the error logging message to
// the application logger.
func (a *LogAdapter) Error(
	e error,
) error {
	// propagate the logging signal to the adapter stored log instance
	return a.logger.Signal(a.channel, a.errorLevel, a.formatter.Error(a.name, e), nil)
}

// Done will format and redirect the termination logging message to
// the application logger.
func (a *LogAdapter) Done() error {
	// propagate the logging signal to the adapter stored log instance
	return a.logger.Signal(a.channel, a.doneLevel, a.formatter.Done(a.name), nil)
}
