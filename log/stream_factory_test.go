package log

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewStreamFactory(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		sut := NewStreamFactory()
		if sut == nil {
			t.Error("didn't returned the expected reference")
		}
	})
}

func Test_StreamFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&StreamFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register the stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockStreamStrategy(ctrl)
		sut := &StreamFactory{}

		if e := sut.Register(strategy); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if (*sut)[0] != strategy {
			t.Error("didn't stored the strategy")
		}
	})
}

func Test_StreamFactory_Create(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		src, e := (&StreamFactory{}).Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("unrecognized stream type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.Partial{}
		strategy := NewMockStreamStrategy(ctrl)
		strategy.EXPECT().Accept(cfg).Return(false).Times(1)

		sut := &StreamFactory{}
		_ = sut.Register(strategy)

		stream, e := sut.Create(cfg)
		switch {
		case stream != nil:
			t.Error("returned a config stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrInvalidStream):
			t.Errorf("returned the (%v) error when expecting (%v)", e, ErrInvalidStream)
		}
	})

	t.Run("create the config stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.Partial{}
		stream := NewMockStream(ctrl)
		strategy := NewMockStreamStrategy(ctrl)
		strategy.EXPECT().Accept(cfg).Return(true).Times(1)
		strategy.EXPECT().Create(cfg).Return(stream, nil).Times(1)

		sut := &StreamFactory{}
		_ = sut.Register(strategy)

		if s, e := sut.Create(cfg); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(s, stream) {
			t.Error("didn't returned the created stream")
		}
	})
}
