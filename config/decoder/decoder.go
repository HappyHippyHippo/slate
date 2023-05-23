package decoder

import (
	"io"

	"github.com/happyhippyhippo/slate/config"
)

// UnderlyingDecoder defines the interface to a decoder underlying
// format decoder
type UnderlyingDecoder interface {
	Decode(partial interface{}) error
}

// Decoder defines a config source content decoder instance.
type Decoder struct {
	Reader            io.Reader
	UnderlyingDecoder UnderlyingDecoder
}

// Close terminate the decoder, closing the associated underlying decoder.
func (d *Decoder) Close() error {
	// check if there is a jsonReader reference
	if d.Reader != nil {
		// check if the underlying decoder implements the closer interface
		if r, ok := d.Reader.(io.Closer); ok {
			// close the underlying decoder
			if e := r.Close(); e != nil {
				return e
			}
		}
		d.Reader = nil
	}
	return nil
}

// Decode parse the associated configuration source encoded content
// into a configuration instance.
func (d *Decoder) Decode() (*config.Partial, error) {
	// decode the referenced data
	data := map[string]interface{}{}
	if e := d.UnderlyingDecoder.Decode(&data); e != nil {
		return nil, e
	}
	// convert the read data into a normalized partial
	p := config.Convert(data).(config.Partial)
	return &p, nil
}
