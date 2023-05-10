package json

import (
	"io"

	"github.com/happyhippyhippo/slate/config"
)

const (
	// Format defines the value to be used to declare
	// a JSON config source format.
	Format = "json"
)

// DecoderStrategy defines a JSON config decoder
// instantiation strategy
type DecoderStrategy struct{}

var _ config.IDecoderStrategy = &DecoderStrategy{}

// NewDecoderStrategy will instantiate a new JSON format
// decoder creation strategy
func NewDecoderStrategy() *DecoderStrategy {
	return &DecoderStrategy{}
}

// Accept will check if the decoder factory strategy can instantiate a
// decoder giving the format and the creation request parameters.
func (DecoderStrategy) Accept(
	format string,
) bool {
	// only accepts JSON format
	return format == Format
}

// Create will instantiate the desired decoder instance with the given jsonReader
// instance as source of the content to decode.
func (DecoderStrategy) Create(
	args ...interface{},
) (config.IDecoder, error) {
	// check for the existence of the mandatory jsonReader argument
	if len(args) == 0 {
		return nil, errNilPointer("args[0]")
	}
	// validate the jsonReader argument
	reader, ok := args[0].(io.Reader)
	if !ok {
		return nil, errConversion(args[0], "io.Reader")
	}
	// create the decoder
	return NewDecoder(reader)
}
