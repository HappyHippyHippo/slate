package yaml

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewDecoder(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		sut, e := NewDecoder(nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new yaml decoder adapter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)

		if sut, e := NewDecoder(reader); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else {
			defer func() { _ = sut.Close() }()
			if e != nil {
				t.Errorf("returned the (%v) error", e)
			} else if sut.Reader != reader {
				t.Error("didn't store the reader reference")
			}
		}
	})
}

func Test_Decoder_Close(t *testing.T) {
	t.Run("error while closing the reader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Return(expected).Times(1)
		sut, _ := NewDecoder(reader)

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
		sut, _ := NewDecoder(reader)

		_ = sut.Close()
		_ = sut.Close()
	})
}

func Test_Decoder_Decode(t *testing.T) {
	t.Run("return decode error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := NewDecoder(reader)
		defer func() { _ = sut.Close() }()
		underlyingDecoder := NewMockUnderlyingDecoder(ctrl)
		underlyingDecoder.EXPECT().Decode(&config.Partial{}).DoAndReturn(func(p *config.Partial) error { return expected }).Times(1)
		sut.UnderlyingDecoder = underlyingDecoder

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

		data := config.Partial{"node": "data"}
		reader := NewMockReader(ctrl)
		reader.EXPECT().Close().Times(1)
		sut, _ := NewDecoder(reader)
		defer func() { _ = sut.Close() }()
		underlyingDecoder := NewMockUnderlyingDecoder(ctrl)
		underlyingDecoder.EXPECT().Decode(&config.Partial{}).DoAndReturn(func(p *config.Partial) error {
			(*p)["node"] = data["node"]
			return nil
		}).Times(1)
		sut.UnderlyingDecoder = underlyingDecoder

		check, e := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case !reflect.DeepEqual(*check, data):
			t.Errorf("returned (%v)", check)
		}
	})

	t.Run("decode yaml string", func(t *testing.T) {
		yaml := "node:\n  sub_node: data"
		expected := config.Partial{"node": config.Partial{"sub_node": "data"}}
		reader := strings.NewReader(yaml)
		sut, _ := NewDecoder(reader)
		defer func() { _ = sut.Close() }()

		check, e := sut.Decode()
		switch {
		case check == nil:
			t.Error("returned a nil data")
		case e != nil:
			t.Errorf("returned the unexpected (%v) error", e)
		case !reflect.DeepEqual(*check, expected):
			t.Errorf("returned (%v) instead of the expected (%v)", *check, expected)
		}
	})
}
