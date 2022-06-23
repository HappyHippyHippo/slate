package gconfig

// DecoderFactory defined the instance used to instantiate a new config
// logStream decoder for a specific encoding format.
type DecoderFactory []DecoderStrategy

// Register will store a new decoder factory strategy to be used
// to evaluate a request of an instance capable to parse a specific format.
// If the strategy accepts the format, then it will be used to instantiate the
// appropriate decoder that will be used to decode the configuration content.
func (f *DecoderFactory) Register(strategy DecoderStrategy) error {
	if strategy == nil {
		return errNilPointer("strategy")
	}

	*f = append(*f, strategy)

	return nil
}

// Create will instantiate the requested new decoder capable to
// parse the formatted content into a usable configuration Partial.
func (f DecoderFactory) Create(format string, args ...interface{}) (Decoder, error) {
	for _, s := range f {
		if s.Accept(format) {
			return s.Create(args...)
		}
	}
	return nil, errInvalidConfigDecoderFormat(format)
}
