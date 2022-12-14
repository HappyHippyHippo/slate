package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewJSONDecoder(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		sut, e := NewJSONDecoder(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("new json decoder adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)

		if sut, e := NewJSONDecoder(reader); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else {
			defer func() { _ = sut.Close() }()
			if e != nil {
				t.Errorf("returned the (%v) error", e)
			} else if sut.reader != reader {
				t.Error("didn't store the reader reference")
			}
		}
	})
}

func Test_JSONDecoder_Close(t *testing.T) {
	t.Run("error while closing the reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Return(expected).Times(1)
		sut, _ := NewJSONDecoder(reader)

		if e := sut.Close(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("call close method on reader only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := NewJSONDecoder(reader)

		_ = sut.Close()
		_ = sut.Close()
	})
}

func Test_JSONDecoder_Decode(t *testing.T) {
	t.Run("return decode error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := NewJSONDecoder(reader)
		defer func() { _ = sut.Close() }()
		json := NewMockJSONReader(ctrl)
		json.EXPECT().Decode(&map[string]interface{}{}).DoAndReturn(func(p *map[string]interface{}) error {
			return expected
		}).Times(1)
		sut.decoder = json

		check, e := sut.Decode()
		switch {
		case check != nil:
			t.Error("returned an reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) err when expecting (%v)", e, expected)
		}
	})

	t.Run("redirect to the underlying decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Config{"node": "data"}
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := NewJSONDecoder(reader)
		defer func() { _ = sut.Close() }()
		json := NewMockJSONReader(ctrl)
		json.EXPECT().Decode(&map[string]interface{}{}).DoAndReturn(func(p *map[string]interface{}) error {
			(*p)["node"] = data["node"]
			return nil
		}).Times(1)
		sut.decoder = json

		check, e := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case !reflect.DeepEqual(*check.(*Config), data):
			t.Errorf("returned (%v)", check)
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("decode json string", func(t *testing.T) {
		json := `{"node": {"sub_node": "data"}}`
		expected := Config{"node": Config{"sub_node": "data"}}
		reader := strings.NewReader(json)
		sut, _ := NewJSONDecoder(reader)
		defer func() { _ = sut.Close() }()

		check, e := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case !reflect.DeepEqual(*check.(*Config), expected):
			t.Errorf("returned (%v)", check)
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}
