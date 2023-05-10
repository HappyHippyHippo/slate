package yaml

import (
	"io"

	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/config/decoder"
	goyaml "gopkg.in/yaml.v3"
)

// Decoder defines a config source YAML content decoder instance.
type Decoder struct {
	decoder.Decoder
}

var _ config.IDecoder = &Decoder{}

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

// Decode parse the associated configuration source reader content
// into a configuration.
func (d Decoder) Decode() (config.IConfig, error) {
	// decode the referenced reader data
	data := config.Config{}
	if e := d.UnderlyingDecoder.Decode(&data); e != nil {
		return nil, e
	}
	// convert the read data into a normalized config
	p := config.Convert(data).(config.Config)
	return &p, nil
}
