package log

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate/err"
)

func Test_NewConsoleStream(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	formatter := NewMockFormatter(ctrl)
	var channels []string
	level := WARNING

	t.Run("nil formatter", func(t *testing.T) {
		sut, e := NewConsoleStream(nil, channels, level)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, err.NilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, err.NilPointer)
		}
	})

	t.Run("new console stream", func(t *testing.T) {
		sut, e := NewConsoleStream(formatter, []string{}, WARNING)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut.writer != os.Stdout:
			t.Error("didn't stored the stdout as the defined writer")
		}
	})
}

func Test_ConsoleStream_Signal(t *testing.T) {
	t.Run("signal message to the writer without context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				channel string
				message string
			}
			callTimes int
			expected  string
		}{
			{ // signal through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
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
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
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
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
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

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message).Return(scenario.expected).Times(scenario.callTimes)

				sut, _ := NewConsoleStream(formatter, scenario.state.channels, scenario.state.level)
				sut.writer = writer

				if e := sut.Signal(scenario.call.channel, scenario.call.level, scenario.call.message); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})

	t.Run("signal message to the writer with context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				channel string
				ctx     Context
				message string
			}
			callTimes int
			expected  string
		}{
			{ // signal through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					channel string
					ctx     Context
					message string
				}{
					level:   FATAL,
					channel: "dummy_channel",
					ctx:     Context{},
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with a filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					channel string
					ctx     Context
					message string
				}{
					level:   DEBUG,
					channel: "dummy_channel",
					ctx:     Context{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with an unregistered channel
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					channel string
					ctx     Context
					message string
				}{
					level:   FATAL,
					channel: "not_a_valid_dummy_channel",
					ctx:     Context{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message, scenario.call.ctx).Return(scenario.expected).Times(scenario.callTimes)

				sut, _ := NewConsoleStream(formatter, scenario.state.channels, scenario.state.level)
				sut.writer = writer

				if e := sut.Signal(scenario.call.channel, scenario.call.level, scenario.call.message, scenario.call.ctx); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})
}

func Test_ConsoleStream_Broadcast(t *testing.T) {
	t.Run("broadcast message to the writer without context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				message string
			}
			callTimes int
			expected  string
		}{
			{ // broadcast through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
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
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					message string
				}{
					level:   DEBUG,
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message).Return(scenario.expected).Times(scenario.callTimes)

				sut, _ := NewConsoleStream(formatter, scenario.state.channels, scenario.state.level)
				sut.writer = writer

				if e := sut.Broadcast(scenario.call.level, scenario.call.message); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})

	t.Run("broadcast message to the writer with context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    Level
			}
			call struct {
				level   Level
				ctx     Context
				message string
			}
			callTimes int
			expected  string
		}{
			{ // broadcast through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					ctx     Context
					message string
				}{
					ctx:     Context{},
					level:   FATAL,
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // broadcast through a valid channel with a filtered level
				state: struct {
					channels []string
					level    Level
				}{
					channels: []string{"dummy_channel"},
					level:    WARNING,
				},
				call: struct {
					level   Level
					ctx     Context
					message string
				}{
					level:   DEBUG,
					ctx:     Context{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
		}

		for _, scenario := range scenarios {
			test := func() {
				ctrl := gomock.NewController(t)
				defer func() { ctrl.Finish() }()

				writer := NewMockWriter(ctrl)
				writer.EXPECT().Write([]byte(scenario.expected + "\n")).Times(scenario.callTimes)
				formatter := NewMockFormatter(ctrl)
				formatter.EXPECT().Format(scenario.call.level, scenario.call.message, scenario.call.ctx).Return(scenario.expected).Times(scenario.callTimes)

				sut, _ := NewConsoleStream(formatter, scenario.state.channels, scenario.state.level)
				sut.writer = writer

				if e := sut.Broadcast(scenario.call.level, scenario.call.message, scenario.call.ctx); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})
}
