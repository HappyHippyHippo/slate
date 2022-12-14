package watchdog

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_LogFormatterFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&LogFormatterFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.NilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("register the stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockLogFormatterStrategy(ctrl)
		sut := &LogFormatterFactory{}

		if e := sut.Register(strategy); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if (*sut)[0] != strategy {
			t.Error("didn't stored the strategy")
		}
	})
}

func Test_LogFormatterFactory_Create(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		src, e := (&LogFormatterFactory{}).Create(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("unrecognized stream type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfig(ctrl)
		strategy := NewMockLogFormatterStrategy(ctrl)
		strategy.EXPECT().Accept(cfg).Return(false).Times(1)

		sut := &LogFormatterFactory{}
		_ = sut.Register(strategy)

		stream, e := sut.Create(cfg)
		switch {
		case stream != nil:
			t.Error("returned a config stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.InvalidWatchdogConfig):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.InvalidWatchdogConfig)
		}
	})

	t.Run("create the config stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfig(ctrl)
		stream := NewMockLogFormatter(ctrl)
		strategy := NewMockLogFormatterStrategy(ctrl)
		strategy.EXPECT().Accept(cfg).Return(true).Times(1)
		strategy.EXPECT().Create(cfg).Return(stream, nil).Times(1)

		sut := &LogFormatterFactory{}
		_ = sut.Register(strategy)

		if s, e := sut.Create(cfg); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(s, stream) {
			t.Error("didn't returned the created stream")
		}
	})
}
