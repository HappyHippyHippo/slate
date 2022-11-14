package config

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	serror "github.com/happyhippyhippo/slate/error"
)

func Test_DecoderStrategyYAML_Accept(t *testing.T) {
	t.Run("accept only yaml format", func(t *testing.T) {
		scenarios := []struct {
			format   string
			expected bool
		}{
			{ // _test yaml format
				format:   DecoderFormatYAML,
				expected: true,
			},
			{ // _test non-yaml format
				format:   DecoderFormatUnknown,
				expected: false,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				if check := (decoderStrategyYAML{}).Accept(scenario.format); check != scenario.expected {
					t.Errorf("returned (%v) when checking (%v) format", check, scenario.format)
				}
			}
			test()
		}
	})
}

func Test_DecoderStrategyYAML_Create(t *testing.T) {
	t.Run("invalid reader instance", func(t *testing.T) {
		if decoder, e := (decoderStrategyYAML{}).Create("string"); decoder != nil {
			t.Error("returned an unexpected valid decoder instance")
		} else if !errors.Is(e, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting : %v", e, serror.ErrConversion)
		}
	})

	t.Run("create the decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		decoder, e := (decoderStrategyYAML{}).Create(NewMockReader(ctrl))
		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case decoder == nil:
			t.Error("didn't returned a valid reference")
		default:
			switch decoder.(type) {
			case *decoderYAML:
			default:
				t.Error("didn't returned a YAML decoder")
			}
		}
	})
}
