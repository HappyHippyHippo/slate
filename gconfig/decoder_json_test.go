package gconfig

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gerror"
	"reflect"
	"strings"
	"testing"
)

func Test_NewDecoderJSON(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		decoder, err := NewDecoderJSON(nil)
		switch {
		case decoder != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("new json decoder adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)

		if decoder, err := NewDecoderJSON(reader); decoder == nil {
			t.Errorf("didn't returned a valid reference")
		} else {
			defer func() { _ = decoder.Close() }()
			if err != nil {
				t.Errorf("returned the (%v) error", err)
			} else if decoder.(*decoderJSON).reader != reader {
				t.Error("didn't store the reader reference")
			}
		}
	})
}

func Test_DecoderJSON_Close(t *testing.T) {
	t.Run("error while closing the reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Return(expected).Times(1)
		decoder, _ := NewDecoderJSON(reader)

		if err := decoder.Close(); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if err.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("call close method on reader only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		decoder, _ := NewDecoderJSON(reader)

		_ = decoder.Close()
		_ = decoder.Close()
	})
}

func Test_DecoderJSON_Decode(t *testing.T) {
	t.Run("return decode error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		decoder, _ := NewDecoderJSON(reader)
		defer func() { _ = decoder.Close() }()
		json := NewMockJsoner(ctrl)
		json.EXPECT().Decode(&map[string]interface{}{}).DoAndReturn(func(p *map[string]interface{}) error {
			return expected
		}).Times(1)
		decoder.(*decoderJSON).decoder = json

		check, err := decoder.Decode()
		switch {
		case check != nil:
			t.Error("returned an reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case err.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", err, expected)
		}
	})

	t.Run("redirect to the underlying decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{"node": "data"}
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		decoder, _ := NewDecoderJSON(reader)
		defer func() { _ = decoder.Close() }()
		json := NewMockJsoner(ctrl)
		json.EXPECT().Decode(&map[string]interface{}{}).DoAndReturn(func(p *map[string]interface{}) error {
			(*p)["node"] = data["node"]
			return nil
		}).Times(1)
		decoder.(*decoderJSON).decoder = json

		check, err := decoder.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case !reflect.DeepEqual(*check.(*Partial), data):
			t.Errorf("returned (%v)", check)
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})

	t.Run("decode json string", func(t *testing.T) {
		json := `{"node": {"subnode": "data"}}`
		expected := Partial{"node": Partial{"subnode": "data"}}
		reader := strings.NewReader(json)
		decoder, _ := NewDecoderJSON(reader)
		defer func() { _ = decoder.Close() }()

		check, err := decoder.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case !reflect.DeepEqual(*check.(*Partial), expected):
			t.Errorf("returned (%v)", check)
		case err != nil:
			t.Errorf("returned the unexpected (%v) error", err)
		}
	})
}
