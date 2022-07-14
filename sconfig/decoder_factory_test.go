package sconfig

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"testing"
)

func Test_DecoderFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if err := (&decoderFactory{}).Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register the decoder dFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockDecoderStrategy(ctrl)
		sut := decoderFactory{}

		if err := sut.Register(strategy); err != nil {
			t.Errorf("returned the (%v) error", err)
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

		check, err := sut.Create(format, reader)
		switch {
		case check != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidConfigDecoderFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidConfigDecoderFormat)
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

		if check, err := sut.Create(format, reader); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !reflect.DeepEqual(check, decoder) {
			t.Error("didn't returned the created strategy")
		}
	})
}
