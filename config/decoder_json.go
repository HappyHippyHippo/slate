package config

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

var _ IDecoder = &decoderJSON{}

// NewDecoderJSON will instantiate a new JSON configuration decoder.
func NewDecoderJSON(reader io.Reader) (IDecoder, error) {
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
			if e := r.Close(); e != nil {
				return e
			}
		}
		d.reader = nil
	}
	return nil
}

// Decode parse the associated configuration source reader content
// into a configuration Partial.
func (d decoderJSON) Decode() (IConfig, error) {
	data := map[string]interface{}{}
	if e := d.decoder.Decode(&data); e != nil {
		return nil, e
	}
	p := (Partial{}).convert(data).(Partial)

	return &p, nil
}
