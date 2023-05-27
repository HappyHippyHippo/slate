package console

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/happyhippyhippo/slate"
	"github.com/happyhippyhippo/slate/log"
)

func Test_NewStream(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	formatter := NewMockFormatter(ctrl)
	var channels []string
	level := log.WARNING

	t.Run("nil formatter", func(t *testing.T) {
		sut, e := NewStream(nil, channels, level)
		switch {
		case sut != nil:
			t.Error("returned a valid reference")
		case e == nil:
			t.Error("didn't returned the expected error")
		case !errors.Is(e, slate.ErrNilPointer):
			t.Errorf("returned the (%v) error when expecting (%v)", e, slate.ErrNilPointer)
		}
	})

	t.Run("new console stream", func(t *testing.T) {
		sut, e := NewStream(formatter, []string{}, log.WARNING)
		switch {
		case sut == nil:
			t.Error("didn't returned a valid reference")
		case e != nil:
			t.Errorf("returned the (%v) error", e)
		case sut.Writer != os.Stdout:
			t.Error("didn't stored the stdout as the defined writer")
		}
	})
}

func Test_Stream_Signal(t *testing.T) {
	t.Run("signal message to the writer without context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    log.Level
			}
			call struct {
				level   log.Level
				channel string
				message string
			}
			callTimes int
			expected  string
		}{
			{ // signal through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					message string
				}{
					level:   log.FATAL,
					channel: "dummy_channel",
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with a filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					message string
				}{
					level:   log.DEBUG,
					channel: "dummy_channel",
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with an unregistered channel
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					message string
				}{
					level:   log.FATAL,
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

				sut, _ := NewStream(formatter, scenario.state.channels, scenario.state.level)
				sut.Writer = writer

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
				level    log.Level
			}
			call struct {
				level   log.Level
				channel string
				ctx     log.Context
				message string
			}
			callTimes int
			expected  string
		}{
			{ // signal through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					ctx     log.Context
					message string
				}{
					level:   log.FATAL,
					channel: "dummy_channel",
					ctx:     log.Context{},
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with a filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					ctx     log.Context
					message string
				}{
					level:   log.DEBUG,
					channel: "dummy_channel",
					ctx:     log.Context{},
					message: "dummy_message",
				},
				callTimes: 0,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // signal through a valid channel with an unregistered channel
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					channel string
					ctx     log.Context
					message string
				}{
					level:   log.FATAL,
					channel: "not_a_valid_dummy_channel",
					ctx:     log.Context{},
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

				sut, _ := NewStream(formatter, scenario.state.channels, scenario.state.level)
				sut.Writer = writer

				if e := sut.Signal(scenario.call.channel, scenario.call.level, scenario.call.message, scenario.call.ctx); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})
}

func Test_Stream_Broadcast(t *testing.T) {
	t.Run("broadcast message to the writer without context", func(t *testing.T) {
		scenarios := []struct {
			state struct {
				channels []string
				level    log.Level
			}
			call struct {
				level   log.Level
				message string
			}
			callTimes int
			expected  string
		}{
			{ // broadcast through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					message string
				}{
					level:   log.FATAL,
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // broadcast through a valid channel with a filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					message string
				}{
					level:   log.DEBUG,
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

				sut, _ := NewStream(formatter, scenario.state.channels, scenario.state.level)
				sut.Writer = writer

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
				level    log.Level
			}
			call struct {
				level   log.Level
				ctx     log.Context
				message string
			}
			callTimes int
			expected  string
		}{
			{ // broadcast through a valid channel with a not filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					ctx     log.Context
					message string
				}{
					ctx:     log.Context{},
					level:   log.FATAL,
					message: "dummy_message",
				},
				callTimes: 1,
				expected:  `{"message" : "dummy_message"}`,
			},
			{ // broadcast through a valid channel with a filtered level
				state: struct {
					channels []string
					level    log.Level
				}{
					channels: []string{"dummy_channel"},
					level:    log.WARNING,
				},
				call: struct {
					level   log.Level
					ctx     log.Context
					message string
				}{
					level:   log.DEBUG,
					ctx:     log.Context{},
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

				sut, _ := NewStream(formatter, scenario.state.channels, scenario.state.level)
				sut.Writer = writer

				if e := sut.Broadcast(scenario.call.level, scenario.call.message, scenario.call.ctx); e != nil {
					t.Errorf("returned the (%v) error", e)
				}
			}
			test()
		}
	})
}
