package json

import (
	gojson "encoding/json"
	"io"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/decoder"
)

// Decoder defines a config source JSON content decoder instance.
type Decoder struct {
	decoder.Decoder
}

var _ config.Decoder = &Decoder{}

// NewDecoder will instantiate a new JSON configuration decoder.
func NewDecoder(
	reader io.Reader,
) (*Decoder, error) {
	// validate the reader reference
	if reader == nil {
		return nil, errNilPointer("jsonReader")
	}
	// return the new decoder reference
	return &Decoder{
		Decoder: decoder.Decoder{
			Reader:            reader,
			UnderlyingDecoder: gojson.NewDecoder(reader),
		},
	}, nil
}
