package gconfig

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/gerror"
	"reflect"
	"testing"
)

func Test_SourceFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if err := (&SourceFactory{}).Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, gerror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("register the source factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockSourceStrategy(ctrl)
		factory := &SourceFactory{}

		if err := factory.Register(strategy); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if (*factory)[0] != strategy {
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
		factory := &SourceFactory{}
		_ = factory.Register(strategy)

		src, err := factory.Create(sourceType, path, format)
		switch {
		case src != nil:
			t.Error("returned an unexpected source")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrInvalidConfigSourceType):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrInvalidConfigSourceType)
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
		factory := &SourceFactory{}
		_ = factory.Register(strategy)

		if check, err := factory.Create(sourceType, path, format); err != nil {
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

		src, err := (&SourceFactory{}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrNilPointer)
		}
	})

	t.Run("error on unrecognized format", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		data := Partial{}
		strategy := NewMockSourceStrategy(ctrl)
		strategy.EXPECT().AcceptFromConfig(&data).Return(false).Times(1)
		factory := &SourceFactory{}
		_ = factory.Register(strategy)

		src, err := factory.CreateFromConfig(&data)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, gerror.ErrInvalidConfigSourcePartial):
			t.Errorf("returned the (%v) error when expecting (%v)", err, gerror.ErrInvalidConfigSourcePartial)
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
		factory := &SourceFactory{}
		_ = factory.Register(strategy)

		if check, err := factory.CreateFromConfig(&data); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !reflect.DeepEqual(check, src) {
			t.Error("didn't returned the created src")
		}
	})
}
