package watchdog

// LogFormatter defines an interface to a watchdog logging message formatter.
type LogFormatter interface {
	Start(service string) string
	Error(service string, e error) string
	Done(service string) string
}
