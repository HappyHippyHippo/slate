package slog

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"testing"
)

func Test_FormatterFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if err := (&FormatterFactory{}).Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockFormatterStrategy(ctrl)
		factory := &FormatterFactory{}

		if err := factory.Register(strategy); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if (*factory)[0] != strategy {
			t.Errorf("didn't stored the s")
		}
	})
}

func Test_FormatterFactory_Create(t *testing.T) {
	t.Run("unrecognized format", func(t *testing.T) {
		format := "invalid format"

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockFormatterStrategy(ctrl)
		strategy.EXPECT().Accept(format).Return(false).Times(1)
		factory := &FormatterFactory{}
		_ = factory.Register(strategy)

		res, err := factory.Create(format)
		switch {
		case res != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidLogFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidLogFormat)
		}
	})

	t.Run("create the formatter", func(t *testing.T) {
		format := "format"

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		formatter := &FormatterJSON{}
		strategy := NewMockFormatterStrategy(ctrl)
		strategy.EXPECT().Accept(format).Return(true).Times(1)
		strategy.EXPECT().Create().Return(formatter, nil).Times(1)
		factory := &FormatterFactory{}
		_ = factory.Register(strategy)

		if res, err := factory.Create(format); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !reflect.DeepEqual(res, formatter) {
			t.Errorf("didn't returned the formatter")
		}
	})
}
