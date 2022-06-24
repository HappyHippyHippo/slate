package slog

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/sconfig"
	"github.com/happyhippyhippo/slate/serror"
	"testing"
)

func Test_NewLoader(t *testing.T) {
	t.Run("error when missing the config", func(t *testing.T) {
		loader, err := NewLoader(nil, NewLogger(), &StreamFactory{})
		switch {
		case loader != nil:
			t.Errorf("return a valid reference")
		case err == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error when missing the logger", func(t *testing.T) {
		loader, err := NewLoader(sconfig.NewConfig(0), nil, &StreamFactory{})
		switch {
		case loader != nil:
			t.Errorf("return a valid reference")
		case err == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("error when missing the logger stream factory", func(t *testing.T) {
		loader, err := NewLoader(sconfig.NewConfig(0), NewLogger(), nil)
		switch {
		case loader != nil:
			t.Errorf("return a valid reference")
		case err == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(err, serror.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrNilPointer)
		}
	})

	t.Run("create loader", func(t *testing.T) {
		if loader, err := NewLoader(sconfig.NewConfig(0), NewLogger(), &StreamFactory{}); loader == nil {
			t.Errorf("didn't returned a valid reference")
		} else if err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})
}

func Test_Loader_Load(t *testing.T) {
	t.Run("no-op if stream list is empty", func(t *testing.T) {
		loader, _ := NewLoader(sconfig.NewConfig(0), NewLogger(), &StreamFactory{})

		if err := loader.Load(); err != nil {
			t.Errorf("returned the (%s) error", err)
		}
	})

	t.Run("invalid config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := sconfig.Partial{"log": sconfig.Partial{"streams": "string"}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		loader, _ := NewLoader(cfg, NewLogger(), &StreamFactory{})

		if err := loader.Load(); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("request config path from global variable", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		path := "path"
		LoaderConfigPath = path
		defer func() { LoaderConfigPath = "log.streams" }()

		partial := sconfig.Partial{path: "string"}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		loader, _ := NewLoader(cfg, NewLogger(), &StreamFactory{})

		if err := loader.Load(); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error getting stream information", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{"string"}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		loader, _ := NewLoader(cfg, NewLogger(), &StreamFactory{})

		if err := loader.Load(); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error getting stream id", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scfg := sconfig.Partial{"id": 123}
		partial := sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		loader, _ := NewLoader(cfg, NewLogger(), &StreamFactory{})

		if err := loader.Load(); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrConversion) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrConversion)
		}
	})

	t.Run("error creating stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scfg := sconfig.Partial{"id": "id", "type": "invalid"}
		partial := sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		loader, _ := NewLoader(cfg, NewLogger(), &StreamFactory{})

		if err := loader.Load(); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrInvalidLogStreamConfig) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrInvalidLogStreamConfig)
		}
	})

	t.Run("error storing stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scfg := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial := sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg, scfg}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		formatterFactory := &FormatterFactory{}
		_ = formatterFactory.Register(&formatterStrategyJSON{})
		streamFactory := &StreamFactory{}
		strategy, _ := newStreamStrategyConsole(formatterFactory)
		_ = streamFactory.Register(strategy)
		loader, _ := NewLoader(cfg, NewLogger(), streamFactory)

		if err := loader.Load(); err == nil {
			t.Errorf("didn't returned the expected error")
		} else if !errors.Is(err, serror.ErrDuplicateLogStream) {
			t.Errorf("returned the (%v) error when expecting (%v)", err, serror.ErrDuplicateLogStream)
		}
	})

	t.Run("register stream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scfg := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial := sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		formatterFactory := &FormatterFactory{}
		_ = formatterFactory.Register(&formatterStrategyJSON{})
		streamFactory := &StreamFactory{}
		strategy, _ := newStreamStrategyConsole(formatterFactory)
		_ = streamFactory.Register(strategy)
		loader, _ := NewLoader(cfg, NewLogger(), streamFactory)

		if err := loader.Load(); err != nil {
			t.Errorf("returned the (%v) error", err)
		}
	})

	t.Run("error on creating the reconfigured log streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scfg1 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial := sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg1}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		formatterFactory := &FormatterFactory{}
		_ = formatterFactory.Register(&formatterStrategyJSON{})
		streamFactory := &StreamFactory{}
		strategy, _ := newStreamStrategyConsole(formatterFactory)
		_ = streamFactory.Register(strategy)
		logger := NewLogger()
		loader, _ := NewLoader(cfg, logger, streamFactory)

		_ = loader.Load()

		partial = sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{"string"}}}
		source = NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		_ = cfg.AddSource("id2", 100, source)
	})

	t.Run("error on storing the reconfigured log streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scfg1 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial := sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg1}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		formatterFactory := &FormatterFactory{}
		_ = formatterFactory.Register(&formatterStrategyJSON{})
		streamFactory := &StreamFactory{}
		strategy, _ := newStreamStrategyConsole(formatterFactory)
		_ = streamFactory.Register(strategy)
		logger := NewLogger()
		loader, _ := NewLoader(cfg, logger, streamFactory)

		_ = loader.Load()

		scfg2 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial = sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg2, scfg2}}}
		source = NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		_ = cfg.AddSource("id2", 100, source)
	})

	t.Run("reconfigured log streams", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		scfg1 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial := sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg1}}}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		cfg := sconfig.NewConfig(0)
		_ = cfg.AddSource("id", 0, source)
		formatterFactory := &FormatterFactory{}
		_ = formatterFactory.Register(&formatterStrategyJSON{})
		streamFactory := &StreamFactory{}
		strategy, _ := newStreamStrategyConsole(formatterFactory)
		_ = streamFactory.Register(strategy)
		logger := NewLogger()
		loader, _ := NewLoader(cfg, logger, streamFactory)

		_ = loader.Load()

		scfg2 := sconfig.Partial{"id": "id", "type": "console", "format": "json", "channels": []interface{}{}, "level": "fatal"}
		partial = sconfig.Partial{"log": sconfig.Partial{"streams": []interface{}{scfg2}}}
		source = NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).MinTimes(1)
		_ = cfg.AddSource("id2", 100, source)
	})
}
