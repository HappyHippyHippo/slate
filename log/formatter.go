package log

// IFormatter interface defines the methods of a logging Formatter instance
// responsible to parse a logging request into the output string.
type IFormatter interface {
	Format(level Level, message string, ctx ...Context) string
}
