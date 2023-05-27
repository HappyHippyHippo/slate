package log

// Formatter interface defines the methods of a logging Formatter instance
// responsible to parse a logging request into the output string.
type Formatter interface {
	Format(level Level, message string, ctx ...Context) string
}
