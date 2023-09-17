package config

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
)

func Test_NewSourceFactory_Register(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		if NewSourceFactory() == nil {
			t.Error("didn't returned the expected reference")
		}
	})
}

func Test_SourceFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&SourceFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrNilPointer) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("register the source dFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockSourceStrategy(ctrl)
		sut := &SourceFactory{}

		if e := sut.Register(strategy); e != nil {
			t.Errorf("unexpected (%v) error", e)
		} else if (*sut)[0] != strategy {
			t.Error("didn't stored the strategy")
		}
	})
}

func Test_SourceFactory_Create(t *testing.T) {
	t.Run("nil partial", func(t *testing.T) {
		sut := &SourceFactory{}
		src, e := sut.Create(nil)
		switch {
		case src != nil:
			t.Error("returned an unexpected source")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error on unrecognized format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourceType := "type"
		path := "path"
		format := "format"
		partial := Partial{
			"type":   sourceType,
			"path":   path,
			"format": format,
		}
		strategy := NewMockSourceStrategy(ctrl)
		strategy.EXPECT().Accept(partial).Return(false).Times(1)
		sut := &SourceFactory{}
		_ = sut.Register(strategy)

		src, e := sut.Create(partial)
		switch {
		case src != nil:
			t.Error("returned an unexpected source")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, ErrInvalidSource):
			t.Errorf("(%v) when expecting (%v)", e, ErrInvalidSource)
		}
	})

	t.Run("create the requested partial source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourceType := "type"
		path := "path"
		format := "format"
		partial := Partial{
			"type":   sourceType,
			"path":   path,
			"format": format,
		}
		src := NewMockSource(ctrl)
		strategy := NewMockSourceStrategy(ctrl)
		strategy.EXPECT().Accept(partial).Return(true).Times(1)
		strategy.EXPECT().Create(partial).Return(src, nil).Times(1)
		sut := &SourceFactory{}
		_ = sut.Register(strategy)

		if check, e := sut.Create(partial); e != nil {
			t.Errorf("unexpected (%v) error", e)
		} else if !reflect.DeepEqual(check, src) {
			t.Error("didn't returned the created source")
		}
	})
}
