package store

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/happyhippyhippo/slate/cache"
)

// Store @todo doc
type Store struct {
	Expiration time.Duration
}

// NormalizeExpire @todo doc
func (s Store) NormalizeExpire(
	expire time.Duration,
) time.Duration {
	switch expire {
	case cache.DEFAULT:
		return s.Expiration
	case cache.FOREVER:
		return time.Duration(0)
	}
	return expire
}

// Serialize @todo doc
func (Store) Serialize(
	value interface{},
) ([]byte, error) {
	// check if the value can be directly converted into an array of bytes
	if b, ok := value.([]byte); ok {
		return b, nil
	}
	// gob encoding of the data
	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)
	if e := encoder.Encode(value); e != nil {
		return nil, e
	}
	// return the encoding result
	return b.Bytes(), nil
}

// Deserialize @todo doc
func (Store) Deserialize(
	byt []byte,
	ptr interface{},
) (e error) {
	// check if the given pointer to an array of bytes
	// meaning that can be directly used to Store the source byte array
	if b, ok := ptr.(*[]byte); ok {
		*b = byt
		return nil
	}
	// gob decoding of the data
	b := bytes.NewBuffer(byt)
	decoder := gob.NewDecoder(b)
	return decoder.Decode(ptr)
}
