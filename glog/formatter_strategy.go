package glog

// FormatterStrategy interface defines the methods of the formatter
// factory strategy that can validate creation requests and instantiation
// of particular decoder.
type FormatterStrategy interface {
	Accept(format string) bool
	Create(args ...interface{}) (Formatter, error)
}
