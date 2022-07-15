package config

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
	"reflect"
	"testing"
)

func Test_DecoderFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&decoderFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("register the decoder dFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockDecoderStrategy(ctrl)
		sut := decoderFactory{}

		if e := sut.Register(strategy); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if sut[0] != strategy {
			t.Error("didn't stored the strategy")
		}
	})
}

func Test_DecoderFactory_Create(t *testing.T) {
	t.Run("error if the format is unrecognized", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		format := DecoderFormatUnknown
		reader := NewMockReader(ctrl)
		strategy := NewMockDecoderStrategy(ctrl)
		strategy.EXPECT().Accept(format).Return(false).Times(1)
		sut := decoderFactory{}
		_ = sut.Register(strategy)

		check, e := sut.Create(format, reader)
		switch {
		case check != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrInvalidConfigDecoderFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrInvalidConfigDecoderFormat)
		}
	})

	t.Run("should create the requested decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		format := DecoderFormatUnknown
		reader := NewMockReader(ctrl)
		decoder := NewMockDecoder(ctrl)
		strategy := NewMockDecoderStrategy(ctrl)
		strategy.EXPECT().Accept(format).Return(true).Times(1)
		strategy.EXPECT().Create(reader).Return(decoder, nil).Times(1)
		sut := decoderFactory{}
		_ = sut.Register(strategy)

		if check, e := sut.Create(format, reader); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(check, decoder) {
			t.Error("didn't returned the created strategy")
		}
	})
}
