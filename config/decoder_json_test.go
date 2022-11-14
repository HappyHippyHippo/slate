package config

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_NewDecoderJSON(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		sut, e := NewDecoderJSON(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serror.ErrNilPointer)
		}
	})

	t.Run("new json decoder adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)

		if sut, e := NewDecoderJSON(reader); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else {
			defer func() { _ = sut.Close() }()
			if e != nil {
				t.Errorf("returned the (%v) error", e)
			} else if sut.(*decoderJSON).reader != reader {
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
		sut, _ := NewDecoderJSON(reader)

		if e := sut.Close(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("call close method on reader only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := NewDecoderJSON(reader)

		_ = sut.Close()
		_ = sut.Close()
	})
}

func Test_DecoderJSON_Decode(t *testing.T) {
	t.Run("return decode error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := NewDecoderJSON(reader)
		defer func() { _ = sut.Close() }()
		json := NewMockJsoner(ctrl)
		json.EXPECT().Decode(&map[string]interface{}{}).DoAndReturn(func(p *map[string]interface{}) error {
			return expected
		}).Times(1)
		sut.(*decoderJSON).decoder = json

		check, e := sut.Decode()
		switch {
		case check != nil:
			t.Error("returned an reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("redirect to the underlying decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{"node": "data"}
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := NewDecoderJSON(reader)
		defer func() { _ = sut.Close() }()
		json := NewMockJsoner(ctrl)
		json.EXPECT().Decode(&map[string]interface{}{}).DoAndReturn(func(p *map[string]interface{}) error {
			(*p)["node"] = data["node"]
			return nil
		}).Times(1)
		sut.(*decoderJSON).decoder = json

		check, e := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case !reflect.DeepEqual(*check.(*Partial), data):
			t.Errorf("returned (%v)", check)
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})

	t.Run("decode json string", func(t *testing.T) {
		json := `{"node": {"sub_node": "data"}}`
		expected := Partial{"node": Partial{"sub_node": "data"}}
		reader := strings.NewReader(json)
		sut, _ := NewDecoderJSON(reader)
		defer func() { _ = sut.Close() }()

		check, e := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case !reflect.DeepEqual(*check.(*Partial), expected):
			t.Errorf("returned (%v)", check)
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}
