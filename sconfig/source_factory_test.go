package sconfig

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"testing"
)

func Test_SourceFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if err := (&sourceFactory{}).Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register the source dFactory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockSourceStrategy(ctrl)
		sFactory := &sourceFactory{}

		if err := sFactory.Register(strategy); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if (*sFactory)[0] != strategy {
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
		sFactory := &sourceFactory{}
		_ = sFactory.Register(strategy)

		src, err := sFactory.Create(sourceType, path, format)
		switch {
		case src != nil:
			t.Error("returned an unexpected source")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidConfigSourceType):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidConfigSourceType)
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
		sFactory := &sourceFactory{}
		_ = sFactory.Register(strategy)

		if check, err := sFactory.Create(sourceType, path, format); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !reflect.DeepEqual(check, src) {
			t.Error("didn't returned the created source")
		}
	})
}

func Test_SourceFactory_CreateConfig(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		src, err := (&sourceFactory{}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error on unrecognized format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{}
		strategy := NewMockSourceStrategy(ctrl)
		strategy.EXPECT().AcceptFromConfig(&data).Return(false).Times(1)
		sFactory := &sourceFactory{}
		_ = sFactory.Register(strategy)

		src, err := sFactory.CreateFromConfig(&data)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidConfigSourcePartial):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidConfigSourcePartial)
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
		sFactory := &sourceFactory{}
		_ = sFactory.Register(strategy)

		if check, err := sFactory.CreateFromConfig(&data); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !reflect.DeepEqual(check, src) {
			t.Error("didn't returned the created src")
		}
	})
}
