package json

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
	t.Run("accept only json format", func(t *testing.T) {
		scenarios := []struct {
			format   string
			expected bool
		}{
			{ // _test json format
				format:   Format,
				expected: true,
			},
			{ // _test non-json format
				format:   config.UnknownDecoderFormat,
				expected: false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				if check := (DecoderStrategy{}).Accept(scenario.format); check != scenario.expected {
					t.Errorf("returned (%v) when checking (%v) format", check, scenario.format)
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
			t.Errorf("returned the (%v) error when expecting : %v", e, slate.ErrNilPointer)
		}
	})

	t.Run("invalid reader instance", func(t *testing.T) {
		if decoder, e := (DecoderStrategy{}).Create("string"); decoder != nil {
			t.Error("returned an unexpected valid decoder instance")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting : %v", e, slate.ErrConversion)
		}
	})

	t.Run("create the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoder, e := (DecoderStrategy{}).Create(NewMockReader(ctrl))
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case decoder == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch decoder.(type) {
			case *Decoder:
			default:
				t.Error("didn't returned a JSON decoder")
			}
		}
	})
}
