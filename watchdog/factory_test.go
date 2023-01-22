package watchdog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/err"
	"github.com/happyhippyhippo/slate/log"
)

func Test_Factory(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, e := NewFactory(nil, l, formatterFactory)

		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("nil log", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfigManager(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, e := NewFactory(cfg, nil, formatterFactory)

		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("nil formatter factory", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfigManager(ctrl)
		l := NewMockLog(ctrl)
		sut, e := NewFactory(cfg, l, nil)

		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("valid instantiation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		cfg := NewMockConfigManager(ctrl)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, e := NewFactory(cfg, l, formatterFactory)

		switch {
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		case sut.(*Factory).config != cfg:
			t.Errorf("didn't store the given config instance")
		case sut.(*Factory).log != l:
			t.Errorf("didn't store the given log instance")
		case sut.(*Factory).formatterFactory != formatterFactory:
			t.Errorf("didn't store the given formatter instance")
		}
	})
}

func Test_Factory_Create(t *testing.T) {
	t.Run("error retrieving config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		service := "service"
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("slate.watchdog.service", gomock.Any()).Return(nil, expected).Times(1)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("try to get config with path from env", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "test"
		defer func() { ConfigPathPrefix = prev }()

		expected := fmt.Errorf("error message")
		service := "service"
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("test.service", gomock.Any()).Return(nil, expected).Times(1)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error populating the watchdog config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "test"
		defer func() { ConfigPathPrefix = prev }()

		expected := fmt.Errorf("error message")
		service := "service"
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).Return(nil, expected).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("test.service", gomock.Any()).Return(cfg, nil).Times(1)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("invalid start message level populated on the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "test"
		defer func() { ConfigPathPrefix = prev }()

		service := "service"
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *watchdogConfig, icase ...bool) (interface{}, error) {
			c.Level.Start = "invalid"
			return nil, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("test.service", gomock.Any()).Return(cfg, nil).Times(1)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("invalid error message level populated on the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "test"
		defer func() { ConfigPathPrefix = prev }()

		service := "service"
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *watchdogConfig, icase ...bool) (interface{}, error) {
			c.Level.Error = "invalid"
			return nil, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("test.service", gomock.Any()).Return(cfg, nil).Times(1)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("invalid done message level populated on the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "test"
		defer func() { ConfigPathPrefix = prev }()

		service := "service"
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *watchdogConfig, icase ...bool) (interface{}, error) {
			c.Level.Done = "invalid"
			return nil, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("test.service", gomock.Any()).Return(cfg, nil).Times(1)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, err.Conversion):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.Conversion)
		}
	})

	t.Run("error creating the log formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "test"
		defer func() { ConfigPathPrefix = prev }()

		expected := fmt.Errorf("error message")
		service := "service"
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).Return(nil, nil).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("test.service", gomock.Any()).Return(cfg, nil).Times(1)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(&config.Config{"type": FormatterDefault}).Return(nil, expected).Times(1)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("error creating the log formatter with type from config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "test"
		defer func() { ConfigPathPrefix = prev }()

		expected := fmt.Errorf("error message")
		service := "service"
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *watchdogConfig, icase ...bool) (interface{}, error) {
			c.Formatter = "my_formatter"
			return nil, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("test.service", gomock.Any()).Return(cfg, nil).Times(1)
		l := NewMockLog(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(&config.Config{"type": "my_formatter"}).Return(nil, expected).Times(1)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("returned the (%v) error when expecting (%v)", e, expected)
		}
	})

	t.Run("create the watchdog", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "test"
		defer func() { ConfigPathPrefix = prev }()

		service := "service"
		cfg := NewMockConfig(ctrl)
		cfg.EXPECT().Populate("", gomock.Any()).DoAndReturn(func(path string, c *watchdogConfig, icase ...bool) (interface{}, error) {
			c.Service = service
			return nil, nil
		}).Times(1)
		cfgManager := NewMockConfigManager(ctrl)
		cfgManager.EXPECT().Config("test.service", gomock.Any()).Return(cfg, nil).Times(1)
		l := NewMockLog(ctrl)
		formatter := NewMockLogFormatter(ctrl)
		formatterFactory := NewMockLogFormatterFactory(ctrl)
		formatterFactory.EXPECT().Create(&config.Config{"type": FormatterDefault}).Return(formatter, nil).Times(1)
		sut, _ := NewFactory(cfgManager, l, formatterFactory)

		wd, e := sut.Create(service)
		switch {
		case e != nil:
			t.Errorf("returned the unexpected error : %v", e)
		case wd == nil:
			t.Errorf("didn't returned the expected watchdog reference")
		case wd.(*Watchdog).log.(*LogAdapter).service != service:
			t.Errorf("didn't stored the expected service name")
		case wd.(*Watchdog).log.(*LogAdapter).channel != LogChannel:
			t.Errorf("didn't stored the expected channel name")
		case wd.(*Watchdog).log.(*LogAdapter).startLevel != log.LevelMap[LogStartLevel]:
			t.Errorf("didn't stored the expected start log message level")
		case wd.(*Watchdog).log.(*LogAdapter).errorLevel != log.LevelMap[LogErrorLevel]:
			t.Errorf("didn't stored the expected error log message level")
		case wd.(*Watchdog).log.(*LogAdapter).doneLevel != log.LevelMap[LogDoneLevel]:
			t.Errorf("didn't stored the expected done log message level")
		case wd.(*Watchdog).log.(*LogAdapter).logger != l:
			t.Errorf("didn't stored the expected logger")
		case wd.(*Watchdog).log.(*LogAdapter).formatter != formatter:
			t.Errorf("didn't stored the expected formatter")
		}
	})
}
