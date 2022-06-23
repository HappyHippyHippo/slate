package gconfig

import (
	"io"
)

// DecoderStrategyJSON defines a strategy used to instantiate
// a JSON config logStream decoder.
type DecoderStrategyJSON struct{}

var _ DecoderStrategy = &DecoderStrategyJSON{}

// Accept will check if the decoder factory strategy can instantiate a
// decoder giving the format and the creation request parameters.
func (DecoderStrategyJSON) Accept(format string) bool {
	return format == DecoderFormatJSON
}

// Create will instantiate the desired decoder instance with the given reader
// instance as source of the content to decode.
func (DecoderStrategyJSON) Create(args ...interface{}) (Decoder, error) {
	reader, ok := args[0].(io.Reader)
	if !ok {
		return nil, errConversion(args[0], "io.Reader")
	}
	return NewDecoderJSON(reader)
}
