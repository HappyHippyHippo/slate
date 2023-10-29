package slate

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_Watchdog_err(t *testing.T) {
	t.Run("errInvalidWatchdogLogWriter", func(t *testing.T) {
		arg := &ConfigPartial{"field": "value"}
		context := map[string]interface{}{"field": "value"}
		message := "&map[field:value] : invalid watchdog log writer"

		t.Run("creation without context", func(t *testing.T) {
			if e := errInvalidWatchdogLogWriter(arg); !errors.Is(e, ErrInvalidWatchdogLogWriter) {
				t.Errorf("error not a instance of ErrInvalidWatchdog")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errInvalidWatchdogLogWriter(arg, context); !errors.Is(e, ErrInvalidWatchdogLogWriter) {
				t.Errorf("error not a instance of ErrInvalidWatchdog")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})

	t.Run("errDuplicateWatchdog", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : duplicate watchdog"

		t.Run("creation without context", func(t *testing.T) {
			if e := errDuplicateWatchdog(arg); !errors.Is(e, ErrDuplicateWatchdog) {
				t.Errorf("error not a instance of ErrDuplicateService")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})

		t.Run("creation with context", func(t *testing.T) {
			if e := errDuplicateWatchdog(arg, context); !errors.Is(e, ErrDuplicateWatchdog) {
				t.Errorf("error not a instance of ErrDuplicateService")
			} else if e.Error() != message {
				t.Errorf("error message (%v) not same as expected (%v)", e, message)
			} else {
				var te *Error
				if !errors.As(e, &te) {
					t.Errorf("didn't returned a slate error instance")
				}
			}
		})
	})
}

func Test_WatchdogLogFormatterFactory(t *testing.T) {
	t.Run("NewWatchdogLogFormatterFactory", func(t *testing.T) {
		t.Run("creation with empty strategy list", func(t *testing.T) {
			sut := NewWatchdogLogFormatterFactory(nil)
			if sut == nil {
				t.Error("didn't returned the expected reference")
			}
		})

		t.Run("creation with strategy list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			creator := NewMockWatchdogLogFormatterCreator(ctrl)

			sut := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{creator})
			if sut == nil {
				t.Error("didn't returned the expected reference")
			} else if (*sut)[0] != creator {
				t.Error("didn't stored the passed strategy")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("nil config", func(t *testing.T) {
			sut := NewWatchdogLogFormatterFactory(nil)
			src, e := sut.Create(nil)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("unrecognized format type", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := &ConfigPartial{}
			strategy := NewMockWatchdogLogFormatterCreator(ctrl)
			strategy.EXPECT().Accept(config).Return(false).Times(1)
			sut := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{strategy})

			stream, e := sut.Create(config)
			switch {
			case stream != nil:
				t.Error("returned a config stream")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidWatchdogLogWriter):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidWatchdogLogWriter)
			}
		})

		t.Run("create the formatter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := &ConfigPartial{}
			formatter := NewMockWatchdogLogFormatter(ctrl)
			strategy := NewMockWatchdogLogFormatterCreator(ctrl)
			strategy.EXPECT().Accept(config).Return(true).Times(1)
			strategy.EXPECT().Create(config).Return(formatter, nil).Times(1)
			sut := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{strategy})

			if s, e := sut.Create(config); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(s, formatter) {
				t.Error("didn't returned the created stream")
			}
		})
	})
}

func Test_WatchdogDefaultLogFormatter(t *testing.T) {
	t.Run("Start", func(t *testing.T) {
		t.Run("format message", func(t *testing.T) {
			service := "service name"
			expected := fmt.Sprintf(WatchdogLogStartMessage, service)

			if check := NewWatchdogDefaultLogFormatter().Start(service); check != expected {
				t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
			}
		})

		t.Run("format message from env", func(t *testing.T) {
			prev := WatchdogLogStartMessage
			message := "test message : %s"
			WatchdogLogStartMessage = message
			defer func() { WatchdogLogStartMessage = prev }()

			service := "service name"
			expected := fmt.Sprintf(WatchdogLogStartMessage, service)

			if check := NewWatchdogDefaultLogFormatter().Start(service); check != expected {
				t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
			}
		})
	})

	t.Run("Error", func(t *testing.T) {
		t.Run("format message", func(t *testing.T) {
			service := "service name"
			err := fmt.Errorf("error message")
			expected := fmt.Sprintf(WatchdogLogErrorMessage, service, err)

			if check := NewWatchdogDefaultLogFormatter().Error(service, err); check != expected {
				t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
			}
		})

		t.Run("format message from env", func(t *testing.T) {
			prev := WatchdogLogErrorMessage
			message := "test message : %s - %e"
			WatchdogLogErrorMessage = message
			defer func() { WatchdogLogErrorMessage = prev }()

			service := "service name"
			err := fmt.Errorf("error message")
			expected := fmt.Sprintf(WatchdogLogErrorMessage, service, err)

			if check := NewWatchdogDefaultLogFormatter().Error(service, err); check != expected {
				t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
			}
		})
	})

	t.Run("Done", func(t *testing.T) {
		t.Run("format message", func(t *testing.T) {
			service := "service name"
			expected := fmt.Sprintf(WatchdogLogDoneMessage, service)

			if check := NewWatchdogDefaultLogFormatter().Done(service); check != expected {
				t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
			}
		})

		t.Run("format message from env", func(t *testing.T) {
			prev := WatchdogLogDoneMessage
			message := "test message : %s"
			WatchdogLogDoneMessage = message
			defer func() { WatchdogLogDoneMessage = prev }()

			service := "service name"
			expected := fmt.Sprintf(WatchdogLogDoneMessage, service)

			if check := NewWatchdogDefaultLogFormatter().Done(service); check != expected {
				t.Errorf("returned the (%v) message when expecting (%v)", check, expected)
			}
		})
	})
}

func Test_WatchdogDefaultLogFormatterCreator(t *testing.T) {
	t.Run("Accept", func(t *testing.T) {
		t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
			if NewWatchdogDefaultLogFormatterCreator().Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept on type retrieval error", func(t *testing.T) {
			if NewWatchdogDefaultLogFormatterCreator().Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept on invalid type", func(t *testing.T) {
			if NewWatchdogDefaultLogFormatterCreator().Accept(&ConfigPartial{"type": "invalid"}) {
				t.Error("returned true")
			}
		})

		t.Run("accept on valid type", func(t *testing.T) {
			if !NewWatchdogDefaultLogFormatterCreator().Accept(&ConfigPartial{"type": "default"}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("new formatter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			stream, e := NewWatchdogDefaultLogFormatterCreator().Create(&ConfigPartial{})
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case stream == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch stream.(type) {
				case *WatchdogDefaultLogFormatter:
				default:
					t.Error("didn't returned a log formatter")
				}
			}
		})
	})
}

