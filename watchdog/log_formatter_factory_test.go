package watchdog

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_LogFormatterFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&LogFormatterFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register the log formatter creator strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockLogFormatterStrategy(ctrl)
		sut := &LogFormatterFactory{}

		if e := sut.Register(strategy); e != nil {
			t.Errorf("unexpected (%v) error", e)
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
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("unrecognized format type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &config.Partial{}
		strategy := NewMockLogFormatterStrategy(ctrl)
		strategy.EXPECT().Accept(cfg).Return(false).Times(1)

		sut := &LogFormatterFactory{}
		_ = sut.Register(strategy)

		formatter, e := sut.Create(cfg)
		switch {
		case formatter != nil:
			t.Error("returned a log formatter")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrInvalidWatchdog):
			t.Errorf("(%v) when expecting (%v)", e, ErrInvalidWatchdog)
		}
	})

	t.Run("create the log formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := &config.Partial{}
		formatter := NewMockLogFormatter(ctrl)
		strategy := NewMockLogFormatterStrategy(ctrl)
		strategy.EXPECT().Accept(cfg).Return(true).Times(1)
		strategy.EXPECT().Create(cfg).Return(formatter, nil).Times(1)

		sut := &LogFormatterFactory{}
		_ = sut.Register(strategy)

		if f, e := sut.Create(cfg); e != nil {
			t.Errorf("unexpected (%v) error", e)
		} else if !reflect.DeepEqual(f, formatter) {
			t.Error("didn't returned the created stream")
		}
	})
}
