package gconfig

import (
	"gopkg.in/yaml.v2"
	"io"
)

type yamler interface {
	Decode(partial interface{}) error
}

type decoderYAML struct {
	reader  io.Reader
	decoder yamler
}

var _ Decoder = &decoderYAML{}

// NewDecoderYAML instantiate a new yaml configuration decoder object
// used to parse a yaml configuration source into a config Partial.
func NewDecoderYAML(reader io.Reader) (Decoder, error) {
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
func (d decoderYAML) Decode() (Config, error) {
	data := Partial{}
	if err := d.decoder.Decode(&data); err != nil {
		return nil, err
	}
	p := (Partial{}).convert(data).(Partial)

	return &p, nil
}
