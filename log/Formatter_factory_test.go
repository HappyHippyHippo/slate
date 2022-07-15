package log

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
	"reflect"
	"testing"
)

func Test_FormatterFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&FormatterFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("register the formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockFormatterStrategy(ctrl)
		sut := &FormatterFactory{}

		if e := sut.Register(strategy); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if (*sut)[0] != strategy {
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
		sut := &FormatterFactory{}
		_ = sut.Register(strategy)

		res, e := sut.Create(format)
		switch {
		case res != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrInvalidLogFormat):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrInvalidLogFormat)
		}
	})

	t.Run("create the formatter", func(t *testing.T) {
		format := "format"

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		formatter := NewMockFormatter(ctrl)
		strategy := NewMockFormatterStrategy(ctrl)
		strategy.EXPECT().Accept(format).Return(true).Times(1)
		strategy.EXPECT().Create().Return(formatter, nil).Times(1)
		sut := &FormatterFactory{}
		_ = sut.Register(strategy)

		if res, e := sut.Create(format); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(res, formatter) {
			t.Errorf("didn't returned the formatter")
		}
	})
}
