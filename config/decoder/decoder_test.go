package decoder

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
)

func Test_BaseDecoder_Close(t *testing.T) {
	t.Run("error while closing the jsonReader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Return(expected).Times(1)
		sut := Decoder{Reader: reader}

		if e := sut.Close(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("call close method on jsonReader only once", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut := Decoder{Reader: reader}

		_ = sut.Close()
		_ = sut.Close()
	})
}

func Test_BaseDecoder_Decode(t *testing.T) {
	t.Run("return decode error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut := Decoder{Reader: reader}
		defer func() { _ = sut.Close() }()
		baseDecoder := NewMockUnderlyingDecoder(ctrl)
		baseDecoder.EXPECT().Decode(&map[string]interface{}{}).DoAndReturn(func(p *map[string]interface{}) error {
			return expected
		}).Times(1)
		sut.UnderlyingDecoder = baseDecoder

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

		data := config.Config{"node": "data"}
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut := Decoder{Reader: reader}
		defer func() { _ = sut.Close() }()
		baseDecoder := NewMockUnderlyingDecoder(ctrl)
		baseDecoder.EXPECT().Decode(&map[string]interface{}{}).DoAndReturn(func(p *map[string]interface{}) error {
			(*p)["node"] = data["node"]
			return nil
		}).Times(1)
		sut.UnderlyingDecoder = baseDecoder

		check, e := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case !reflect.DeepEqual(*check.(*config.Config), data):
			t.Errorf("returned (%v)", check)
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		}
	})
}
