package watchdog

// ILogFormatter defines an interface to a watchdog logging message formatter.
type ILogFormatter interface {
	Start(service string) string
	Error(service string, e error) string
	Done(service string) string
}
