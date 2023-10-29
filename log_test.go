package slate

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"sort"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func Test_log_err(t *testing.T) {
	t.Run("errInvalidLogFormat", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : invalid log format"

		t.Run("creation without context", func(t *testing.T) {
			if e := errInvalidLogFormat(arg); !errors.Is(e, ErrInvalidLogFormat) {
				t.Errorf("error not a instance of ErrInvalidFormat")
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
			if e := errInvalidLogFormat(arg, context); !errors.Is(e, ErrInvalidLogFormat) {
				t.Errorf("error not a instance of ErrInvalidFormat")
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

	t.Run("errInvalidLogLevel", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : invalid log level"

		t.Run("creation without context", func(t *testing.T) {
			if e := errInvalidLogLevel(arg); !errors.Is(e, ErrInvalidLogLevel) {
				t.Errorf("error not a instance of ErrInvalidLevel")
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
			if e := errInvalidLogLevel(arg, context); !errors.Is(e, ErrInvalidLogLevel) {
				t.Errorf("error not a instance of ErrInvalidLevel")
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

	t.Run("errInvalidLogConfig", func(t *testing.T) {
		arg := ConfigPartial{"field": "value"}
		context := map[string]interface{}{"field": "value"}
		message := "map[field:value] : invalid log writer config"

		t.Run("creation without context", func(t *testing.T) {
			if e := errInvalidLogConfig(arg); !errors.Is(e, ErrInvalidLogConfig) {
				t.Errorf("error not a instance of ErrInvalidStream")
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
			if e := errInvalidLogConfig(arg, context); !errors.Is(e, ErrInvalidLogConfig) {
				t.Errorf("error not a instance of ErrInvalidStream")
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

	t.Run("errLogWriterNotFound", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : log writer not found"

		t.Run("creation without context", func(t *testing.T) {
			if e := errLogWriterNotFound(arg); !errors.Is(e, ErrLogWriterNotFound) {
				t.Errorf("error not a instance of ErrStreamNotFound")
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
			if e := errLogWriterNotFound(arg, context); !errors.Is(e, ErrLogWriterNotFound) {
				t.Errorf("error not a instance of ErrStreamNotFound")
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

	t.Run("errDuplicateLogWriter", func(t *testing.T) {
		arg := "dummy argument"
		context := map[string]interface{}{"field": "value"}
		message := "dummy argument : log writer already registered"

		t.Run("creation without context", func(t *testing.T) {
			if e := errDuplicateLogWriter(arg); !errors.Is(e, ErrDuplicateLogWriter) {
				t.Errorf("error not a instance of ErrDuplicateStream")
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
			if e := errDuplicateLogWriter(arg, context); !errors.Is(e, ErrDuplicateLogWriter) {
				t.Errorf("error not a instance of ErrDuplicateStream")
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

func Test_LogFormatterFactory(t *testing.T) {
	t.Run("NewLogFormatterFactory", func(t *testing.T) {
		t.Run("creation with empty creators list", func(t *testing.T) {
			sut := NewLogFormatterFactory(nil)
			if sut == nil {
				t.Error("didn't returned the expected reference")
			}
		})

		t.Run("creation with creators list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			creator := NewMockLogFormatterCreator(ctrl)

			sut := NewLogFormatterFactory([]LogFormatterCreator{creator})
			if sut == nil {
				t.Error("didn't returned the expected reference")
			} else if (*sut)[0] != creator {
				t.Error("didn't stored the passed creator")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("unrecognized format", func(t *testing.T) {
			format := "invalid format"

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			creator := NewMockLogFormatterCreator(ctrl)
			creator.EXPECT().Accept(format).Return(false).Times(1)
			sut := NewLogFormatterFactory([]LogFormatterCreator{creator})

			res, e := sut.Create(format)
			switch {
			case res != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidLogFormat):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidLogFormat)
			}
		})

		t.Run("create the formatter", func(t *testing.T) {
			format := "format"

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			formatter := NewMockLogFormatter(ctrl)
			creator := NewMockLogFormatterCreator(ctrl)
			creator.EXPECT().Accept(format).Return(true).Times(1)
			creator.EXPECT().Create().Return(formatter, nil).Times(1)
			sut := NewLogFormatterFactory([]LogFormatterCreator{creator})

			if res, e := sut.Create(format); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(res, formatter) {
				t.Errorf("didn't returned the formatter")
			}
		})
	})
}

func Test_LogJSONEncoder(t *testing.T) {
	t.Run("NewLogJSONEncoder", func(t *testing.T) {
		t.Run("creation", func(t *testing.T) {
			if NewLogJSONEncoder() == nil {
				t.Error("didn't returned the expected reference")
			}
		})
	})

	t.Run("Format", func(t *testing.T) {
		t.Run("correctly format the message", func(t *testing.T) {
			scenarios := []struct {
				level    LogLevel
				ctx      LogContext
				message  string
				expected string
			}{
				{ // _test level FATAL
					level:    FATAL,
					ctx:      nil,
					message:  "",
					expected: `"level"\s*\:\s*"FATAL"`,
				},
				{ // _test level ERROR
					level:    ERROR,
					ctx:      nil,
					message:  "",
					expected: `"level"\s*\:\s*"ERROR"`,
				},
				{ // _test level WARNING
					level:    WARNING,
					ctx:      nil,
					message:  "",
					expected: `"level"\s*\:\s*"WARNING"`,
				},
				{ // _test level NOTICE
					level:    NOTICE,
					ctx:      nil,
					message:  "",
					expected: `"level"\s*\:\s*"NOTICE"`,
				},
				{ // _test level INFO
					level:    INFO,
					ctx:      nil,
					message:  "",
					expected: `"level"\s*\:\s*"INFO"`,
				},
				{ // _test level DEBUG
					level:    DEBUG,
					ctx:      nil,
					message:  "",
					expected: `"level"\s*\:\s*"DEBUG"`,
				},
				{ // _test ctx (single value)
					level:    DEBUG,
					ctx:      LogContext{"field1": "value1"},
					message:  "",
					expected: `"field1"\s*\:\s*"value1"`,
				},
				{ // _test ctx (multiple value)
					level:    DEBUG,
					ctx:      LogContext{"field1": "value1", "field2": "value2"},
					message:  "",
					expected: `"field1"\s*\:\s*"value1"|"field2"\s*\:\s*"value2"`,
				},
				{ // _test message
					level:    DEBUG,
					ctx:      nil,
					message:  "My_message",
					expected: `"message"\s*\:\s*"My_message"`,
				},
			}

			for _, s := range scenarios {
				check := NewLogJSONEncoder().Format(s.level, s.message, s.ctx)
				match, _ := regexp.Match(s.expected, []byte(check))
				if !match {
					t.Errorf("didn't validated (%s) output", check)
				}
			}
		})
	})
}

func Test_LogJSONEncoderCreator(t *testing.T) {
	t.Run("NewLogJSONEncoderCreator", func(t *testing.T) {
		t.Run("creation", func(t *testing.T) {
			if NewLogJSONEncoderCreator() == nil {
				t.Error("didn't returned the expected reference")
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		t.Run("accept only json format", func(t *testing.T) {
			scenarios := []struct {
				format   string
				expected bool
			}{
				{ // _test json format
					format:   LogFormatJSON,
					expected: true,
				},
				{ // _test non-json format
					format:   "unknown",
					expected: false,
				},
			}

			for _, s := range scenarios {
				if check := NewLogJSONEncoderCreator().Accept(s.format); check != s.expected {
					t.Errorf("(%v) for the (%s) format", check, s.format)
				}
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("create json formatter", func(t *testing.T) {
			sut, e := NewLogJSONEncoderCreator().Create()
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch sut.(type) {
				case *LogJSONEncoder:
				default:
					t.Errorf("didn't returned a new json formatter")
				}
			}
		})
	})
}

func Test_LogWriterFactory(t *testing.T) {
	t.Run("NewLogWriterFactory", func(t *testing.T) {
		t.Run("creation with empty creator list", func(t *testing.T) {
			sut := NewLogWriterFactory(nil)
			if sut == nil {
				t.Error("didn't returned the expected reference")
			}
		})

		t.Run("creation with creator list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			creator := NewMockLogWriterCreator(ctrl)

			sut := NewLogWriterFactory([]LogWriterCreator{creator})
			if sut == nil {
				t.Error("didn't returned the expected reference")
			} else if (*sut)[0] != creator {
				t.Error("didn't stored the passed creator")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("nil config", func(t *testing.T) {
			src, e := NewLogWriterFactory(nil).Create(nil)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("unrecognized writer type", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := ConfigPartial{}
			creator := NewMockLogWriterCreator(ctrl)
			creator.EXPECT().Accept(&config).Return(false).Times(1)

			sut := NewLogWriterFactory([]LogWriterCreator{creator})

			writer, e := sut.Create(&config)
			switch {
			case writer != nil:
				t.Error("returned a config writer")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidLogConfig):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidLogConfig)
			}
		})

		t.Run("create the config writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := ConfigPartial{}
			writer := NewMockLogWriter(ctrl)
			creator := NewMockLogWriterCreator(ctrl)
			creator.EXPECT().Accept(&config).Return(true).Times(1)
			creator.EXPECT().Create(&config).Return(writer, nil).Times(1)

			sut := NewLogWriterFactory([]LogWriterCreator{creator})

			if s, e := sut.Create(&config); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if !reflect.DeepEqual(s, writer) {
				t.Error("didn't returned the created writer")
			}
		})
	})
}

func Test_Stream(t *testing.T) {
	t.Run("Close", func(t *testing.T) {
		t.Run("should call flush on closing the stream", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() { ctrl.Finish() }()

			writer := NewMockWriter(ctrl)
			gomock.InOrder(
				writer.EXPECT().Write([]byte("formatted 1\n")).Return(0, nil),
				writer.EXPECT().Write([]byte("formatted 2\n")).Return(0, nil),
				writer.EXPECT().Write([]byte("formatted 3\n")).Return(0, nil),
			)
			formatter := NewMockLogFormatter(ctrl)
			gomock.InOrder(
				formatter.EXPECT().Format(FATAL, "message 1").Return("formatted 1"),
				formatter.EXPECT().Format(FATAL, "message 2").Return("formatted 2"),
				formatter.EXPECT().Format(FATAL, "message 3").Return("formatted 3"),
			)

			sut := &LogStream{
				Formatter: formatter,
				Channels:  []string{},
				Level:     DEBUG,
				Writer:    writer,
			}

			_ = sut.Broadcast(FATAL, "message 1")
			_ = sut.Broadcast(FATAL, "message 2")
			_ = sut.Broadcast(FATAL, "message 3")

			if e := sut.Close(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("HasChannel", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut := &LogStream{
			Formatter: NewMockLogFormatter(ctrl),
			Channels:  []string{"channel.1", "channel.2"},
			Level:     WARNING,
			Writer:    nil,
		}

		t.Run("check the channel registration", func(t *testing.T) {
			switch {
			case !sut.HasChannel("channel.1"):
				t.Error("'channel.1' channel was not found")
			case !sut.HasChannel("channel.2"):
				t.Error("'channel.2' channel was not found")
			case sut.HasChannel("channel.3"):
				t.Error("'channel.3' channel was found")
			}
		})
	})

	t.Run("ListChannels", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		channels := []string{"channel.1", "channel.2"}
		sut := &LogStream{
			Formatter: NewMockLogFormatter(ctrl),
			Channels:  channels,
			Level:     WARNING,
			Writer:    nil,
		}

		t.Run("list the registered Channels", func(t *testing.T) {
			if check := sut.ListChannels(); !reflect.DeepEqual(check, channels) {
				t.Errorf("returned the (%v) list of Channels", check)
			}
		})
	})

	t.Run("AddChannel", func(t *testing.T) {
		t.Run("register a new channel", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				channel  string
				expected []string
			}{
				{ // adding into an empty list
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{},
						level:    DEBUG,
					},
					channel:  "channel.1",
					expected: []string{"channel.1"},
				},
				{ // adding should keep sorting
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"channel.1", "channel.3"},
						level:    DEBUG,
					},
					channel:  "channel.2",
					expected: []string{"channel.1", "channel.2", "channel.3"},
				},
				{ // adding an already existent should result in a no-op
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"channel.1", "channel.2", "channel.3"},
						level:    DEBUG,
					},
					channel:  "channel.2",
					expected: []string{"channel.1", "channel.2", "channel.3"},
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					sut := &LogStream{
						Formatter: NewMockLogFormatter(ctrl),
						Channels:  s.state.channels,
						Level:     s.state.level,
						Writer:    nil,
					}
					sut.AddChannel(s.channel)

					if check := sut.ListChannels(); !reflect.DeepEqual(check, s.expected) {
						t.Errorf("returned the (%v) list of Channels", check)
					}
				}
				test()
			}
		})
	})

	t.Run("RemoveChannel", func(t *testing.T) {
		t.Run("unregister a channel", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				channel  string
				expected []string
			}{
				{ // removing from an empty list
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{},
						level:    DEBUG,
					},
					channel:  "channel.1",
					expected: []string{},
				},
				{ // removing a non-existing channel
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"channel.1", "channel.3"},
						level:    DEBUG,
					},
					channel:  "channel.2",
					expected: []string{"channel.1", "channel.3"},
				},
				{ // removing an existing channel
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"channel.1", "channel.2", "channel.3"},
						level:    DEBUG,
					},
					channel:  "channel.2",
					expected: []string{"channel.1", "channel.3"},
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					sut := &LogStream{
						Formatter: NewMockLogFormatter(ctrl),
						Channels:  s.state.channels,
						Level:     s.state.level,
						Writer:    nil,
					}
					sut.RemoveChannel(s.channel)

					if check := sut.ListChannels(); !reflect.DeepEqual(check, s.expected) {
						t.Errorf("returned the (%v) list of Channels", check)
					}
				}
				test()
			}
		})
	})

	t.Run("Format", func(t *testing.T) {
		t.Run("return message if there is no formatter", func(t *testing.T) {
			msg := "message"
			level := WARNING
			sut := &LogStream{Formatter: nil,
				Channels: []string{},
				Level:    level,
				Writer:   nil,
			}

			if check := sut.Format(level, msg, LogContext{"field": "value"}); check != msg {
				t.Errorf("returned the (%v) formatted message", check)
			}
		})

		t.Run("return formatter response if formatter is present without context", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			msg := "message"
			expected := "formatted message"
			level := WARNING
			formatter := NewMockLogFormatter(ctrl)
			formatter.EXPECT().Format(level, msg).Return(expected).Times(1)
			sut := &LogStream{
				Formatter: formatter,
				Channels:  []string{},
				Level:     level,
				Writer:    nil,
			}

			if check := sut.Format(level, msg); check != expected {
				t.Errorf("(%v) formatted message", check)
			}
		})

		t.Run("return formatter response if formatter is present with context", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			msg := "message"
			ctx := LogContext{"field": "value"}
			expected := "formatted message"
			level := WARNING
			formatter := NewMockLogFormatter(ctrl)
			formatter.EXPECT().Format(level, msg, ctx).Return(expected).Times(1)
			sut := &LogStream{
				Formatter: formatter,
				Channels:  []string{},
				Level:     level,
				Writer:    nil,
			}

			if check := sut.Format(level, msg, ctx); check != expected {
				t.Errorf("(%v) formatted message", check)
			}
		})
	})

	t.Run("Signal", func(t *testing.T) {
		t.Run("buffer output messages", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() { ctrl.Finish() }()

			writer := NewMockWriter(ctrl)
			formatter := NewMockLogFormatter(ctrl)
			formatter.
				EXPECT().
				Format(FATAL, "message").
				Return("formatted").
				Times(1)

			sut := &LogStream{
				Formatter: formatter,
				Channels:  []string{"channel"},
				Level:     DEBUG,
				Writer:    writer,
			}

			if e := sut.Signal("channel", FATAL, "message"); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if len(sut.Buffer) == 0 {
				t.Error("didn't buffered the logging message")
			} else if sut.Buffer[0] != "formatted" {
				t.Error("buffered the wrong message")
			}
		})

		t.Run("signal message to the writer without context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					channel string
					message string
				}
				callTimes int
				expected  string
			}{
				{ // signal through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   FATAL,
						channel: "dummy_channel",
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   DEBUG,
						channel: "dummy_channel",
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with an unregistered channel
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   FATAL,
						channel: "not_a_valid_dummy_channel",
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message).
						Return(s.expected).
						Times(s.callTimes)

					sut := &LogStream{
						Formatter: formatter,
						Channels:  s.state.channels,
						Level:     s.state.level,
						Writer:    writer,
					}

					if e := sut.Signal(s.call.channel, s.call.level, s.call.message); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})

		t.Run("signal message to the writer with context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					channel string
					ctx     LogContext
					message string
				}
				callTimes int
				expected  string
			}{
				{ // signal through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   FATAL,
						channel: "dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   DEBUG,
						channel: "dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with an unregistered channel
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   FATAL,
						channel: "not_a_valid_dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message, s.call.ctx).
						Return(s.expected).
						Times(s.callTimes)

					sut := &LogStream{
						Formatter: formatter,
						Channels:  s.state.channels,
						Level:     s.state.level,
						Writer:    writer,
					}

					if e := sut.Signal(s.call.channel, s.call.level, s.call.message, s.call.ctx); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})
	})

	t.Run("Broadcast", func(t *testing.T) {
		t.Run("buffer output messages", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() { ctrl.Finish() }()

			writer := NewMockWriter(ctrl)
			formatter := NewMockLogFormatter(ctrl)
			formatter.
				EXPECT().
				Format(FATAL, "message").
				Return("formatted").
				Times(1)

			sut := &LogStream{
				Formatter: formatter,
				Channels:  []string{"channel"},
				Level:     DEBUG,
				Writer:    writer,
			}

			if e := sut.Broadcast(FATAL, "message"); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if len(sut.Buffer) == 0 {
				t.Error("didn't buffered the logging message")
			} else if sut.Buffer[0] != "formatted" {
				t.Error("buffered the wrong message")
			}
		})

		t.Run("broadcast message to the writer without context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					message string
				}
				callTimes int
				expected  string
			}{
				{ // broadcast through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						message string
					}{
						level:   FATAL,
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // broadcast through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						message string
					}{
						level:   DEBUG,
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message).
						Return(s.expected).
						Times(s.callTimes)

					sut := &LogStream{
						Formatter: formatter,
						Channels:  s.state.channels,
						Level:     s.state.level,
						Writer:    writer,
					}

					if e := sut.Broadcast(s.call.level, s.call.message); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})

		t.Run("broadcast message to the writer with context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					ctx     LogContext
					message string
				}
				callTimes int
				expected  string
			}{
				{ // broadcast through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						ctx     LogContext
						message string
					}{
						ctx:     LogContext{},
						level:   FATAL,
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // broadcast through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						ctx     LogContext
						message string
					}{
						level:   DEBUG,
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message, s.call.ctx).
						Return(s.expected).
						Times(s.callTimes)

					sut := &LogStream{
						Formatter: formatter,
						Channels:  s.state.channels,
						Level:     s.state.level,
						Writer:    writer,
					}

					if e := sut.Broadcast(s.call.level, s.call.message, s.call.ctx); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})
	})

	t.Run("Flush", func(t *testing.T) {
		t.Run("no-op if no buffered message", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() { ctrl.Finish() }()

			writer := NewMockWriter(ctrl)
			formatter := NewMockLogFormatter(ctrl)

			sut := &LogStream{
				Formatter: formatter,
				Channels:  []string{},
				Level:     DEBUG,
				Writer:    writer,
			}

			if e := sut.Flush(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("write buffered messages", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() { ctrl.Finish() }()

			writer := NewMockWriter(ctrl)
			gomock.InOrder(
				writer.EXPECT().Write([]byte("formatted 1\n")).Return(0, nil),
				writer.EXPECT().Write([]byte("formatted 2\n")).Return(0, nil),
				writer.EXPECT().Write([]byte("formatted 3\n")).Return(0, nil),
			)
			formatter := NewMockLogFormatter(ctrl)
			gomock.InOrder(
				formatter.EXPECT().Format(FATAL, "message 1").Return("formatted 1"),
				formatter.EXPECT().Format(FATAL, "message 2").Return("formatted 2"),
				formatter.EXPECT().Format(FATAL, "message 3").Return("formatted 3"),
			)

			sut := &LogStream{
				Formatter: formatter,
				Channels:  []string{},
				Level:     DEBUG,
				Writer:    writer,
			}

			_ = sut.Broadcast(FATAL, "message 1")
			_ = sut.Broadcast(FATAL, "message 2")
			_ = sut.Broadcast(FATAL, "message 3")

			if e := sut.Flush(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error while flushing", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer func() { ctrl.Finish() }()

			expected := fmt.Errorf("error message")
			writer := NewMockWriter(ctrl)
			writer.EXPECT().Write([]byte("formatted 1\n")).Return(0, expected)
			formatter := NewMockLogFormatter(ctrl)
			formatter.EXPECT().Format(FATAL, "message 1").Return("formatted 1")

			sut := &LogStream{
				Formatter: formatter,
				Channels:  []string{},
				Level:     DEBUG,
				Writer:    writer,
			}

			_ = sut.Broadcast(FATAL, "message 1")

			if e := sut.Flush(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})
	})
}

func Test_LogConsoleStream(t *testing.T) {
	t.Run("NewLogConsoleStream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		formatter := NewMockLogFormatter(ctrl)
		var channels []string
		level := WARNING

		t.Run("nil formatter", func(t *testing.T) {
			sut, e := NewLogConsoleStream(level, nil, channels)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new console stream", func(t *testing.T) {
			sut, e := NewLogConsoleStream(WARNING, formatter, []string{})
			switch {
			case sut == nil:
				t.Error("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case sut.Writer != os.Stdout:
				t.Error("didn't stored the stdout as the defined writer")
			}
		})
	})

	t.Run("Signal", func(t *testing.T) {
		t.Run("signal message to the writer without context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					channel string
					message string
				}
				callTimes int
				expected  string
			}{
				{ // signal through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   FATAL,
						channel: "dummy_channel",
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   DEBUG,
						channel: "dummy_channel",
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with an unregistered channel
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   FATAL,
						channel: "not_a_valid_dummy_channel",
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message).
						Return(s.expected).
						Times(s.callTimes)

					sut, _ := NewLogConsoleStream(s.state.level, formatter, s.state.channels)
					sut.Writer = writer

					if e := sut.Signal(s.call.channel, s.call.level, s.call.message); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})

		t.Run("signal message to the writer with context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					channel string
					ctx     LogContext
					message string
				}
				callTimes int
				expected  string
			}{
				{ // signal through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   FATAL,
						channel: "dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   DEBUG,
						channel: "dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with an unregistered channel
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   FATAL,
						channel: "not_a_valid_dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message, s.call.ctx).
						Return(s.expected).
						Times(s.callTimes)

					sut, _ := NewLogConsoleStream(s.state.level, formatter, s.state.channels)
					sut.Writer = writer

					if e := sut.Signal(s.call.channel, s.call.level, s.call.message, s.call.ctx); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})
	})

	t.Run("Broadcast", func(t *testing.T) {
		t.Run("broadcast message to the writer without context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					message string
				}
				callTimes int
				expected  string
			}{
				{ // broadcast through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						message string
					}{
						level:   FATAL,
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // broadcast through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						message string
					}{
						level:   DEBUG,
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message).
						Return(s.expected).
						Times(s.callTimes)

					sut, _ := NewLogConsoleStream(s.state.level, formatter, s.state.channels)
					sut.Writer = writer

					if e := sut.Broadcast(s.call.level, s.call.message); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})

		t.Run("broadcast message to the writer with context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					ctx     LogContext
					message string
				}
				callTimes int
				expected  string
			}{
				{ // broadcast through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						ctx     LogContext
						message string
					}{
						ctx:     LogContext{},
						level:   FATAL,
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // broadcast through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						ctx     LogContext
						message string
					}{
						level:   DEBUG,
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)
					defer func() { ctrl.Finish() }()

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message, s.call.ctx).
						Return(s.expected).
						Times(s.callTimes)

					sut, _ := NewLogConsoleStream(s.state.level, formatter, s.state.channels)
					sut.Writer = writer

					if e := sut.Broadcast(s.call.level, s.call.message, s.call.ctx); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})
	})
}

func Test_LogConsoleStreamCreator(t *testing.T) {
	t.Run("NewLogConsoleStreamCreator", func(t *testing.T) {
		t.Run("nil formatter factory", func(t *testing.T) {
			sut, e := NewLogConsoleStreamCreator(nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new console stream factory creator", func(t *testing.T) {
			if sut, e := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil)); sut == nil {
				t.Errorf("didn't returned a valid reference")
			} else if e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
			sut, _ := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil))

			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept on type retrieval error", func(t *testing.T) {
			partial := ConfigPartial{"type": ConfigPartial{}}
			sut, _ := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil))

			if sut.Accept(&partial) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept on invalid type", func(t *testing.T) {
			partial := ConfigPartial{"type": "invalid type"}
			sut, _ := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil))

			if sut.Accept(&partial) {
				t.Error("returned true")
			}
		})

		t.Run("accept on valid type", func(t *testing.T) {
			partial := ConfigPartial{"type": LogTypeConsole}
			sut, _ := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil))

			if !sut.Accept(&partial) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			sut, _ := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil))

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

		t.Run("non-string format", func(t *testing.T) {
			partial := ConfigPartial{
				"type":   LogTypeConsole,
				"format": 123,
			}
			sut, _ := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil))

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-log level name level", func(t *testing.T) {
			partial := ConfigPartial{
				"type":   LogTypeConsole,
				"format": LogFormatJSON,
				"level":  "invalid",
			}
			sut, _ := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil))

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidLogLevel):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidLogLevel)
			}
		})

		t.Run("error creating the formatter", func(t *testing.T) {
			partial := ConfigPartial{
				"type":   LogTypeConsole,
				"format": LogFormatJSON,
				"level":  "fatal",
			}
			sut, _ := NewLogConsoleStreamCreator(NewLogFormatterFactory(nil))

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidLogFormat):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidLogFormat)
			}
		})

		t.Run("new stream", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":     LogTypeConsole,
				"format":   "format",
				"level":    "fatal",
				"channels": []interface{}{"channel1", "channel2"}}
			formatter := NewMockLogFormatter(ctrl)
			formatterCreator := NewMockLogFormatterCreator(ctrl)
			formatterCreator.EXPECT().Accept("format").Return(true).Times(1)
			formatterCreator.EXPECT().Create(gomock.Any()).Return(formatter, nil).Times(1)
			formatterFactory := NewLogFormatterFactory([]LogFormatterCreator{formatterCreator})
			sut, _ := NewLogConsoleStreamCreator(formatterFactory)

			stream, e := sut.Create(&partial)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case stream == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch s := stream.(type) {
				case *LogConsoleStream:
					switch {
					case s.Level != FATAL:
						t.Errorf("invalid level (%s)", LogLevelMapName[s.Level])
					case len(s.Channels) != 2:
						t.Errorf("invalid channel list (%v)", s.Channels)
					}
				default:
					t.Error("didn't returned a new console stream")
				}
			}
		})
	})
}

