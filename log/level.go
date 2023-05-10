package log

// Level identifies a value type that describes a logging Level.
type Level int

const (
	// FATAL defines a fatal logging Level.
	FATAL Level = 1 + iota
	// ERROR defines a error logging Level.
	ERROR
	// WARNING defines a warning logging Level.
	WARNING
	// NOTICE defines a notice logging Level.
	NOTICE
	// INFO defines a info logging Level.
	INFO
	// DEBUG defines a debug logging Level.
	DEBUG
)

// LevelMap defines a relation between a human-readable string
// and a code Level identifier of a logging Level.
var LevelMap = map[string]Level{
	"fatal":   FATAL,
	"error":   ERROR,
	"warning": WARNING,
	"notice":  NOTICE,
	"info":    INFO,
	"debug":   DEBUG,
}

// LevelMapName defines a relation between a code Level identifier of a
// logging Level and human-readable string representation of that Level.
var LevelMapName = map[Level]string{
	FATAL:   "fatal",
	ERROR:   "error",
	WARNING: "warning",
	NOTICE:  "notice",
	INFO:    "info",
	DEBUG:   "debug",
}
