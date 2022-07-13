package sconfig

import (
	"io"
)

type decoderStrategyYAML struct{}

var _ IDecoderStrategy = &decoderStrategyYAML{}

// Accept will check if the decoder factory strategy can instantiate a
// decoder giving the format and the creation request parameters.
func (decoderStrategyYAML) Accept(format string) bool {
	return format == DecoderFormatYAML
}

// Create will instantiate the desired decoder instance with the given reader
// instance as source of the content to decode.
func (decoderStrategyYAML) Create(args ...interface{}) (IDecoder, error) {
	reader, ok := args[0].(io.Reader)
	if !ok {
		return nil, errConversion(args[0], "io.Reader")
	}
	return newDecoderYAML(reader)
}