func Test_LogFileStream(t *testing.T) {
	t.Run("NewLogFileStream", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		writer := NewMockWriter(ctrl)
		writer.EXPECT().Close().Times(1)
		formatter := NewMockLogFormatter(ctrl)
		var channels []string
		level := WARNING

		t.Run("nil writer", func(t *testing.T) {
			sut, e := NewLogFileStream(level, formatter, channels, nil)
			switch {
			case sut != nil:
				_ = sut.Close()
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil formatter", func(t *testing.T) {
			sut, e := NewLogFileStream(level, nil, channels, writer)
			switch {
			case sut != nil:
				_ = sut.Close()
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new file stream", func(t *testing.T) {
			sut, e := NewLogFileStream(WARNING, formatter, []string{}, writer)
			switch {
			case sut == nil:
				t.Error("didn't returned a valid reference")
			default:
				_ = sut.Close()
				if e != nil {
					t.Errorf("unexpected (%v) error", e)
				}
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("call the close on the writer only once", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			writer := NewMockWriter(ctrl)
			writer.EXPECT().Close().Times(1)
			sut, _ := NewLogFileStream(WARNING, NewMockLogFormatter(ctrl), []string{}, writer)

			_ = sut.Close()
			_ = sut.Close()
		})
	})

	t.Run("Signal", func(t *testing.T) {
		t.Run("signal message to the writer with context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					channel string
					message string
				}
				callTimes int
				expected  string
			}{
				{ // signal through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   FATAL,
						channel: "dummy_channel",
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   DEBUG,
						channel: "dummy_channel",
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with an unregistered channel
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						message string
					}{
						level:   FATAL,
						channel: "not_a_valid_dummy_channel",
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Close().Times(1)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message).
						Return(s.expected).
						Times(s.callTimes)

					sut, _ := NewLogFileStream(s.state.level, formatter, s.state.channels, writer)
					defer func() {
						_ = sut.Close()
						ctrl.Finish()
					}()

					if e := sut.Signal(s.call.channel, s.call.level, s.call.message); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})

		t.Run("signal message to the writer with context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					channel string
					ctx     LogContext
					message string
				}
				callTimes int
				expected  string
			}{
				{ // signal through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   FATAL,
						channel: "dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   DEBUG,
						channel: "dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // signal through a valid channel with an unregistered channel
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						channel string
						ctx     LogContext
						message string
					}{
						level:   FATAL,
						channel: "not_a_valid_dummy_channel",
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Close().Times(1)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message, s.call.ctx).
						Return(s.expected).
						Times(s.callTimes)

					sut, _ := NewLogFileStream(s.state.level, formatter, s.state.channels, writer)
					defer func() {
						_ = sut.Close()
						ctrl.Finish()
					}()

					if e := sut.Signal(s.call.channel, s.call.level, s.call.message, s.call.ctx); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})
	})

	t.Run("Broadcast", func(t *testing.T) {
		t.Run("broadcast message to the writer without context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					message string
				}
				callTimes int
				expected  string
			}{
				{ // broadcast through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						message string
					}{
						level:   FATAL,
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // broadcast through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						message string
					}{
						level:   DEBUG,
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Close().Times(1)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message).
						Return(s.expected).
						Times(s.callTimes)

					sut, _ := NewLogFileStream(s.state.level, formatter, s.state.channels, writer)
					defer func() {
						_ = sut.Close()
						ctrl.Finish()
					}()

					if e := sut.Broadcast(s.call.level, s.call.message); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})

		t.Run("broadcast message to the writer with context", func(t *testing.T) {
			scenarios := []struct {
				state struct {
					channels []string
					level    LogLevel
				}
				call struct {
					level   LogLevel
					ctx     LogContext
					message string
				}
				callTimes int
				expected  string
			}{
				{ // broadcast through a valid channel with a not filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						ctx     LogContext
						message string
					}{
						level:   FATAL,
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 1,
					expected:  `{"message" : "dummy_message"}`,
				},
				{ // broadcast through a valid channel with a filtered level
					state: struct {
						channels []string
						level    LogLevel
					}{
						channels: []string{"dummy_channel"},
						level:    WARNING,
					},
					call: struct {
						level   LogLevel
						ctx     LogContext
						message string
					}{
						level:   DEBUG,
						ctx:     LogContext{},
						message: "dummy_message",
					},
					callTimes: 0,
					expected:  `{"message" : "dummy_message"}`,
				},
			}

			for _, s := range scenarios {
				test := func() {
					ctrl := gomock.NewController(t)

					writer := NewMockWriter(ctrl)
					writer.EXPECT().Close().Times(1)
					writer.EXPECT().Write([]byte(s.expected + "\n")).Times(s.callTimes)
					formatter := NewMockLogFormatter(ctrl)
					formatter.
						EXPECT().
						Format(s.call.level, s.call.message, s.call.ctx).
						Return(s.expected).
						Times(s.callTimes)

					sut, _ := NewLogFileStream(s.state.level, formatter, s.state.channels, writer)
					defer func() {
						_ = sut.Close()
						ctrl.Finish()
					}()

					if e := sut.Broadcast(s.call.level, s.call.message, s.call.ctx); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
					if e := sut.Flush(); e != nil {
						t.Errorf("unexpected (%v) error", e)
					}
				}
				test()
			}
		})
	})
}

