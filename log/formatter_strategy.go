package log

const (
	// UnknownFormatter defines the value to be used to declare a
	// unknown Log Formatter format.
	UnknownFormatter = "unknown"
)

// FormatterStrategy interface defines the methods of the Formatter
// factory strategy that can validate creation requests and instantiation
// of particular decoder.
type FormatterStrategy interface {
	Accept(format string) bool
	Create(args ...interface{}) (Formatter, error)
}
