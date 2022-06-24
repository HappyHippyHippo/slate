package sconfig

import (
	"io"
)

type decoderStrategyJSON struct{}

var _ DecoderStrategy = &decoderStrategyJSON{}

// Accept will check if the decoder factory strategy can instantiate a
// decoder giving the format and the creation request parameters.
func (decoderStrategyJSON) Accept(format string) bool {
	return format == DecoderFormatJSON
}

// Create will instantiate the desired decoder instance with the given reader
// instance as source of the content to decode.
func (decoderStrategyJSON) Create(args ...interface{}) (Decoder, error) {
	reader, ok := args[0].(io.Reader)
	if !ok {
		return nil, errConversion(args[0], "io.Reader")
	}
	return newDecoderJSON(reader)
}
