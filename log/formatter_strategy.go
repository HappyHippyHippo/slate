package log

// IFormatterStrategy interface defines the methods of the formatter
// factory strategy that can validate creation requests and instantiation
// of particular decoder.
type IFormatterStrategy interface {
	Accept(format string) bool
	Create(args ...interface{}) (IFormatter, error)
}
