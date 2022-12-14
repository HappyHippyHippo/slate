package config

import (
	"encoding/json"
	"io"
)

type jsonReader interface {
	Decode(partial interface{}) error
}

// JSONDecoder defines a config source JSON content decoder instance.
type JSONDecoder struct {
	reader  io.Reader
	decoder jsonReader
}

var _ IDecoder = &JSONDecoder{}

// NewJSONDecoder will instantiate a new JSON configuration decoder.
func NewJSONDecoder(
	reader io.Reader,
) (*JSONDecoder, error) {
	// validate the reader reference
	if reader == nil {
		return nil, errNilPointer("reader")
	}
	// return the new decoder reference
	return &JSONDecoder{
		reader:  reader,
		decoder: json.NewDecoder(reader),
	}, nil
}

// Close terminate the decoder, closing the associated reader.
func (d *JSONDecoder) Close() error {
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
// into a configuration instance.
func (d JSONDecoder) Decode() (IConfig, error) {
	// decode the referenced reader data
	data := map[string]interface{}{}
	if e := d.decoder.Decode(&data); e != nil {
		return nil, e
	}
	// convert the read data into a normalized config
	p := (Config{}).convert(data).(Config)
	return &p, nil
}
