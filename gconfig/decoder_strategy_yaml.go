package gconfig

import (
	"io"
)

// DecoderStrategyYAML defines a strategy used to instantiate
// a YAML config logStream decoder.
type DecoderStrategyYAML struct{}

var _ DecoderStrategy = &DecoderStrategyYAML{}

// Accept will check if the decoder factory strategy can instantiate a
// decoder giving the format and the creation request parameters.
func (DecoderStrategyYAML) Accept(format string) bool {
	return format == DecoderFormatYAML
}

// Create will instantiate the desired decoder instance with the given reader
// instance as source of the content to decode.
func (DecoderStrategyYAML) Create(args ...interface{}) (Decoder, error) {
	reader, ok := args[0].(io.Reader)
	if !ok {
		return nil, errConversion(args[0], "io.Reader")
	}
	return NewDecoderYAML(reader)
}
