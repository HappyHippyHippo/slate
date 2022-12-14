package config

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_YAMLDecoderStrategy_Accept(t *testing.T) {
	t.Run("accept only yaml format", func(t *testing.T) {
		scenarios := []struct {
			format   string
			expected bool
		}{
			{ // _test yaml format
				format:   FormatYAML,
				expected: true,
			},
			{ // _test non-yaml format
				format:   FormatUnknown,
				expected: false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				if check := (YAMLDecoderStrategy{}).Accept(scenario.format); check != scenario.expected {
					t.Errorf("returned (%v) when checking (%v) format", check, scenario.format)
				}
			}
			test()
		}
	})
}

func Test_YAMLDecoderStrategy_Create(t *testing.T) {
	t.Run("nil reader", func(t *testing.T) {
		if decoder, e := (YAMLDecoderStrategy{}).Create(); decoder != nil {
			t.Error("returned an unexpected valid decoder instance")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) err when expecting : %v", e, err.NilPointer)
		}
	})

	t.Run("invalid reader instance", func(t *testing.T) {
		if decoder, e := (YAMLDecoderStrategy{}).Create("string"); decoder != nil {
			t.Error("returned an unexpected valid decoder instance")
		} else if !errors.Is(e, err.Conversion) {
			t.Errorf("returned the (%v) err when expecting : %v", e, err.Conversion)
		}
	})

	t.Run("create the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoder, e := (YAMLDecoderStrategy{}).Create(NewMockReader(ctrl))
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case decoder == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch decoder.(type) {
			case *YAMLDecoder:
			default:
				t.Error("didn't returned a YAML decoder")
			}
		}
	})
}
