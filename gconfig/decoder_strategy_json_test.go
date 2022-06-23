package gconfig

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gerror"
	"testing"
)

func Test_DecoderStrategyJSON_Accept(t *testing.T) {
	t.Run("accept only json format", func(t *testing.T) {
		scenarios := []struct {
			format   string
			expected bool
		}{
			{ // _test json format
				format:   DecoderFormatJSON,
				expected: true,
			},
			{ // _test non-json format
				format:   DecoderFormatUnknown,
				expected: false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				if check := (DecoderStrategyJSON{}).Accept(scenario.format); check != scenario.expected {
					t.Errorf("returned (%v) when checking (%v) format", check, scenario.format)
				}
			}
			test()
		}
	})
}

func Test_DecoderStrategyJSON_Create(t *testing.T) {
	t.Run("invalid reader instance", func(t *testing.T) {
		if decoder, err := (DecoderStrategyJSON{}).Create("string"); decoder != nil {
			t.Error("returned an unexpected valid decoder instance")
		} else if !errors.Is(err, gerror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting : %v", err, gerror.ErrConversion)
		}
	})

	t.Run("create the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoder, err := (DecoderStrategyJSON{}).Create(NewMockReader(ctrl))
		switch {
		case err != nil:
			t.Errorf("returned the (%v) error", err)
		case decoder == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch decoder.(type) {
			case *decoderJSON:
			default:
				t.Error("didn't returned a JSON decoder")
			}
		}
	})
}
