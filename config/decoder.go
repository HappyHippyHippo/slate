package config

import (
	"io"
)

const (
	// FormatUnknown defines the value to be used to
	// declare an unknown config source format.
	FormatUnknown = "unknown"

	// FormatYAML defines the value to be used to declare
	// a YAML config source format.
	FormatYAML = "yaml"

	// FormatJSON defines the value to be used to declare
	// a JSON config source format.
	FormatJSON = "json"
)

// IDecoder interface defines the interaction methods to a config
// content decoder used to parse the source content into an application
// usable configuration Config instance.
type IDecoder interface {
	io.Closer

	Decode() (IConfig, error)
}
