package yaml

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewDecoderStrategy(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		if NewDecoderStrategy() == nil {
			t.Error("didn't returned the expected reference")
		}
	})
}

func Test_DecoderStrategy_Accept(t *testing.T) {
	t.Run("accept only yaml format", func(t *testing.T) {
		scenarios := []struct {
			format   string
			expected bool
		}{
			{ // _test yaml format
				format:   Format,
				expected: true,
			},
			{ // _test non-yaml format
				format:   config.UnknownDecoder,
				expected: false,
			},
		}

		for _, s := range scenarios {
			test := func() {
				if check := (DecoderStrategy{}).Accept(s.format); check != s.expected {
					t.Errorf("returned (%v) when checking (%v)", check, s.format)
				}
			}
			test()
		}
	})
}

func Test_DecoderStrategy_Create(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		if decoder, e := (DecoderStrategy{}).Create(); decoder != nil {
			t.Error("returned an unexpected valid decoder instance")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("invalid reader instance", func(t *testing.T) {
		if decoder, e := (DecoderStrategy{}).Create("string"); decoder != nil {
			t.Error("returned an unexpected valid decoder instance")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("create the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoder, e := (DecoderStrategy{}).Create(NewMockReader(ctrl))
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case decoder == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch decoder.(type) {
			case *Decoder:
			default:
				t.Error("didn't returned a YAML decoder")
			}
		}
	})
}