func Test_LogFileStreamCreator(t *testing.T) {
	t.Run("NewLogFileStreamCreator", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			sut, e := NewLogFileStreamCreator(nil, NewLogFormatterFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil formatter factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewLogFileStreamCreator(NewMockFs(ctrl), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new file stream factory creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			if sut, e := NewLogFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil)); sut == nil {
				t.Errorf("didn't returned a valid reference")
			} else if e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewLogFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

		t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept on type retrieval error", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept on invalid type", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": "invalid"}) {
				t.Error("returned true")
			}
		})

		t.Run("accept on valid type", func(t *testing.T) {
			if !sut.Accept(&ConfigPartial{"type": LogTypeFile}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewLogFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

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

		t.Run("non-string format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":   LogTypeFile,
				"format": 123,
			}
			sut, _ := NewLogFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

			src, e := sut.Create(&partial)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-log level name level", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":   LogTypeFile,
				"format": "format",
				"level":  "invalid",
			}
			sut, _ := NewLogFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidLogLevel):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidLogLevel)
			}
		})

		t.Run("error on creating the formatter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":   LogTypeFile,
				"format": "format",
				"level":  "fatal",
			}
			sut, _ := NewLogFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				_ = stream.(io.Closer).Close()
				t.Error("returned a valid stream")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidLogFormat):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidLogFormat)
			}
		})

		t.Run("error on opening the file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			partial := ConfigPartial{
				"type":     LogTypeFile,
				"format":   "format",
				"level":    "fatal",
				"path":     "path",
				"channels": []interface{}{"channel1", "channel2"}}
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(nil, expected).
				Times(1)
			formatter := NewMockLogFormatter(ctrl)
			formatterCreator := NewMockLogFormatterCreator(ctrl)
			formatterCreator.EXPECT().Accept("format").Return(true).Times(1)
			formatterCreator.EXPECT().Create(gomock.Any()).Return(formatter, nil).Times(1)
			formatterFactory := NewLogFormatterFactory([]LogFormatterCreator{formatterCreator})

			sut, _ := NewLogFileStreamCreator(fileSystem, formatterFactory)

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				_ = stream.(io.Closer).Close()
				t.Error("returned a valid stream")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("new stream", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":     LogTypeFile,
				"format":   "format",
				"level":    "fatal",
				"path":     "path",
				"channels": []interface{}{"channel1", "channel2"}}
			file := NewMockFile(ctrl)
			file.EXPECT().Close().Return(nil).Times(1)
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile("path", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)
			formatter := NewMockLogFormatter(ctrl)
			formatterCreator := NewMockLogFormatterCreator(ctrl)
			formatterCreator.EXPECT().Accept("format").Return(true).Times(1)
			formatterCreator.EXPECT().Create(gomock.Any()).Return(formatter, nil).Times(1)
			formatterFactory := NewLogFormatterFactory([]LogFormatterCreator{formatterCreator})

			sut, _ := NewLogFileStreamCreator(fileSystem, formatterFactory)

			stream, e := sut.Create(&partial)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case stream == nil:
				t.Error("didn't returned a valid reference")
			default:
				_ = stream.(io.Closer).Close()
				switch s := stream.(type) {
				case *LogFileStream:
					switch {
					case s.Level != FATAL:
						t.Errorf("invalid level (%s)", LogLevelMapName[s.Level])
					case len(s.Channels) != 2:
						t.Errorf("invalid channel list (%v)", s.Channels)
					}
				default:
					t.Error("didn't returned a new file stream")
				}
			}
		})
	})
}

