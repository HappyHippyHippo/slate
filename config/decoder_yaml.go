package config

import (
	"io"

	"gopkg.in/yaml.v2"
)

type yamler interface {
	Decode(partial interface{}) error
}

type decoderYAML struct {
	reader  io.Reader
	decoder yamler
}

var _ IDecoder = &decoderYAML{}

// NewDecoderYAML will instantiate a new YAML configuration decoder.
func NewDecoderYAML(reader io.Reader) (IDecoder, error) {
	if reader == nil {
		return nil, errNilPointer("reader")
	}

	return &decoderYAML{
		reader:  reader,
		decoder: yaml.NewDecoder(reader),
	}, nil
}

// Close terminate the decoder, closing the associated reader.
func (d *decoderYAML) Close() error {
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
func (d decoderYAML) Decode() (IConfig, error) {
	data := Partial{}
	if e := d.decoder.Decode(&data); e != nil {
		return nil, e
	}
	p := (Partial{}).convert(data).(Partial)

	return &p, nil
}