func Test_WatchdogLogAdapter(t *testing.T) {
	t.Run("NewWatchdogLogAdapter", func(t *testing.T) {
		t.Run("nil log", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			formatter := NewMockWatchdogLogFormatter(ctrl)

			sut, e := NewWatchdogLogAdapter("service", "channel", FATAL, FATAL, FATAL, nil, formatter)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil formatter", func(t *testing.T) {
			sut, e := NewWatchdogLogAdapter("service", "channel", FATAL, FATAL, FATAL, NewLog(), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("valid instantiation", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			logger := NewLog()
			formatter := NewMockWatchdogLogFormatter(ctrl)

			sut, e := NewWatchdogLogAdapter("service", "channel", FATAL, ERROR, WARNING, logger, formatter)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case sut.name != "service":
				t.Errorf("invalid service name : %v", sut.name)
			case sut.channel != "channel":
				t.Errorf("invalid channel name : %v", sut.channel)
			case sut.startLevel != FATAL:
				t.Errorf("invalid start log level : %v", LogLevelMapName[sut.startLevel])
			case sut.errorLevel != ERROR:
				t.Errorf("invalid error log level : %v", LogLevelMapName[sut.errorLevel])
			case sut.doneLevel != WARNING:
				t.Errorf("invalid done log level : %v", LogLevelMapName[sut.doneLevel])
			case sut.logger != logger:
				t.Errorf("invalid log instance")
			case sut.formatter != formatter:
				t.Errorf("invalid formatter instance")
			}
		})
	})

	t.Run("Start", func(t *testing.T) {
		t.Run("error while logging", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "service name"
			channel := "channel name"
			message := "formatter message"
			expected := fmt.Errorf("error message")

			logger := NewMockWatchdogLogger(ctrl)
			logger.EXPECT().Signal(channel, FATAL, message).Return(expected).Times(1)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			formatter.EXPECT().Start(service).Return(message).Times(1)
			sut, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			sut.logger = logger

			chk := sut.Start()
			switch {
			case chk == nil:
				t.Errorf("didn't returned the expected error")
			case chk.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", chk, expected)
			}
		})

		t.Run("success logging", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "service name"
			channel := "channel name"
			message := "formatter message"

			logger := NewMockWatchdogLogger(ctrl)
			logger.EXPECT().Signal(channel, FATAL, message).Return(nil).Times(1)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			formatter.EXPECT().Start(service).Return(message).Times(1)
			sut, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			sut.logger = logger

			if chk := sut.Start(); chk != nil {
				t.Errorf("unexpected (%v) error", chk)
			}
		})
	})

	t.Run("Error", func(t *testing.T) {
		t.Run("error while logging", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := fmt.Errorf("error")
			service := "service name"
			channel := "channel name"
			message := "formatter message"
			expected := fmt.Errorf("error message")

			logger := NewMockWatchdogLogger(ctrl)
			logger.EXPECT().Signal(channel, ERROR, message).Return(expected).Times(1)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			formatter.EXPECT().Error(service, e).Return(message).Times(1)
			sut, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			sut.logger = logger

			chk := sut.Error(e)
			switch {
			case chk == nil:
				t.Errorf("didn't returned the expected error")
			case chk.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", chk, expected)
			}
		})

		t.Run("success logging", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := fmt.Errorf("error")
			service := "service name"
			channel := "channel name"
			message := "formatter message"

			logger := NewMockWatchdogLogger(ctrl)
			logger.EXPECT().Signal(channel, ERROR, message).Return(nil).Times(1)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			formatter.EXPECT().Error(service, e).Return(message).Times(1)
			sut, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			sut.logger = logger

			if chk := sut.Error(e); chk != nil {
				t.Errorf("unexpected (%v) error", chk)
			}
		})
	})

	t.Run("Done", func(t *testing.T) {
		t.Run("error while logging", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "service name"
			channel := "channel name"
			message := "formatter message"
			expected := fmt.Errorf("error message")

			logger := NewMockWatchdogLogger(ctrl)
			logger.EXPECT().Signal(channel, WARNING, message).Return(expected).Times(1)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			formatter.EXPECT().Done(service).Return(message).Times(1)
			sut, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			sut.logger = logger

			chk := sut.Done()
			switch {
			case chk == nil:
				t.Errorf("didn't returned the expected error")
			case chk.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", chk, expected)
			}
		})

		t.Run("success logging", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "service name"
			channel := "channel name"
			message := "formatter message"

			logger := NewMockWatchdogLogger(ctrl)
			logger.EXPECT().Signal(channel, WARNING, message).Return(nil).Times(1)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			formatter.EXPECT().Done(service).Return(message).Times(1)
			sut, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			sut.logger = logger

			if chk := sut.Done(); chk != nil {
				t.Errorf("unexpected (%v) error", chk)
			}
		})
	})
}

