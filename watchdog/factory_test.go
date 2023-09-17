package watchdog

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/config"
	"github.com/happyhippyhippo/slate/log"
)

func Test_Factory(t *testing.T) {
	t.Run("nil config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewFactory(nil, log.NewLog(), NewLogFormatterFactory())

		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil log", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewFactory(config.NewConfig(), nil, NewLogFormatterFactory())

		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("nil formatter creator", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewFactory(config.NewConfig(), log.NewLog(), nil)

		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("valid instantiation", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, e := NewFactory(config.NewConfig(), log.NewLog(), NewLogFormatterFactory())

		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case sut == nil:
			t.Errorf("didn't returned a valid reference")
		}
	})
}

func Test_Factory_Create(t *testing.T) {
	t.Run("error retrieving config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "test"
		expected := fmt.Errorf("error message")
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.watchdog.services.test", config.Partial{}).
			Return(nil, expected).
			Times(1)

		sut, _ := NewFactory(config.NewConfig(), log.NewLog(), NewLogFormatterFactory())
		sut.configurer = configurer

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("try to get config with path from env", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		prev := ConfigPathPrefix
		ConfigPathPrefix = "path"
		defer func() { ConfigPathPrefix = prev }()

		service := "test"
		expected := fmt.Errorf("error message")
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("path.test", config.Partial{}).
			Return(nil, expected).
			Times(1)

		sut, _ := NewFactory(config.NewConfig(), log.NewLog(), NewLogFormatterFactory())
		sut.configurer = configurer

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("error populating the watchdog configurer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "test"
		watchdogCfg := config.Partial{"formatter": 123}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.watchdog.services.test", config.Partial{}).
			Return(watchdogCfg, nil).
			Times(1)

		sut, _ := NewFactory(config.NewConfig(), log.NewLog(), NewLogFormatterFactory())
		sut.configurer = configurer

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("invalid start message level populated on the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "test"
		watchdogCfg := config.Partial{"level": config.Partial{"start": "invalid"}}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.watchdog.services.test", config.Partial{}).
			Return(watchdogCfg, nil).
			Times(1)

		sut, _ := NewFactory(config.NewConfig(), log.NewLog(), NewLogFormatterFactory())
		sut.configurer = configurer

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("invalid error message level populated on the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "test"
		watchdogCfg := config.Partial{"level": config.Partial{"error": "invalid"}}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.watchdog.services.test", config.Partial{}).
			Return(watchdogCfg, nil).
			Times(1)

		sut, _ := NewFactory(config.NewConfig(), log.NewLog(), NewLogFormatterFactory())
		sut.configurer = configurer

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("invalid done message level populated on the config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "test"
		watchdogCfg := config.Partial{"level": config.Partial{"done": "invalid"}}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.watchdog.services.test", config.Partial{}).
			Return(watchdogCfg, nil).
			Times(1)

		sut, _ := NewFactory(config.NewConfig(), log.NewLog(), NewLogFormatterFactory())
		sut.configurer = configurer

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case !errors.Is(e, slate.ErrConversion):
			t.Errorf("(%v) when expecting (%v)", e, slate.ErrConversion)
		}
	})

	t.Run("error creating the log formatter", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		expected := fmt.Errorf("error message")
		service := "test"
		watchdogCfg := config.Partial{"formatter": "my formatter"}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.watchdog.services.test", config.Partial{}).
			Return(watchdogCfg, nil).
			Times(1)
		logFormatterCreator := NewMockLogFormatterCreator(ctrl)
		logFormatterCreator.
			EXPECT().
			Create(&config.Partial{"type": "my formatter"}).
			Return(nil, expected).
			Times(1)

		sut, _ := NewFactory(config.NewConfig(), log.NewLog(), NewLogFormatterFactory())
		sut.configurer = configurer
		sut.logFormatterCreator = logFormatterCreator

		chk, e := sut.Create(service)
		switch {
		case chk != nil:
			t.Errorf("returned an unexpected watchdog reference")
		case e == nil:
			t.Errorf("didn't returned the expected error")
		case e.Error() != expected.Error():
			t.Errorf("(%v) when expecting (%v)", e, expected)
		}
	})

	t.Run("create the watchdog", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "test"
		watchdogCfg := config.Partial{"name": service}
		configurer := NewMockConfigurer(ctrl)
		configurer.
			EXPECT().
			Partial("slate.watchdog.services.test", config.Partial{}).
			Return(watchdogCfg, nil).
			Times(1)
		formatter := NewMockLogFormatter(ctrl)
		logFormatterCreator := NewMockLogFormatterCreator(ctrl)
		logFormatterCreator.
			EXPECT().
			Create(&config.Partial{"type": "simple"}).
			Return(formatter, nil).Times(1)
		logger := log.NewLog()

		sut, _ := NewFactory(config.NewConfig(), logger, NewLogFormatterFactory())
		sut.configurer = configurer
		sut.logFormatterCreator = logFormatterCreator

		wd, e := sut.Create(service)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case wd == nil:
			t.Errorf("didn't returned the expected watchdog reference")
		case wd.logAdapter.(*LogAdapter).name != service:
			t.Errorf("didn't stored the expected service name")
		case wd.logAdapter.(*LogAdapter).channel != LogChannel:
			t.Errorf("didn't stored the expected channel name")
		case wd.logAdapter.(*LogAdapter).startLevel != log.LevelMap[LogStartLevel]:
			t.Errorf("didn't stored the expected start log message level")
		case wd.logAdapter.(*LogAdapter).errorLevel != log.LevelMap[LogErrorLevel]:
			t.Errorf("didn't stored the expected error log message level")
		case wd.logAdapter.(*LogAdapter).doneLevel != log.LevelMap[LogDoneLevel]:
			t.Errorf("didn't stored the expected done log message level")
		case wd.logAdapter.(*LogAdapter).logger != logger:
			t.Errorf("didn't stored the expected logger")
		case wd.logAdapter.(*LogAdapter).formatter != formatter:
			t.Errorf("didn't stored the expected formatter")
		}
	})
}