func Test_LogRotatingFileWriter(t *testing.T) {
	t.Run("NewLogRotatingFileWriter", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			sut, e := NewLogRotatingFileWriter(nil, "path")
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error opening file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(nil, expected).
				Times(1)

			sut, e := NewLogRotatingFileWriter(fileSystem, path)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("create valid writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			file := NewMockFile(ctrl)
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)

			if sut, e := NewLogRotatingFileWriter(fileSystem, path); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if sut == nil {
				t.Error("didn't returned the expected writer reference")
			}
		})
	})

	t.Run("Write", func(t *testing.T) {
		t.Run("error while writing", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			output := []byte("message")
			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			file := NewMockFile(ctrl)
			file.EXPECT().Write(output).Return(0, expected).Times(1)
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)

			sut, _ := NewLogRotatingFileWriter(fileSystem, path)

			if _, e := sut.Write(output); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("no rotation if write done in same day of opened file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			output := []byte("message")
			count := 123
			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			file := NewMockFile(ctrl)
			file.EXPECT().Write(output).Return(count, nil).Times(1)
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)

			sut, _ := NewLogRotatingFileWriter(fileSystem, path)

			if written, e := sut.Write(output); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if written != count {
				t.Errorf("unexpected number of written elements of (%v) when expecting (%v)", written, count)
			}
		})

		t.Run("error while closing rotated file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			output := []byte("message")
			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			file := NewMockFile(ctrl)
			file.EXPECT().Close().Return(expected)
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)

			sut, _ := NewLogRotatingFileWriter(fileSystem, path)
			sut.(*LogRotatingFileWriter).day++

			if _, e := sut.Write(output); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error while opening rotating file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			output := []byte("message")
			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			file := NewMockFile(ctrl)
			file.EXPECT().Close().Return(nil)
			fileSystem := NewMockFs(ctrl)
			gomock.InOrder(
				fileSystem.
					EXPECT().
					OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
					Return(file, nil),
				fileSystem.
					EXPECT().
					OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
					Return(nil, expected),
			)

			sut, _ := NewLogRotatingFileWriter(fileSystem, path)
			sut.(*LogRotatingFileWriter).day++

			if _, e := sut.Write(output); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("rotate file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			output := []byte("message")
			count := 123
			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			file1 := NewMockFile(ctrl)
			file1.EXPECT().Close().Return(nil)
			file2 := NewMockFile(ctrl)
			file2.EXPECT().Write(output).Return(count, nil).Times(1)
			fileSystem := NewMockFs(ctrl)
			gomock.InOrder(
				fileSystem.
					EXPECT().
					OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
					Return(file1, nil),
				fileSystem.
					EXPECT().
					OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
					Return(file2, nil),
			)

			sut, _ := NewLogRotatingFileWriter(fileSystem, path)
			sut.(*LogRotatingFileWriter).day++

			if written, e := sut.Write(output); e != nil {
				t.Errorf("unexpected (%v) error", e)
			} else if written != count {
				t.Errorf("unexpected number of written elements of (%v) when expecting (%v)", written, count)
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("call close on the underlying file pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			file := NewMockFile(ctrl)
			file.EXPECT().Close().Return(expected)
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)

			sut, _ := NewLogRotatingFileWriter(fileSystem, path)

			if e := sut.(io.Closer).Close(); e == nil {
				t.Error("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})
	})
}

func Test_NewRotatingFileStreamStrategy(t *testing.T) {
	t.Run("NewLogRotatingFileStreamCreator", func(t *testing.T) {
		t.Run("nil file system adapter", func(t *testing.T) {
			sut, e := NewLogRotatingFileStreamCreator(nil, NewLogFormatterFactory(nil))
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("nil formatter factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewLogRotatingFileStreamCreator(NewMockFs(ctrl), nil)
			switch {
			case sut != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("new file stream factory creator", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewLogRotatingFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))
			switch {
			case sut == nil:
				t.Errorf("didn't returned a valid reference")
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("Accept", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		sut, _ := NewLogRotatingFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

		t.Run("don't accept if config is a nil pointer", func(t *testing.T) {
			if sut.Accept(nil) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept on type retrieval error", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": 123}) {
				t.Error("returned true")
			}
		})

		t.Run("don't accept on invalid type", func(t *testing.T) {
			if sut.Accept(&ConfigPartial{"type": "invalid"}) {
				t.Error("returned true")
			}
		})

		t.Run("accept on valid type", func(t *testing.T) {
			if !sut.Accept(&ConfigPartial{"type": LogTypeRotatingFile}) {
				t.Error("returned false")
			}
		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("error on nil config pointer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, _ := NewLogRotatingFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

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

		t.Run("non-string format", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":   LogTypeRotatingFile,
				"format": 123,
			}
			sut, _ := NewLogRotatingFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

			src, e := sut.Create(&partial)
			switch {
			case src != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrConversion):
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("non-log level name level", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":   LogTypeRotatingFile,
				"format": "format",
				"level":  "invalid",
			}
			sut, _ := NewLogRotatingFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				t.Error("returned a valid reference")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidLogLevel):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidLogLevel)
			}
		})

		t.Run("error on creating the formatter", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{
				"type":   LogTypeRotatingFile,
				"format": "format",
				"level":  "fatal",
			}
			sut, _ := NewLogRotatingFileStreamCreator(NewMockFs(ctrl), NewLogFormatterFactory(nil))

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				_ = stream.(io.Closer).Close()
				t.Error("returned a valid stream")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrInvalidLogFormat):
				t.Errorf("(%v) when expecting (%v)", e, ErrInvalidLogFormat)
			}
		})

		t.Run("error on opening the file", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			expected := fmt.Errorf("error message")
			partial := ConfigPartial{
				"type":     LogTypeRotatingFile,
				"format":   "format",
				"level":    "fatal",
				"path":     path,
				"channels": []interface{}{"channel1", "channel2"}}
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(nil, expected).
				Times(1)
			formatter := NewMockLogFormatter(ctrl)
			formatterCreator := NewMockLogFormatterCreator(ctrl)
			formatterCreator.EXPECT().Accept("format").Return(true).Times(1)
			formatterCreator.EXPECT().Create(gomock.Any()).Return(formatter, nil).Times(1)
			formatterFactory := NewLogFormatterFactory([]LogFormatterCreator{formatterCreator})

			sut, _ := NewLogRotatingFileStreamCreator(fileSystem, formatterFactory)

			stream, e := sut.Create(&partial)
			switch {
			case stream != nil:
				_ = stream.(io.Closer).Close()
				t.Error("returned a valid stream")
			case e == nil:
				t.Error("didn't returned the expected error")
			case e.Error() != expected.Error():
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("new stream", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			path := "path-%s"
			expectedPath := fmt.Sprintf(path, time.Now().Format("2006-01-02"))
			partial := ConfigPartial{
				"type":     LogTypeRotatingFile,
				"format":   "format",
				"level":    "fatal",
				"path":     path,
				"channels": []interface{}{"channel1", "channel2"}}
			file := NewMockFile(ctrl)
			file.EXPECT().Close().Return(nil).Times(1)
			fileSystem := NewMockFs(ctrl)
			fileSystem.
				EXPECT().
				OpenFile(expectedPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0o644)).
				Return(file, nil).
				Times(1)
			formatter := NewMockLogFormatter(ctrl)
			formatterCreator := NewMockLogFormatterCreator(ctrl)
			formatterCreator.EXPECT().Accept("format").Return(true).Times(1)
			formatterCreator.EXPECT().Create(gomock.Any()).Return(formatter, nil).Times(1)
			formatterFactory := NewLogFormatterFactory([]LogFormatterCreator{formatterCreator})

			sut, _ := NewLogRotatingFileStreamCreator(fileSystem, formatterFactory)

			stream, e := sut.Create(&partial)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case stream == nil:
				t.Error("didn't returned a valid reference")
			default:
				_ = stream.(io.Closer).Close()
				switch s := stream.(type) {
				case *LogFileStream:
					switch {
					case s.Level != FATAL:
						t.Errorf("invalid level (%s)", LogLevelMapName[s.Level])
					case len(s.Channels) != 2:
						t.Errorf("invalid channel list (%v)", s.Channels)
					}
				default:
					t.Error("didn't returned a new file stream")
				}
			}
		})
	})
}

