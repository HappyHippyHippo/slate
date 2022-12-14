package config

import (
	"io"

	"gopkg.in/yaml.v3"
)

type yamlReader interface {
	Decode(partial interface{}) error
}

// YAMLDecoder defines a config source YAML content decoder instance.
type YAMLDecoder struct {
	reader  io.Reader
	decoder yamlReader
}

var _ IDecoder = &YAMLDecoder{}

// NewYAMLDecoder will instantiate a new YAML configuration decoder.
func NewYAMLDecoder(
	reader io.Reader,
) (*YAMLDecoder, error) {
	// validate the reader reference
	if reader == nil {
		return nil, errNilPointer("reader")
	}
	// return the new decoder reference
	return &YAMLDecoder{
		reader:  reader,
		decoder: yaml.NewDecoder(reader),
	}, nil
}

// Close terminate the decoder, closing the associated reader.
func (d *YAMLDecoder) Close() error {
	// check if there is a reader reference
	if d.reader != nil {
		// check if the reader implements the closer interface
		if r, ok := d.reader.(io.Closer); ok {
			// close the reader
			if e := r.Close(); e != nil {
				return e
			}
		}
		d.reader = nil
	}
	return nil
}

// Decode parse the associated configuration source reader content
// into a configuration.
func (d YAMLDecoder) Decode() (IConfig, error) {
	// decode the referenced reader data
	data := Config{}
	if e := d.decoder.Decode(&data); e != nil {
		return nil, e
	}
	// convert the read data into a normalized config
	p := (Config{}).convert(data).(Config)
	return &p, nil
}
