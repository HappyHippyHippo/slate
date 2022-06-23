package slog

// Formatter interface defines the methods of a logging formatter instance
// responsible to parse a logging request into the output string.
type Formatter interface {
	Format(level Level, message string, ctx map[string]interface{}) string
}
