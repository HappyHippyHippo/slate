package gconfig

import (
	"encoding/json"
	"io"
)

type jsoner interface {
	Decode(partial interface{}) error
}

type decoderJSON struct {
	reader  io.Reader
	decoder jsoner
}

var _ Decoder = &decoderJSON{}

// NewDecoderJSON instantiate a new yaml configuration decoder object
// used to parse a yaml configuration source into a config Partial.
func NewDecoderJSON(reader io.Reader) (Decoder, error) {
	if reader == nil {
		return nil, errNilPointer("reader")
	}

	return &decoderJSON{
		reader:  reader,
		decoder: json.NewDecoder(reader),
	}, nil
}

// Close terminate the decoder, closing the associated reader.
func (d *decoderJSON) Close() error {
	if d.reader != nil {
		if r, ok := d.reader.(io.Closer); ok {
			if err := r.Close(); err != nil {
				return err
			}
		}
		d.reader = nil
	}
	return nil
}

// Decode parse the associated configuration source reader content
// into a configuration Partial.
func (d decoderJSON) Decode() (Config, error) {
	data := map[string]interface{}{}
	if err := d.decoder.Decode(&data); err != nil {
		return nil, err
	}
	p := (Partial{}).convert(data).(Partial)

	return &p, nil
}
