package log

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"testing"
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
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
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
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
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
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("create loader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory()); sut == nil {
			t.Errorf("didn't returned a valid reference")
		} else if e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})
}

func Test_Loader_Load(t *testing.T) {
	t.Run("error retrieving stream list", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().Partial("slate.logger.streams", config.Partial{}).Return(nil, expected).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.cfg = cfg

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("no-op if stream list is empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := &config.Partial{}
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().Partial("slate.logger.streams", config.Partial{}).Return(partial, nil).Times(1)
		cfg.EXPECT().AddObserver("slate.logger.streams", gomock.Any()).Return(nil).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.cfg = cfg

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

		partial := &config.Partial{}
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().Partial("path", config.Partial{}).Return(partial, nil).Times(1)
		cfg.EXPECT().AddObserver("path", gomock.Any()).Return(nil).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.cfg = cfg

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%s) error", e)
		}
	})

	t.Run("error getting stream information", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := &config.Partial{"id": 1}
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().Partial("slate.logger.streams", config.Partial{}).Return(partial, nil).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.cfg = cfg

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, slate.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error creating stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		partial := &config.Partial{"id": config.Partial{}}
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().Partial("slate.logger.streams", config.Partial{}).Return(partial, nil).Times(1)
		streamFactory := NewMockStreamFactory(ctrl)
		streamFactory.EXPECT().Create(&config.Partial{}).Return(nil, expected).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.cfg = cfg
		sut.streamFactory = streamFactory

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error storing stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		partial := &config.Partial{"id": config.Partial{}}
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().Partial("slate.logger.streams", config.Partial{}).Return(partial, nil).Times(1)
		stream := NewMockStream(ctrl)
		streamFactory := NewMockStreamFactory(ctrl)
		streamFactory.EXPECT().Create(&config.Partial{}).Return(stream, nil).Times(1)
		log := NewMockLogger(ctrl)
		log.EXPECT().AddStream("id", stream).Return(expected).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.cfg = cfg
		sut.log = log
		sut.streamFactory = streamFactory

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("register stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := &config.Partial{"id": config.Partial{}}
		cfg := NewMockConfigurer(ctrl)
		cfg.EXPECT().Partial("slate.logger.streams", config.Partial{}).Return(partial, nil).Times(1)
		cfg.EXPECT().AddObserver("slate.logger.streams", gomock.Any()).Return(nil).Times(1)
		stream := NewMockStream(ctrl)
		streamFactory := NewMockStreamFactory(ctrl)
		streamFactory.EXPECT().Create(&config.Partial{}).Return(stream, nil).Times(1)
		log := NewMockLogger(ctrl)
		log.EXPECT().AddStream("id", stream).Return(nil).Times(1)

		sut, _ := NewLoader(config.NewConfig(), NewLog(), NewStreamFactory())
		sut.cfg = cfg
		sut.log = log
		sut.streamFactory = streamFactory

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("error on creating the reconfigured logger streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config1 := config.Partial{"type": "console", "format": "json", "channels": []interface{}{}, "Level": "fatal"}
		partial1 := config.Partial{"slate": config.Partial{"logger": config.Partial{"streams": config.Partial{"id": config1}}}}
		partial2 := config.Partial{"slate": config.Partial{"logger": config.Partial{"streams": "string"}}}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).Times(2)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(partial2, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 1, source1)
		stream := NewMockStream(ctrl)
		streamFactory := NewMockStreamFactory(ctrl)
		streamFactory.EXPECT().Create(&config1).Return(stream, nil).Times(1)
		log := NewMockLogger(ctrl)
		log.EXPECT().AddStream("id", stream).Return(nil).Times(1)

		sut, _ := NewLoader(cfg, NewLog(), NewStreamFactory())
		sut.streamFactory = streamFactory
		sut.log = log
		_ = sut.Load()

		_ = cfg.AddSource("id2", 100, source2)
	})

	t.Run("reconfigured logger streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config1 := config.Partial{"type": "console", "format": "json", "Channels": []interface{}{}, "Level": "fatal"}
		config2 := config.Partial{"type": "console", "format": "json", "Channels": []interface{}{}, "Level": "fatal"}
		partial1 := config.Partial{"slate": config.Partial{"logger": config.Partial{"streams": config.Partial{"id1": config1}}}}
		partial2 := config.Partial{"slate": config.Partial{"logger": config.Partial{"streams": config.Partial{"id2": config2}}}}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).Times(2)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(partial2, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id1", 1, source1)
		stream1 := NewMockStream(ctrl)
		stream2 := NewMockStream(ctrl)
		streamFactory := NewMockStreamFactory(ctrl)
		gomock.InOrder(
			streamFactory.EXPECT().Create(&config1).Return(stream1, nil),
			streamFactory.EXPECT().Create(&config1).Return(stream1, nil),
			streamFactory.EXPECT().Create(&config1).Return(stream2, nil),
		)
		log := NewMockLogger(ctrl)
		log.EXPECT().RemoveAllStreams().Times(1)
		gomock.InOrder(
			log.EXPECT().AddStream("id1", stream1).Return(nil),
			log.EXPECT().AddStream("id1", stream1).Return(nil),
			log.EXPECT().AddStream("id2", stream2).Return(nil),
		)

		sut, _ := NewLoader(cfg, NewLog(), NewStreamFactory())
		sut.streamFactory = streamFactory
		sut.log = log
		_ = sut.Load()

		_ = cfg.AddSource("id2", 100, source2)
	})
}
