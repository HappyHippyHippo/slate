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
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewLoader(nil, NewMockLog(ctrl), NewMockStreamFactory(ctrl))
		switch {
		case sut != nil:
			t.Errorf("return a valid reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error when missing the log", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewLoader(NewMockConfigManager(ctrl), nil, NewMockStreamFactory(ctrl))
		switch {
		case sut != nil:
			t.Errorf("return a valid reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error when missing the log stream factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewLoader(NewMockConfigManager(ctrl), NewMockLog(ctrl), nil)
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

		if sut, e := NewLoader(NewMockConfigManager(ctrl), NewMockLog(ctrl), NewMockStreamFactory(ctrl)); sut == nil {
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
		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().Config("slate.log.streams", config.Config{}).Return(nil, expected).Times(1)

		sut, _ := NewLoader(manager, NewMockLog(ctrl), NewMockStreamFactory(ctrl))

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("no-op if stream list is empty", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().Config("slate.log.streams", config.Config{}).Return(&config.Config{}, nil).Times(1)
		manager.EXPECT().AddObserver("slate.log.streams", gomock.Any()).Return(nil).Times(1)

		sut, _ := NewLoader(manager, NewMockLog(ctrl), NewMockStreamFactory(ctrl))

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

		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().Config("path", config.Config{}).Return(&config.Config{}, nil).Times(1)
		manager.EXPECT().AddObserver("path", gomock.Any()).Return(nil).Times(1)

		sut, _ := NewLoader(manager, NewMockLog(ctrl), NewMockStreamFactory(ctrl))

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%s) error", e)
		}
	})

	t.Run("error getting stream information", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().Config("slate.log.streams", config.Config{}).Return(&config.Config{"id": "string"}, nil).Times(1)

		sut, _ := NewLoader(manager, NewMockLog(ctrl), NewMockStreamFactory(ctrl))

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
		partial := config.Config{"type": "invalid"}
		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().Config("slate.log.streams", config.Config{}).Return(&config.Config{"id": partial}, nil).Times(1)
		sFactory := NewMockStreamFactory(ctrl)
		sFactory.EXPECT().Create(&partial).Return(nil, expected).Times(1)

		sut, _ := NewLoader(manager, NewMockLog(ctrl), sFactory)

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
		partial := config.Config{"type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().Config("slate.log.streams", config.Config{}).Return(&config.Config{"id": partial}, nil).Times(1)
		stream1 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		sFactory.EXPECT().Create(&partial).Return(stream1, nil).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().AddStream("id", stream1).Return(expected).Times(1)

		sut, _ := NewLoader(manager, logger, sFactory)

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("register stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Config{"type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().Config("slate.log.streams", config.Config{}).Return(&config.Config{"id": partial}, nil).Times(1)
		manager.EXPECT().AddObserver("slate.log.streams", gomock.Any()).Return(nil).Times(1)
		stream1 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		sFactory.EXPECT().Create(&partial).Return(stream1, nil).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().AddStream("id", stream1).Return(nil).Times(1)

		sut, _ := NewLoader(manager, logger, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("error on creating the reconfigured log streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config1 := config.Config{"type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial := config.Config{"slate": config.Config{"log": config.Config{"streams": config.Config{"id": config1}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		manager := config.NewManager(0)
		_ = manager.AddSource("id", 0, source)
		stream1 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		sFactory.EXPECT().Create(&config1).Return(stream1, nil).Times(1)
		logger := NewMockLog(ctrl)
		logger.EXPECT().AddStream("id", stream1).Return(nil).Times(1)
		logger.EXPECT().Signal("exec", ERROR, "reloading log streams error", gomock.Any()).Return(nil).Times(1)

		sut, _ := NewLoader(manager, logger, sFactory)
		_ = sut.Load()

		partial = config.Config{"slate": config.Config{"log": config.Config{"streams": "string"}}}
		source = NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		_ = manager.AddSource("id2", 100, source)
	})

	t.Run("error on storing the reconfigured log streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		config1 := config.Config{"type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		config2 := config.Config{"type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial1 := config.Config{"slate": config.Config{"log": config.Config{"streams": config.Config{"id": config1}}}}
		partial2 := config.Config{"slate": config.Config{"log": config.Config{"streams": config.Config{"id": config2}}}}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(partial2, nil).MinTimes(1)
		manager := config.NewManager(0)
		_ = manager.AddSource("id", 0, source1)
		stream1 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		gomock.InOrder(
			sFactory.EXPECT().Create(&config1).Return(stream1, nil),
			sFactory.EXPECT().Create(&config2).Return(nil, expected),
		)
		logger := NewMockLog(ctrl)
		logger.EXPECT().AddStream("id", stream1).Return(nil).Times(1)
		logger.EXPECT().RemoveAllStreams().Times(1)
		logger.EXPECT().Signal("exec", ERROR, "reloading log streams error", gomock.Any()).Return(nil).Times(1)

		sut, _ := NewLoader(manager, logger, sFactory)
		_ = sut.Load()

		_ = manager.AddSource("id2", 100, source2)
	})

	t.Run("reconfigured log streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config1 := config.Config{"type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		config2 := config.Config{"type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial1 := config.Config{"slate": config.Config{"log": config.Config{"streams": config.Config{"id": config1}}}}
		partial2 := config.Config{"slate": config.Config{"log": config.Config{"streams": config.Config{"id": config2}}}}
		source1 := NewMockConfigSource(ctrl)
		source1.EXPECT().Get("").Return(partial1, nil).MinTimes(1)
		source2 := NewMockConfigSource(ctrl)
		source2.EXPECT().Get("").Return(partial2, nil).MinTimes(1)
		manager := config.NewManager(0)
		_ = manager.AddSource("id", 0, source1)
		stream1 := NewMockStream(ctrl)
		stream2 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		gomock.InOrder(
			sFactory.EXPECT().Create(&config1).Return(stream1, nil),
			sFactory.EXPECT().Create(&config2).Return(stream2, nil),
		)
		logger := NewMockLog(ctrl)
		gomock.InOrder(
			logger.EXPECT().AddStream("id", stream1).Return(nil),
			logger.EXPECT().AddStream("id", stream2).Return(nil),
		)
		logger.EXPECT().RemoveAllStreams().Times(1)

		sut, _ := NewLoader(manager, logger, sFactory)
		_ = sut.Load()

		_ = manager.AddSource("id2", 100, source2)
	})
}