func Test_Watchdog(t *testing.T) {
	t.Run("NewWatchdog", func(t *testing.T) {
		t.Run("nil logAdapter", func(t *testing.T) {
			sut, e := NewWatchdog(nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("valid instantiation", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "service name"
			channel := "channel name"
			formatter := NewMockWatchdogLogFormatter(ctrl)
			logAdapter, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)

			sut, e := NewWatchdog(logAdapter)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case sut.logAdapter != logAdapter:
				t.Errorf("didn't store the given logAdapter instance")
			}
		})
	})

	t.Run("Run", func(t *testing.T) {
		t.Run("simple execution", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "service"
			channel := "channel"
			logger := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger.EXPECT().Signal(channel, FATAL, "start formatted message").Return(nil).Times(1),
				logger.EXPECT().Signal(channel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter.EXPECT().Start(service).Return("start formatted message").Times(1),
				formatter.EXPECT().Done(service).Return("done formatted message").Times(1),
			)
			logAdapter, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			logAdapter.logger = logger

			sut, _ := NewWatchdog(logAdapter)

			count := 0
			p, _ := NewWatchdogProcess(service, func() error {
				count++
				return nil
			})

			chk := sut.Run(p)
			switch {
			case count != 1:
				t.Errorf("didn't executed the process method")
			case chk != nil:
				t.Errorf("unexpected (%v) error", chk)
			}
		})

		t.Run("simple execution but return an error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := fmt.Errorf("error message")
			service := "service"
			channel := "channel"
			logger := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger.EXPECT().Signal(channel, FATAL, "start formatted message").Return(nil).Times(1),
				logger.EXPECT().Signal(channel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter.EXPECT().Start(service).Return("start formatted message").Times(1),
				formatter.EXPECT().Done(service).Return("done formatted message").Times(1),
			)
			logAdapter, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			logAdapter.logger = logger

			sut, _ := NewWatchdog(logAdapter)

			count := 0
			p, _ := NewWatchdogProcess(service, func() error {
				count++
				return e
			})

			chk := sut.Run(p)
			switch {
			case count != 1:
				t.Errorf("didn't executed the process method")
			case chk == nil:
				t.Error("didn't returned the expected error")
			case chk.Error() != e.Error():
				t.Errorf("(%v) when expecting (%v)", chk, e)
			}
		})

		t.Run("panic recovery on execution", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := fmt.Errorf("error message")
			service := "service"
			channel := "channel"
			logger := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger.EXPECT().Signal(channel, FATAL, "start formatted message").Return(nil).Times(1),
				logger.EXPECT().Signal(channel, ERROR, "error formatted message").Return(nil).Times(1),
				logger.EXPECT().Signal(channel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter.EXPECT().Start(service).Return("start formatted message").Times(1),
				formatter.EXPECT().Error(service, e).Return("error formatted message").Times(1),
				formatter.EXPECT().Done(service).Return("done formatted message").Times(1),
			)
			logAdapter, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			logAdapter.logger = logger

			sut, _ := NewWatchdog(logAdapter)

			count := 0
			p, _ := NewWatchdogProcess(service, func() error {
				count++
				if count == 1 {
					panic(e)
				}
				return nil
			})

			chk := sut.Run(p)
			switch {
			case count != 2:
				t.Errorf("didn't executed the process method two times")
			case chk != nil:
				t.Errorf("unexpected (%v) error", chk)
			}
		})

		t.Run("panic recovery on execution and error return", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			e := fmt.Errorf("error message")
			service := "service"
			channel := "channel"
			logger := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger.EXPECT().Signal(channel, FATAL, "start formatted message").Return(nil).Times(1),
				logger.EXPECT().Signal(channel, ERROR, "error formatted message").Return(nil).Times(1),
				logger.EXPECT().Signal(channel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter.EXPECT().Start(service).Return("start formatted message").Times(1),
				formatter.EXPECT().Error(service, e).Return("error formatted message").Times(1),
				formatter.EXPECT().Done(service).Return("done formatted message").Times(1),
			)
			logAdapter, _ := NewWatchdogLogAdapter(service, channel, FATAL, ERROR, WARNING, NewLog(), formatter)
			logAdapter.logger = logger

			sut, _ := NewWatchdog(logAdapter)

			count := 0
			p, _ := NewWatchdogProcess(service, func() error {
				count++
				if count == 1 {
					panic(e)
				}
				return e
			})

			chk := sut.Run(p)
			switch {
			case count != 2:
				t.Errorf("didn't executed the process method two times")
			case chk == nil:
				t.Error("didn't returned the expected error")
			case chk.Error() != e.Error():
				t.Errorf("(%v) when expecting (%v)", chk, e)
			}
		})
	})
}

