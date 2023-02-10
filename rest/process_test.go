package rest

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewProcess(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		logger := NewMockLog(ctrl)
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(nil, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil logger", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgManager := NewMockConfigManager(ctrl)
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfgManager, nil, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil engine", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfgManager := NewMockConfigManager(ctrl)
		logger := NewMockLog(ctrl)

		sut, e := NewProcess(cfgManager, logger, nil)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("error while retrieving configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(nil, expected).Times(1)
		logger := NewMockLog(ctrl)
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfgManager, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error while retrieving configuration from env path", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPath
		ConfigPath = "test"
		defer func() { ConfigPath = prev }()

		expected := fmt.Errorf("error message")
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(nil, expected).Times(1)
		logger := NewMockLog(ctrl)
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfgManager, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error while populating configuration", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).Return(nil, expected).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(cfg, nil).Times(1)
		logger := NewMockLog(ctrl)
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfgManager, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid log level", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *processConfig, icase ...bool) (interface{}, error) {
			c.Log.Level = "invalid"
			return nil, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(cfg, nil).Times(1)
		logger := NewMockLog(ctrl)
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfgManager, logger, engine)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("successful process creation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *processConfig, icase ...bool) (interface{}, error) {
			return c, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(cfg, nil).Times(1)
		logger := NewMockLog(ctrl)
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfgManager, logger, engine)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case sut.Service() != WatchdogName:
			t.Error("didn't returned the expected valid reference")
		case e != nil:
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("successful process creation with name from env", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := WatchdogName
		WatchdogName = "test"
		defer func() { WatchdogName = prev }()

		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *processConfig, icase ...bool) (interface{}, error) {
			return c, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(cfg, nil).Times(1)
		logger := NewMockLog(ctrl)
		engine := NewMockEngine(ctrl)

		sut, e := NewProcess(cfgManager, logger, engine)
		switch {
		case sut == nil:
			t.Error("didn't returned the expected valid reference")
		case sut.Service() != WatchdogName:
			t.Error("didn't returned the expected valid reference")
		case e != nil:
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("successful process run", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *processConfig, icase ...bool) (interface{}, error) {
			return c, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(cfg, nil).Times(1)
		logger := NewMockLog(ctrl)
		gomock.InOrder(
			logger.EXPECT().Signal(LogChannel, log.INFO, LogStartMessage, log.Context{"port": 80}),
			logger.EXPECT().Signal(LogChannel, log.INFO, LogEndMessage),
		)
		engine := NewMockEngine(ctrl)
		engine.EXPECT().Run(":80").Return(nil).Times(1)

		sut, _ := NewProcess(cfgManager, logger, engine)
		if e := sut.Runner()(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
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
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *processConfig, icase ...bool) (interface{}, error) {
			c.Watchdog = name
			c.Port = port
			c.Log.Level = log.LevelMapName[logLevel]
			c.Log.Channel = logChannel
			c.Log.Message.Start = logStartMessage
			c.Log.Message.End = logEndMessage
			return c, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(cfg, nil).Times(1)
		logger := NewMockLog(ctrl)
		gomock.InOrder(
			logger.EXPECT().Signal(logChannel, logLevel, logStartMessage, log.Context{"port": port}),
			logger.EXPECT().Signal(logChannel, logLevel, logEndMessage),
		)
		engine := NewMockEngine(ctrl)
		engine.EXPECT().Run(fmt.Sprintf(":%d", port)).Return(nil).Times(1)

		sut, _ := NewProcess(cfgManager, logger, engine)
		if chk := sut.Service(); chk != name {
			t.Errorf("returned the unexpected watchdog service name (%v) when expected (%v)", chk, name)
		} else if e := sut.Runner()(); e != nil {
			t.Errorf("returned the unexpected error : %v", e)
		}
	})

	t.Run("failure when running process", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		errorMessage := "error message"
		expected := fmt.Errorf("%s", errorMessage)
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *processConfig, icase ...bool) (interface{}, error) {
			return c, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(cfg, nil).Times(1)
		logger := NewMockLog(ctrl)
		gomock.InOrder(
			logger.EXPECT().Signal(LogChannel, log.INFO, LogStartMessage, log.Context{"port": 80}),
			logger.EXPECT().Signal(LogChannel, log.FATAL, LogErrorMessage, log.Context{"error": errorMessage}),
		)
		engine := NewMockEngine(ctrl)
		engine.EXPECT().Run(":80").Return(expected).Times(1)

		sut, _ := NewProcess(cfgManager, logger, engine)
		e := sut.Runner()()
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("failure when running process with values from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		errorMessage := "error message"
		expected := fmt.Errorf("%s", errorMessage)
		name := "watchdog name"
		port := 1234
		logLevel := log.ERROR
		logChannel := "test channel"
		logStartMessage := "start message"
		logErrorMessage := "error message"
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *processConfig, icase ...bool) (interface{}, error) {
			c.Watchdog = name
			c.Port = port
			c.Log.Level = log.LevelMapName[logLevel]
			c.Log.Channel = logChannel
			c.Log.Message.Start = logStartMessage
			c.Log.Message.Error = logErrorMessage
			return c, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config(ConfigPath, gomock.Any()).Return(cfg, nil).Times(1)
		logger := NewMockLog(ctrl)
		gomock.InOrder(
			logger.EXPECT().Signal(logChannel, log.ERROR, logStartMessage, log.Context{"port": port}),
			logger.EXPECT().Signal(logChannel, log.FATAL, logErrorMessage, log.Context{"error": errorMessage}),
		)
		engine := NewMockEngine(ctrl)
		engine.EXPECT().Run(fmt.Sprintf(":%d", port)).Return(expected).Times(1)

		sut, _ := NewProcess(cfgManager, logger, engine)
		e := sut.Runner()()
		switch {
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, expected):
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})
}
