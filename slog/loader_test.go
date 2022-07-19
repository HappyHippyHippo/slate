package slog

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serr"
	"testing"
)

func Test_NewLoader(t *testing.T) {
	t.Run("error when missing the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newLoader(nil, NewMockLogger(ctrl), NewMockStreamFactory(ctrl))
		switch {
		case sut != nil:
			t.Errorf("return a valid reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("error when missing the logger", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newLoader(NewMockConfigManager(ctrl), nil, NewMockStreamFactory(ctrl))
		switch {
		case sut != nil:
			t.Errorf("return a valid reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("error when missing the logger stream factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := newLoader(NewMockConfigManager(ctrl), NewMockLogger(ctrl), nil)
		switch {
		case sut != nil:
			t.Errorf("return a valid reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, serr.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrNilPointer)
		}
	})

	t.Run("create loader", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		if sut, e := newLoader(NewMockConfigManager(ctrl), NewMockLogger(ctrl), NewMockStreamFactory(ctrl)); sut == nil {
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
		manager.EXPECT().List("slate.log.streams", []interface{}{}).Return(nil, expected).Times(1)

		sut, _ := newLoader(manager, NewMockLogger(ctrl), NewMockStreamFactory(ctrl))

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
		manager.EXPECT().List("slate.log.streams", []interface{}{}).Return([]interface{}{}, nil).Times(1)
		manager.EXPECT().AddObserver("slate.log.streams", gomock.Any()).Return(nil).Times(1)

		sut, _ := newLoader(manager, NewMockLogger(ctrl), NewMockStreamFactory(ctrl))

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%s) error", e)
		}
	})

	t.Run("request config path from global variable", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		LoaderConfigPath = path
		defer func() { LoaderConfigPath = "slate.log.streams" }()

		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().List(path, []interface{}{}).Return([]interface{}{"string"}, nil).Times(1)

		sut, _ := newLoader(manager, NewMockLogger(ctrl), NewMockStreamFactory(ctrl))

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error getting stream information", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().List("slate.log.streams", []interface{}{}).Return([]interface{}{"string"}, nil).Times(1)

		sut, _ := newLoader(manager, NewMockLogger(ctrl), NewMockStreamFactory(ctrl))

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error getting stream id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := sconfig.Partial{"id": 123}
		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().List("slate.log.streams", []interface{}{}).Return([]interface{}{partial}, nil).Times(1)

		sut, _ := newLoader(manager, NewMockLogger(ctrl), NewMockStreamFactory(ctrl))

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(e, serr.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", e, serr.ErrConversion)
		}
	})

	t.Run("error creating stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		partial := sconfig.Partial{"id": "id", "type": "invalid"}
		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().List("slate.log.streams", []interface{}{}).Return([]interface{}{partial}, nil).Times(1)
		sFactory := NewMockStreamFactory(ctrl)
		sFactory.EXPECT().CreateFromConfig(&partial).Return(nil, expected).Times(1)

		sut, _ := newLoader(manager, NewMockLogger(ctrl), sFactory)

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
		partial := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().List("slate.log.streams", []interface{}{}).Return([]interface{}{partial}, nil).Times(1)
		stream1 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		sFactory.EXPECT().CreateFromConfig(&partial).Return(stream1, nil).Times(1)
		logger := NewMockLogger(ctrl)
		logger.EXPECT().AddStream("id", stream1).Return(expected).Times(1)

		sut, _ := newLoader(manager, logger, sFactory)

		if e := sut.Load(); e == nil {
			t.Errorf("didn't returned the expected error")
		} else if e.Error() != expected.Error() {
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("register stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		manager := NewMockConfigManager(ctrl)
		manager.EXPECT().List("slate.log.streams", []interface{}{}).Return([]interface{}{partial}, nil).Times(1)
		manager.EXPECT().AddObserver("slate.log.streams", gomock.Any()).Return(nil).Times(1)
		stream1 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		sFactory.EXPECT().CreateFromConfig(&partial).Return(stream1, nil).Times(1)
		logger := NewMockLogger(ctrl)
		logger.EXPECT().AddStream("id", stream1).Return(nil).Times(1)

		sut, _ := newLoader(manager, logger, sFactory)

		if e := sut.Load(); e != nil {
			t.Errorf("returned the (%v) error", e)
		}
	})

	t.Run("error on creating the reconfigured slog streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config1 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial := sconfig.Partial{"slate": sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{config1}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		manager := sconfig.NewManager(0)
		_ = manager.AddSource("id", 0, source)
		stream1 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		sFactory.EXPECT().CreateFromConfig(&config1).Return(stream1, nil).Times(1)
		logger := NewMockLogger(ctrl)
		logger.EXPECT().AddStream("id", stream1).Return(nil).Times(1)
		logger.EXPECT().Signal("exec", ERROR, "reloading slog streams error", gomock.Any()).Return(nil).Times(1)

		sut, _ := newLoader(manager, logger, sFactory)
		_ = sut.Load()

		partial = sconfig.Partial{"slate": sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{"string"}}}}
		source = NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		_ = manager.AddSource("id2", 100, source)
	})

	t.Run("error on storing the reconfigured slog streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		config1 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial := sconfig.Partial{"slate": sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{config1}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		manager := sconfig.NewManager(0)
		_ = manager.AddSource("id", 0, source)
		stream1 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		gomock.InOrder(
			sFactory.EXPECT().CreateFromConfig(&config1).Return(stream1, nil),
			sFactory.EXPECT().CreateFromConfig(&config1).Return(nil, expected),
		)
		logger := NewMockLogger(ctrl)
		logger.EXPECT().AddStream("id", stream1).Return(nil).Times(1)
		logger.EXPECT().RemoveAllStreams().Times(1)
		logger.EXPECT().Signal("exec", ERROR, "reloading slog streams error", gomock.Any()).Return(nil).Times(1)

		sut, _ := newLoader(manager, logger, sFactory)
		_ = sut.Load()

		config2 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial = sconfig.Partial{"slate": sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{config2, config2}}}}
		source = NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		_ = manager.AddSource("id2", 100, source)
	})

	t.Run("reconfigured slog streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		config1 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		config2 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "error"}
		partial := sconfig.Partial{"slate": sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{config1}}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		manager := sconfig.NewManager(0)
		_ = manager.AddSource("id", 0, source)
		stream1 := NewMockStream(ctrl)
		stream2 := NewMockStream(ctrl)
		sFactory := NewMockStreamFactory(ctrl)
		gomock.InOrder(
			sFactory.EXPECT().CreateFromConfig(&config1).Return(stream1, nil),
			sFactory.EXPECT().CreateFromConfig(&config2).Return(stream2, nil),
		)
		logger := NewMockLogger(ctrl)
		gomock.InOrder(
			logger.EXPECT().AddStream("id", stream1).Return(nil),
			logger.EXPECT().AddStream("id", stream2).Return(nil),
		)
		logger.EXPECT().RemoveAllStreams().Times(1)

		sut, _ := newLoader(manager, logger, sFactory)
		_ = sut.Load()

		partial = sconfig.Partial{"slate": sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{config2}}}}
		source = NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		_ = manager.AddSource("id2", 100, source)
	})
}
