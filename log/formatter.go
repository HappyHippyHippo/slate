package log

const (
	// FormatUnknown defines the value to be used to declare an unknown
	// Log formatter format.
	FormatUnknown = "unknown"

	// FormatJSON defines the value to be used to declare a JSON
	// Log formatter format.
	FormatJSON = "json"
)

// IFormatter interface defines the methods of a logging formatter instance
// responsible to parse a logging request into the output string.
type IFormatter interface {
	Format(level Level, message string, ctx map[string]interface{}) string
}
