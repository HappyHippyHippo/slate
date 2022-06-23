package sconfig

import "io"

// Decoder interface defines the interaction methods to a config
// content decoder used to parse the source content into an application
// usable configuration Partial instance.
type Decoder interface {
	io.Closer
	Decode() (Config, error)
}
