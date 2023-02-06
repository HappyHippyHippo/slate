package config

import (
	"io"
)

// IDecoder interface defines the interaction methods to a config
// content decoder used to parse the source content into an application
// usable configuration Config instance.
type IDecoder interface {
	io.Closer

	Decode() (IConfig, error)
}
