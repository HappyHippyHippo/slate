package store

import (
	"reflect"
	"testing"
	"time"

	"github.com/happyhippyhippo/slate/cache"
)

func Test_normalizeExpire(t *testing.T) {
	t.Run("set default value if is equals to global default", func(t *testing.T) {
		defaultValue := time.Duration(100)
		sut := Store{Expiration: defaultValue}
		if check := sut.NormalizeExpire(cache.DEFAULT); check != defaultValue {
			t.Errorf("normalized to (%v) instead of the expected (%d)", check, defaultValue)
		}
	})

	t.Run("set zero if is equals to global forever", func(t *testing.T) {
		defaultValue := time.Duration(100)
		sut := Store{Expiration: defaultValue}
		if check := sut.NormalizeExpire(cache.FOREVER); check != time.Duration(0) {
			t.Errorf("normalized to (%v) instead of the expected (%d)", check, time.Duration(0))
		}
	})

	t.Run("no-op if not default or forever", func(t *testing.T) {
		defaultValue := time.Duration(100)
		value := time.Duration(200)
		sut := Store{Expiration: defaultValue}
		if check := sut.NormalizeExpire(value); check != value {
			t.Errorf("normalized to (%v) instead of the expected (%d)", check, value)
		}
	})
}

func Test_serialize(t *testing.T) {
	t.Run("no-op if value is a byte array", func(t *testing.T) {
		value := []byte("string")
		check, e := (Store{}).Serialize(value)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(check, value):
			t.Errorf("didn't returned the exepcted same byte array")
		}
	})

	t.Run("valid serialization", func(t *testing.T) {
		scenarios := []struct {
			value    interface{}
			expected []byte
		}{
			{ // int
				value:    123456,
				expected: []byte{6, 4, 0, 253, 3, 196, 128},
			},
			{ // int8
				value:    int8(123),
				expected: []byte{4, 4, 0, 255, 246},
			},
			{ // int16
				value:    int16(1234),
				expected: []byte{5, 4, 0, 254, 9, 164},
			},
			{ // int32
				value:    int32(123456),
				expected: []byte{6, 4, 0, 253, 3, 196, 128},
			},
			{ // int64
				value:    int64(123456123456),
				expected: []byte{8, 4, 0, 251, 57, 125, 29, 228, 128},
			},
			{ // uint
				value:    uint(123456),
				expected: []byte{6, 6, 0, 253, 1, 226, 64},
			},
			{ // uint8
				value:    uint8(123),
				expected: []byte{3, 6, 0, 123},
			},
			{ // uint16
				value:    uint16(1234),
				expected: []byte{5, 6, 0, 254, 4, 210},
			},
			{ // uint32
				value:    uint32(123456),
				expected: []byte{6, 6, 0, 253, 1, 226, 64},
			},
			{ // uint64
				value:    uint64(123456123456),
				expected: []byte{8, 6, 0, 251, 28, 190, 142, 242, 64},
			},
			{ // float32
				value:    float32(1.1),
				expected: []byte{8, 8, 0, 251, 160, 153, 153, 241, 63},
			},
			{ // float64
				value:    1.1,
				expected: []byte{11, 8, 0, 248, 154, 153, 153, 153, 153, 153, 241, 63},
			},
			{ // string
				value:    "test string",
				expected: []byte{14, 12, 0, 11, 116, 101, 115, 116, 32, 115, 116, 114, 105, 110, 103},
			},
			{ // struct
				value:    struct{ I int }{I: 123},
				expected: []byte{18, 255, 129, 3, 1, 2, 255, 130, 0, 1, 1, 1, 1, 73, 1, 4, 0, 0, 0, 6, 255, 130, 1, 255, 246, 0},
			},
		}

		for _, scenario := range scenarios {
			check, e := (Store{}).Serialize(scenario.value)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !reflect.DeepEqual(scenario.expected, check):
				t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", scenario.value, scenario.expected, check)
			}
		}
	})

	t.Run("error while trying to serialize the value", func(t *testing.T) {
		check, e := (Store{}).Serialize(nil)
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case check != nil:
			t.Error("unexpected valid reference to a serialized byte array")
		}
	})
}

func Test_deserialize(t *testing.T) {
	t.Run("no-op if value is a byte array", func(t *testing.T) {
		value := []byte("string")
		check := &[]byte{}
		e := (Store{}).Deserialize(value, check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(*check, value):
			t.Errorf("didn't returned the exepcted same byte array")
		}
	})

	t.Run("valid deserialization - byte", func(t *testing.T) {
		value := []byte{6, 4, 0, 253, 3, 196, 128}
		check := 0
		expected := 123456

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - int", func(t *testing.T) {
		value := []byte{6, 4, 0, 253, 3, 196, 128}
		check := 0
		expected := 123456

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - int8", func(t *testing.T) {
		value := []byte{4, 4, 0, 255, 246}
		check := int8(0)
		expected := int8(123)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - int16", func(t *testing.T) {
		value := []byte{5, 4, 0, 254, 9, 164}
		check := int16(0)
		expected := int16(1234)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - int32", func(t *testing.T) {
		value := []byte{6, 4, 0, 253, 3, 196, 128}
		check := int32(0)
		expected := int32(123456)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - int64", func(t *testing.T) {
		value := []byte{8, 4, 0, 251, 57, 125, 29, 228, 128}
		check := int64(0)
		expected := int64(123456123456)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - uint", func(t *testing.T) {
		value := []byte{6, 6, 0, 253, 1, 226, 64}
		check := uint(0)
		expected := uint(123456)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - uint8", func(t *testing.T) {
		value := []byte{3, 6, 0, 123}
		check := uint8(0)
		expected := uint8(123)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - uint16", func(t *testing.T) {
		value := []byte{5, 6, 0, 254, 4, 210}
		check := uint16(0)
		expected := uint16(1234)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - uint32", func(t *testing.T) {
		value := []byte{6, 6, 0, 253, 1, 226, 64}
		check := uint32(0)
		expected := uint32(123456)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - uint64", func(t *testing.T) {
		value := []byte{8, 6, 0, 251, 28, 190, 142, 242, 64}
		check := uint64(0)
		expected := uint64(123456123456)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - float32", func(t *testing.T) {
		value := []byte{8, 8, 0, 251, 160, 153, 153, 241, 63}
		check := float32(0)
		expected := float32(1.1)

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - float64", func(t *testing.T) {
		value := []byte{11, 8, 0, 248, 154, 153, 153, 153, 153, 153, 241, 63}
		check := float64(0)
		expected := 1.1

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - string", func(t *testing.T) {
		value := []byte{14, 12, 0, 11, 116, 101, 115, 116, 32, 115, 116, 114, 105, 110, 103}
		check := ""
		expected := "test string"

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("valid deserialization - struct", func(t *testing.T) {
		value := []byte{18, 255, 129, 3, 1, 2, 255, 130, 0, 1, 1, 1, 1, 73, 1, 4, 0, 0, 0, 6, 255, 130, 1, 255, 246, 0}
		check := struct{ I int }{}
		expected := struct{ I int }{I: 123}

		e := (Store{}).Deserialize(value, &check)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case !reflect.DeepEqual(expected, check):
			t.Errorf("didn't serialized (%v) to the expected (%v) : (%v)", value, expected, check)
		}
	})

	t.Run("error while trying to deserialize the value", func(t *testing.T) {
		if e := (Store{}).Deserialize(nil, nil); e == nil {
			t.Error("didn't returned the expected error")
		}
	})
}
