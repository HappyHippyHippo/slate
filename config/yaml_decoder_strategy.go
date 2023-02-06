package config

import (
	"io"
)

const (
	// DecoderFormatYAML defines the value to be used to declare
	// a YAML config source format.
	DecoderFormatYAML = "yaml"
)

// YAMLDecoderStrategy defines a YAML config decoder
// instantiation strategy
type YAMLDecoderStrategy struct{}

var _ IDecoderStrategy = &YAMLDecoderStrategy{}

// Accept will check if the decoder factory strategy can instantiate a
// decoder giving the format and the creation request parameters.
func (YAMLDecoderStrategy) Accept(
	format string,
) bool {
	// only accepts YAML format
	return format == DecoderFormatYAML
}

// Create will instantiate the desired decoder instance with the given reader
// instance as source of the content to decode.
func (YAMLDecoderStrategy) Create(
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
	return NewYAMLDecoder(reader)
}
