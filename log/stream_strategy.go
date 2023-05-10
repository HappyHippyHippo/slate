package log

import (
	"github.com/happyhippyhippo/slate/config"
)

const (
	// UnknownStream defines the value to be used to declare an
	// unknown Log stream type.
	UnknownStream = "unknown"
)

// IStreamStrategy interface defines the methods of the stream
// factory strategy that can validate creation requests and instantiation
// of particular type of stream.
type IStreamStrategy interface {
	Accept(cfg config.IConfig) bool
	Create(cfg config.IConfig) (IStream, error)
}
