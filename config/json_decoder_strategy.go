package config

import (
	"io"
)

// JSONDecoderStrategy defines a JSON config decoder
// instantiation strategy
type JSONDecoderStrategy struct{}

var _ IDecoderStrategy = &JSONDecoderStrategy{}

// Accept will check if the decoder factory strategy can instantiate a
// decoder giving the format and the creation request parameters.
func (JSONDecoderStrategy) Accept(
	format string,
) bool {
	// only accepts JSON format
	return format == FormatJSON
}

// Create will instantiate the desired decoder instance with the given reader
// instance as source of the content to decode.
func (JSONDecoderStrategy) Create(
	args ...interface{},
) (IDecoder, error) {
	// check for the existence of the mandatory reader argument
	if len(args) == 0 {
		return nil, errNilPointer("args[0]")
	}
	// validate the reader argument
	reader, ok := args[0].(io.Reader)
	if !ok {
		return nil, errConversion(args[0], "io.Reader")
	}
	// create the decoder
	return NewJSONDecoder(reader)
}
