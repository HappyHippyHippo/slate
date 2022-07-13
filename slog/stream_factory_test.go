package slog

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serror"
	"reflect"
	"testing"
)

func Test_StreamFactory_Register(t *testing.T) {
	t.Run("nil strategy", func(t *testing.T) {
		if err := (&StreamFactory{}).Register(nil); err == nil {
			t.Error("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrNilPointer) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("register the stream factory strategy", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		strategy := NewMockStreamStrategy(ctrl)
		factory := &StreamFactory{}

		if err := factory.Register(strategy); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if (*factory)[0] != strategy {
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
		factory := &StreamFactory{}
		_ = factory.Register(strategy)

		stream, err := factory.Create(streamType, path, format)
		switch {
		case stream != nil:
			t.Error("returned an valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidLogStreamType):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidLogStreamType)
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
		factory := &StreamFactory{}
		_ = factory.Register(strategy)

		if s, err := factory.Create(sourceType, path, format); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !reflect.DeepEqual(s, stream) {
			t.Error("didn't returned the created stream")
		}
	})
}

func Test_StreamFactory_CreateFromConfig(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		src, err := (&StreamFactory{}).CreateFromConfig(nil)
		switch {
		case src != nil:
			t.Error("returned a valid reference")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("unrecognized stream type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := sconfig.NewManager(0)
		strategy := NewMockStreamStrategy(ctrl)
		strategy.EXPECT().AcceptFromConfig(cfg).Return(false).Times(1)
		factory := &StreamFactory{}
		_ = factory.Register(strategy)

		stream, err := factory.CreateFromConfig(cfg)
		switch {
		case stream != nil:
			t.Error("returned a config stream")
		case err == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(err, serror.ErrInvalidLogStreamConfig):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidLogStreamConfig)
		}
	})

	t.Run("create the config stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := sconfig.NewManager(0)
		stream := NewMockStream(ctrl)
		strategy := NewMockStreamStrategy(ctrl)
		strategy.EXPECT().AcceptFromConfig(cfg).Return(true).Times(1)
		strategy.EXPECT().CreateFromConfig(cfg).Return(stream, nil).Times(1)
		factory := &StreamFactory{}
		_ = factory.Register(strategy)

		if s, err := factory.CreateFromConfig(cfg); err != nil {
			t.Errorf("returned the (%v) error", err)
		} else if !reflect.DeepEqual(s, stream) {
			t.Error("didn't returned the created stream")
		}
	})
}