func Test_Log(t *testing.T) {
	t.Run("NewLog", func(t *testing.T) {
		t.Run("new Log", func(t *testing.T) {
			if sut := NewLog(); sut == nil {
				t.Error("didn't returned a valid reference")
			}
		})
	})

	t.Run("Close", func(t *testing.T) {
		t.Run("execute close process", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id1 := "writer.1"
			id2 := "writer.2"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Close().Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Close().Times(1)
			sut := NewLog()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)
			_ = sut.Close()

			if sut.HasWriter(id1) {
				t.Error("didn't removed the writer")
			}
			if sut.HasWriter(id2) {
				t.Error("didn't removed the writer")
			}
		})
	})

	t.Run("Signal", func(t *testing.T) {
		t.Run("propagate to all writers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			channel := "channel"
			level := WARNING
			message := "message"
			ctx := LogContext{"field": "value"}
			id1 := "writer.1"
			id2 := "writer.2"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Signal(channel, level, message, ctx).Return(nil).Times(1)
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Signal(channel, level, message, ctx).Return(nil).Times(1)
			writer2.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)

			if e := sut.Signal(channel, level, message, ctx); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("return on the first error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			channel := "channel"
			level := WARNING
			message := "message"
			ctx := LogContext{"field": "value"}
			expected := fmt.Errorf("error message")
			id1 := "writer.1"
			id2 := "writer.2"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Signal(channel, level, message, ctx).Return(expected).AnyTimes()
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Signal(channel, level, message, ctx).Return(nil).AnyTimes()
			writer2.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)

			if e := sut.Signal(channel, level, message, ctx); e == nil {
				t.Error("didn't returned the expected  error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})
	})

	t.Run("Broadcast", func(t *testing.T) {
		t.Run("propagate to all writers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			level := WARNING
			message := "message"
			ctx := LogContext{"field": "value"}
			id1 := "writer.1"
			id2 := "writer.2"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Broadcast(level, message, ctx).Return(nil).Times(1)
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Broadcast(level, message, ctx).Return(nil).Times(1)
			writer2.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)

			if e := sut.Broadcast(level, message, ctx); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("return on the first error", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			level := WARNING
			ctx := LogContext{"field": "value"}
			message := "message"
			expected := fmt.Errorf("error")
			id1 := "writer.1"
			id2 := "writer.2"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Broadcast(level, message, ctx).Return(expected).AnyTimes()
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Broadcast(level, message, ctx).Return(nil).AnyTimes()
			writer2.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)

			if e := sut.Broadcast(level, message, ctx); e == nil {
				t.Error("didn't returned the expected  error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})
	})

	t.Run("Flush", func(t *testing.T) {
		t.Run("error on flushing", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			id := "writer.1"
			writer := NewMockLogWriter(ctrl)
			writer.EXPECT().Flush().Return(expected).Times(1)
			writer.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id, writer)

			if e := sut.Flush(); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !reflect.DeepEqual(e, expected) {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("propagate to all writers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id1 := "writer.1"
			id2 := "writer.2"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Flush().Return(nil).Times(1)
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Flush().Return(nil).Times(1)
			writer2.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)

			_ = sut.Flush()
		})
	})

	t.Run("HasWriter", func(t *testing.T) {
		t.Run("check the registration of a writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id1 := "writer.1"
			id2 := "writer.2"
			id3 := "writer.3"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)

			if !sut.HasWriter(id1) {
				t.Errorf("returned false")
			}
			if !sut.HasWriter(id2) {
				t.Errorf("returned false")
			}
			if sut.HasWriter(id3) {
				t.Errorf("returned true")
			}
		})
	})

	t.Run("ListWriters", func(t *testing.T) {
		t.Run("retrieve the list of registered writers id's", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id1 := "writer.1"
			id2 := "writer.2"
			id3 := "writer.3"
			expected := []string{id1, id2, id3}
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Close().Return(nil).Times(1)
			writer3 := NewMockLogWriter(ctrl)
			writer3.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)
			_ = sut.AddWriter(id3, writer3)

			writers := sut.ListWriters()
			sort.Slice(writers, func(i, j int) bool {
				return writers[i] <= writers[j]
			})

			if sort.Search(len(writers), func(i int) bool {
				return writers[i] >= "id1"
			}) >= len(writers) {
				t.Errorf("{%v} when expecting {%v}", writers, expected)
			}
			if sort.Search(len(writers), func(i int) bool {
				return writers[i] >= "id2"
			}) >= len(writers) {
				t.Errorf("{%v} when expecting {%v}", writers, expected)
			}
			if sort.Search(len(writers), func(i int) bool {
				return writers[i] >= "id3"
			}) >= len(writers) {
				t.Errorf("{%v} when expecting {%v}", writers, expected)
			}
		})
	})

	t.Run("AddWriter", func(t *testing.T) {
		t.Run("error if nil writer", func(t *testing.T) {
			sut := NewLog()
			defer func() { _ = sut.Close() }()

			if e := sut.AddWriter("id", nil); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error if id is duplicate", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "writer"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id, writer1)

			if e := sut.AddWriter(id, writer2); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !errors.Is(e, ErrDuplicateLogWriter) {
				t.Errorf("(%v) when expecting (%v)", e, ErrDuplicateLogWriter)
			}
		})

		t.Run("register a new writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "writer"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()

			if e := sut.AddWriter(id, writer1); e != nil {
				t.Errorf("resulted the (%v) error", e)
			} else if check, e := sut.Writer(id); !reflect.DeepEqual(check, writer1) {
				t.Errorf("didn't stored the writer")
			} else if e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("RemoveWriter", func(t *testing.T) {
		t.Run("unregister a writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "writer"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id, writer1)
			sut.RemoveWriter(id)

			if sut.HasWriter(id) {
				t.Errorf("dnd't removed the writer")
			}
		})
	})

	t.Run("RemoveAllWriters", func(t *testing.T) {
		t.Run("remove all registered writers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id1 := "writer.1"
			id2 := "writer.2"
			id3 := "writer.3"
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writer2.EXPECT().Close().Return(nil).Times(1)
			writer3 := NewMockLogWriter(ctrl)
			writer3.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id1, writer1)
			_ = sut.AddWriter(id2, writer2)
			_ = sut.AddWriter(id3, writer3)
			sut.RemoveAllWriters()

			if check := sut.ListWriters(); len(check) != 0 {
				t.Errorf("{%v} id's list instead of an empty list", check)
			}
		})
	})

	t.Run("Writer", func(t *testing.T) {
		t.Run("non-existing writer", func(t *testing.T) {
			sut := NewLog()
			defer func() { _ = sut.Close() }()

			result, e := sut.Writer("invalid id")
			switch {
			case result != nil:
				t.Errorf("returned a valid writer")
			case e == nil:
				t.Error("didn't returned the expected error")
			case !errors.Is(e, ErrLogWriterNotFound):
				t.Errorf("(%v) when expecting (%v)", e, ErrLogWriterNotFound)
			}
		})

		t.Run("retrieve the requested writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := "writer"
			writer := NewMockLogWriter(ctrl)
			writer.EXPECT().Close().Return(nil).Times(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id, writer)

			if check, e := sut.Writer(id); !reflect.DeepEqual(check, writer) {
				t.Errorf("didn0t retrieved the stored writer")
			} else if e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("running", func(t *testing.T) {
		t.Run("no flusher if zero frequency", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			prev := LogFlushFrequency
			LogFlushFrequency = 0
			defer func() { LogFlushFrequency = prev }()

			id := "writer"
			writer := NewMockLogWriter(ctrl)
			writer.EXPECT().Close().Return(nil).Times(1)
			writer.EXPECT().Broadcast(FATAL, "message").Return(nil).Times(1)
			writer.EXPECT().Flush().Return(nil).Times(0)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id, writer)

			_ = sut.Broadcast(FATAL, "message")

			time.Sleep(20 * time.Millisecond)
		})

		t.Run("flusher run all writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			prev := LogFlushFrequency
			LogFlushFrequency = 10
			defer func() { LogFlushFrequency = prev }()

			id := "writer"
			writer := NewMockLogWriter(ctrl)
			writer.EXPECT().Close().Return(nil).Times(1)
			writer.EXPECT().Broadcast(FATAL, "message").Return(nil).Times(1)
			writer.EXPECT().Flush().Return(nil).MinTimes(1)
			sut := NewLog()
			defer func() { _ = sut.Close() }()
			_ = sut.AddWriter(id, writer)

			_ = sut.Broadcast(FATAL, "message")

			time.Sleep(30 * time.Millisecond)
		})
	})
}

