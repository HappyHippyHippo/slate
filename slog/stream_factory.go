package slog

import (
	"github.com/happyhippyhippo/slate/sconfig"
)

// IStreamFactory defined the interface of a slog stream factory instance.
type IStreamFactory interface {
	Register(strategy IStreamStrategy) error
	Create(streamType string, args ...interface{}) (IStream, error)
	CreateFromConfig(cfg sconfig.IConfig) (IStream, error)
}

// StreamFactory is a logging stream generator based on a
// registered list of stream generation strategies.
type StreamFactory []IStreamStrategy

var _ IStreamFactory = &StreamFactory{}

// Register will register a new stream factory strategy to be used
// on creation requests.
func (f *StreamFactory) Register(strategy IStreamStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	*f = append(*f, strategy)

	return nil
}

// Create will instantiate and return a new sconfig stream.
func (f StreamFactory) Create(streamType string, args ...interface{}) (IStream, error) {
	for _, s := range f {
		if s.Accept(streamType) {
			return s.Create(args...)
		}
	}
	return nil, errInvalidStreamType("streamType")
}

// CreateFromConfig will instantiate and return a new sconfig stream loaded
// by a configuration instance.
func (f StreamFactory) CreateFromConfig(cfg sconfig.IConfig) (IStream, error) {
	if cfg == nil {
		return nil, errNilPointer("cfg")
	}

	for _, s := range f {
		if s.AcceptFromConfig(cfg) {
			return s.CreateFromConfig(cfg)
		}
	}
	return nil, errInvalidStreamConfig(cfg)
}
