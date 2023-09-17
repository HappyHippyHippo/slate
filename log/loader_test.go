package log

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
)

func Test_NewLoader(t *testing.T) {
	t.Run("error when missing the config", func(t *testing.T) {
		sut, e := NewLoader(nil, NewLog(), NewStreamFactory())
		switch {
		case sut != nil:
			t.Errorf("return a valid reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error when missing the logger", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewLoader(config.NewConfig(), nil, NewStreamFactory())
		switch {
		case sut != nil:
			t.Errorf("return a valid reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error when missing the logger stream factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewLoader(config.NewConfig(), NewLog(), nil)
		switch {
		case sut != nil:
			t.Errorf("return a valid reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("create loader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory()); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})
}

func Test_Loader_Load(t *testing.T) {
	t.Run("error retrieving stream list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.logger.streams", config.Partial{}).
			Return(nil, expected).
			Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.configurer = configurer

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("no-op if stream list is empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.logger.streams", config.Partial{}).
			Return(partial, nil).
			Times(1)
		configurer.
			EXPECT().
			AddObserver("slate.logger.streams", gomock.Any()).
			Return(nil).
			Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.configurer = configurer

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%s) error", e)
		}
	})

	t.Run("request config path from global variable", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := LoaderConfigPath
		LoaderConfigPath = "path"
		defer func() { LoaderConfigPath = prev }()

		partial := config.Partial{}
		configurer := NewMockConfigurer(ctrl)
		configurer.EXPECT().Partial("path", config.Partial{}).Return(partial, nil).Times(1)
		configurer.EXPECT().AddObserver("path", gomock.Any()).Return(nil).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.configurer = configurer

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%s) error", e)
		}
	})

	t.Run("error getting stream information", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"id": 1}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.logger.streams", config.Partial{}).
			Return(partial, nil).
			Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.configurer = configurer

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error creating stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		partial := config.Partial{"id": config.Partial{}}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.logger.streams", config.Partial{}).
			Return(partial, nil).
			Times(1)
		streamCreator := NewMockStreamCreator(ctrl)
		streamCreator.EXPECT().Create(config.Partial{}).Return(nil, expected).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.configurer = configurer
		sut.streamCreator = streamCreator

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("error storing stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		partial := config.Partial{"id": config.Partial{}}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.logger.streams", config.Partial{}).
			Return(partial, nil).
			Times(1)
		stream := NewMockStream(ctrl)
		streamCreator := NewMockStreamCreator(ctrl)
		streamCreator.EXPECT().Create(config.Partial{}).Return(stream, nil).Times(1)
		logger := NewMockLogger(ctrl)
		logger.EXPECT().AddStream("id", stream).Return(expected).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.configurer = configurer
		sut.logger = logger
		sut.streamCreator = streamCreator

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("register stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{"id": config.Partial{}}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.logger.streams", config.Partial{}).
			Return(partial, nil).
			Times(1)
		configurer.
			EXPECT().
			AddObserver("slate.logger.streams", gomock.Any()).
			Return(nil).
			Times(1)
		stream := NewMockStream(ctrl)
		streamCreator := NewMockStreamCreator(ctrl)
		streamCreator.EXPECT().Create(config.Partial{}).Return(stream, nil).Times(1)
		logger := NewMockLogger(ctrl)
		logger.EXPECT().AddStream("id", stream).Return(nil).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.configurer = configurer
		sut.logger = logger
		sut.streamCreator = streamCreator

		if e := sut.Load(); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("error on creating the reconfigured logger streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config1 := config.Partial{
			"type":     "console",
			"format":   "json",
			"channels": []interface{}{},
			"Level":    "fatal",
		}
		partial1 := config.Partial{}
		_, _ = partial1.Set("slate.logger.streams.id", config1)
		partial2 := config.Partial{}
		_, _ = partial2.Set("slate.logger.streams", "string")
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).Times(2)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(partial2, nil).Times(1)
		configurer := config.NewConfig()
		_ = configurer.AddSource("id", 1, source1)
		stream := NewMockStream(ctrl)
		streamCreator := NewMockStreamCreator(ctrl)
		streamCreator.EXPECT().Create(config1).Return(stream, nil).Times(1)
		logger := NewMockLogger(ctrl)
		logger.EXPECT().AddStream("id", stream).Return(nil).Times(1)

		sut, _ := NewLoader(configurer, NewLog(), NewStreamFactory())
		sut.streamCreator = streamCreator
		sut.logger = logger
		_ = sut.Load()

		_ = configurer.AddSource("id2", 100, source2)
	})

	t.Run("reconfigured logger streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config1 := config.Partial{
			"type":     "console",
			"format":   "json",
			"Channels": []interface{}{},
			"Level":    "fatal",
		}
		config2 := config.Partial{
			"type":     "console",
			"format":   "json",
			"Channels": []interface{}{},
			"Level":    "fatal",
		}
		partial1 := config.Partial{}
		_, _ = partial1.Set("slate.logger.streams.id1", config1)
		partial2 := config.Partial{}
		_, _ = partial2.Set("slate.logger.streams.id2", config2)
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).Times(2)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(partial2, nil).Times(1)
		configurer := config.NewConfig()
		_ = configurer.AddSource("id1", 1, source1)
		stream1 := NewMockStream(ctrl)
		stream2 := NewMockStream(ctrl)
		streamCreator := NewMockStreamCreator(ctrl)
		streamCreator.EXPECT().Create(config1).Return(stream1, nil).Times(2)
		streamCreator.EXPECT().Create(config1).Return(stream2, nil).Times(1)
		logger := NewMockLogger(ctrl)
		logger.EXPECT().RemoveAllStreams().Times(1)
		logger.EXPECT().AddStream("id1", stream1).Return(nil).Times(2)
		logger.EXPECT().AddStream("id2", stream2).Return(nil).Times(1)

		sut, _ := NewLoader(configurer, NewLog(), NewStreamFactory())
		sut.streamCreator = streamCreator
		sut.logger = logger
		_ = sut.Load()

		_ = configurer.AddSource("id2", 100, source2)
	})
}
