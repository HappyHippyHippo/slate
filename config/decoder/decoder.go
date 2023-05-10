package decoder

import (
	"io"

	"github.com/happyhippyhippo/slate/config"
)

// IUnderlyingDecoder defines the interface to a decoder underlying
// format decoder
type IUnderlyingDecoder interface {
	Decode(partial interface{}) error
}

// Decoder defines a config source JSON content decoder instance.
type Decoder struct {
	Reader            io.Reader
	UnderlyingDecoder IUnderlyingDecoder
}

// Close terminate the decoder, closing the associated jsonReader.
func (d *Decoder) Close() error {
	// check if there is a jsonReader reference
	if d.Reader != nil {
		// check if the jsonReader implements the closer interface
		if r, ok := d.Reader.(io.Closer); ok {
			// close the jsonReader
			if e := r.Close(); e != nil {
				return e
			}
		}
		d.Reader = nil
	}
	return nil
}

// Decode parse the associated configuration source jsonReader content
// into a configuration instance.
func (d *Decoder) Decode() (config.IConfig, error) {
	// decode the referenced jsonReader data
	data := map[string]interface{}{}
	if e := d.UnderlyingDecoder.Decode(&data); e != nil {
		return nil, e
	}
	// convert the read data into a normalized config
	p := config.Convert(data).(config.Config)
	return &p, nil
}
