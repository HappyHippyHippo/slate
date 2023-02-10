package log

import (
	"github.com/happyhippyhippo/slate/config"
)

// IStreamFactory defined the interface of a log stream factory instance.
type IStreamFactory interface {
	Register(strategy IStreamStrategy) error
	Create(cfg config.IConfig) (IStream, error)
}

// StreamFactory is a logging stream generator based on a
// registered list of stream generation strategies.
type StreamFactory []IStreamStrategy

var _ IStreamFactory = &StreamFactory{}

// NewStreamFactory @todo doc
func NewStreamFactory() IStreamFactory {
	return &StreamFactory{}
}

// Register will register a new stream factory strategy to be used
// on creation requests.
func (f *StreamFactory) Register(
	strategy IStreamStrategy,
) error {
	// check the strategy argument reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// add the strategy to the stream factory strategy pool
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate and return a new config stream loaded
// by a configuration instance.
func (f StreamFactory) Create(
	cfg config.IConfig,
) (IStream, error) {
	// check config argument reference
	if cfg == nil {
		return nil, errNilPointer("config")
	}
	// search in the factory strategy pool for one that would accept
	// to generate the requested stream with the requested format defined
	// in the given config
	for _, s := range f {
		if s.Accept(cfg) {
			// return the creation of the requested stream
			return s.Create(cfg)
		}
	}
	return nil, errInvalidStream(cfg)
}
