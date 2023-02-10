package config

// IDecoderFactory defined the interface of a config decoder
// factory instance.
type IDecoderFactory interface {
	Register(strategy IDecoderStrategy) error
	Create(format string, args ...interface{}) (IDecoder, error)
}

// DecoderFactory defines a decoder instantiation factory.
type DecoderFactory []IDecoderStrategy

var _ IDecoderFactory = &DecoderFactory{}

// NewDecoderFactory £todo doc
func NewDecoderFactory() IDecoderFactory {
	return &DecoderFactory{}
}

// Register will store a new decoder factory strategy to be used
// to evaluate a request of an instance capable to parse a specific format.
// If the strategy accepts the format, then it will be used to instantiate the
// appropriate decoder that will be used to decode the configuration content.
func (f *DecoderFactory) Register(
	strategy IDecoderStrategy,
) error {
	// check for a valid strategy reference
	if strategy == nil {
		return errNilPointer("strategy")
	}
	// store the strategy reference
	*f = append(*f, strategy)
	return nil
}

// Create will instantiate the requested new decoder capable to
// parse the formatted content into a usable configuration.
func (f DecoderFactory) Create(
	format string,
	args ...interface{},
) (IDecoder, error) {
	// find a stored strategy that will accept the requested format
	for _, s := range f {
		if s.Accept(format) {
			// return the decoder instantiation
			return s.Create(args...)
		}
	}
	// signal that no strategy was found that would accept the requested format
	return nil, errInvalidFormat(format)
}