func Test_WatchdogFactory(t *testing.T) {
	t.Run("NewWatchdogFactory", func(t *testing.T) {
		t.Run("nil config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewWatchdogFactory(nil, NewLog(), NewWatchdogLogFormatterFactory(nil))

			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil log", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewWatchdogFactory(NewConfig(), nil, NewWatchdogLogFormatterFactory(nil))

			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil formatter creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewWatchdogFactory(NewConfig(), NewLog(), nil)

			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("valid instantiation", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewWatchdogFactory(NewConfig(), NewLog(), NewWatchdogLogFormatterFactory(nil))

			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error retrieving config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "test"
			supplier := NewMockConfigSupplier(ctrl)
			supplier.
				EXPECT().
				Get("").
				Return(ConfigPartial{"slate": ConfigPartial{"watchdog": ConfigPartial{"services": ConfigPartial{"test": 123}}}}, nil).
				Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)

			sut, e := NewWatchdogFactory(config, NewLog(), NewWatchdogLogFormatterFactory(nil))

			chk, e := sut.Create(service)
			switch {
			case chk != nil:
				t.Errorf("returned an unexpected watchdog reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("try to get config with path from env", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			prev := WatchdogConfigPathPrefix
			WatchdogConfigPathPrefix = "path"
			defer func() { WatchdogConfigPathPrefix = prev }()

			service := "test"
			supplier := NewMockConfigSupplier(ctrl)
			supplier.
				EXPECT().
				Get("").
				Return(ConfigPartial{"path": ConfigPartial{"test": 123}}, nil).
				Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)

			sut, e := NewWatchdogFactory(config, NewLog(), NewWatchdogLogFormatterFactory(nil))

			chk, e := sut.Create(service)
			switch {
			case chk != nil:
				t.Errorf("returned an unexpected watchdog reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("error populating the watchdog config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "test"
			cfgData := ConfigPartial{"name": 123}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.
				EXPECT().
				Get("").
				Return(ConfigPartial{"slate": ConfigPartial{"watchdog": ConfigPartial{"services": ConfigPartial{"test": cfgData}}}}, nil).
				Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)

			sut, e := NewWatchdogFactory(config, NewLog(), NewWatchdogLogFormatterFactory(nil))

			chk, e := sut.Create(service)
			switch {
			case chk != nil:
				t.Errorf("returned an unexpected watchdog reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid start message level populated on the config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "test"
			cfgData := ConfigPartial{"level": ConfigPartial{"start": "invalid"}}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.
				EXPECT().
				Get("").
				Return(ConfigPartial{"slate": ConfigPartial{"watchdog": ConfigPartial{"services": ConfigPartial{"test": cfgData}}}}, nil).
				Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)

			sut, e := NewWatchdogFactory(config, NewLog(), NewWatchdogLogFormatterFactory(nil))

			chk, e := sut.Create(service)
			switch {
			case chk != nil:
				t.Errorf("returned an unexpected watchdog reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid error message level populated on the config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "test"
			cfgData := ConfigPartial{"level": ConfigPartial{"error": "invalid"}}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.
				EXPECT().
				Get("").
				Return(ConfigPartial{"slate": ConfigPartial{"watchdog": ConfigPartial{"services": ConfigPartial{"test": cfgData}}}}, nil).
				Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)

			sut, e := NewWatchdogFactory(config, NewLog(), NewWatchdogLogFormatterFactory(nil))

			chk, e := sut.Create(service)
			switch {
			case chk != nil:
				t.Errorf("returned an unexpected watchdog reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid done message level populated on the config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "test"
			cfgData := ConfigPartial{"level": ConfigPartial{"done": "invalid"}}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.
				EXPECT().
				Get("").
				Return(ConfigPartial{"slate": ConfigPartial{"watchdog": ConfigPartial{"services": ConfigPartial{"test": cfgData}}}}, nil).
				Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)

			sut, e := NewWatchdogFactory(config, NewLog(), NewWatchdogLogFormatterFactory(nil))

			chk, e := sut.Create(service)
			switch {
			case chk != nil:
				t.Errorf("returned an unexpected watchdog reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("error creating the log formatter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "test"
			cfgData := ConfigPartial{"formatter": "invalid"}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.
				EXPECT().
				Get("").
				Return(ConfigPartial{"slate": ConfigPartial{"watchdog": ConfigPartial{"services": ConfigPartial{"test": cfgData}}}}, nil).
				Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)

			sut, e := NewWatchdogFactory(config, NewLog(), NewWatchdogLogFormatterFactory(nil))

			chk, e := sut.Create(service)
			switch {
			case chk != nil:
				t.Errorf("returned an unexpected watchdog reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidWatchdogLogWriter):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidWatchdogLogWriter)
			}
		})

		t.Run("create the watchdog", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "test"
			formatterConfig := ConfigPartial{"type": "my_formatter"}
			watchdogConfig := ConfigPartial{"formatter": "my_formatter"}
			supplier := NewMockConfigSupplier(ctrl)
			supplier.
				EXPECT().
				Get("").
				Return(ConfigPartial{"slate": ConfigPartial{"watchdog": ConfigPartial{"services": ConfigPartial{"test": watchdogConfig}}}}, nil).
				Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 0, supplier)
			logger := NewLog()
			formatter := NewMockWatchdogLogFormatter(ctrl)
			formatterCreator := NewMockWatchdogLogFormatterCreator(ctrl)
			formatterCreator.EXPECT().Accept(&formatterConfig).Return(true).Times(1)
			formatterCreator.EXPECT().Create(&formatterConfig).Return(formatter, nil).Times(1)
			formatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{formatterCreator})

			sut, e := NewWatchdogFactory(config, logger, formatterFactory)

			wd, e := sut.Create(service)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case wd == nil:
				t.Errorf("didn't returned the expected watchdog reference")
			case wd.logAdapter.name != service:
				t.Errorf("didn't stored the expected service name")
			case wd.logAdapter.channel != WatchdogLogChannel:
				t.Errorf("didn't stored the expected channel name")
			case wd.logAdapter.startLevel != LogLevelMap[WatchdogLogStartLevel]:
				t.Errorf("didn't stored the expected start log message level")
			case wd.logAdapter.errorLevel != LogLevelMap[WatchdogLogErrorLevel]:
				t.Errorf("didn't stored the expected error log message level")
			case wd.logAdapter.doneLevel != LogLevelMap[WatchdogLogDoneLevel]:
				t.Errorf("didn't stored the expected done log message level")
			case wd.logAdapter.logger != logger:
				t.Errorf("didn't stored the expected logger")
			case wd.logAdapter.formatter != formatter:
				t.Errorf("didn't stored the expected formatter")
			}
		})
	})

	t.Run("override watchdog name by config", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		service := "test"
		otherName := "other name"
		formatterConfig := ConfigPartial{"type": "my_formatter"}
		watchdogConfig := ConfigPartial{"name": otherName, "formatter": "my_formatter"}
		supplier := NewMockConfigSupplier(ctrl)
		supplier.
			EXPECT().
			Get("").
			Return(ConfigPartial{"slate": ConfigPartial{"watchdog": ConfigPartial{"services": ConfigPartial{"test": watchdogConfig}}}}, nil).
			Times(1)
		config := NewConfig()
		_ = config.AddSupplier("id", 0, supplier)
		logger := NewLog()
		formatter := NewMockWatchdogLogFormatter(ctrl)
		formatterCreator := NewMockWatchdogLogFormatterCreator(ctrl)
		formatterCreator.EXPECT().Accept(&formatterConfig).Return(true).Times(1)
		formatterCreator.EXPECT().Create(&formatterConfig).Return(formatter, nil).Times(1)
		formatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{formatterCreator})

		sut, e := NewWatchdogFactory(config, logger, formatterFactory)

		wd, e := sut.Create(service)
		switch {
		case e != nil:
			t.Errorf("unexpected (%v) error", e)
		case wd == nil:
			t.Errorf("didn't returned the expected watchdog reference")
		case wd.logAdapter.name != otherName:
			t.Errorf("didn't stored the expected service name")
		case wd.logAdapter.channel != WatchdogLogChannel:
			t.Errorf("didn't stored the expected channel name")
		case wd.logAdapter.startLevel != LogLevelMap[WatchdogLogStartLevel]:
			t.Errorf("didn't stored the expected start log message level")
		case wd.logAdapter.errorLevel != LogLevelMap[WatchdogLogErrorLevel]:
			t.Errorf("didn't stored the expected error log message level")
		case wd.logAdapter.doneLevel != LogLevelMap[WatchdogLogDoneLevel]:
			t.Errorf("didn't stored the expected done log message level")
		case wd.logAdapter.logger != logger:
			t.Errorf("didn't stored the expected logger")
		case wd.logAdapter.formatter != formatter:
			t.Errorf("didn't stored the expected formatter")
		}
	})
}