func Test_LogLoader(t *testing.T) {
	t.Run("NewLogLoader", func(t *testing.T) {
		t.Run("error when missing the config", func(t *testing.T) {
			sut, e := NewLogLoader(nil, NewLog(), NewLogWriterFactory(nil))
			switch {
			case sut != nil:
				t.Errorf("return a valid reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error when missing the logger", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewLogLoader(NewConfig(), nil, NewLogWriterFactory(nil))
			switch {
			case sut != nil:
				t.Errorf("return a valid reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("error when missing the logger writer factory", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sut, e := NewLogLoader(NewConfig(), NewLog(), nil)
			switch {
			case sut != nil:
				t.Errorf("return a valid reference")
			case e == nil:
				t.Errorf("didn't returned the expected error")
			case !errors.Is(e, ErrNilPointer):
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("create loader", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			if sut, e := NewLogLoader(NewConfig(), NewLog(), NewLogWriterFactory(nil)); sut == nil {
				t.Errorf("didn't returned a valid reference")
			} else if e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})

	t.Run("Load", func(t *testing.T) {
		t.Run("error retrieving writers list", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{}
			_, _ = partial.Set("slate.log.writers", "string")
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier", 0, supplier)

			sut, _ := NewLogLoader(config, NewLog(), NewLogWriterFactory(nil))

			if e := sut.Load(); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("no-op if writers list is empty", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{}
			_, _ = partial.Set("slate.log.writers", ConfigPartial{})
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier", 0, supplier)

			sut, _ := NewLogLoader(config, NewLog(), NewLogWriterFactory(nil))

			if e := sut.Load(); e != nil {
				t.Errorf("returned the (%s) error", e)
			}
		})

		t.Run("request config path from global variable", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			prev := LogLoaderConfigPath
			LogLoaderConfigPath = "path"
			defer func() { LogLoaderConfigPath = prev }()

			partial := ConfigPartial{}
			_, _ = partial.Set("path", ConfigPartial{})
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier", 0, supplier)

			sut, _ := NewLogLoader(config, NewLog(), NewLogWriterFactory(nil))

			if e := sut.Load(); e != nil {
				t.Errorf("returned the (%s) error", e)
			}
		})

		t.Run("error getting writer information", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{}
			_, _ = partial.Set("slate.log.writers", ConfigPartial{"id": 1})
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier", 0, supplier)

			sut, _ := NewLogLoader(config, NewLog(), NewLogWriterFactory(nil))

			if e := sut.Load(); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("error creating writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			expected := fmt.Errorf("error message")
			partial := ConfigPartial{}
			_, _ = partial.Set("slate.log.writers", ConfigPartial{"id": ConfigPartial{}})
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier", 0, supplier)
			writerCreator := NewMockLogWriterCreator(ctrl)
			writerCreator.EXPECT().Accept(&ConfigPartial{}).Return(true).Times(1)
			writerCreator.EXPECT().Create(&ConfigPartial{}).Return(nil, expected).Times(1)
			writerFactory := NewLogWriterFactory([]LogWriterCreator{writerCreator})

			sut, _ := NewLogLoader(config, NewLog(), writerFactory)

			if e := sut.Load(); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if e.Error() != expected.Error() {
				t.Errorf("(%v) when expecting (%v)", e, expected)
			}
		})

		t.Run("error storing writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{}
			_, _ = partial.Set("slate.log.writers", ConfigPartial{"id": ConfigPartial{}})
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier", 0, supplier)
			writer := NewMockLogWriter(ctrl)
			writerCreator := NewMockLogWriterCreator(ctrl)
			writerCreator.EXPECT().Accept(&ConfigPartial{}).Return(true).Times(1)
			writerCreator.EXPECT().Create(&ConfigPartial{}).Return(writer, nil).Times(1)
			writerFactory := NewLogWriterFactory([]LogWriterCreator{writerCreator})
			log := NewLog()
			_ = log.AddWriter("id", writer)

			sut, _ := NewLogLoader(config, log, writerFactory)

			if e := sut.Load(); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !errors.Is(e, ErrDuplicateLogWriter) {
				t.Errorf("(%v) when expecting (%v)", e, ErrDuplicateLogWriter)
			}
		})

		t.Run("register writer", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			partial := ConfigPartial{}
			_, _ = partial.Set("slate.log.writers", ConfigPartial{"id": ConfigPartial{}})
			supplier := NewMockConfigSupplier(ctrl)
			supplier.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier", 0, supplier)
			writer := NewMockLogWriter(ctrl)
			writerCreator := NewMockLogWriterCreator(ctrl)
			writerCreator.EXPECT().Accept(&ConfigPartial{}).Return(true).Times(1)
			writerCreator.EXPECT().Create(&ConfigPartial{}).Return(writer, nil).Times(1)
			writerFactory := NewLogWriterFactory([]LogWriterCreator{writerCreator})

			sut, _ := NewLogLoader(config, NewLog(), writerFactory)

			if e := sut.Load(); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error on creating the reconfigured logger writers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config1 := ConfigPartial{
				"type":     "console",
				"format":   "json",
				"channels": []interface{}{},
				"Level":    "fatal",
			}
			partial1 := ConfigPartial{}
			_, _ = partial1.Set("slate.log.writers.id", config1)
			partial2 := ConfigPartial{}
			_, _ = partial2.Set("slate.log.writers", "string")
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(partial1, nil).Times(2)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(partial2, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier.1", 1, supplier1)
			writer := NewMockLogWriter(ctrl)
			writerCreator := NewMockLogWriterCreator(ctrl)
			writerCreator.EXPECT().Accept(&config1).Return(true).Times(1)
			writerCreator.EXPECT().Create(&config1).Return(writer, nil).Times(1)
			writerFactory := NewLogWriterFactory([]LogWriterCreator{writerCreator})

			sut, _ := NewLogLoader(config, NewLog(), writerFactory)
			_ = sut.Load()

			_ = config.AddSupplier("supplier.2", 100, supplier2)
		})

		t.Run("reconfigured logger writers", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config1 := ConfigPartial{
				"type":     "console",
				"format":   "json",
				"Channels": []interface{}{},
				"Level":    "fatal",
			}
			config2 := ConfigPartial{
				"type":     "console",
				"format":   "json",
				"Channels": []interface{}{},
				"Level":    "debug",
			}
			partial1 := ConfigPartial{}
			_, _ = partial1.Set("slate.log.writers.id", config1)
			partial2 := ConfigPartial{}
			_, _ = partial2.Set("slate.log.writers.id", config2)
			supplier1 := NewMockConfigSupplier(ctrl)
			supplier1.EXPECT().Get("").Return(partial1, nil).Times(2)
			supplier2 := NewMockConfigSupplier(ctrl)
			supplier2.EXPECT().Get("").Return(partial2, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("supplier.1", 1, supplier1)
			writer1 := NewMockLogWriter(ctrl)
			writer1.EXPECT().Close().Return(nil).Times(1)
			writer2 := NewMockLogWriter(ctrl)
			writerCreator := NewMockLogWriterCreator(ctrl)
			gomock.InOrder(
				writerCreator.EXPECT().Accept(&config1).Return(true),
				writerCreator.EXPECT().Accept(&config2).Return(true),
			)
			gomock.InOrder(
				writerCreator.EXPECT().Create(&config1).Return(writer1, nil),
				writerCreator.EXPECT().Create(&config2).Return(writer2, nil),
			)
			writerFactory := NewLogWriterFactory([]LogWriterCreator{writerCreator})

			sut, _ := NewLogLoader(config, NewLog(), writerFactory)
			_ = sut.Load()

			_ = config.AddSupplier("supplier.2", 100, supplier2)
		})
	})
}

func Test_LogServiceRegister(t *testing.T) {
	t.Run("NewLogServiceRegister", func(t *testing.T) {
		t.Run("create", func(t *testing.T) {
			if NewLogServiceRegister() == nil {
				t.Error("didn't returned a valid reference")
			}
		})

		t.Run("create with app reference", func(t *testing.T) {
			app := NewApp()
			if sut := NewLogServiceRegister(app); sut == nil {
				t.Error("didn't returned a valid reference")
			} else if sut.App != app {
				t.Error("didn't stored the app reference")
			}
		})
	})

	t.Run("Provide", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewLogServiceRegister(nil).Provide(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("register components", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewConfigServiceRegister(nil).Provide(container)
			sut := NewLogServiceRegister()

			e := sut.Provide(container)
			switch {
			case e != nil:
				t.Errorf("unexpected (%v) error", e)
			case !container.Has(LogJSONEncoderCreatorContainerID):
				t.Errorf("no JSON formatter creator : %v", sut)
			case !container.Has(LogAllFormatterCreatorsContainerID):
				t.Errorf("no formatter creators aggregator : %v", sut)
			case !container.Has(LogFormatterFactoryContainerID):
				t.Error("no logger formatter factory", e)
			case !container.Has(LogConsoleStreamCreatorContainerID):
				t.Errorf("no console writer creator : %v", sut)
			case !container.Has(LogFileStreamCreatorContainerID):
				t.Errorf("no file writer creator : %v", sut)
			case !container.Has(LogRotatingFileStreamCreatorContainerID):
				t.Errorf("no rotating file writer creator : %v", sut)
			case !container.Has(LogAllWriterCreatorsContainerID):
				t.Errorf("no writer creators aggregator : %v", sut)
			case !container.Has(LogWriterFactoryContainerID):
				t.Error("no logger writer factory", e)
			case !container.Has(LogContainerID):
				t.Error("no logger", e)
			case !container.Has(LogLoaderContainerID):
				t.Error("no logger loader", e)
			}
		})

		t.Run("retrieving logger JSON encoder creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewLogServiceRegister().Provide(container)

			factory, e := container.Get(LogJSONEncoderCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *LogJSONEncoderCreator:
				default:
					t.Error("didn't return a JSON encoder creator instance")
				}
			}
		})

		t.Run("retrieving formatter strategies", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewLogServiceRegister().Provide(container)

			formatterCreator := NewMockLogFormatterCreator(ctrl)
			_ = container.Add("formatter.id", func() LogFormatterCreator {
				return formatterCreator
			}, LogFormatterCreatorTag)

			creators, e := container.Get(LogAllFormatterCreatorsContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creators == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := creators.(type) {
				case []LogFormatterCreator:
					found := false
					for _, creator := range c {
						if creator == formatterCreator {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a formatter creator slice populated with the expected creator instance")
					}
				default:
					t.Error("didn't return a formatter creator slice")
				}
			}
		})

		t.Run("retrieving logger formatter factory", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewLogServiceRegister().Provide(container)

			factory, e := container.Get(LogFormatterFactoryContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *LogFormatterFactory:
				default:
					t.Error("didn't return a formatter factory")
				}
			}
		})

		t.Run("retrieving logger console stream creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewLogServiceRegister().Provide(container)

			factory, e := container.Get(LogConsoleStreamCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *LogConsoleStreamCreator:
				default:
					t.Error("didn't return a console stream creator instance")
				}
			}
		})

		t.Run("retrieving logger file stream creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewLogServiceRegister().Provide(container)

			factory, e := container.Get(LogFileStreamCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *LogFileStreamCreator:
				default:
					t.Error("didn't return a file stream creator instance")
				}
			}
		})

		t.Run("retrieving logger rotating file stream creator", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister(nil).Provide(container)
			_ = NewLogServiceRegister().Provide(container)

			factory, e := container.Get(LogRotatingFileStreamCreatorContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *LogRotatingFileStreamCreator:
				default:
					t.Error("didn't return a rotating file stream creator instance")
				}
			}
		})

		t.Run("retrieving stream strategies", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)

			writerCreator := NewMockLogWriterCreator(ctrl)
			_ = container.Add("writer.id", func() LogWriterCreator {
				return writerCreator
			}, LogWriterCreatorTag)

			creators, e := container.Get(LogAllWriterCreatorsContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case creators == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch c := creators.(type) {
				case []LogWriterCreator:
					found := false
					for _, creator := range c {
						if creator == writerCreator {
							found = true
						}
					}
					if !found {
						t.Error("didn't return a writer creator slice populated with the expected creator instance")
					}
				default:
					t.Error("didn't return a writer creator slice")
				}
			}
		})

		t.Run("retrieving logger stream factory", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewLogServiceRegister().Provide(container)

			factory, e := container.Get(LogWriterFactoryContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case factory == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch factory.(type) {
				case *LogWriterFactory:
				default:
					t.Error("didn't return a stream factory")
				}
			}
		})

		t.Run("retrieving logger", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewLogServiceRegister().Provide(container)

			log, e := container.Get(LogContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case log == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch log.(type) {
				case *Log:
				default:
					t.Error("didn't return a log")
				}
			}
		})

		t.Run("error retrieving config on retrieving logger loader", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewLogServiceRegister().Provide(container)
			_ = container.Add(ConfigContainerID, func() (*Config, error) {
				return nil, expected
			})

			if _, e := container.Get(LogLoaderContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving logger on retrieving logger loader", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)
			_ = container.Add(LogContainerID, func() (*Log, error) {
				return nil, expected
			})

			if _, e := container.Get(LogLoaderContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("error retrieving stream factory on retrieving logger loader", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)
			_ = container.Add(LogWriterFactoryContainerID, func() (*LogWriterFactory, error) {
				return nil, expected
			})

			if _, e := container.Get(LogLoaderContainerID); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("retrieving logger loader", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			_ = NewLogServiceRegister().Provide(container)

			loader, e := container.Get(LogLoaderContainerID)
			switch {
			case e != nil:
				t.Errorf("unexpected error (%v)", e)
			case loader == nil:
				t.Error("didn't returned a valid reference")
			default:
				switch loader.(type) {
				case *LogLoader:
				default:
					t.Error("didn't return a loader reference")
				}
			}
		})
	})

	t.Run("Boot", func(t *testing.T) {
		t.Run("nil container", func(t *testing.T) {
			if e := NewLogServiceRegister().Boot(nil); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrNilPointer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrNilPointer)
			}
		})

		t.Run("valid simple boot with strategies (no loader)", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			LogLoaderActive = false
			defer func() { LogLoaderActive = true }()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			sut := NewLogServiceRegister()
			_ = sut.Provide(container)

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected error (%v)", e)
			}
		})

		t.Run("don't run loader if globally configured so", func(t *testing.T) {
			LogLoaderActive = false
			defer func() { LogLoaderActive = true }()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			sut := NewLogServiceRegister()
			_ = sut.Provide(container)
			_ = container.Add(LogLoaderContainerID, func() (interface{}, error) {
				panic(fmt.Errorf("error message"))
			})

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})

		t.Run("error retrieving loader", func(t *testing.T) {
			expected := fmt.Errorf("error message")
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			sut := NewLogServiceRegister()
			_ = sut.Provide(container)
			_ = container.Add(LogLoaderContainerID, func() (*LogLoader, error) {
				return nil, expected
			})

			if e := sut.Boot(container); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrServiceContainer) {
				t.Errorf("(%v) when expecting (%v)", e, ErrServiceContainer)
			}
		})

		t.Run("invalid loader", func(t *testing.T) {
			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			sut := NewLogServiceRegister()
			_ = sut.Provide(container)
			_ = container.Add(LogLoaderContainerID, func() interface{} {
				return "string"
			})

			if e := sut.Boot(container); e == nil {
				t.Error("didn't returned the expected error")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("invalid logger entry config", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			sut := NewLogServiceRegister()
			_ = sut.Provide(container)

			partial := ConfigPartial{}
			_, _ = partial.Set("slate.log.writers", "string")
			source := NewMockConfigSupplier(ctrl)
			source.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 1, source)
			_ = container.Add(ConfigContainerID, func() *Config {
				return config
			})

			if e := sut.Boot(container); e == nil {
				t.Errorf("didn't returned the expected error")
			} else if !errors.Is(e, ErrConversion) {
				t.Errorf("(%v) when expecting (%v)", e, ErrConversion)
			}
		})

		t.Run("correct boot", func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			container := NewServiceContainer()
			_ = NewFileSystemServiceRegister().Provide(container)
			_ = NewConfigServiceRegister().Provide(container)
			sut := NewLogServiceRegister()
			_ = sut.Provide(container)

			partial := ConfigPartial{}
			_, _ = partial.Set("slate.log.writers", ConfigPartial{})
			source := NewMockConfigSupplier(ctrl)
			source.EXPECT().Get("").Return(partial, nil).Times(1)
			config := NewConfig()
			_ = config.AddSupplier("id", 1, source)
			_ = container.Add(ConfigContainerID, func() *Config {
				return config
			})

			if e := sut.Boot(container); e != nil {
				t.Errorf("unexpected (%v) error", e)
			}
		})
	})
}
