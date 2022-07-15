package config

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
	"reflect"
	"testing"
)

func Test_SourceFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&sourceFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, err.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("register the source dFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockSourceStrategy(ctrl)
		sut := &sourceFactory{}

		if e := sut.Register(strategy); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if (*sut)[0] != strategy {
			t.Error("didn't stored the strategy")
		}
	})
}

func Test_SourceFactory_Create(t *testing.T) {
	t.Run("error on unrecognized format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourceType := "type"
		path := "path"
		format := "format"
		strategy := NewMockSourceStrategy(ctrl)
		strategy.EXPECT().Accept(sourceType).Return(false).Times(1)
		sut := &sourceFactory{}
		_ = sut.Register(strategy)

		src, e := sut.Create(sourceType, path, format)
		switch {
		case src != nil:
			t.Error("returned an unexpected source")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrInvalidConfigSourceType):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrInvalidConfigSourceType)
		}
	})

	t.Run("create the requested config source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourceType := "type"
		path := "path"
		format := "format"
		src := &source{}
		strategy := NewMockSourceStrategy(ctrl)
		strategy.EXPECT().Accept(sourceType).Return(true).Times(1)
		strategy.EXPECT().Create(path, format).Return(src, nil).Times(1)
		sut := &sourceFactory{}
		_ = sut.Register(strategy)

		if check, e := sut.Create(sourceType, path, format); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(check, src) {
			t.Error("didn't returned the created source")
		}
	})
}

func Test_SourceFactory_CreateConfig(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src, e := (&sourceFactory{}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrNilPointer)
		}
	})

	t.Run("error on unrecognized format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{}
		strategy := NewMockSourceStrategy(ctrl)
		strategy.EXPECT().AcceptFromConfig(&data).Return(false).Times(1)
		sut := &sourceFactory{}
		_ = sut.Register(strategy)

		src, e := sut.CreateFromConfig(&data)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.ErrInvalidConfigSourcePartial):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.ErrInvalidConfigSourcePartial)
		}
	})

	t.Run("create the config source", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{}
		src := NewMockSource(ctrl)
		strategy := NewMockSourceStrategy(ctrl)
		strategy.EXPECT().AcceptFromConfig(&data).Return(true).Times(1)
		strategy.EXPECT().CreateFromConfig(&data).Return(src, nil).Times(1)
		sut := &sourceFactory{}
		_ = sut.Register(strategy)

		if check, e := sut.CreateFromConfig(&data); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(check, src) {
			t.Error("didn't returned the created src")
		}
	})
}
