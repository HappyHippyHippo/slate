package rest

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewProcess(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := log.NewLog()
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(nil, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil logger", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := config.NewConfig()
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfg, nil, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil engine", func(t *testing.T) {
		cfg := config.NewConfig()
		logger := log.NewLog()

		sut, e := NewProcess(cfg, logger, nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error while retrieving configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest", "string")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logger := log.NewLog()
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfg, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error while retrieving configuration from env path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPath
		ConfigPath = "test"
		defer func() { ConfigPath = prev }()

		partial := config.Partial{"test": "string"}
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logger := log.NewLog()
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfg, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error while populating configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.watchdog", 123)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logger := log.NewLog()
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfg, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("invalid log level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.log.level", "invalid")
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logger := log.NewLog()
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfg, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("successful process creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest", config.Partial{})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logger := log.NewLog()
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfg, logger, engine)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case sut.Service() != WatchdogName:
			t.Error("didn't returned the expected valid reference")
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("successful process creation with name from env", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := WatchdogName
		WatchdogName = "test"
		defer func() { WatchdogName = prev }()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest", config.Partial{})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logger := log.NewLog()
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfg, logger, engine)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case sut.Service() != WatchdogName:
			t.Error("didn't returned the expected valid reference")
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("successful process run", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest", config.Partial{})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		gomock.InOrder(
			logStream.
				EXPECT().
				Signal("rest", log.INFO, "[service:rest] service starting ...", log.Context{"port": 80}).
				Return(nil),
			logStream.
				EXPECT().
				Signal("rest", log.INFO, "[service:rest] service terminated").
				Return(nil),
		)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		engine := NewMockEngine(ctrl)
		engine.EXPECT().Run(":80").Return(nil).Times(1)

		sut, _ := NewProcess(cfg, logger, engine)
		if e := sut.Runner()(); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("successful process run with config values", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		name := "watchdog name"
		port := 1234
		logLevel := log.FATAL
		logChannel := "test channel"
		logStartMessage := "start message"
		logEndMessage := "end message"

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.watchdog", name)
		_, _ = partial.Set("slate.api.rest.port", port)
		_, _ = partial.Set("slate.api.rest.log.level", log.LevelMapName[logLevel])
		_, _ = partial.Set("slate.api.rest.log.channel", logChannel)
		_, _ = partial.Set("slate.api.rest.log.message.start", logStartMessage)
		_, _ = partial.Set("slate.api.rest.log.message.end", logEndMessage)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		gomock.InOrder(
			logStream.
				EXPECT().
				Signal(logChannel, logLevel, logStartMessage, log.Context{"port": port}).
				Return(nil),
			logStream.
				EXPECT().
				Signal(logChannel, logLevel, logEndMessage).
				Return(nil),
		)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		engine := NewMockEngine(ctrl)
		engine.EXPECT().Run(":1234").Return(nil).Times(1)

		sut, _ := NewProcess(cfg, logger, engine)
		if chk := sut.Service(); chk != name {
			t.Errorf("(%v) when expected (%v)", chk, name)
		} else if e := sut.Runner()(); e != nil {
			t.Errorf("unexpected (%v) error", e)
		}
	})

	t.Run("failure when running process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		errorMessage := "error message"
		expected := fmt.Errorf("%s", errorMessage)

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest", config.Partial{})
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		gomock.InOrder(
			logStream.
				EXPECT().
				Signal("rest", log.INFO, "[service:rest] service starting ...", log.Context{"port": 80}).
				Return(nil),
			logStream.
				EXPECT().
				Signal("rest", log.FATAL, "[service:rest] service error", log.Context{"error": errorMessage}).
				Return(nil),
		)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		engine := NewMockEngine(ctrl)
		engine.EXPECT().Run(":80").Return(expected).Times(1)

		sut, _ := NewProcess(cfg, logger, engine)
		e := sut.Runner()()
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("failure when running process with values from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		errorMessage := "error message"
		expected := fmt.Errorf("%s", errorMessage)

		name := "watchdog name"
		port := 1234
		logLevel := log.FATAL
		logChannel := "test channel"
		logStartMessage := "start message"
		logErrorMessage := "error message"

		partial := config.Partial{}
		_, _ = partial.Set("slate.api.rest.watchdog", name)
		_, _ = partial.Set("slate.api.rest.port", port)
		_, _ = partial.Set("slate.api.rest.log.level", log.LevelMapName[logLevel])
		_, _ = partial.Set("slate.api.rest.log.channel", logChannel)
		_, _ = partial.Set("slate.api.rest.log.message.start", logStartMessage)
		_, _ = partial.Set("slate.api.rest.log.message.error", logErrorMessage)
		source := NewMockConfigSource(ctrl)
		source.EXPECT().Get("").Return(partial, nil).Times(1)
		cfg := config.NewConfig()
		_ = cfg.AddSource("id", 0, source)
		logStream := NewMockLogStream(ctrl)
		gomock.InOrder(
			logStream.
				EXPECT().
				Signal(logChannel, logLevel, logStartMessage, log.Context{"port": port}).
				Return(nil),
			logStream.
				EXPECT().
				Signal(logChannel, logLevel, logErrorMessage, log.Context{"error": errorMessage}).
				Return(nil),
		)
		logger := log.NewLog()
		_ = logger.AddStream("id", logStream)
		engine := NewMockEngine(ctrl)
		engine.EXPECT().Run(":1234").Return(expected).Times(1)

		sut, _ := NewProcess(cfg, logger, engine)
		e := sut.Runner()()
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})
}