func Test_NewWatchdogProcess(t *testing.T) {
	t.Run("nil runner", func(t *testing.T) {
		t.Run("nil runner", func(t *testing.T) {
			sut, e := NewWatchdogProcess("service", nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new process", func(t *testing.T) {
			service := "service name"
			runner := func() error { return nil }

			sut, e := NewWatchdogProcess(service, runner)
			switch {
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut.service != service:
				t.Errorf("(%v) service when expecting (%v)", sut.service, service)
			case fmt.Sprintf("%p", sut.runner) != fmt.Sprintf("%p", runner):
				t.Errorf("(%p) runner when expecting (%p)", sut.runner, runner)
			}
		})
	})

	t.Run("Service", func(t *testing.T) {
		t.Run("retrieve the service name", func(t *testing.T) {
			service := "service name"
			runner := func() error { return nil }
			sut, _ := NewWatchdogProcess(service, runner)

			if chk := sut.Service(); chk != service {
				t.Errorf("(%v) service when expecting (%v)", sut.service, service)
			}
		})
	})

	t.Run("Runner", func(t *testing.T) {
		t.Run("retrieve the runner method", func(t *testing.T) {
			service := "service name"
			runner := func() error { return nil }
			sut, _ := NewWatchdogProcess(service, runner)

			if chk := sut.Runner(); fmt.Sprintf("%p", chk) != fmt.Sprintf("%p", runner) {
				t.Errorf("(%p) runner when expecting (%p)", sut.runner, runner)
			}
		})
	})
}

func Test_WatchdogKennel(t *testing.T) {
	t.Run("NewWatchdogKennel", func(t *testing.T) {
		t.Run("nil factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewWatchdogKennel(nil, nil)

			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("duplicate process name", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			watchdogLogFormatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{NewWatchdogDefaultLogFormatterCreator()})
			watchdogFactory, _ := NewWatchdogFactory(NewConfig(), NewLog(), watchdogLogFormatterFactory)
			process1, _ := NewWatchdogProcess("process", func() error { return nil })
			process2, _ := NewWatchdogProcess("process", func() error { return nil })
			sut, e := NewWatchdogKennel(watchdogFactory, []WatchdogProcessor{process1, process2})

			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrDuplicateWatchdog):
				t.Errorf("(%v) when expecting (%v)", e, ErrDuplicateWatchdog)
			}
		})

		t.Run("error generating the watchdog", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			watchdogLogFormatterFactory := NewWatchdogLogFormatterFactory(nil)
			watchdogFactory, _ := NewWatchdogFactory(NewConfig(), NewLog(), watchdogLogFormatterFactory)
			process, _ := NewWatchdogProcess("process", func() error { return nil })
			sut, e := NewWatchdogKennel(watchdogFactory, []WatchdogProcessor{process})

			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidWatchdogLogWriter):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidWatchdogLogWriter)
			}
		})

		t.Run("valid instantiation", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			watchdogLogFormatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{NewWatchdogDefaultLogFormatterCreator()})
			watchdogFactory, _ := NewWatchdogFactory(NewConfig(), NewLog(), watchdogLogFormatterFactory)
			sut, e := NewWatchdogKennel(watchdogFactory, nil)

			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case sut.watchdogFactory != watchdogFactory:
				t.Errorf("didn't store the given creator instance")
			case sut.regs == nil:
				t.Errorf("didn't initialize the kennel registration map")
			}
		})

		t.Run("instantiation is processes", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			watchdogLogFormatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{NewWatchdogDefaultLogFormatterCreator()})
			watchdogFactory, _ := NewWatchdogFactory(NewConfig(), NewLog(), watchdogLogFormatterFactory)
			process, _ := NewWatchdogProcess("process", func() error { return nil })
			sut, e := NewWatchdogKennel(watchdogFactory, []WatchdogProcessor{process})

			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case sut.watchdogFactory != watchdogFactory:
				t.Errorf("didn't store the given creator instance")
			case len(sut.regs) != 1:
				t.Errorf("didn't initialize the kennel registration map with the given process")
			}
		})
	})

	t.Run("Run", func(t *testing.T) {
		t.Run("no-op if no process was registered", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			watchdogLogFormatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{NewWatchdogDefaultLogFormatterCreator()})
			watchdogFactory, _ := NewWatchdogFactory(NewConfig(), NewLog(), watchdogLogFormatterFactory)
			sut, _ := NewWatchdogKennel(watchdogFactory, nil)

			if e := sut.Run(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("simple process", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := "service"
			logger := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger.EXPECT().Signal(WatchdogLogChannel, FATAL, "start formatted message").Return(nil).Times(1),
				logger.EXPECT().Signal(WatchdogLogChannel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter.EXPECT().Start(service).Return("start formatted message").Times(1),
				formatter.EXPECT().Done(service).Return("done formatted message").Times(1),
			)
			logAdapter, _ := NewWatchdogLogAdapter(service, WatchdogLogChannel, FATAL, ERROR, WARNING, NewLog(), formatter)
			logAdapter.logger = logger

			watchdogLogFormatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{NewWatchdogDefaultLogFormatterCreator()})
			watchdogFactory, _ := NewWatchdogFactory(NewConfig(), NewLog(), watchdogLogFormatterFactory)
			process, _ := NewWatchdogProcess(service, func() error { return nil })
			sut, _ := NewWatchdogKennel(watchdogFactory, []WatchdogProcessor{process})
			sut.regs[service].watchdog.logAdapter = logAdapter

			if e := sut.Run(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("multiple processes", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service1 := "service.1"
			logger1 := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger1.EXPECT().Signal(WatchdogLogChannel, FATAL, "start formatted message").Return(nil).Times(1),
				logger1.EXPECT().Signal(WatchdogLogChannel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter1 := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter1.EXPECT().Start(service1).Return("start formatted message").Times(1),
				formatter1.EXPECT().Done(service1).Return("done formatted message").Times(1),
			)
			logAdapter1, _ := NewWatchdogLogAdapter(service1, WatchdogLogChannel, FATAL, ERROR, WARNING, NewLog(), formatter1)
			logAdapter1.logger = logger1
			process1, _ := NewWatchdogProcess(service1, func() error { return nil })

			service2 := "service.2"
			logger2 := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger2.EXPECT().Signal(WatchdogLogChannel, FATAL, "start formatted message").Return(nil).Times(1),
				logger2.EXPECT().Signal(WatchdogLogChannel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter2 := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter2.EXPECT().Start(service2).Return("start formatted message").Times(1),
				formatter2.EXPECT().Done(service2).Return("done formatted message").Times(1),
			)
			logAdapter2, _ := NewWatchdogLogAdapter(service2, WatchdogLogChannel, FATAL, ERROR, WARNING, NewLog(), formatter2)
			logAdapter2.logger = logger2
			process2, _ := NewWatchdogProcess(service2, func() error { return nil })

			service3 := "service.3"
			logger3 := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger3.EXPECT().Signal(WatchdogLogChannel, FATAL, "start formatted message").Return(nil).Times(1),
				logger3.EXPECT().Signal(WatchdogLogChannel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter3 := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter3.EXPECT().Start(service3).Return("start formatted message").Times(1),
				formatter3.EXPECT().Done(service3).Return("done formatted message").Times(1),
			)
			logAdapter3, _ := NewWatchdogLogAdapter(service3, WatchdogLogChannel, FATAL, ERROR, WARNING, NewLog(), formatter3)
			logAdapter3.logger = logger3
			process3, _ := NewWatchdogProcess(service3, func() error { return nil })

			watchdogLogFormatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{NewWatchdogDefaultLogFormatterCreator()})
			watchdogFactory, _ := NewWatchdogFactory(NewConfig(), NewLog(), watchdogLogFormatterFactory)
			sut, _ := NewWatchdogKennel(watchdogFactory, []WatchdogProcessor{process1, process2, process3})
			sut.regs[service1].watchdog.logAdapter = logAdapter1
			sut.regs[service2].watchdog.logAdapter = logAdapter2
			sut.regs[service3].watchdog.logAdapter = logAdapter3

			if e := sut.Run(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("return process error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			panicError := fmt.Errorf("panic error")
			service1 := "service.1"
			logger1 := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger1.EXPECT().Signal(WatchdogLogChannel, FATAL, "start formatted message").Return(nil).Times(1),
				logger1.EXPECT().Signal(WatchdogLogChannel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter1 := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter1.EXPECT().Start(service1).Return("start formatted message").Times(1),
				formatter1.EXPECT().Done(service1).Return("done formatted message").Times(1),
			)
			logAdapter1, _ := NewWatchdogLogAdapter(service1, WatchdogLogChannel, FATAL, ERROR, WARNING, NewLog(), formatter1)
			logAdapter1.logger = logger1
			process1, _ := NewWatchdogProcess(service1, func() error { return nil })

			service2 := "service.2"
			logger2 := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger2.EXPECT().Signal(WatchdogLogChannel, FATAL, "start formatted message").Return(nil).Times(1),
				logger2.EXPECT().Signal(WatchdogLogChannel, ERROR, "error formatted message").Return(nil).Times(1),
				logger2.EXPECT().Signal(WatchdogLogChannel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter2 := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter2.EXPECT().Start(service2).Return("start formatted message").Times(1),
				formatter2.EXPECT().Error(service2, panicError).Return("error formatted message").Times(1),
				formatter2.EXPECT().Done(service2).Return("done formatted message").Times(1),
			)
			logAdapter2, _ := NewWatchdogLogAdapter(service2, WatchdogLogChannel, FATAL, ERROR, WARNING, NewLog(), formatter2)
			logAdapter2.logger = logger2
			count := 0
			process2, _ := NewWatchdogProcess(service2, func() error {
				count++
				if count == 1 {
					panic(panicError)
				}
				return expected
			})

			service3 := "service.3"
			logger3 := NewMockWatchdogLogger(ctrl)
			gomock.InOrder(
				logger3.EXPECT().Signal(WatchdogLogChannel, FATAL, "start formatted message").Return(nil).Times(1),
				logger3.EXPECT().Signal(WatchdogLogChannel, WARNING, "done formatted message").Return(nil).Times(1),
			)
			formatter3 := NewMockWatchdogLogFormatter(ctrl)
			gomock.InOrder(
				formatter3.EXPECT().Start(service3).Return("start formatted message").Times(1),
				formatter3.EXPECT().Done(service3).Return("done formatted message").Times(1),
			)
			logAdapter3, _ := NewWatchdogLogAdapter(service3, WatchdogLogChannel, FATAL, ERROR, WARNING, NewLog(), formatter3)
			logAdapter3.logger = logger3
			process3, _ := NewWatchdogProcess(service3, func() error { return nil })

			watchdogLogFormatterFactory := NewWatchdogLogFormatterFactory([]WatchdogLogFormatterCreator{NewWatchdogDefaultLogFormatterCreator()})
			watchdogFactory, _ := NewWatchdogFactory(NewConfig(), NewLog(), watchdogLogFormatterFactory)
			sut, _ := NewWatchdogKennel(watchdogFactory, []WatchdogProcessor{process1, process2, process3})
			sut.regs[service1].watchdog.logAdapter = logAdapter1
			sut.regs[service2].watchdog.logAdapter = logAdapter2
			sut.regs[service3].watchdog.logAdapter = logAdapter3

			e := sut.Run()
			switch {
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})
	})
}

func Test_WatchdogServiceRegister(t *testing.T) {
	t.Run("NewWatchdogServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if NewWatchdogServiceRegister() == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewWatchdogServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewWatchdogServiceRegister().Provide(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expected (%v)", e, ErrNilPointer)
			}
		})

		t.Run("register components", func(t *testing.T) {
			container := NewServiceContainer()
			sut := NewWatchdogServiceRegister()

			e := sut.Provide(container)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !container.Has(WatchdogDefaultLogFormatterCreatorContainerID):
				t.Errorf("no default log formatter creator : %v", sut)
			case !container.Has(WatchdogAllLogFormatterCreatorsContainerID):
				t.Errorf("no log formatter aggreatate list : %v", sut)
			case !container.Has(WatchdogLogFormatterFactoryContainerID):
				t.Errorf("no log formatter factory : %v", sut)
			case !container.Has(WatchdogFactoryContainerID):
				t.Errorf("no watchdog creator : %v", sut)
			case !container.Has(WatchdogAllProcessesContainerID):
				t.Errorf("no process aggreatate list : %v", sut)
			case !container.Has(WatchdogContainerID):
				t.Errorf("no kannel : %v", sut)
			}
		})

		t.Run("retrieving log formatter creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)
			_ = NewWatchdogServiceRegister().Provide(container)

			sut, e := container.Get(WatchdogDefaultLogFormatterCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case sut == nil:
				t.Error("didn't returned a reference to service")
			default:
				switch sut.(type) {
				case *WatchdogDefaultLogFormatterCreator:
				default:
					t.Error("didn't returned the default log formatter creator")
				}
			}
		})

		t.Run("retrieving log formatter creators", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)
			_ = NewWatchdogServiceRegister().Provide(container)

			logFormatterCreator := NewMockWatchdogLogFormatterCreator(ctrl)
			_ = container.Add("creator.id", func() WatchdogLogFormatterCreator {
				return logFormatterCreator
			}, WatchdogLogFormatterCreatorTag)

			creators, e := container.Get(WatchdogAllLogFormatterCreatorsContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creators == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := creators.(type) {
				case []WatchdogLogFormatterCreator:
					found := false
					for _, creator := range c {
						if creator == logFormatterCreator {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a watchdog log formatter creator slice populated with the expected creator instance")
					}
				default:
					t.Error("didn't return a watchdog log formatter creator slice")
				}
			}
		})

		t.Run("retrieving log formatter factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)
			_ = NewWatchdogServiceRegister().Provide(container)

			sut, e := container.Get(WatchdogLogFormatterFactoryContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case sut == nil:
				t.Error("didn't returned a reference to service")
			default:
				switch sut.(type) {
				case *WatchdogLogFormatterFactory:
				default:
					t.Error("didn't returned the log formatter factory")
				}
			}
		})

		t.Run("retrieving watchdog factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)
			_ = NewWatchdogServiceRegister().Provide(container)

			sut, e := container.Get(WatchdogFactoryContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case sut == nil:
				t.Error("didn't returned a reference to service")
			default:
				switch sut.(type) {
				case *WatchdogFactory:
				default:
					t.Error("didn't returned the watchdog factory")
				}
			}
		})

		t.Run("retrieving all processes", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)
			_ = NewWatchdogServiceRegister().Provide(container)

			process, _ := NewWatchdogProcess("service", func() error { return nil })
			_ = container.Add("process.id", func() WatchdogProcessor {
				return process
			}, WatchdogProcessTag)

			processes, e := container.Get(WatchdogAllProcessesContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case processes == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := processes.(type) {
				case []WatchdogProcessor:
					found := false
					for _, p := range c {
						if p == process {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a watchdog processor slice populated with the expected process instance")
					}
				default:
					t.Error("didn't return a watchdog processor slice")
				}
			}
		})

		t.Run("retrieving watchdog kennel", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)
			_ = NewWatchdogServiceRegister().Provide(container)

			sut, e := container.Get(WatchdogContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case sut == nil:
				t.Error("didn't returned a reference to service")
			default:
				switch sut.(type) {
				case *WatchdogKennel:
				default:
					t.Error("didn't returned the watchdog kennel instance")
				}
			}
		})
	})
}
