package slog

import "github.com/happyhippyhippo/slate/sconfig"

// StreamFactory @todo doc
type StreamFactory []StreamStrategy

// Register will register a new stream factory strategy to be used
// on creation requests.
func (f *StreamFactory) Register(strategy StreamStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	*f = append(*f, strategy)

	return nil
}

// Create will instantiate and return a new config stream.
func (f StreamFactory) Create(streamType string, args ...interface{}) (Stream, error) {
	for _, s := range f {
		if s.Accept(streamType) {
			return s.Create(args...)
		}
	}
	return nil, errInvalidStreamType("streamType")
}

// CreateFromConfig will instantiate and return a new config stream loaded
// by a configuration instance.
func (f StreamFactory) CreateFromConfig(cfg sconfig.Config) (Stream, error) {
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
