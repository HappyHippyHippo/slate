package serror

import (
	"fmt"
)

var (
	// ErrNilPointer defines a nil pointer argument error
	ErrNilPointer = fmt.Errorf("invalid nil pointer")

	// ErrConversion defines a type conversion error
	ErrConversion = fmt.Errorf("invalid type conversion")

	// ErrServiceNotFound defines a service not found on the container
	ErrServiceNotFound = fmt.Errorf("service not found")

	// ErrConfigSourceNotFound defines a source config source not found error.
	ErrConfigSourceNotFound = fmt.Errorf("config source not found")

	// ErrDuplicateConfigSource defines a duplicate config source
	// registration attempt.
	ErrDuplicateConfigSource = fmt.Errorf("config source already registered")

	// ErrConfigPathNotFound defines a path in Partial not found error.
	ErrConfigPathNotFound = fmt.Errorf("config path not found")

	// ErrConfigRemotePathNotFound defines a config path not found error.
	ErrConfigRemotePathNotFound = fmt.Errorf("remote path not found")

	// ErrInvalidConfigDecoderFormat defines an error that signal an
	// unexpected/unknown config source decoder format.
	ErrInvalidConfigDecoderFormat = fmt.Errorf("invalid config decoder format")

	// ErrInvalidConfigSourceType defines an error that signal an
	// unexpected/unknown config source type.
	ErrInvalidConfigSourceType = fmt.Errorf("invalid config source type")

	// ErrInvalidConfigSourcePartial defines an error that signal an
	// invalid source configuration Partial.
	ErrInvalidConfigSourcePartial = fmt.Errorf("invalid config source config")

	// ErrInvalidLogFormat @todo doc
	ErrInvalidLogFormat = fmt.Errorf("invalid output format")

	// ErrInvalidLogLevel @todo doc
	ErrInvalidLogLevel = fmt.Errorf("invalid logger level")

	// ErrDuplicateLogStream @todo doc
	ErrDuplicateLogStream = fmt.Errorf("stream already registered")

	// ErrInvalidLogStreamType @todo doc
	ErrInvalidLogStreamType = fmt.Errorf("invalid stream type")

	// ErrInvalidLogStreamConfig @todo doc
	ErrInvalidLogStreamConfig = fmt.Errorf("invalid log stream config")

	// ErrDatabaseConfigNotFound @todo doc
	ErrDatabaseConfigNotFound = fmt.Errorf("database config not found")

	// ErrUnknownDatabaseDialect @todo doc
	ErrUnknownDatabaseDialect = fmt.Errorf("unknown database dialect")

	// ErrTranslatorNotFound @todo doc
	ErrTranslatorNotFound = fmt.Errorf("translator not found")
)
