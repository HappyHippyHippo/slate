package yaml

import (
	"io"

	goyaml "gopkg.in/yaml.v3"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/decoder"
)

// Decoder defines a config source YAML content decoder instance.
type Decoder struct {
	decoder.Decoder
}

var _ config.Decoder = &Decoder{}

// NewDecoder will instantiate a new YAML configuration decoder.
func NewDecoder(
	reader io.Reader,
) (*Decoder, error) {
	// validate the reader reference
	if reader == nil {
		return nil, errNilPointer("reader")
	}
	// return the new decoder reference
	return &Decoder{
		Decoder: decoder.Decoder{
			Reader:            reader,
			UnderlyingDecoder: goyaml.NewDecoder(reader),
		},
	}, nil
}

// Decode parse the associated configuration source encoded content
// into a configuration instance.
func (d Decoder) Decode() (*config.Partial, error) {
	// decode the referenced reader data
	data := config.Partial{}
	if e := d.UnderlyingDecoder.Decode(&data); e != nil {
		return nil, e
	}
	// convert the read data into a normalized config
	p := config.Convert(data).(config.Partial)
	return &p, nil
}
