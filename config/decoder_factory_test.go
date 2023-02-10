package config

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_DecoderFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&DecoderFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("register the decoder factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockDecoderStrategy(ctrl)
		sut := DecoderFactory{}

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

		format := UnknownDecoderFormat
		reader := NewMockReader(ctrl)
		strategy := NewMockDecoderStrategy(ctrl)
		strategy.EXPECT().Accept(format).Return(false).Times(1)
		sut := DecoderFactory{}
		_ = sut.Register(strategy)

		check, e := sut.Create(format, reader)
		switch {
		case check != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.InvalidConfigFormat):
			t.Errorf("returned the (%v) err when expecting (%v)", e, err.InvalidConfigFormat)
		}
	})

	t.Run("should create the requested decoder", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		format := UnknownDecoderFormat
		reader := NewMockReader(ctrl)
		decoder := NewMockDecoder(ctrl)
		strategy := NewMockDecoderStrategy(ctrl)
		strategy.EXPECT().Accept(format).Return(true).Times(1)
		strategy.EXPECT().Create(reader).Return(decoder, nil).Times(1)
		sut := DecoderFactory{}
		_ = sut.Register(strategy)

		if check, e := sut.Create(format, reader); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(check, decoder) {
			t.Error("didn't returned the created strategy")
		}
	})
}
