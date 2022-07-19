package slog

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/serr"
	"reflect"
	"testing"
)

func Test_StreamFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if e := (&StreamFactory{}).Register(nil); e == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
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
	t.Run("unrecognized stream type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		streamType := "type"
		path := "path"
		format := "format"
		strategy := NewMockStreamStrategy(ctrl)
		strategy.EXPECT().Accept(streamType).Return(false).Times(1)

		sut := &StreamFactory{}
		_ = sut.Register(strategy)

		stream, e := sut.Create(streamType, path, format)
		switch {
		case stream != nil:
			t.Error("returned an valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrInvalidLogStreamType):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrInvalidLogStreamType)
		}
	})

	t.Run("create the config stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sourceType := "type"
		path := "path"
		format := "format"
		stream := NewMockStream(ctrl)
		strategy := NewMockStreamStrategy(ctrl)
		strategy.EXPECT().Accept(sourceType).Return(true).Times(1)
		strategy.EXPECT().Create(path, format).Return(stream, nil).Times(1)

		sut := &StreamFactory{}
		_ = sut.Register(strategy)

		if s, e := sut.Create(sourceType, path, format); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(s, stream) {
			t.Error("didn't returned the created stream")
		}
	})
}

func Test_StreamFactory_CreateFromConfig(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		src, e := (&StreamFactory{}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("unrecognized stream type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfigManager(ctrl)
		strategy := NewMockStreamStrategy(ctrl)
		strategy.EXPECT().AcceptFromConfig(cfg).Return(false).Times(1)

		sut := &StreamFactory{}
		_ = sut.Register(strategy)

		stream, e := sut.CreateFromConfig(cfg)
		switch {
		case stream != nil:
			t.Error("returned a config stream")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, serr.ErrInvalidLogStreamConfig):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrInvalidLogStreamConfig)
		}
	})

	t.Run("create the config stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfigManager(ctrl)
		stream := NewMockStream(ctrl)
		strategy := NewMockStreamStrategy(ctrl)
		strategy.EXPECT().AcceptFromConfig(cfg).Return(true).Times(1)
		strategy.EXPECT().CreateFromConfig(cfg).Return(stream, nil).Times(1)

		sut := &StreamFactory{}
		_ = sut.Register(strategy)

		if s, e := sut.CreateFromConfig(cfg); e != nil {
			t.Errorf("returned the (%v) error", e)
		} else if !reflect.DeepEqual(s, stream) {
			t.Error("didn't returned the created stream")
		}
	})
}
